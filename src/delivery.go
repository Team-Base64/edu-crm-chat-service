package chat

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"os"

	e "main/domain/errors"
	"main/domain/model"
	m "main/domain/model"

	"github.com/google/uuid"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var chatFilesPath = "/chat"
var homeworkFilesPath = "/homework"
var solutionFilesPath = "/solution"

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
	for _, path := range []string{chatFilesPath, homeworkFilesPath, solutionFilesPath} {
		if err := os.MkdirAll(fs+path, os.ModePerm); err != nil {
			log.Fatalln(e.StacktraceError(err))
		}
	}

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
// @Param type query string true "type: homework or solution or chat"
// @Success 200 {object} model.Response "ok"
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal Server Error - Request is valid but operation failed at server side"
// @Router /attach [post]
func (api *Handler) UploadFile(w http.ResponseWriter, r *http.Request) {
	typeS := r.URL.Query().Get("type")

	filePath := ""
	switch typeS {
	case "homework":
		filePath = api.filestorage + chatFilesPath
	case "solution":
		filePath = api.filestorage + solutionFilesPath
	case "chat":
		filePath = api.filestorage + chatFilesPath
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

	json.NewEncoder(w).Encode(&model.UploadAttachResponse{File: "http://127.0.0.1:8080/" + fileName[1:]})
}

// SetOAUTH2Token godoc
// @Summary Sets teacher's OAUTH2Token
// @Description Sets teacher's OAUTH2Token
// @ID SetOAUTH2Token
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /oauth [post]
func (api *Handler) SetOAUTH2Token(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var tok model.OAUTH2Token
	if err := decoder.Decode(&tok); err != nil {
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}
	json.NewEncoder(w).Encode(&model.Response{})
}

// CreateCalendar godoc
// @Summary Creates teacher's calendar
// @Description Creates teacher's calendar
// @ID CreateCalendar
// @Accept  json
// @Produce  json
// @Success 200 {object} model.CreateCalendarResponse
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar [post]
func (api *Handler) CreateCalendar(w http.ResponseWriter, r *http.Request) {
	// decoder := json.NewDecoder(r.Body)
	// var tok model.OAUTH2Token
	// if err := decoder.Decode(&tok); err != nil {
	// 	ReturnErrorJSON(w, e.ErrBadRequest400)
	// 	return
	// }
	mockTeackerID := 1
	// if _, err := api.store.GetTokenDB(mockTeackerID); err != nil {
	// 	log.Println(err)
	// 	ReturnErrorJSON(w, e.ErrUnauthorized401)
	// 	return
	// }

	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}

	client, err := getClient(config)
	if err != nil {
		log.Println("Unable to get client from token: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}

	newCal := &calendar.Calendar{TimeZone: "Europe/Moscow", Summary: "EDUCRM Calendar"}
	cal, err := srv.Calendars.Insert(newCal).Do()
	if err != nil {
		log.Println("Unable to create calendar: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}
	//fmt.Printf("Created Cal: %s\n", cal.Id)
	innerID, err := api.store.CreateCalendarDB(mockTeackerID, cal.Id)
	if err != nil {
		log.Println("DB err: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}

	Acl := &calendar.AclRule{Scope: &calendar.AclRuleScope{Type: "default"}, Role: "reader"}
	_, err = srv.Acl.Insert(cal.Id, Acl).Do()
	if err != nil {
		log.Println("Unable to create ACL: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}
	//fmt.Printf("Event created: %s\n", event.Id)

	json.NewEncoder(w).Encode(&model.CreateCalendarResponse{ID: innerID, IDInGoogle: cal.Id})
}

// CreateCalendarEvent godoc
// @Summary Creates teacher's calendar event
// @Description Creates teacher's calendar event
// @ID CreateCalendarEvent
// @Accept  json
// @Produce  json
// @Success 200 {object} model.Response
// @Failure 401 {object} model.Error "unauthorized - Access token is missing or invalid"
// @Failure 500 {object} model.Error "internal server error - Request is valid but operation failed at server side"
// @Router /calendar/addevent [post]
func (api *Handler) CreateCalendarEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		return
	}
	mockTeackerID := 1

	decoder := json.NewDecoder(r.Body)
	var req model.CreateCalendarEvent
	if err := decoder.Decode(&req); err != nil {
		ReturnErrorJSON(w, e.ErrBadRequest400)
		return
	}

	// if _, err := api.store.GetTokenDB(mockTeackerID); err != nil {
	// 	ReturnErrorJSON(w, e.ErrUnauthorized401)
	// 	return
	// }

	ctx := context.Background()
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Println("Unable to read client secret file: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
		return
	}

	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Println("Unable to parse client secret file to config: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}

	client, err := getClient(config)
	if err != nil {
		log.Println("Unable to get client from token: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Println("Unable to retrieve calendar Client: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}

	event := &calendar.Event{
		Summary:     req.Title,
		Description: req.Description,
		Start: &calendar.EventDateTime{
			DateTime: req.StartDate,
			TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: req.EndDate,
			TimeZone: "Europe/Moscow",
		},
		Visibility: "public",
	}
	calendarID, err := api.store.GetCalendarGoogleID(mockTeackerID)
	if err != nil {
		log.Println("DB err: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}
	log.Println(calendarID, req.StartDate, req.EndDate)
	event, err = srv.Events.Insert(calendarID, event).Do()
	if err != nil {
		log.Println("Unable to create event: ", err)
		ReturnErrorJSON(w, e.ErrServerError500)
	}

	MockClassID := 1
	ids, err := api.store.GetChatsByClassID(MockClassID)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(" Broadcast for event chats: ", ids)
	for _, id := range *ids {
		type1, err := api.store.GetTypeByChatID(id)
		if err != nil {
			log.Println("err with mes into chat ", id, " : ", err)
			//return &proto.Nothing{}, err
		}
		switch type1 {
		case "tg":
			api.hub.MessagesToTGBot <- &m.MessageWebsocket{ChatID: int32(id), Text: "Новое событие!" + "\n" + req.Title + "\n" + req.Description + "\n" + "Начало: " + req.StartDate + "\n" + "Окончание: " + req.EndDate, AttachmentURLs: []string{}}
		case "vk":
			api.hub.MessagesToVKBot <- &m.MessageWebsocket{ChatID: int32(id), Text: "Новое событие!" + "\n" + req.Title + "\n" + req.Description + "\n" + "Начало: " + req.StartDate + "\n" + "Окончание: " + req.EndDate, AttachmentURLs: []string{}}
		default:
		}
		// err = sm.store.AddMessage(&m.CreateMessage{Text: req.Title + "\n" + req.Description, ChatID: id, IsAuthorTeacher: true, IsRead: false, AttachmentURLs: req.AttachmentURLs})
		// if err != nil {
		// 	log.Println("err with mes into chat ", id, " : ", err)
		// 	//return err
		// }

	}

	json.NewEncoder(w).Encode(&model.Response{})
}
