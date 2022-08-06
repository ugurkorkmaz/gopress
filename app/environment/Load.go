package environment

import (
	"errors"

	"github.com/joho/godotenv"
)

func Load() error {
	err := godotenv.Load()
	if err != nil {
		return errors.New("error loading .env file")
	}
	return nil
}
