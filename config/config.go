package BaseConfig

var PG_USER = "POSTGRES_USER"
var PG_HOST = "POSTGRES_HOST"
var PG_PWD = "POSTGRES_PASSWORD"
var PG_PORT = "POSTGRES_PORT"
var PG_DB = "POSTGRES_DB"
var FilestoragePath = "FILESTORAGE_PATH"
var UrlDomain = "URL_DOMAIN"

var PortWS = ":8081"
var PortGRPC = ":8082"

var BaseUrl = "/apichat"
var PathWS = BaseUrl + "/ws"

var PathDocs = BaseUrl + "/docs"

var Headers = map[string]string{
	"Access-Control-Allow-Origin":      "http://127.0.0.1:8001",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept, csrf",
	"Access-Control-Allow-Methods":     "GET, POST, DELETE, OPTIONS",
	"Content-Type":                     "application/json",
}
