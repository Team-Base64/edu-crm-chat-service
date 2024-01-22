package chat

import (
	"time"

	e "main/domain/errors"
	m "main/domain/model"
)

func (uc *ChatUsecase) ValidateToken(token string) (int, error) {
	id, err := uc.dataStore.ValidateToken(token)
	if err != nil {
		return -1, e.StacktraceError(err)
	}

	return id, nil
}

func (uc *ChatUsecase) CreateStudent(student *m.Student) (int, error) {
	id, err := uc.dataStore.CreateStudent(student)
	if err != nil {
		return -1, e.StacktraceError(err)
	}

	return id, nil
}

func (uc *ChatUsecase) CreateChat(chat *m.Chat) (int, error) {
	teacherID, chatID, err := uc.dataStore.CreateChat(chat)
	if err != nil {
		return -1, e.StacktraceError(err)
	}

	tLogin, err := uc.dataStore.GetTeacherLoginById(teacherID)
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
	hws, err := uc.dataStore.GetHomeworksByChatID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return hws, nil
}

func (uc *ChatUsecase) GetEvents(classID int) ([]m.CalendarEvent, error) {
	teacherID, err := uc.dataStore.GetTeacherIDByClassID(classID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	events, err := uc.calendar.GetCalendarEvents(teacherID)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	return events, nil
}
