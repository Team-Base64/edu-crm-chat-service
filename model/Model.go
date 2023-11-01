package model

type CreateMessage struct {
	ChatID          int    `json:"chatid,omitempty"`
	Text            string `json:"text"`
	IsAuthorTeacher bool   `json:"isAuthorTeacher"`
	IsRead          bool   `json:"isread"`
}
