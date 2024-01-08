package model

import "time"

type CalendarEvent struct {
	Title       string
	Description string
	StartDate   time.Time
	EndDate     time.Time
	ClassID     int
	ID          string
}
