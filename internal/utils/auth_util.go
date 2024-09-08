package utils

import (
	"gogo/internal/app/entities"
	"log"
	"os"
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/golang-jwt/jwt/v5"
)

type Password struct {
	Plaintext *string
	Hash      string
}

func (p *Password) Set(plaintextPassword string) error {
	hash, err := argon2id.CreateHash(plaintextPassword, argon2id.DefaultParams)
	if err != nil {
		return err
	}
	p.Plaintext = &plaintextPassword
	p.Hash = hash
	return nil
}

func (p *Password) Matches(plaintextPassword string) (bool, error) {
	match, err := argon2id.ComparePasswordAndHash(plaintextPassword, p.Hash)
	if err != nil {
		log.Fatal(err)
	}

	return match, nil
}

func CreateToken(userData *entities.User) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":    userData.Username,
		"userId": userData.Id,                           // Subject (user identifier)
		"aud":    "user",                                // Audience (user role)
		"exp":    time.Now().Add(time.Hour * 60).Unix(), // Expiration time
		"iat":    time.Now().Unix(),                     // Issued at
	})

	secretKey := os.Getenv("SECRET_KEY")
	jwtKey := []byte(secretKey)
	tokenString, err := claims.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
