package models

import "github.com/guregu/null"

type ResetPasswordRequest struct {
	Email null.String `json:"email"`
}
