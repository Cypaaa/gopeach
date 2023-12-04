package gopeach

import (
	"crypto/ed25519"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// Generate generates a JWT token. It takes a jwtSecretKey string and a struct (interface{})
// jwtSecretKey must be EdDSA private key
// and returns a token and an error.
func GenerateJWT(jwtSecretKey ed25519.PrivateKey, mc jwtv5.MapClaims) (string, error) {
	token := jwtv5.New(jwtv5.SigningMethodEdDSA)
	token.Claims = mc
	return token.SignedString(jwtSecretKey)
}

// Verify verifies a JWT token. It takes a jwtSecretKey string and a token string
func VerifyJWT(jwtSecretKey ed25519.PublicKey, token string, opts ...jwtv5.ParserOption) (jwtv5.MapClaims, error) {
	claims := jwtv5.MapClaims{}
	t, err := jwtv5.ParseWithClaims(
		token,
		claims,
		func(t *jwtv5.Token) (interface{}, error) {
			// check the signing method and get the expiration time from the token
			_, ok := t.Method.(*jwtv5.SigningMethodEd25519)
			if !ok {
				return nil, jwtv5.ErrSignatureInvalid
			}
			return jwtSecretKey, nil
		},
		opts...)
	if t != nil && t.Valid {
	}
	return claims, err
}
