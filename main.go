package main

import (
	"log"
	"net"
	"net/http"
	"time"

	src "main/src"
	proto "main/src/proto"

	conf "main/config"

	"github.com/gorilla/mux"

	"database/sql"

	_ "github.com/jackc/pgx/v4/stdlib"
	"google.golang.org/grpc"

	"google.golang.org/grpc/keepalive"
)

func main() {
	myRouter := mux.NewRouter()

	urlDB := "postgres://" + conf.DBSPuser + ":" + conf.DBPassword + "@" + conf.DBHost + ":" + conf.DBPort + "/" + conf.DBName
	//urlDB := "postgres://" + os.Getenv("TEST_POSTGRES_USER") + ":" + os.Getenv("TEST_POSTGRES_PASSWORD") + "@" + os.Getenv("TEST_DATABASE_HOST") + ":" + os.Getenv("DB_PORT") + "/" + os.Getenv("TEST_POSTGRES_DB")
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
	myRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { log.Println("main page") })
	myRouter.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) { src.ServeWs(hub, w, r) })

	lis, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Println("cant listen grpc port", err)
	}
	server := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024),
		grpc.MaxConcurrentStreams(35),
		grpc.KeepaliveParams(keepalive.ServerParameters{Time: 1 * time.Second, Timeout: 5 * time.Second}),
	)
	proto.RegisterBotChatServer(server, src.NewChatManager(db, hub))
	log.Println("starting grpc server at :8082")
	go server.Serve(lis)

	log.Println("starting web server at :8081")
	err = http.ListenAndServe(":8081", myRouter)

	if err != nil {
		log.Println("cant serve", err)
	}

}
