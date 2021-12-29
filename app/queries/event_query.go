package queries

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/fabiotavarespr/liveProject-asynchronous-event-handling/app/models"
)

// BookQueries struct for queries from Book model.
type EventsQueries struct {
	*sqlx.DB
}

// GetBook method for getting one book by given ID.
func (q *EventsQueries) EventExists(event events.Event) (bool, error) {
	// Define book variable.
	book := models.Book{}

	// Define query string.
	query := `SELECT * FROM books WHERE id = $1`

	// Send query to database.
	err := q.Get(&book, query, id)
	if err != nil {
		// Return empty object and error.
		return book, err
	}

	// Return query result.
	return book, nil
}

// CreateBook method for creating book by given Book object.
func (q *EventsQueries) CreateEventevent events.Event) error {
	// Define query string.
	query := `INSERT INTO books VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	// Send query to database.
	_, err := q.Exec(query, b.ID, b.CreatedAt, b.UpdatedAt, b.UserID, b.Title, b.Author, b.BookStatus, b.BookAttrs)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}