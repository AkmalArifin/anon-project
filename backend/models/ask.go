package models

import (
	"time"

	"example.com/anon-project/db"
	"github.com/guregu/null"
)

type Ask struct {
	ID        int64       `json:"id"`
	UserID    int64       `json:"user_id"`
	Question  null.String `json:"question"`
	IsRead    null.Bool   `json:"is_read"`
	CreatedAt NullTime    `json:"created_at"`
}

func GetAllAskLog() ([]Ask, error) {
	var asks []Ask
	query := "SELECT * FROM ask_log"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ask Ask
		err = rows.Scan(&ask.ID, &ask.UserID, &ask.Question, &ask.CreatedAt)

		if err != nil {
			return nil, err
		}

		asks = append(asks, ask)
	}

	return asks, nil
}

func GetAskByID(askID int64) (Ask, error) {
	query := "SELECT * FROM ask_log WHERE id = ?"
	row := db.DB.QueryRow(query, askID)

	var ask Ask
	err := row.Scan(&ask.ID, &ask.UserID, &ask.Question, &ask.CreatedAt)

	if err != nil {
		return Ask{}, err
	}

	return ask, nil
}

func GetAskByUserID(userID int64) ([]Ask, error) {
	var asks []Ask
	query := "SELECT * FROM ask_log WHERE user_id = ?"
	rows, err := db.DB.Query(query, userID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ask Ask
		err = rows.Scan(&ask.ID, &ask.UserID, &ask.Question, &ask.IsRead, &ask.CreatedAt)

		if err != nil {
			return nil, err
		}

		asks = append(asks, ask)
	}

	return asks, nil
}

func (a *Ask) Save() error {
	query := `
	INSERT INTO ask_log(user_id, question, is_read, created_at)
	VALUES (?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	a.CreatedAt.SetValue(time.Now())
	a.IsRead.SetValid(false)

	results, err := stmt.Exec(a.UserID, a.Question, a.IsRead, a.CreatedAt)

	if err != nil {
		return err
	}

	a.ID, err = results.LastInsertId()

	return err
}
