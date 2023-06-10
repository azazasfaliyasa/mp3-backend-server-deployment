package auth

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

type Payload struct {
	Username string
	Password string
}

type JWTClaim struct {
	Username string
	jwt.RegisteredClaims
}
