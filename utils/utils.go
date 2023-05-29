package utils

import (
	"strings"

	"github.com/imJayanth/go-modules/errors"
)

/*
 TODO:
 	1. Maybe convert to a switch case
	2. Add more error check conditions
*/
func MapErrorFromGorm(err error, msg ...string) errors.RestAPIError {

	if strings.Contains(err.Error(), "record not found") {
		return errors.NewNotFoundError("Invalid ID")
	} else if strings.Contains(err.Error(), "Duplicate entry") {
		return errors.NewDuplicateRecord("Duplicate Record")
	} else if strings.Contains(err.Error(), "Error 1452:") {
		return errors.NewInternalServerError(msg[0])
	}

	return errors.NewInternalServerError("something unexpected happen")

}
