package models

import (
	"time"

	"example.com/anon-project/db"
	"example.com/anon-project/utils"
	"github.com/guregu/null"
)

type ResetPassword struct {
	ID        int64       `json:"id"`
	UserID    null.Int    `json:"user_id"`
	Token     null.String `json:"token"`
	ExpiresAt NullTime    `json:"expires_at"`
	CreatedAt NullTime    `json:"created_at"`
}

func GetResetPasswordByID(id int64) (ResetPassword, error) {
	query := `SELECT * FROM password_resets WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var resetPassword ResetPassword
	err := row.Scan(&resetPassword.ID, &resetPassword.UserID, &resetPassword.Token, &resetPassword.ExpiresAt, &resetPassword.CreatedAt)

	if err != nil {
		return ResetPassword{}, err
	}

	return resetPassword, nil
}

func (pr *ResetPassword) Save() error {
	query := `
	INSERT INTO password_resets (user_id, token, expires_at, created_at)
	VALUES (?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	token := utils.GenerateResetPasswordToken(pr.UserID.ValueOrZero())

	pr.Token.SetValid(token)
	pr.ExpiresAt.SetValue(time.Now().Add(30 * time.Minute))
	pr.CreatedAt.SetValue(time.Now())

	result, err := stmt.Exec(pr.UserID, pr.Token, pr.ExpiresAt, pr.CreatedAt)

	if err != nil {
		return err
	}

	pr.ID, err = result.LastInsertId()

	return err
}
