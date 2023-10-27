package main

import (
	"log"
	"net"
	"net/http"
	"time"

	conf "main/config"
	src "main/src"
	proto "main/src/proto"

	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"

	_ "main/docs"

	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc/keepalive"
)

func loggingAndCORSHeadersMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method)
		for header := range conf.Headers {
			w.Header().Set(header, conf.Headers[header])
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	myRouter := mux.NewRouter()

	urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	//config, _ := sql.Open("pgx", os.Getenv(conf.UrlDB))
	db, err := sql.Open("pgx", urlDB)
	if err != nil {
		log.Println("could not connect to database")
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Println("unable to reach database ", err)
	} else {
		log.Println("database is reachable")
	}

	hub := src.NewHub()
	go hub.Run()

	Store := src.NewStore(db)
	Handler := src.NewHandler(Store, hub)

	//myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { log.Println("main page") })
	myRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { src.ServeWs(hub, w, r) })
	myRouter.HandleFunc("/api/attach", Handler.UploadFile).Methods(http.MethodPost, http.MethodOptions)
	myRouter.PathPrefix("/api/docs").Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	lis, err := net.Listen("tcp", conf.PortGRPC)
	if err != nil {
		log.Println("cant listen grpc port", err)
	}
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024),
		grpc.MaxConcurrentStreams(35),
		grpc.KeepaliveParams(keepalive.ServerParameters{Time: 1 * time.Second, Timeout: 5 * time.Second}),
	)
	proto.RegisterBotChatServer(server, src.NewChatManager(Store, hub))
	log.Println("starting grpc server at :8082")
	go server.Serve(lis)

	log.Println("starting web server at " + conf.PortWS)
	err = http.ListenAndServe(conf.PortWS, myRouter)

	if err != nil {
		log.Println("cant serve", err)
	}

}
