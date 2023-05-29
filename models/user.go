package models

import (
	"fmt"

	"github.com/imJayanth/go-modules/helpers"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

type UserRole string

const (
	UNVERIFIED UserRole = "unverified"
	BASIC      UserRole = "basic"
	ADMIN      UserRole = "admin"
	MERCHANT   UserRole = "merchant"
	GUEST      UserRole = "guest"
)

type User struct {
	ID             uint                          `json:"id" gorm:"primaryKey;autoIncrement"`
	Name           string                        `json:"name" gorm:"not null;type:varchar(255)"`
	Mobile         *helpers.MobileNumber         `json:"mobile" gorm:"uniqueIndex:idx_mobile_country_code;type:varchar(15)"`
	CountryCode    *string                       `json:"country_code" gorm:"uniqueIndex:idx_mobile_country_code;type:varchar(5)"`
	MobileVerified bool                          `json:"mobile_verified" gorm:"uniqueIndex:idx_mobile_country_code"`
	Email          *helpers.Email                `json:"email" gorm:"uniqueIndex:idx_email;type:varchar(255)"`
	EmailVerified  bool                          `json:"email_verified" gorm:"uniqueIndex:idx_email"`
	Password       string                        `json:"password" gorm:"type:varchar(255)"`
	Country        *helpers.Country              `json:"country"`
	Roles          datatypes.JSONSlice[UserRole] `json:"roles"`
	Base
}

func (u *User) GetMobile() helpers.MobileNumber {
	if u.Mobile != nil {
		return *u.Mobile
	}
	return ""
}

func (u *User) GetCountryCode() string {
	if u.CountryCode != nil {
		return *u.CountryCode
	}
	return ""
}

func (u *User) GetMobileWithCountryCode() string {
	mobile := u.GetMobile()
	countryCode := u.GetCountryCode()
	if mobile != "" && countryCode != "" {
		return fmt.Sprintf("%v-%v", mobile, countryCode)
	}
	return ""
}

func (u *User) GetEmail() helpers.Email {
	if u.Email != nil {
		return *u.Email
	}
	return ""
}

func (u *User) SetPassword(password string) error {
	if e := helpers.ValidatePassword(password); e != nil {
		return e
	}
	hash, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if e != nil {
		return e
	}
	u.Password = string(hash)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

func (u *User) Mask() {
	if u.Password != "" {
		u.Password = "--masked--"
	}
}
