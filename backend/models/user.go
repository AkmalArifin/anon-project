package models

import (
	"errors"
	"time"

	"example.com/anon-project/db"
	"example.com/anon-project/utils"
	"github.com/guregu/null"
)

type User struct {
	ID           int64       `json:"id"`
	Username     null.String `json:"username"`
	Email        null.String `json:"email"`
	Password     null.String `json:"password"`
	PhotoProfile null.String `json:"photo_profile"`
	CreatedAt    NullTime    `json:"created_at"`
	UpdatedAt    NullTime    `json:"updated_at"`
	DeletedAt    NullTime    `json:"deleted_at"`
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
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

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
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	query := `SELECT * FROM users WHERE username = ?`
	row := db.DB.QueryRow(query, username)

	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.PhotoProfile, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (u *User) Save() error {
	query := `
	INSERT INTO users(username, email, password, photo_profile, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?)
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

	u.CreatedAt.SetValue(time.Now())
	u.UpdatedAt.SetValue(time.Now())

	results, err := stmt.Exec(u.Username, u.Email, u.Password, u.PhotoProfile, u.CreatedAt, u.UpdatedAt)

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

	_, err = stmt.Exec(u.Username, u.Email, u.Password, u.PhotoProfile, u.UpdatedAt, u.ID)

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

func (u *User) ValidateCredentials() error {
	query := "SELECT id, username, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var id int64
	var username string
	var retrievedPassword string
	err := row.Scan(&id, &username, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid")
	}

	isPasswordValid := utils.CompareHashPassword(u.Password.String, retrievedPassword)

	if !isPasswordValid {
		return errors.New("credentials invalid")
	}

	u.Username.String = username
	u.ID = id

	return nil
}
