package grpc

import (
	"context"
	d "main/delivery"
	proto "main/delivery/grpc/proto"
	e "main/domain/errors"
	m "main/domain/model"
	"time"
)

type CalendarService struct {
	client proto.CalendarControllerClient
}

func NewCalendarService(client proto.CalendarControllerClient) d.CalendarInterface {
	return &CalendarService{
		client: client,
	}
}

func (c *CalendarService) GetCalendarEvents(teacherID int) ([]m.CalendarEvent, error) {
	protoEvents, err := c.client.GetEventsCalendar(
		context.Background(),
		&proto.GetEventsRequestCalendar{TeacherID: int32(teacherID)},
	)
	if err != nil {
		return nil, e.StacktraceError(err)
	}

	events := []m.CalendarEvent{}
	for _, protoEvent := range protoEvents.Events {
		startDate, err := time.Parse(time.RFC3339, protoEvent.StartDate)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		endDate, err := time.Parse(time.RFC3339, protoEvent.EndDate)
		if err != nil {
			return nil, e.StacktraceError(err)
		}

		events = append(events, m.CalendarEvent{
			Title:       protoEvent.Title,
			Description: protoEvent.Description,
			StartDate:   startDate,
			EndDate:     endDate,
			ClassID:     int(protoEvent.ClassID),
			ID:          protoEvent.Id,
		})
	}

	return events, nil
}
