package chat

import (
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	d "main/delivery"
	e "main/domain/errors"
	m "main/domain/model"
	rep "main/repository"
	uc "main/usecase"

	"github.com/google/uuid"
)

type ChatUsecase struct {
	hub             d.HubInterface
	store           rep.StoreInterface
	calendar        d.CalendarInterface
	filestoragePath string
	urlDomain       string
}

func NewChatUsecase(
	hud d.HubInterface,
	store rep.StoreInterface,
	calendar d.CalendarInterface,
	fsPath string,
	urlDomain string,
) uc.UsecaseInterface {
	return &ChatUsecase{
		hub:             hud,
		store:           store,
		calendar:        calendar,
		filestoragePath: fsPath,
		urlDomain:       urlDomain,
	}
}

func (uc *ChatUsecase) SaveMsg(msg *m.MessageWebsocket) error {
	if err := uc.store.AddMessage(&m.CreateMessage{
		Text:            msg.Text,
		ChatID:          int(msg.ChatID),
		IsAuthorTeacher: true,
		IsRead:          true,
		CreateTime:      time.Now(),
		AttachmentURLs:  msg.AttachmentURLs,
	}); err != nil {
		return e.StacktraceError(err)
	}
	return nil
}

func (uc *ChatUsecase) GetMsgForTG() (*m.MessageWebsocket, error) {
	msg := uc.hub.GetMsgForTG()

	if msg.Text == "" && msg.AttachmentURLs == nil {
		return nil, nil
	}

	if msg.IsSavedToDB {
		if err := uc.SaveMsg(msg); err != nil {
			return nil, e.StacktraceError(err)
		}
	}

	return msg, nil
}

func (uc *ChatUsecase) GetMsgForVK() (*m.MessageWebsocket, error) {
	msg := uc.hub.GetMsgForVK()

	if msg.Text == "" && msg.AttachmentURLs == nil {
		return nil, nil
	}

	if msg.IsSavedToDB {
		if err := uc.SaveMsg(msg); err != nil {
			return nil, e.StacktraceError(err)
		}
	}

	return msg, nil
}

func (uc *ChatUsecase) SendMsgToClient(msg *m.Message, social string) error {
	createTime := time.Now()
	if err := uc.store.AddMessage(&m.CreateMessage{
		Text:            msg.Text,
		ChatID:          msg.ChatID,
		IsAuthorTeacher: false,
		IsRead:          false,
		CreateTime:      createTime,
		AttachmentURLs:  msg.AttachmentURLs,
	}); err != nil {
		return e.StacktraceError(err)
	}

	login, err := uc.store.GetTeacherLoginByChatId(msg.ChatID)
	if err != nil {
		return e.StacktraceError(err)
	}

	uc.hub.AddMsgForClient(&m.MessageWebsocket{
		Text:           msg.Text,
		ChatID:         msg.ChatID,
		Channel:        "chat",
		AttachmentURLs: msg.AttachmentURLs,
		CreateTime:     createTime,
		IsSavedToDB:    true,
		SocialType:     social,
		TeacherLogin:   login,
	})
	return nil
}

