package gc

import (
	"google.golang.org/api/calendar/v3"
	"time"
)

func (gc *CalendarClient) GetEvents() (*calendar.Events, error) {
	now := time.Now()
	return gc.svc.Events.
		List("primary").
		ShowDeleted(false).
		SingleEvents(true).
		TimeMin(time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).Format(time.RFC3339)).
		TimeMax(time.Date(now.Year(), now.Month(), now.Day()+7, 0, 0, 0, 0, time.Local).Add(-1 * time.Second).Format(time.RFC3339)).
		OrderBy("startTime").
		Do()
}
