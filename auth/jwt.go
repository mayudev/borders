package auth

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/melsincostan/borders/auth/key"
)

const (
	Validity  = 1 * time.Hour
	CookieKey = "borders-admin-auth"
	issuer    = "borders-app-auth"
)

var (
	ErrUnparseableClaims = errors.New("claims could not be cast into the expected format")
	ErrExpired           = errors.New("the token expired")
	ErrTooNew            = errors.New("the token cannot be used yet")
	ErrWrongIssuer       = errors.New("the issuer doesn't match the expected value")
	ErrIssuedInTheFuture = errors.New("the token was issued in the future")
	ErrValidityTooLong   = errors.New("the validity period of the token is too long")
)

type customClaims struct {
	UID uint `json:"uid"`
	jwt.RegisteredClaims
}

func Create(userID uint) (signedToken string, err error) {
	ct := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": issuer,
		"iat": ct.Unix(),
		"nbf": ct.Add(-(5 * time.Second)).Unix(),
		"exp": ct.Add(Validity).Unix(),
		"uid": userID,
	})
	return token.SignedString(key.Get())
}

func Check(signedToken string) (sub uint, exp time.Time, err error) {
	token, err := jwt.ParseWithClaims(signedToken, &customClaims{}, key.Func())
	if err != nil || !token.Valid {
		return 0, exp, fmt.Errorf("could not parse token: %w", err)
	}
	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return 0, exp, ErrUnparseableClaims
	}

	texp, err := claims.GetExpirationTime()
	if err != nil {
		return 0, exp, fmt.Errorf("%w (exp)", err)
	}
	if texp.Before(time.Now()) {
		return 0, exp, ErrExpired
	}

	nbf, err := claims.GetNotBefore()
	if err != nil {
		return 0, exp, fmt.Errorf("%w (nbf)", err)
	}
	if nbf.After(time.Now()) {
		return 0, exp, ErrTooNew
	}

	iat, err := claims.GetIssuedAt()
	if err != nil {
		return 0, exp, fmt.Errorf("%w (nbf)", err)
	}
	if iat.After(time.Now()) {
		return 0, exp, ErrIssuedInTheFuture
	}

	iss, err := claims.GetIssuer()
	if err != nil {
		return 0, exp, fmt.Errorf("%w (iss)", err)
	}
	if iss != issuer {
		return 0, exp, ErrWrongIssuer
	}

	if exp.Sub(iat.Time) > Validity {
		return 0, exp, ErrValidityTooLong
	}

	exp = texp.Time
	sub = claims.UID

	return
}
