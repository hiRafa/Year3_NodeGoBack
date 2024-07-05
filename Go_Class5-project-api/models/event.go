package models

import (
	"time"

	"api.com/sqldb"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int64     `json:"userId"`
}

// var events []Event = []Event{}
var events = []Event{}

func (e *Event) Save() error {
	query := `
		INSERT INTO events(name, description, location, dateTime, user_id)
		VALUES (?, ?, ?, ?, ?)
	`
	stmt, err := sqldb.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	e.ID = id // Assign the last insert ID to e.ID

	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, error := sqldb.DB.Query(query)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		rows.Scan(
			&event.ID,
			&event.Name,
			&event.Description,
			&event.Location,
			&event.DateTime,
			&event.UserID,
		)

		if error != nil {
			return nil, error
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := sqldb.DB.QueryRow(query, id)

	var event Event
	error := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if error != nil {
		return nil, error
	}

	return &event, nil
}

func (event *Event) Update() error {
	query := `
		UPDATE events 
		SET name = ?, description = ?, location = ?, dateTime = ?
		WHERE id = ?
	`
	stmt, error := sqldb.DB.Prepare(query)

	if error != nil {
		return error
	}

	defer stmt.Close()

	_, error = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return error
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, error := sqldb.DB.Prepare(query)

	if error != nil {
		return error
	}

	defer stmt.Close()

	_, error = stmt.Exec(event.ID)
	return error
}

func (e Event) AddEventUser(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES (?,?)"
	stmt, error := sqldb.DB.Prepare(query)

	if error != nil {
		return error
	}

	defer stmt.Close()

	_, error = stmt.Exec(e.ID, userId)

	return error
}

func (e Event) DeleteEventUser(userId int64) error {
	query := "DELETE FROM registrations WHERE event_id = ? AND user_id = ?"
	stmt, error := sqldb.DB.Prepare(query)

	if error != nil {
		return error
	}

	defer stmt.Close()

	_, error = stmt.Exec(e.ID, userId)

	return error
}
