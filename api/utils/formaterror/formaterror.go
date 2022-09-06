package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("nickname Already Taken")
	}

	if strings.Contains(err, "email") {
		return errors.New("email Already Taken")
	}

	if strings.Contains(err, "tittle") {
		return errors.New("tittle Already Taken")
	}

	if strings.Contains(err, "hashedPassword") {
		return errors.New("incorrect password")
	}
	return errors.New(err)
}
