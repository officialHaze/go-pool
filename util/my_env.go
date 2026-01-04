package util

import (
	"os"
	"strings"
)

func InDevMode() bool {
	env := os.Getenv("ENV")
	env = strings.ToLower(env)

	return strings.Contains(env, "dev") || strings.Contains(env, "local")
}