func (uc *ChatUsecase) LoadFile(mimeType string, fileURL string, dest string) (string, error) {
	fileExt := ""
	switch mimeType {
	case "image/jpeg":
		fileExt = ".jpg"
	case "image/png":
		fileExt = ".png"
	case "image/svg+xml":
		fileExt = ".svg"
	case "application/pdf":
		fileExt = ".pdf"
	default:
		return "", e.StacktraceError(errors.New("error: " + mimeType + " is not allowed file extension"))
	}

	homeworkNum := uuid.New().String()
	fileName := uc.filestoragePath + "/" + dest + "/" + homeworkNum + fileExt

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return "", e.StacktraceError(err)
	}
	defer f.Close()

	resp, err := http.Get(fileURL)
	if err != nil {
		return "", e.StacktraceError(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(f, resp.Body)
	if err != nil {
		return "", e.StacktraceError(err)
	}

	return uc.urlDomain + fileName, nil
}

func (uc *ChatUsecase) SendBroadcastMsg(bcMsg *m.BroadcastMessage) error {
	ids, err := uc.store.GetChatsByClassID(bcMsg.ChatID)
	if err != nil {
		return e.StacktraceError(err)
	}

	for _, id := range ids {
		socialType, err := uc.store.GetTypeByChatID(id)
		if err != nil {
			return e.StacktraceError(err)
		}
		msg := m.MessageWebsocket{
			ChatID:         id,
			Text:           bcMsg.Title + "\n" + bcMsg.Description,
			AttachmentURLs: bcMsg.AttachmentURLs,
			IsSavedToDB:    false,
		}
		switch socialType {
		case "tg":
			uc.hub.AddMsgForTG(&msg)
		case "vk":
			uc.hub.AddMsgForVK(&msg)
		default:
			return e.StacktraceError(errors.New("unknow social type: " + socialType))
		}
	}

	return nil
}

func (uc *ChatUsecase) ValidateToken(token string) (int, error) {
	id, err := uc.store.ValidateToken(token)
	if err != nil {
		return -1, e.StacktraceError(err)
	}

	return id, nil
}

func (uc *ChatUsecase) CreateStudent(student *m.Student) (int, error) {
	id, err := uc.store.CreateStudent(student)
	if err != nil {
		return -1, e.StacktraceError(err)
	}

	return id, nil
}

func (uc *ChatUsecase) CreateChat(chat *m.Chat) (int, error) {
	teacherID, chatID, err := uc.store.CreateChat(chat)
	if err != nil {
		return -1, e.StacktraceError(err)
	}

	tLogin, err := uc.store.GetTeacherLoginById(teacherID)
	if err != nil {
		return -1, e.StacktraceError(err)
	}
	msg := m.MessageWebsocket{
		Text:         "",
		ChatID:       chatID,
		Channel:      "newchat",
		IsSavedToDB:  false,
		TeacherLogin: tLogin,
		CreateTime:   time.Now(),
	}

	uc.hub.AddMsgForClient(&msg)

	return chatID, nil
}

func (uc *ChatUsecase) GetHomeworks(classID int) ([]m.Homework, error) {
	hws, err := uc.store.GetHomeworksByChatID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return hws, nil
}

func (uc *ChatUsecase) SendSolution(sol *m.Solution) error {
	if err := uc.store.CreateSolution(sol); err != nil {
		return e.StacktraceError(err)
	}

	tLogin, err := uc.store.GetTeacherLoginByHomeworkId(sol.HomeworkID)
	if err != nil {
		return e.StacktraceError(err)
	}

	msg := m.MessageWebsocket{
		Text:         "",
		ChatID:       -1,
		Channel:      "newsolution",
		IsSavedToDB:  false,
		TeacherLogin: tLogin,
		CreateTime:   time.Now(),
	}
	uc.hub.AddMsgForClient(&msg)
	return nil
}

func (uc *ChatUsecase) SendNotification(msg *m.Message) error {
	socialType, err := uc.store.GetTypeByChatID(msg.ChatID)
	if err != nil {
		return e.StacktraceError(err)
	}
	wsMsg := m.MessageWebsocket{
		ChatID:         msg.ChatID,
		Text:           msg.Text,
		AttachmentURLs: msg.AttachmentURLs,
		IsSavedToDB:    false,
	}
	switch socialType {
	case "tg":
		uc.hub.AddMsgForTG(&wsMsg)
	case "vk":
		uc.hub.AddMsgForVK(&wsMsg)
	default:
		return e.StacktraceError(errors.New("unknow social type: " + socialType))
	}
	return nil
}

func (uc *ChatUsecase) GetEvents(classID int) ([]m.CalendarEvent, error) {
	teacherID, err := uc.store.GetTeacherIDByClassID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	events, err := uc.calendar.GetCalendarEvents(teacherID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return events, nil
}
