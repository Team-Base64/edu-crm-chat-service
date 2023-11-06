package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	conf "main/config"
	src "main/src"
	proto "main/src/proto"

	"github.com/gorilla/mux"

	"github.com/jackc/pgx/v5"
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

	db, err := pgx.Connect(context.Background(), os.Getenv(conf.UrlDB))
	if err != nil {
		log.Fatalln("could not connect to database")
	}
	defer db.Close(context.Background())

	if err := db.Ping(context.Background()); err != nil {
		log.Fatalln("unable to reach database ", err)
	}
	log.Println("database is reachable")

	hub := src.NewHub()
	go hub.Run()

	Store := src.NewStore(db)
	Handler := src.NewHandler(Store, hub, os.Getenv(conf.BaseFilestorage))

	myRouter.HandleFunc(conf.PathWS, func(w http.ResponseWriter, r *http.Request) { src.ServeWs(hub, w, r) })
	myRouter.HandleFunc(conf.PathAttach, Handler.UploadFile).Methods(http.MethodPost, http.MethodOptions)
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
	proto.RegisterBotChatServer(server, src.NewChatManager(Store, hub))
	log.Println("starting grpc server at " + conf.PortGRPC)
	go server.Serve(lis)

	log.Println("starting web server at " + conf.PortWS)
	err = http.ListenAndServe(conf.PortWS, myRouter)

	if err != nil {
		log.Fatalln("cant serve", err)
	}

}
