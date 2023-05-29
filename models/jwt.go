package models

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/imJayanth/go-modules/helpers"
	"gorm.io/datatypes"
)

const (
	TOKEN_EXPIRE_TIME_IN_HR = 72
)

type JWTClaims struct {
	jwt.StandardClaims
	ID             uint                          `json:"id"`
	Name           string                        `json:"name"`
	Mobile         helpers.MobileNumber          `json:"mobile"`
	CountryCode    string                        `json:"country_code"`
	MobileVerified bool                          `json:"mobile_verified"`
	Email          helpers.Email                 `json:"email"`
	EmailVerified  bool                          `json:"email_verified"`
	Roles          datatypes.JSONSlice[UserRole] `json:"roles"`
}

type JwtToken struct {
	Token string `json:"token"`
}
