package models

import (
	"fmt"
	"time"

	"example.com/restapi/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userId"`
}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime, user_id) 
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %v", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return fmt.Errorf("error executing query: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return fmt.Errorf("error retrieving last insert ID: %v", err)
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT id, name, description, location, dateTime, user_id FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error fetching events: %v", err)
	}
	defer rows.Close()

	var eventList []Event
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			return nil, fmt.Errorf("error scanning row: %v", err)
		}
		eventList = append(eventList, event)
	}

	return eventList, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT id, name, description, location, dateTime, user_id FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		return nil, fmt.Errorf("event not found: %v", err)
	}

	return &event, nil
}

func UpdateEvent(id int64, updatedEvent Event) error {
	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedEvent.Name, updatedEvent.Description, updatedEvent.Location, updatedEvent.DateTime, updatedEvent.UserID, id)
	if err != nil {
		return fmt.Errorf("error executing update query: %v", err)
	}

	return nil
}

func DeleteEvent(id int64) error {
	query := "DELETE FROM events WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error preparing delete query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return fmt.Errorf("error executing delete query: %v", err)
	}

	return nil
}
