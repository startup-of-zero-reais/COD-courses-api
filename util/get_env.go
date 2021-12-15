package util

import "os"

func GetEnv(key, _default string) string {
	if e := os.Getenv(key); e != "" {
		return e
	}

	return _default
}
