package models

import (
	"fmt"

	"github.com/melsincostan/argon2id"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserBase
}

type UserBase struct {
	Name     string
	Password string
}

func (u *User) Check(password string) (err error) {
	pw, err := argon2id.Parse(u.Password)
	if err != nil {
		return fmt.Errorf("could not parse password (user '%s'): %w", u.Name, err)
	}

	if !pw.Compare(password) {
		return fmt.Errorf("Wrong password (user '%s')", u.Name)
	}
	return
}
