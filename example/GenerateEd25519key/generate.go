package main

import (
	"crypto/ed25519"
	"encoding/base64"
	"os"
)

const seedFileName = "seed.pem"

func main() {
	_, privateKey, _ := ed25519.GenerateKey(nil)

	// Write a readable seed to the file
	// making the seed readable is optional
	f1, _ := os.Create(seedFileName)
	defer f1.Close()
	f1.Write([]byte(base64.StdEncoding.EncodeToString(privateKey.Seed())))

	// Read the seed from the file
	privFile, _ := os.ReadFile(seedFileName)
	seed, _ := base64.StdEncoding.DecodeString(string(privFile))
	// Panics if error
	ed25519.NewKeyFromSeed([]byte(seed))

	// regenPrivateKey, _ := ed25519.NewKeyFromSeed([]byte(seed))
	// regenPrivateKey must be equal to privateKey
}
