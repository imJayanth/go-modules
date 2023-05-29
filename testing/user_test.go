package testing

import (
	"log"
	"testing"

	"github.com/imJayanth/go-modules/models"
)

func TestGetMobileNumberWithCountryCode(t *testing.T) {
	u := models.User{}
	log.Println(u.GetMobileWithCountryCode())
}
