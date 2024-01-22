package model

import "time"

type Homework struct {
	HomeworkID   int
	Title        string
	Description  string
	CreateDate   time.Time
	DeadlineDate time.Time
	Tasks        []Task
}

type Task struct {
	Description    string
	AttachmentURLs []string
}
