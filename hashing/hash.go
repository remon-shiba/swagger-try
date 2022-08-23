package hashing

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func VerifyHashPassword(password, hash string) bool {
	fmt.Println("(Monitor) hash.go | Password:", password, "Hash:", hash)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetVerifiedHash(password string) (hashPassword string) {
	c := *&fiber.Ctx{}
	hashPw, hashErr := GenerateHashPassword(password)
	if hashErr != nil {
		c.Status(http.StatusUnprocessableEntity).JSON(
			&fiber.Map{"message": "request failed to register user",
				"error": hashErr})
	}

	fmt.Println("Password:", password, "hash:", hashPw)
	isMatch := VerifyHashPassword(password, hashPw)
	fmt.Println("(Monitor) hash.go | Match:   ", isMatch)
	return hashPw
}
