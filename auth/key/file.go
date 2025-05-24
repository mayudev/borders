package key

import (
	"crypto/rand"
	"fmt"
	"io"
	"os"
)

func Create(kc *Config) (err error) {
	kb := make([]byte, kc.Length)
	if _, err := rand.Read(kb); err != nil {
		return fmt.Errorf("couldn't create a signing key: %w", err)
	}

	file, err := os.OpenFile(kc.File, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		return fmt.Errorf("couldn't create/open the key file: %w", err)
	}
	defer file.Close()

	if _, err := file.Write(kb); err != nil {
		return fmt.Errorf("couldn't write key to the key file: %w", err)
	}
	return
}

func Load(kc *Config) (err error) {
	file, err := os.Open(kc.File)
	if err != nil {
		return fmt.Errorf("couldn't open the key file: %w", err)
	}
	defer file.Close()
	key, err = io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("couldn't read the key: %w", err)
	}
	return
}
