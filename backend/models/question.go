package models

import "github.com/guregu/null"

type Question struct {
	Username null.String `json:"username"`
	Question null.String `json:"question"`
}
