package chat

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	e "main/domain/errors"
	"main/domain/model"
	m "main/domain/model"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second
	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10
	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan *m.MessageWebsocket
}

func returnErrorJSON(w http.ResponseWriter, err error) {
	errCode, errText := e.CheckError(err)
	w.WriteHeader(errCode)
	json.NewEncoder(w).Encode(&model.Error{Error: errText})
}

func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {

		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println(e.StacktraceError(err))
			}
			break
		}
		var req *m.MessageWebsocket
		err = json.Unmarshal(message, &req)
		if err != nil {
			log.Println(e.StacktraceError(err))
			return
		}
		log.Println("Received mes from websocket: ", "text:", req.Text, "chatid:", req.ChatID, "attaches: ", req.AttachmentURLs)
		c.hub.chats[req.ChatID] = c
		c.hub.clientChats[c] = append(c.hub.clientChats[c], req.ChatID)

		req.IsSavedToDB = true
		switch type1 := req.SocialType; type1 {
		case "tg":
			c.hub.MessagesToTGBot <- req
		case "vk":
			c.hub.MessagesToVKBot <- req
		}

	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				log.Println(e.StacktraceError(err))
				return
			}
			req, err := json.Marshal(message)
			if err != nil {
				log.Println(e.StacktraceError(err))
				return
			}

			w.Write(req)
			c.hub.chats[message.ChatID] = c
			c.hub.clientChats[c] = append(c.hub.clientChats[c], message.ChatID)
			log.Println("send mes to websocket: ", message)

			if err := w.Close(); err != nil {
				log.Println(e.StacktraceError(err))
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func (api *Handler) ServeWs(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(e.StacktraceError(err))
		return
	}

	// session, err := r.Cookie("session_id")
	// if err == http.ErrNoCookie {
	// 	log.Println(e.StacktraceError(err))
	// 	returnErrorJSON(w, e.ErrUnauthorized401)
	// 	return
	// }
	// usLogin, err := api.store.CheckSession(session.Value)
	// if errors.Is(err, sql.ErrNoRows) {
	// 	log.Println(e.StacktraceError(err, errors.New("no sess: ")))
	// 	returnErrorJSON(w, e.ErrUnauthorized401)
	// 	return
	// }
	// if err != nil {
	// 	log.Println(e.StacktraceError(err))
	// 	returnErrorJSON(w, e.ErrServerError500)
	// 	return
	// }

	client := &Client{hub: api.hub, conn: conn, send: make(chan *m.MessageWebsocket)}
	client.hub.register <- client

	//usLogin := r.URL.Query().Get("teacherLogin")
	// curChats, err := api.store.GetAllUserChatIDs(usLogin)
	// if err != nil {
	// 	log.Println(e.StacktraceError(err))
	// 	returnErrorJSON(w, e.ErrServerError500)
	// 	return
	// }
	curChats := []int32{1, 2, 3, 4, 5}
	for _, ch := range curChats {
		api.hub.chats[ch] = client
		api.hub.clientChats[client] = append(api.hub.clientChats[client], ch)
	}

	log.Println("opened websocket")
	go client.writePump()
	go client.readPump()
}
