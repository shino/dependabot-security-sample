package main

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Create a new token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": "example",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	// Sign the token with a secret
	tokenString, err := token.SignedString([]byte("my-secret-key"))
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}

	fmt.Println("Generated JWT token:", tokenString)
}
