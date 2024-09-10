package models

import (
	"time"

	"example.com/anon-project/db"
	"example.com/anon-project/utils"
)

type User struct {
	ID        int64    `json:"id"`
	Username  string   `json:"username"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	CreatedAt NullTime `json:"created_at"`
	UpdatedAt NullTime `json:"updated_at"`
	DeletedAt NullTime `json:"deleted_at"`
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

	u.CreatedAt.Time = time.Now()
	u.UpdatedAt.Time = time.Now()

	u.Password, err = utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

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

	u.UpdatedAt.Time = time.Now()

	_, err = stmt.Exec(u.Username, u.Email, u.Password, u.UpdatedAt, u.DeletedAt, u.ID)

	return err
}

func (u *User) Delete() error {
	query := `
	UPDATE users
	SET deleted_at = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	u.DeletedAt.Time = time.Now()

	_, err = stmt.Exec(u.DeletedAt, u.ID)

	return err
}
