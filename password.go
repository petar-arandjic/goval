package validator

import "unicode"

type passwordValidate struct {
	number  bool
	upper   bool
	lower   bool
	special bool
	letters int16
}

func IsStrongPassword(s string) bool {
	var password passwordValidate
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			password.number = true
			password.letters++
		case unicode.IsUpper(c):
			password.upper = true
			password.letters++
		case unicode.IsLower(c):
			password.lower = true
			password.letters++
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			password.special = true
			password.letters++
		default:
			password.letters++
		}
	}

	if password.number && password.upper && password.special && password.letters > 8 {
		return true
	}
	return false
}
