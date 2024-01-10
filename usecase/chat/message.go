package chat

import (
	"errors"
	"time"

	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *ChatUsecase) SaveMsg(msg *m.MessageWebsocket) error {
	if err := uc.dataStore.AddMessage(&m.CreateMessage{
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
	if err := uc.dataStore.AddMessage(&m.CreateMessage{
		Text:            msg.Text,
		ChatID:          msg.ChatID,
		IsAuthorTeacher: false,
		IsRead:          false,
		CreateTime:      createTime,
		AttachmentURLs:  msg.AttachmentURLs,
	}); err != nil {
		return e.StacktraceError(err)
	}

	login, err := uc.dataStore.GetTeacherLoginByChatId(msg.ChatID)
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

func (uc *ChatUsecase) SaveFile(file *m.Attach) (string, error) {
	fileName, err := uc.fileStore.UploadFile(file)
	if err != nil {
		return "", e.StacktraceError(err)
	}

	return uc.urlDomain + fileName, nil
}

func (uc *ChatUsecase) SendBroadcastMsg(bcMsg *m.BroadcastMessage) error {
	ids, err := uc.dataStore.GetChatsByClassID(bcMsg.ChatID)
	if err != nil {
		return e.StacktraceError(err)
	}

	for _, id := range ids {
		socialType, err := uc.dataStore.GetSocialTypeByChatID(id)
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

func (uc *ChatUsecase) SendNotification(msg *m.Message) error {
	socialType, err := uc.dataStore.GetSocialTypeByChatID(msg.ChatID)
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

func (uc *ChatUsecase) SendSolution(sol *m.Solution) error {
	newSol := m.CreateSolution{
		HomeworkID:     sol.HomeworkID,
		StudentID:      sol.StudentID,
		Text:           sol.Text,
		CreateDate:     time.Now(),
		AttachmentURLs: sol.AttachmentURLs,
	}
	if err := uc.dataStore.CreateSolution(&newSol); err != nil {
		return e.StacktraceError(err)
	}

	tLogin, err := uc.dataStore.GetTeacherLoginByHomeworkId(sol.HomeworkID)
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
