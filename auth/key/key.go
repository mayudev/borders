package key

import (
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

var key = []byte{}

var (
	ErrWrongAlgorithm = errors.New("token uses the wrong algorithm for the key")
)

func Get() []byte {
	if len(key) < 1 {
		panic(fmt.Errorf("trying to use an uninitialized or empty JWT signing key"))
	}
	return key
}

func Func() jwt.Keyfunc {
	return func(t *jwt.Token) (interface{}, error) {
		if t.Header["alg"] != jwt.SigningMethodHS256.Alg() {
			return nil, ErrWrongAlgorithm
		}
		return Get(), nil
	}
}

type Config struct {
	File   string `env:"JWT_KEY_FILE" default:"./key.bin"`
	Length uint   `env:"JWT_KEY_LENGTH" default:"16"`
}
