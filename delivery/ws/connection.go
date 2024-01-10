package ws

import (
	"encoding/json"
	"log"
	"time"

	e "main/domain/errors"
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

type ConnectionWS struct {
	hub  *Hub
	conn *websocket.Conn
	send chan *m.MessageWebsocket
}

func (c *ConnectionWS) readPump() {
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
		log.Println("Received mes from websocket: ", "type: ", req.SocialType, " text:", req.Text, " chatid:", req.ChatID, " attaches: ", req.AttachmentURLs)
		req.IsSavedToDB = true
		switch req.SocialType {
		case "tg":
			c.hub.msgToTG <- req
		case "vk":
			c.hub.msgToVK <- req
		}

	}
}

func (c *ConnectionWS) writePump() {
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
