package events

import (
	"time"

	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/topics"
	"github.com/google/uuid"
)

// Notification represents a notification that is being requested to be sent
type Notification struct {
	EventBase BaseEvent
	EventBody models.Notification
}

// ID returns the unique identifier of the event
func (n Notification) ID() uuid.UUID {
	return n.EventBase.EventID
}

// Name returns the name of the event
func (n Notification) Name() string {
	return topics.TopicNotification
}

// Timestamp returns the unique timestamp of the event
func (n Notification) Timestamp() time.Time {
	return n.EventBase.EventTimestamp
}

// Body returns the body content of the event
func (n Notification) Body() interface{} {
	return n.EventBody
}
