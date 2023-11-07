package BaseConfig

var UrlDB = "URL_DB"

var BaseFilestorage = "BASE_FILESTORAGE"
var PrefixFilestorage = "PREFIX_FILESTORAGE"

var PortWS = ":8081"
var PortGRPC = ":8082"

var BaseUrl = "/apichat"
var PathWS = BaseUrl + "/ws"
var PathAttach = BaseUrl + "/attach"
var PathDocs = BaseUrl + "/docs"

var Headers = map[string]string{
	"Access-Control-Allow-Origin":      "http://127.0.0.1:8001",
	"Access-Control-Allow-Credentials": "true",
	"Access-Control-Allow-Headers":     "Origin, Content-Type, accept, csrf",
	"Access-Control-Allow-Methods":     "GET, POST, DELETE, OPTIONS",
	"Content-Type":                     "application/json",
}
