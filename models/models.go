package models

const (
	TABLE_USERS    = "users"
	TABLE_USER_OTP = "user_otp"
)

func (User) TableName() string {
	return TABLE_USERS
}

func (UserOtp) TableName() string {
	return TABLE_USER_OTP
}
