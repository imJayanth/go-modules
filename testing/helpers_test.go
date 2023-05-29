package testing

import (
	"log"
	"testing"

	"github.com/imJayanth/go-modules/helpers"
)

func TestValidatePassword(t *testing.T) {
	log.Println(helpers.ValidatePassword("Password1!"))
}
