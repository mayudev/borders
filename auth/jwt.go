package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	validity = 1 * time.Hour
)

func Create() (token jwt.Token) {
	return
}
