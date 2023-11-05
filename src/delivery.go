package chat

import (
	"encoding/hex"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"

	e "main/domain/errors"
	"main/domain/model"
	m "main/domain/model"

	"crypto/sha256"
)

var mockTeacherID = 1

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
// @BasePath  /api

type Handler struct {
	store StoreInterface
	hub   *Hub
}

func NewHandler(store StoreInterface, hub *Hub) *Handler {
	return &Handler{
		store: store,
		hub:   hub,
	}
}

func ReturnErrorJSON(w http.ResponseWriter, err error, errCode int) {
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: err.Error()})
}

// SetAttach godoc
// @Summary Set teacher's attach
// @Description sets teacher's attach
// @ID setAttach
// @Accept  multipart/form-data
// @Produce  json
// @Param file formData file true "teacher's attach"
// @Param text query string false "text"
// @Param chatid query string true "Chat id"
// @Success 200 {object} model.Response "ok"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal Server Error - Request is valid but operation failed at server side"
// @Router /attach [post]
func (api *Handler) SetAttach(w http.ResponseWriter, r *http.Request) {
	textS := r.URL.Query().Get("text")
	chatIDs := r.URL.Query().Get("chatid")

	text, err := url.QueryUnescape(textS)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, e.ErrBadRequest400, 400)
	}

	chatID, err := strconv.Atoi(chatIDs)
	if err != nil {
		log.Println("error: ", err)
		ReturnErrorJSON(w, e.ErrBadRequest400, 400)
		return
	}

	file, _, err := r.FormFile("file")
	if err != nil {
		log.Println("error parse file")
		ReturnErrorJSON(w, e.ErrServerError500, 500)
		return
	}
	defer file.Close()

	s := chatIDs + time.Now().Format("2006.01.02 15:04:05")

	h := sha256.New()

	h.Write([]byte(s))
	attachNum := hex.EncodeToString(h.Sum(nil))

	fileName := "./filestorage/attaches/attach" + attachNum + ".jpg"
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Println("error create/open file")
		ReturnErrorJSON(w, e.ErrServerError500, 500)
		return
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		log.Println("error copy to new file")
		ReturnErrorJSON(w, e.ErrServerError500, 500)
		return
	}

	fileAddr := "http://127.0.0.1:8081/filestorage/attaches/attach" + attachNum + ".jpg"
	mes := &m.MessageWebsocket{Text: text + "\n" + fileAddr, ChatID: int32(chatID)}
	log.Println("Sending mes with attach to bot: ", "text:", mes.Text, "chatid:", mes.ChatID)

	if mes.ChatID == 1 {
		api.hub.MessagesToTGBot <- mes
	} else if mes.ChatID == 2 {
		api.hub.MessagesToVKBot <- mes
	}

	err = api.store.AddMessage(&m.CreateMessage{Text: mes.Text, ChatID: int(mes.ChatID), IsAuthorTeacher: true, IsRead: false})
	if err != nil {
		log.Println(err)
		ReturnErrorJSON(w, e.ErrServerError500, 500)
		return
	}

	json.NewEncoder(w).Encode(&model.Response{})
}
