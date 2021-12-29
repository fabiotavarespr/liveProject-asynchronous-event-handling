package queries

import (
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/events"
	"github.com/jmoiron/sqlx"
)

// EventsQueries struct for queries from Events model.
type EventsQueries struct {
	*sqlx.DB
}

// EventExists method check if event has processed.
func (q *EventsQueries) EventExists(event events.Event) (bool, error) {
	// Define result string
	var id string

	// Define query string.
	query := `SELECT id FROM events.processed_events WHERE id=$1 AND event_name=$2`

	// Send query to database.
	err := q.QueryRow(query, event.ID(), event.Name()).Scan(&id)
	if err != nil {
		// Return empty object and error.
		return false, err
	}

	// Return query result.
	return len(id) > 0, nil
}

// CreateProcessedEvent method for creating processed_events by given event object.
func (q *EventsQueries) CreateProcessedEvent(event events.Event) error {
	// Define query string.
	query := `INSERT INTO events.processed_events (id, event_name, processed_timestamp) VALUES ($1, $2, $3)`

	// Send query to database.
	_, err := q.Exec(query, event.ID(), event.Name(), event.Timestamp())
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}
