package utils

import "fmt"

func ErrJSON(message string) map[string]string {
	return map[string]string{
		"error": fmt.Sprintf("%s - please check the server logs", message),
	}
}
