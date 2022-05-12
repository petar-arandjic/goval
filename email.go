package validator

import (
	"net/mail"
)

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)

	if err != nil {
		return false
	}

	return true
}
