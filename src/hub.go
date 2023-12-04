package chat

import (
	m "main/domain/model"
)

type Hub struct {
	clients         map[*Client]bool
	SendToFront     chan *m.MessageWebsocket
	MessagesToTGBot chan *m.MessageWebsocket
	MessagesToVKBot chan *m.MessageWebsocket
	register        chan *Client
	unregister      chan *Client
	//chats           map[int32]*Client // соединение по id чата
	//teachers map[string]*Client
	teacherClients map[string]map[*Client]struct{}
	//clientChats     map[*Client][]int32
	clientTeacher map[*Client]string
}

func NewHub() *Hub {
	//map1 := make(map[string]map[*Client]struct{}, 0)
	//map2 := make(map[*Client]struct{})
	return &Hub{
		SendToFront:     make(chan *m.MessageWebsocket, 100),
		MessagesToTGBot: make(chan *m.MessageWebsocket, 100),
		MessagesToVKBot: make(chan *m.MessageWebsocket, 100),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		//clientChats:     make(map[*Client][]int32),
		clientTeacher: make(map[*Client]string),
		clients:       make(map[*Client]bool),
		//chats:           make(map[int32]*Client),
		//teachers:        make(map[string]*Client),
		teacherClients: make(map[string]map[*Client]struct{}),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				// for _, cl := range h.clientChats[client] {
				// 	delete(h.chats, cl)
				// }
				t := h.clientTeacher[client]
				delete(h.teacherClients, t)
				delete(h.clientTeacher, client)

				//delete(h.clientChats, client)
				delete(h.clients, client)

				close(client.send)
			}
		case mes := <-h.SendToFront:
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
