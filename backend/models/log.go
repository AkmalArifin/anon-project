package models

import (
	"time"

	"example.com/anon-project/db"
	"github.com/guregu/null"
)

type Log struct {
	ID        int64       `json:"id"`
	UserID    int64       `json:"user_id"`
	Message   null.String `json:"message"`
	IsStarred null.Int    `json:"is_starred"`
	CreatedAt NullTime    `json:"created_at"`
}

func GetAllLog() ([]Log, error) {
	var logs []Log
	query := "SELECT * FROM log"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var log Log
		err = rows.Scan(&log.ID, &log.UserID, &log.Message, &log.IsStarred, &log.CreatedAt)

		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}

func GetLogByID(logID int64) (Log, error) {
	query := "SELECT * FROM log WHERE id = ?"
	row := db.DB.QueryRow(query, logID)

	var log Log
	err := row.Scan(&log.ID, &log.UserID, &log.Message, &log.IsStarred, &log.CreatedAt)

	if err != nil {
		return Log{}, err
	}

	return log, nil
}

func GetLogByUserID(userID int64) ([]Log, error) {
	var logs []Log
	query := "SELECT * FROM log WHERE user_id = ?"
	rows, err := db.DB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var log Log
		err = rows.Scan(&log.ID, &log.UserID, &log.Message, &log.IsStarred, &log.CreatedAt)

		if err != nil {
			return nil, err
		}

		logs = append(logs, log)
	}

	return logs, nil
}

func (l *Log) Save() error {
	query := `
	INSERT INTO log(user_id, message, is_starred, created_at)
	VALUES (?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	l.CreatedAt.SetValue(time.Now())
	l.IsStarred.SetValid(0)

	results, err := stmt.Exec(l.UserID, l.Message, l.IsStarred, l.CreatedAt)

	if err != nil {
		return err
	}

	l.ID, err = results.LastInsertId()

	return err
}

func (l *Log) Update() error {
	query := `
	UPDATE log
	SET is_starred = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(l.IsStarred, l.ID)

	return err
}

func (l *Log) Delete() error {
	query := `
	DELETE FROM log WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(l.ID)

	return err
}
