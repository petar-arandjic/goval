package main

import (
	"fmt"
	validator "github.com/petar-arandjic/goval"
)

func main() {
	var letters int16 = 8
	err := validator.IsStrongPassword("n(Ikojij9", validator.PasswordValidate{
		Number:  true,
		Upper:   true,
		Lower:   true,
		Special: true,
		Letters: &letters,
	})

	fmt.Println(err)
}
