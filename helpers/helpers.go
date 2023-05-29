package helpers

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"regexp"
	"strings"

	"errors"

	"github.com/nyaruka/phonenumbers"
)

const (
	otpChars = "1234567890"
)

func ValidateMobileNumber(number, countryCode string) bool {
	num, err := phonenumbers.Parse(fmt.Sprintf("%v-%v", countryCode, number), "")
	if err != nil {
		return false
	}
	return phonenumbers.IsValidNumber(num)
}

func GetEmailDomain(email string) (string, error) {
	// Check if the email address is valid
	if !isValidEmail(email) {
		return "", fmt.Errorf("invalid email address: %s", email)
	}

	// Split the email address into two parts: the local part and the domain part
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid email address: %s", email)
	}

	// Return the domain part of the email address
	return parts[1], nil
}

func ValidateEmailDomain(email string) error {
	domain, err := GetEmailDomain(email)
	if err != nil {
		return err
	}

	// Use DNS lookup to check if the domain has a valid mail server
	_, err = net.LookupMX(domain)
	if err != nil {
		return err
	}

	return nil
}

func isValidEmail(email string) bool {
	// Define a regular expression to check if the email address is valid
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	// Check if the email address matches the regular expression
	return regex.MatchString(email)
}

func JsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "\t")
	if err != nil {
		return in
	}
	return out.String()
}

func ToJson(o interface{}) string {
	js, serr := json.Marshal(o)
	if serr != nil {
		log.Println("Error while marshalling: ", serr)
	}
	return string(js)
}

func GenerateOTP(length int) string {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		panic(err)
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer)
}

func ValidatePassword(password string) error {
	// Define the regular expression to match the password criteria
	pattern := `^[a-zA-Z0-9~!@#\\$%^&*()_+]{8,15}.*[~!@#\\$%^&*()_+].*$`
	// Compile the regular expression
	regex, err := regexp.Compile(pattern)
	if err != nil {
		panic(err)
	}

	// Test if the password matches the regular expression
	if !regex.MatchString(password) {
		return errors.New("password must be between 8 and 15 characters long and contain at least one lowercase letter, one uppercase letter, one digit, and one special character ~!@#$%^&*()_+")
	}

	// Password is valid
	return nil
}
