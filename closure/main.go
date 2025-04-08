package main

import (
	"fmt"
	"unicode"
)

// Closure döndüren validator factory fonksiyonu
func passwordValidator(minLength int, requireUpper bool, requireNumber bool) func(string) bool {
	return func(password string) bool {
		if len(password) < minLength {
			return false
		}

		hasUpper := false
		hasNumber := false

		for _, ch := range password {
			if unicode.IsUpper(ch) {
				hasUpper = true
			}
			if unicode.IsDigit(ch) {
				hasNumber = true
			}
		}

		if requireUpper && !hasUpper {
			return false
		}
		if requireNumber && !hasNumber {
			return false
		}

		return true
	}
}

func main() {
	// Doğrulayıcıyı oluştur
	validate := passwordValidator(8, true, true)

	passwords := []string{"123456", "Password", "Pass1234", "pass1234", "PASSWORD123"}

	for _, p := range passwords {
		if validate(p) {
			fmt.Printf("%s geçerli\n", p)
		} else {
			fmt.Printf("%s geçersiz\n", p)
		}
	}
}
