package key

import (
	"fmt"
)

var key = []byte{}

func Get() []byte {
	if len(key) < 1 {
		panic(fmt.Errorf("trying to use an uninitialized or empty JWT signing key"))
	}
	return key
}

type Config struct {
	File   string `env:"JWT_KEY_FILE" default:"./key.bin"`
	Length uint   `env:"JWT_KEY_LENGTH" default:"16"`
}
