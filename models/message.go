package models

import (
	"database/sql"
	"errors"
	"time"

	"example.com/gg-messages/db"
)

type Message struct {
	Id        int64
	Message   string `binding:"required"`
	CreatedAt time.Time
}

func (m *Message) Save() error {
	query := `INSERT INTO messages(message, created_at) VALUES(?,?)`
	statement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer statement.Close()

	result, err := statement.Exec(m.Message, time.Now())
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	m.Id = id
	return err
}

func GetAllMessages() ([]string, error) {
	query := "SELECT message FROM messages"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []string

	for rows.Next() {
		var message string
		err := rows.Scan(&message)

		if err != nil {
			return nil, err
		}

		messages = append(messages, message)
	}

	return messages, nil
}

var ErrMessageNotFound = errors.New("message not found")

func GetMessage(id int64) (string, error) {
	query := "SELECT message from messages where id = ?"
	row := db.DB.QueryRow(query, id)

	var messageText string
	err := row.Scan(&messageText)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", ErrMessageNotFound
		} else {
			return "", err
		}
	}

	return messageText, nil
}
