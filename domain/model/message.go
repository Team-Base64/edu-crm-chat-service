package model

import "time"

type Message struct {
	ChatID         int
	Text           string
	AttachmentURLs []string
}

type CreateMessage struct {
	ChatID          int
	Text            string
	IsAuthorTeacher bool
	IsRead          bool
	CreateTime      time.Time
	AttachmentURLs  []string
}

type MessageWebsocket struct {
	Text           string    `json:"text"`
	ChatID         int       `json:"chatID"`
	Channel        string    `json:"channel,omitempty"`
	AttachmentURLs []string  `json:"attaches,omitempty"`
	CreateTime     time.Time `json:"date,omitempty"`
	IsSavedToDB    bool      `json:"-"`
	SocialType     string    `json:"socialType"`
	TeacherLogin   string    `json:"-"`
}

type BroadcastMessage struct {
	ChatID         int
	Title          string
	Description    string
	AttachmentURLs []string
}
