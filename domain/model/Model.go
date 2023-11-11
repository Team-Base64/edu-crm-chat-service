package model

import "time"

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Response struct {
	Body interface{} `json:"body,omitempty"`
}

type UploadAttachResponse struct {
	File string `json:"file,omitempty"`
}

type CreateMessage struct {
	ChatID          int      `json:"chatid,omitempty"`
	Text            string   `json:"text"`
	IsAuthorTeacher bool     `json:"isAuthorTeacher"`
	IsRead          bool     `json:"isread"`
	AttachmentURLs  []string `json:"attaches,omitempty"`
}

type MessageWebsocket struct {
	Text           string    `json:"text"`
	ChatID         int32     `json:"chatid"`
	Channel        string    `json:"channel,omitempty"`
	AttachmentURLs []string  `json:"attaches,omitempty"`
	CreateTime     time.Time `json:"date,omitempty"`
}

type CreateBroadcastMessage struct {
	ChatID int32  `json:"chatid"`
	Type   string `json:"type"`
}

type OAUTH2Token struct {
	Token string `json:"token"`
}

type CreateCalendarResponse struct {
	ID         int    `json:"id"`
	IDInGoogle string `json:"googleid"`
}

type CreateCalendarEvent struct {
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
	StartDate   string `json:"startDate"`
	EndDate     string `json:"endDate"`
}
