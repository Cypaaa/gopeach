package main

import (
	"crypto/ed25519"
	"fmt"
	"log"

	"github.com/Cypaaa/gopeach"
)

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func main() {
	publicKey, privateKey, _ := ed25519.GenerateKey(nil)

	// Generate JWT
	u := User{ID: 1, Name: "gopeach"}
	tokenString, err := gopeach.GenerateJWT(privateKey, u)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(tokenString)

	// Verify JWT
	claims, err := gopeach.VerifyJWT(publicKey, tokenString)
	if err != nil {
		panic(err)
	}

	// check if values are the same
	fmt.Println(claims["ID"].(float64) == float64(u.ID))
	fmt.Println(claims["Name"] == u.Name)
}
