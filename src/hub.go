package chat

type MessageWebsocket struct {
	Text    string `json:"text"`
	ChatID  int32  `json:"chatid"`
	Channel string `json:"channel,omitempty"`
}

type Hub struct {
	clients         map[*Client]bool
	Broadcast       chan *MessageWebsocket
	MessagesToTGBot chan *MessageWebsocket
	MessagesToVKBot chan *MessageWebsocket
	register        chan *Client
	unregister      chan *Client
	chats           map[int32]*Client // соединение по id чата
	clientChats     map[*Client][]int32
}

func NewHub() *Hub {
	return &Hub{
		Broadcast:       make(chan *MessageWebsocket, 100),
		MessagesToTGBot: make(chan *MessageWebsocket, 100),
		MessagesToVKBot: make(chan *MessageWebsocket, 100),
		register:        make(chan *Client),
		unregister:      make(chan *Client),
		clientChats:     make(map[*Client][]int32),
		clients:         make(map[*Client]bool),
		chats:           make(map[int32]*Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				for _, cl := range h.clientChats[client] {
					delete(h.chats, cl)
				}
				delete(h.clientChats, client)
				delete(h.clients, client)
				close(client.send)
			}
		case mes := <-h.Broadcast:

			conn := h.chats[mes.ChatID]
			if conn == nil {
				break
			}
			select {
			case conn.send <- mes:
			default:
				close(conn.send)
				delete(h.chats, mes.ChatID)
			}

		}
	}
}
