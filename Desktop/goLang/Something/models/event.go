package models

import "time"

type Event struct {
	ID          int
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int
}

var events = []Event{}

func (e Event) Save() {
	//store it in db

	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}

func NewEvent(id int, name string, decs string, location string, userid int) Event {
	event := Event{
		ID:          id,
		Name:        name,
		Description: decs,
		Location:    location,
		UserId:      userid,
	}

	events = append(events, event)
	return event
}
