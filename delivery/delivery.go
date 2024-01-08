package delivery

import (
	"net/http"

	m "main/domain/model"
)

type HandlerInterface interface {
	ServeWs(w http.ResponseWriter, r *http.Request)
}

type ClientInterface interface {
}

type HubInterface interface {
	AddMsgForTG(msg *m.MessageWebsocket)
	AddMsgForVK(msg *m.MessageWebsocket)
	AddMsgForClient(msg *m.MessageWebsocket)
	GetMsgForTG() *m.MessageWebsocket
	GetMsgForVK() *m.MessageWebsocket
	GetMsgForClient() *m.MessageWebsocket
}

type HubExtendedInterface interface {
	HubInterface
	AddConnection(w http.ResponseWriter, r *http.Request)
	Run()
}

type CalendarInterface interface {
	GetCalendarEvents(teacherID int) ([]m.CalendarEvent, error)
}
