package model

import "time"

type Solution struct {
	HomeworkID     int
	StudentID      int
	Text           string
	AttachmentURLs []string
}

type CreateSolution struct {
	HomeworkID     int
	StudentID      int
	Text           string
	CreateDate     time.Time
	AttachmentURLs []string
}
