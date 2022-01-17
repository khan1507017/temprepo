package helper

import (
	"errors"
	"tempApi/dto"
)

func ValidateInput(payload dto.Info) error {
	if payload.Name == "" || payload.Branch == "" || payload.Designation == "" {
		return errors.New("invalid payload, fields cannot be empty")
	}
	return nil
}
