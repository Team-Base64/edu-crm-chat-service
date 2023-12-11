package model

import "time"

type Error struct {
	Error interface{} `json:"error,omitempty"`
}

type Response struct {
	Body interface{} `json:"body,omitempty"`
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
	ChatID         int32     `json:"chatID"`
	Channel        string    `json:"channel,omitempty"`
	AttachmentURLs []string  `json:"attaches,omitempty"`
	CreateTime     time.Time `json:"date,omitempty"`
	IsSavedToDB    bool      `json:"-"`
	SocialType     string    `json:"socialType"`
	TeacherLogin   string    `json:"-"`
}

type CreateBroadcastMessage struct {
	ChatID int32  `json:"chatid"`
	Type   string `json:"type"`
}

type CalendarParams struct {
	ID         int    `json:"id"`
	IDInGoogle string `json:"googleid"`
}

type CalendarEvent struct {
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	StartDate   time.Time `json:"startDate"`
	EndDate     time.Time `json:"endDate"`
	ClassID     int       `json:"classid"`
	ID          string    `json:"id,omitempty"`
}
