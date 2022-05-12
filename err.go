package validator

import (
	"fmt"
	errs "github.com/petar-arandjic/goerr"
)

var (
	NewPasswordNumberErr = func() *errs.Error {
		return &errs.Error{
			Key:     "VALIDATION_PASSWORD_NUMBER_ERR",
			Message: "password must contain least one number",
			Errors:  nil,
			Status:  nil,
		}
	}
	NewPasswordLowerLetterErr = func() *errs.Error {
		return &errs.Error{
			Key:     "VALIDATION_LOWER_LETTER_ERR",
			Message: "password must contain least one lower letter",
			Errors:  nil,
			Status:  nil,
		}
	}
	NewPasswordUpperLetterErr = func() *errs.Error {
		return &errs.Error{
			Key:     "VALIDATION_UPPER_LETTER_ERR",
			Message: "password must contain least one upper letter",
			Errors:  nil,
			Status:  nil,
		}
	}
	NewPasswordSpecialCharacterErr = func() *errs.Error {
		return &errs.Error{
			Key:     "VALIDATION_SPECIAL_CHARACTER_ERR",
			Message: "password must contain least one special character",
			Errors:  nil,
			Status:  nil,
		}
	}
	NewPasswordCharacterLengthErr = func(len int16) *errs.Error {
		return &errs.Error{
			Key:     "VALIDATION_CHARACTER_LENGTH_ERR",
			Message: fmt.Sprintf("password must be at least %v character long", len),
			Errors:  nil,
			Status:  nil,
		}
	}
)
