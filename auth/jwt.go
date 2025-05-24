package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/melsincostan/borders/auth/key"
)

const (
	validity = 1 * time.Hour
	issuer   = "borders-app-auth"
)

var (
	ErrUnparseableClaims = errors.New("claims could not be cast into the expected format")
	ErrExpired           = errors.New("the token expired")
	ErrTooNew            = errors.New("the token cannot be used yet")
	ErrWrongIssuer       = errors.New("the issuer doesn't match the expected value")
	ErrIssuedInTheFuture = errors.New("the token was issued in the future")
	ErrValidityTooLong   = errors.New("the validity period of the token is too long")
)

func Create() (signedToken string, err error) {
	ct := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issuer,
		"iat": ct.Unix(),
		"nbf": ct.Add(-(5 * time.Second)).Unix(),
		"exp": ct.Add(validity).Unix(),
	})
	return token.SignedString(key.Get())
}

func Check(signedToken string) (err error) {
	token, err := jwt.Parse(signedToken, key.Func())
	if err != nil || !token.Valid {
		return fmt.Errorf("could not parse token: %w", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return ErrUnparseableClaims
	}

	exp, err := claims.GetExpirationTime()
	if err != nil {
		return fmt.Errorf("%w (exp)", err)
	}
	if exp.Before(time.Now()) {
		return ErrExpired
	}

	nbf, err := claims.GetNotBefore()
	if err != nil {
		return fmt.Errorf("%w (nbf)", err)
	}
	if nbf.Before(time.Now()) {
		return ErrTooNew
	}

	iat, err := claims.GetIssuedAt()
	if err != nil {
		return fmt.Errorf("%w (nbf)", err)
	}
	if iat.Before(time.Now()) {
		return ErrIssuedInTheFuture
	}

	iss, err := claims.GetIssuer()
	if err != nil {
		return fmt.Errorf("%w (iss)", err)
	}
	if iss != issuer {
		return ErrWrongIssuer
	}

	if exp.Sub(iat.Time) > validity {
		return ErrValidityTooLong
	}

	return
}
