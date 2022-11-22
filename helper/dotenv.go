package helper

import (
	"os"

	"github.com/joho/godotenv"
)

func Env(env string) string {
	err := godotenv.Load()
	IfError(err)

	return os.Getenv(env)
}
