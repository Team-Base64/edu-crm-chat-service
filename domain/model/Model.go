package model

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Response struct {
	Body interface{} `json:"body,omitempty"`
}

type CreateMessage struct {
	ChatID          int    `json:"chatid,omitempty"`
	Text            string `json:"text"`
	IsAuthorTeacher bool   `json:"isAuthorTeacher"`
	IsRead          bool   `json:"isread"`
}

type MessageWebsocket struct {
	Text    string `json:"text"`
	ChatID  int32  `json:"chatid"`
	Channel string `json:"channel,omitempty"`
}
