package models

import "example.com/anon-project/db"

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
