package models

import "time"

type OtpMedium string

const (
	OTP_LENGTH                       = 6
	OTP_EXPIRE_TIME_IN_MIN           = 15
	OTP_MEDIUM_MOBILE      OtpMedium = "mobile"
	OTP_MEDIUM_MAIL        OtpMedium = "mail"
)

func (o OtpMedium) IsMobile() bool {
	return o == OTP_MEDIUM_MOBILE
}

func (o OtpMedium) IsMail() bool {
	return o == OTP_MEDIUM_MAIL
}

type UserOtp struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    uint      `json:"use_id"`
	Receipent string    `json:"receipent"`
	Otp       string    `json:"otp"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
