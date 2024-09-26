package models

import "github.com/guregu/null"

type Message struct {
	Username null.String `json:"username"`
	Message  null.String `json:"message"`
}
