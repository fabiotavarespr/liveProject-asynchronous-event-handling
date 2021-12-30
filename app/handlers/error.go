package handlers

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/events"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/topics"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/pkg/logger/attributes"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/platform/producers"
	"github.com/google/uuid"
	"time"
)

// HandleError will publish an error event to Kafka
func HandleError(event events.Event) {
	var err error

	e := translateToErrorEvent(event)
	if err = producers.ProducerEvent(e, topics.TopicError); err != nil {
		logger.Error("an issue ocurred publishing an error event to Kafka", attributes.New().WithError(err).WithField("topic", topics.TopicError))
	}
}

func translateToErrorEvent(event events.Event) events.Event {
	return events.Error{
		EventBase: events.BaseEvent{
			EventID:        uuid.New(),
			EventTimestamp: time.Now(),
		},
		EventBody: event,
	}
}
