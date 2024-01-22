package usecase

import (
	m "main/domain/model"
)

type UsecaseInterface interface {
	// MESSAGE
	SaveMsg(msg *m.MessageWebsocket) error
	GetMsgForTG() (*m.MessageWebsocket, error)
	GetMsgForVK() (*m.MessageWebsocket, error)
	SendMsgToClient(msg *m.Message, social string) error
	SaveFile(file *m.Attach) (string, error)
	SendBroadcastMsg(bcMsg *m.BroadcastMessage) error
	SendNotification(msg *m.Message) error
	SendSolution(sol *m.Solution) error
	// CLASS
	ValidateToken(token string) (int, error)
	CreateStudent(student *m.Student) (int, error)
	CreateChat(chat *m.Chat) (int, error)
	GetHomeworks(classID int) ([]m.Homework, error)
	GetEvents(classID int) ([]m.CalendarEvent, error)
}
