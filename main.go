package main

import (
	"database/sql"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	conf "main/config"
	grpcCalendar "main/delivery/grpc/calendar"
	protoCalendar "main/delivery/grpc/calendar/proto"
	grpcChat "main/delivery/grpc/chat"
	protoChat "main/delivery/grpc/chat/proto"
	ws "main/delivery/ws"
	localStore "main/repository/local-storage"
	pgstore "main/repository/pg"
	chatusecase "main/usecase/chat"

	_ "main/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

var urlDB string
var filestoragePath string
var chatFilesPath string
var homeworkFilesPath string
var solutionFilesPath string
var urlDomain string
var calendarGrpcUrl string

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

	calendarGrpcUrl, exist = os.LookupEnv(conf.CALENDAR_GRPC_URL)
	if !exist || len(calendarGrpcUrl) == 0 {
		log.Fatalln("could not get calendar grpc url from env")
	}

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

	filestoragePath, exist = os.LookupEnv(conf.FILESTORAGE_PATH)
	if !exist || len(filestoragePath) == 0 {
		log.Fatalln("could not get filestorage path from env")
	}

	chatFilesPath, exist = os.LookupEnv(conf.CHAT_FILES_PATH)
	if !exist || len(chatFilesPath) == 0 {
		log.Fatalln("could not get chat files path from env")
	}

	homeworkFilesPath, exist = os.LookupEnv(conf.HOMEWORK_FILES_PATH)
	if !exist || len(homeworkFilesPath) == 0 {
		log.Fatalln("could not get homework files path from env")
	}

	solutionFilesPath, exist = os.LookupEnv(conf.SOLUTION_FILES_PATH)
	if !exist || len(solutionFilesPath) == 0 {
		log.Fatalln("could not get solution files path from env")
	}

	urlDomain, exist = os.LookupEnv(conf.URL_DOMAIN)
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

	hub := ws.NewHub()
	go hub.Run()

	myRouter.HandleFunc(conf.PathWS, hub.AddConnection).Methods(http.MethodGet, http.MethodOptions)

	myRouter.PathPrefix(conf.PathDocs).Handler(httpSwagger.WrapHandler)
	myRouter.Use(loggingAndCORSHeadersMiddleware)

	grcpConnCalendar, err := grpc.Dial(
		calendarGrpcUrl,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("cant create connecter to grpc calendar")
	}
	log.Println("connecter to grpc calendar service is created")
	defer grcpConnCalendar.Close()

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

	dataStore := pgstore.NewPostgreSqlStore(db)
	fileStore := localStore.NewLocalStore(
		chatFilesPath,
		homeworkFilesPath,
		solutionFilesPath,
		filestoragePath,
	)

	calendar := grpcCalendar.NewCalendarService(
		protoCalendar.NewCalendarClient(grcpConnCalendar),
	)

	usecase := chatusecase.NewChatUsecase(
		hub,
		dataStore,
		fileStore,
		calendar,
		filestoragePath,
		urlDomain,
	)

	grpcHandler := grpcChat.NewChatGrpcHander(usecase)
	protoChat.RegisterChatServer(server, grpcHandler)

	log.Println("starting grpc server at " + conf.PortGRPC)
	go server.Serve(lis)

	log.Println("starting web server at " + conf.PortWS)
	err = http.ListenAndServe(conf.PortWS, myRouter)

	if err != nil {
		log.Fatalln("cant serve", err)
	}
}
