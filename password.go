package validator

import (
	errs "github.com/petar-arandjic/goerr"
	"unicode"
)

type PasswordValidate struct {
	Number  bool
	Upper   bool
	Lower   bool
	Special bool
	Letters *int16
}

func IsStrongPassword(s string, options PasswordValidate) *errs.Error {
	var letters int16 = 0
	password := PasswordValidate{
		Number:  false,
		Upper:   false,
		Lower:   false,
		Special: false,
		Letters: &letters,
	}
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			password.Number = true
			*password.Letters++
		case unicode.IsUpper(c):
			password.Upper = true
			*password.Letters++
		case unicode.IsLower(c):
			password.Lower = true
			*password.Letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			password.Special = true
			*password.Letters++
		default:
			*password.Letters++
		}
	}

	var validationErr []errs.Error

	if options.Number && !password.Number {
		validationErr = append(validationErr, *NewPasswordNumberErr())
	}

	if options.Lower && !password.Lower {
		validationErr = append(validationErr, *NewPasswordLowerLetterErr())
	}

	if options.Upper && !password.Upper {
		validationErr = append(validationErr, *NewPasswordNumberErr())
	}

	if options.Special && !password.Special {
		validationErr = append(validationErr, *NewPasswordSpecialCharacterErr())
	}

	if options.Letters != nil && *options.Letters > *password.Letters {
		validationErr = append(validationErr, *NewPasswordCharacterLengthErr(*options.Letters))
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
