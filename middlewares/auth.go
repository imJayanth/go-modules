package middlewares

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/imJayanth/go-modules/config"
	"github.com/imJayanth/go-modules/errors"
	"github.com/imJayanth/go-modules/models"
	"github.com/imJayanth/go-modules/response"
)

func RequireBasicAuth(authConfig *config.AuthConfig) fiber.Handler {
	return basicauth.New(basicauth.Config{
		Users: map[string]string{
			authConfig.Username: authConfig.Password,
		},
		Unauthorized: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusUnauthorized).JSON(errors.NewUnAuthorizedError("Invalid credentials"))
		},
	})
}

func RequireTokenAuth(authConfig *config.AuthConfig) fiber.Handler {
	return func(c *fiber.Ctx) error {
		if err := validateToken(c, authConfig); err.IsNotNull() {
			return response.RespondUnAuthorised(c, err.Message)
		}
		return c.Next()
	}
}

func GenerateJwtToken(authConfig *config.AuthConfig, user *models.User) (models.JwtToken, errors.RestAPIError) {
	var jwtClaims models.JWTClaims
	inrec, _ := json.Marshal(user)
	json.Unmarshal(inrec, &jwtClaims)
	jwtClaims.StandardClaims.ExpiresAt = time.Now().Add(time.Hour * 72).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtClaims)
	tokenString, err := token.SignedString([]byte(authConfig.JwtSecret))
	if err != nil {
		fmt.Println(err)
		return models.JwtToken{}, errors.NewInternalServerError("Something unexpected happened")
	}

	return models.JwtToken{Token: tokenString}, errors.NO_ERROR()
}

func validateToken(c *fiber.Ctx, authConfig *config.AuthConfig) errors.RestAPIError {
	authHeaderByte := c.Request().Header.Peek("Authorization")
	authHeader := string(authHeaderByte)
	if authHeader == "" {
		return errors.NewUnAuthorizedError("User not signed in.")
	}

	bearerToken := strings.Split(authHeader, " ")

	if len(bearerToken) != 2 || bearerToken[1] == "" {
		return errors.NewUnAuthorizedError("Invalid Token.")
	}

	authToken := bearerToken[1]

	parsedToken, parsedError := jwt.ParseWithClaims(authToken, &models.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an error parsing the token")
		}

		return []byte(authConfig.JwtSecret), nil
	})

	if parsedError != nil {
		fmt.Println("Error while Parsing the token: ", parsedError.Error())
		return errors.NewUnAuthorizedError("Invalid Token.")
	}

	if jwtClaims, ok := parsedToken.Claims.(*models.JWTClaims); ok && parsedToken.Valid {
		c.Locals("loggedInUser", models.User{
			ID:             jwtClaims.ID,
			Name:           jwtClaims.Name,
			Mobile:         &jwtClaims.Mobile,
			CountryCode:    &jwtClaims.CountryCode,
			MobileVerified: jwtClaims.MobileVerified,
			Email:          &jwtClaims.Email,
			EmailVerified:  jwtClaims.EmailVerified,
			Roles:          jwtClaims.Roles,
		})
		return errors.NO_ERROR()
	}

	return errors.NewUnAuthorizedError("Invalid Token")
}
