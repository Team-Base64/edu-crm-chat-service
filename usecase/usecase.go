package usecase

import (
	m "main/domain/model"
)

type UsecaseInterface interface {
	GetMsgForTG() (*m.MessageWebsocket, error)
	GetMsgForVK() (*m.MessageWebsocket, error)
	SendMsgToClient(msg *m.Message, social string) error
	LoadFile(mimeType string, fileURL string, dest string) (string, error)
	SendBroadcastMsg(bcMsg *m.BroadcastMessage) error
	ValidateToken(token string) (int, error)
	CreateStudent(student *m.Student) (int, error)
	CreateChat(chat *m.Chat) (int, error)
	GetHomeworks(classID int) ([]m.Homework, error)
	SendSolution(sol *m.Solution) error
	SendNotification(msg *m.Message) error
	GetEvents(classID int) ([]m.CalendarEvent, error)
}
