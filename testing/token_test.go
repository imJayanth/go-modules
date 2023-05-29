package testing

import (
	"log"
	"testing"

	"github.com/imJayanth/go-modules/config"
	"github.com/imJayanth/go-modules/middlewares"
	"github.com/imJayanth/go-modules/models"
)

func TestGenerateToken(t *testing.T) {
	appConfig := config.SetupConfig()
	token, err := middlewares.GenerateJwtToken(&appConfig.AuthConfig, &models.User{
		ID:   1,
		Name: "name",
	})
	if err.IsNotNull() {
		log.Println(err)
	} else {
		log.Println(token)
	}
}
