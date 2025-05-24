package setup

import "errors"

const version = 1

var (
	ErrNotSetup     = errors.New("database isn't set up")
	ErrWrongVersion = errors.New("program version doesn't match the saved version")
	ErrAlreadySetup = errors.New("database is already setup (to delete all data and start anew, remove the .db file)")
)
