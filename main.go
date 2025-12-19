package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	// Create a new token object
	token := jwt.New(jwt.SigningMethodHS256)
	
	// Set some claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "John Doe"
	claims["admin"] = true
	
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		fmt.Println("Error creating token:", err)
		return
	}
	
	fmt.Println("Token:", tokenString)
}
