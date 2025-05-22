package models

import "time"

type Event struct {
	Id int
	Name string
	Description string
	Location string
	DateTime  time.Time
	UserId int
}

var events = []Event{}

func (e Event) Save()  {
	// later: add it to a database
	events = append(events, e)
}