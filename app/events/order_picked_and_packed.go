package events

import (
	"time"

	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/topics"
	"github.com/google/uuid"
)

// OrderPickedAndPacked represents an event when customers order has been picked and packed by the warehouse
type OrderPickedAndPacked struct {
	EventBase BaseEvent
	EventBody models.Order
}

// ID returns the unique identifier of the event
func (opp OrderPickedAndPacked) ID() uuid.UUID {
	return opp.EventBase.EventID
}

// Name returns the name of the event
func (opp OrderPickedAndPacked) Name() string {
	return topics.TopicOrderPickedAndPacked
}

// Timestamp returns the unique timestamp of the event
func (opp OrderPickedAndPacked) Timestamp() time.Time {
	return opp.EventBase.EventTimestamp
}

// Body returns the body content of the event
func (opp OrderPickedAndPacked) Body() interface{} {
	return opp.EventBody
}
