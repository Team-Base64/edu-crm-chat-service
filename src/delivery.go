package chat

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	e "main/domain/errors"
	"main/domain/model"

	"github.com/google/uuid"
)

// @title TCRA API
// @version 1.0
// @description EDUCRM back chat server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8081
// @BasePath  /apichat

type Handler struct {
	store       StoreInterface
	hub         *Hub
	filestorage string
	prefix      string
}

func NewHandler(store StoreInterface, hub *Hub, fs string, pf string) *Handler {
	return &Handler{
		store:       store,
		hub:         hub,
		filestorage: fs,
		prefix:      pf,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error) {
	errCode, errText := e.CheckError(err)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: errText})
}

// UploadAttach godoc
// @Summary Upload attach
// @Description Upload attach
// @ID uploadAttach
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "attach"
// @Param type query string true "type"
// @Success 200 {object} model.Response "ok"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal Server Error - Request is valid but operation failed at server side"
// @Router /attach [post]
func (api *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	typeS := r.URL.Query().Get("type")

	filePath := ""
	switch typeS {
	case "homework":
		filePath = api.filestorage + "/homeworks"
	case "solution":
		filePath = api.filestorage + "/solutions"
	case "chat":
		filePath = api.filestorage + "/chat"
	default:
		log.Println(e.StacktraceError(errors.New("error wrong type query param")))
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	defer file.Close()

	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	log.Println(http.DetectContentType(fileHeader))
	fileExt := ""
	switch http.DetectContentType(fileHeader) {
	case "image/jpeg":
		fileExt = ".jpg"
	case "image/png":
		fileExt = ".png"
	case "application/pdf":
		fileExt = ".pdf"
	default:
		log.Println("error not allowed file extension")
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	attachNum := uuid.New().String()

	fileName := filePath + "/" + attachNum + fileExt
	if err := os.MkdirAll(filePath, os.ModePerm); err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}
	f, err := os.OpenFile(api.prefix+fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println(e.StacktraceError(err))
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}

	json.NewEncoder(w).Encode(&model.UploadAttachResponse{File: fileName})
}
