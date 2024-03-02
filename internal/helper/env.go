package helper

import "os"

func IsLocal() bool {
	e := os.Getenv("ENV")
	switch e {
	case "local":
		return true
	default:
		return false
	}
}
