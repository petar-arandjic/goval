package validator

import (
	errs "github.com/petar-arandjic/goerr"
	"unicode"
)

type PasswordValidate struct {
	number  bool
	upper   bool
	lower   bool
	special bool
	letters *int16
}

func IsStrongPassword(s string, options PasswordValidate) *errs.Error {
	var password PasswordValidate
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			password.number = true
			*password.letters++
		case unicode.IsUpper(c):
			password.upper = true
			*password.letters++
		case unicode.IsLower(c):
			password.lower = true
			*password.letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			password.special = true
			*password.letters++
		default:
			*password.letters++
		}
	}

	var validationErr []errs.Error

	if options.number && !password.number {
		validationErr = append(validationErr, *NewPasswordNumberErr())
	}

	if options.lower && !password.lower {
		validationErr = append(validationErr, *NewPasswordLowerLetterErr())
	}

	if options.upper && !password.upper {
		validationErr = append(validationErr, *NewPasswordNumberErr())
	}

	if options.special && !password.special {
		validationErr = append(validationErr, *NewPasswordSpecialCharacterErr())
	}

	if options.letters != nil && *options.letters <= *password.letters {
		validationErr = append(validationErr, *NewPasswordCharacterLengthErr(*options.letters))
	}

	if len(validationErr) > 0 {
		var status int16 = 404
		return errs.NewCustom(
			"PASSWORD_VALIDATION_ERR",
			"password validation failed",
			&validationErr,
			&status,
		)
	}

	return nil
}
