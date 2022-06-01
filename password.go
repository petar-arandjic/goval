package validator

import (
	"fmt"
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

func IsStrongPassword(s string, options PasswordValidate) *errs.SubError {
	var letters int16 = 0

	password := PasswordValidate{
		Number:  false,
		Upper:   false,
		Lower:   false,
		Special: false,
		Letters: &letters,
	}

	// check what password contains
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

	err := errs.SubError{
		Key:     "STRONG_PASSWORD_VALIDATION_FAILED",
		Message: "password must contain | ",
	}
	isValid := true

	if options.Number && !password.Number {
		err.Message = fmt.Sprintf("%s number | ", err.Message)
		isValid = false
	}

	if options.Lower && !password.Lower {
		err.Message = fmt.Sprintf("%s lower letter | ", err.Message)
		isValid = false
	}

	if options.Upper && !password.Upper {
		err.Message = fmt.Sprintf("%s uppder letter | ", err.Message)
		isValid = false
	}

	if options.Special && !password.Special {
		err.Message = fmt.Sprintf("%s special character | ", err.Message)
		isValid = false
	}

	if options.Letters != nil && *options.Letters > *password.Letters {
		err.Message = fmt.Sprintf("%s %v characters long | ", err.Message, *options.Letters)
		isValid = false
	}

	if !isValid {
		return &err
	}

	return nil
}
