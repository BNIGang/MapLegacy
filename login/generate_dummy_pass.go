package login // change to main if wanna run directly

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func generateHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func main() {
	res, _ := generateHash("pass") // change this as needed
	fmt.Print(res)
}
