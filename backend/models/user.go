package models

import (
	"fmt"
	"time"

	"example.com/anon-project/db"
	"example.com/anon-project/utils"
	"github.com/guregu/null"
)

type User struct {
	ID        int64       `json:"id"`
	Username  null.String `json:"username"`
	Email     null.String `json:"email"`
	Password  null.String `json:"password"`
	CreatedAt NullTime    `json:"created_at"`
	UpdatedAt NullTime    `json:"updated_at"`
	DeletedAt NullTime    `json:"deleted_at"`
}

func GetAllUsers() ([]User, error) {
	var users []User
	query := `SELECT * FROM users`

	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(userID int64) (User, error) {
	query := `SELECT * FROM users WHERE id = ?`
	row := db.DB.QueryRow(query, userID)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(username, email, password, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	u.Password.String, err = utils.HashPassword(u.Password.String)

	if err != nil {
		return err
	}

	fmt.Println(time.Now())

	u.CreatedAt.SetValue(time.Now())
	u.UpdatedAt.SetValue(time.Now())

	results, err := stmt.Exec(u.Username, u.Email, u.Password, u.CreatedAt, u.UpdatedAt)

	if err != nil {
		return err
	}

	u.ID, err = results.LastInsertId()
	return err
}

func (u *User) Update() error {
	query := `
	UPDATE users
	SET username = ?, email = ?, password = ?, updated_at = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	u.UpdatedAt.SetValue(time.Now())

	_, err = stmt.Exec(u.Username, u.Email, u.Password, u.UpdatedAt, u.ID)

	return err
}

func (u *User) Delete() error {
	query := `
	UPDATE users
	SET updated_at = ?, deleted_at = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	u.UpdatedAt.SetValue(time.Now())
	u.DeletedAt.SetValue(time.Now())

	_, err = stmt.Exec(u.UpdatedAt, u.DeletedAt, u.ID)

	return err
}
