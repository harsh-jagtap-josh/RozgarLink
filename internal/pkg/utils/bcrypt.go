package utils

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {

	err := godotenv.Load()
	if err != nil {
		return "", err
	}

	costStr := os.Getenv("BCRYPT_COST")
	cost, err := strconv.Atoi(costStr)
	if err != nil {
		return "", err
	}
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
