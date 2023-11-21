package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	conf "main/config"
	src "main/src"
	proto "main/src/proto"

	"github.com/gorilla/mux"

	"google.golang.org/grpc"

	_ "main/docs"

	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc/keepalive"
)

var urlDB string
var filestoragePath string
var urlDomain string

func loggingAndCORSHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		for header := range conf.Headers {
			w.Header().Set(header, conf.Headers[header])
		}

		next.ServeHTTP(w, r)
	})
}

func init() {
	var exist bool

	pgUser, exist := os.LookupEnv(conf.PG_USER)
	if !exist || len(pgUser) == 0 {
		log.Fatalln("could not get database host from env")
	}
	pgPwd, exist := os.LookupEnv(conf.PG_PWD)
	if !exist || len(pgPwd) == 0 {
		log.Fatalln("could not get database password from env")
	}
	pgHost, exist := os.LookupEnv(conf.PG_HOST)
	if !exist || len(pgHost) == 0 {
		log.Fatalln("could not get database host from env")
	}
	pgPort, exist := os.LookupEnv(conf.PG_PORT)
	if !exist || len(pgPort) == 0 {
		log.Fatalln("could not get database port from env")
	}
	pgDB, exist := os.LookupEnv(conf.PG_DB)
	if !exist || len(pgDB) == 0 {
		log.Fatalln("could not get database name from env")
	}

	urlDB = "postgres://" + pgUser + ":" + pgPwd + "@" + pgHost + ":" + pgPort + "/" + pgDB

	filestoragePath, exist = os.LookupEnv(conf.FilestoragePath)
	if !exist || len(filestoragePath) == 0 {
		log.Fatalln("could not get filestorage path from env")
	}

	urlDomain, exist = os.LookupEnv(conf.UrlDomain)
	if !exist || len(urlDomain) == 0 {
		log.Fatalln("could not get url domain from env")
	}

}

func main() {
	myRouter := mux.NewRouter()

	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Fatalln("could not connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln("unable to reach database ", err)
	}
	log.Println("database is reachable")

	hub := src.NewHub()
	go hub.Run()

	Store := src.NewStore(db)

	myRouter.HandleFunc(conf.PathWS, func(w http.ResponseWriter, r *http.Request) { src.ServeWs(hub, w, r) })
	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	lis, err := net.Listen("tcp", conf.PortGRPC)
	if err != nil {
		log.Fatalln("cant listen grpc port", err)
	}
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024),
		grpc.MaxConcurrentStreams(35),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			Time:    1 * time.Second,
			Timeout: 5 * time.Second,
		}),
	)
	proto.RegisterBotChatServer(
		server,
		src.NewChatManager(
			Store,
			hub,
			filestoragePath,
			urlDomain,
		),
	)
	log.Println("starting grpc server at " + conf.PortGRPC)
	go server.Serve(lis)

	log.Println("starting web server at " + conf.PortWS)
	err = http.ListenAndServe(conf.PortWS, myRouter)

	if err != nil {
		log.Fatalln("cant serve", err)
	}

}
