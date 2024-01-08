package ws

import (
	"log"
	d "main/delivery"
	e "main/domain/errors"
	m "main/domain/model"
	"net/http"

	"github.com/gorilla/websocket"
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

type Hub struct {
	clients        map[*ConnectionWS]bool
	msgToClient    chan *m.MessageWebsocket
	msgToTG        chan *m.MessageWebsocket
	msgToVK        chan *m.MessageWebsocket
	register       chan *ConnectionWS
	unregister     chan *ConnectionWS
	teacherClients map[string]map[*ConnectionWS]struct{}
	clientTeacher  map[*ConnectionWS]string
	upgrader       websocket.Upgrader
}

func NewHub() d.HubExtendedInterface {
	return &Hub{
		msgToClient:    make(chan *m.MessageWebsocket, 100),
		msgToTG:        make(chan *m.MessageWebsocket, 100),
		msgToVK:        make(chan *m.MessageWebsocket, 100),
		register:       make(chan *ConnectionWS),
		unregister:     make(chan *ConnectionWS),
		clientTeacher:  make(map[*ConnectionWS]string),
		clients:        make(map[*ConnectionWS]bool),
		teacherClients: make(map[string]map[*ConnectionWS]struct{}),
		upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

// AddConnection godoc
// @Summary Connect websocket
// @Description Connect websocket
// @ID addConnection
// @Accept  json
// @Produce  json
// @Param teacherLogin query string true "teacherLogin"
// @Router /ws [get]
func (hub *Hub) AddConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := hub.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return
	}

	usLogin := r.URL.Query().Get("teacherLogin")
	//usLogin := "test1"
	client := &ConnectionWS{hub: hub, conn: conn, send: make(chan *m.MessageWebsocket)}
	client.hub.register <- client

	if _, ok := hub.teacherClients[usLogin]; !ok {
		hub.teacherClients[usLogin] = make(map[*ConnectionWS]struct{})
	}

	hub.teacherClients[usLogin][client] = struct{}{}
	hub.clientTeacher[client] = usLogin

	log.Println("opened websocket")
	go client.writePump()
	go client.readPump()
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {

				t := h.clientTeacher[client]
				delete(h.teacherClients, t)
				delete(h.clientTeacher, client)

				delete(h.clients, client)

				close(client.send)
			}
		case mes := <-h.msgToClient:
			for cl := range h.teacherClients[mes.TeacherLogin] {
				conn := cl
				if conn == nil {
					break
				}
				select {
				case conn.send <- mes:
				default:
					close(conn.send)
					delete(h.teacherClients[mes.TeacherLogin], cl)
					//delete(h.chats, mes.ChatID)
				}

			}

		}
	}
}

func (h *Hub) AddMsgForTG(msg *m.MessageWebsocket) {
	h.msgToTG <- msg
}

func (h *Hub) AddMsgForVK(msg *m.MessageWebsocket) {
	h.msgToVK <- msg
}

func (h *Hub) AddMsgForClient(msg *m.MessageWebsocket) {
	h.msgToClient <- msg
}

func (h *Hub) GetMsgForTG() *m.MessageWebsocket {
	return <-h.msgToTG
}

func (h *Hub) GetMsgForVK() *m.MessageWebsocket {
	return <-h.msgToVK
}

func (h *Hub) GetMsgForClient() *m.MessageWebsocket {
	return <-h.msgToClient
}
