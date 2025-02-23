package utils

import (
	"errors"
	"regexp"
)

// ValidateMobileNumber checks if the mobile number is exactly 10 digits. and only contains numbers
func ValidateMobileNumber(mobile string) error {
	mobilePattern := `^[6-9]\d{9}$` // Starts with 6-9 and has 10 digits (Indian mobile format)
	matched, _ := regexp.MatchString(mobilePattern, mobile)
	if !matched {
		return errors.New("invalid mobile number: must be 10 digits and start with 6-9")
	}
	return nil
}

// ValidateEmail checks if the email is in a proper format.
func ValidateEmail(email string) error {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	matched, _ := regexp.MatchString(emailPattern, email)
	if !matched {
		return errors.New("invalid email address format")
	}
	return nil
}

// ValidateName ensures the name contains only alphabets and spaces, and has a proper length, not too big not too small.
func ValidateName(name string) error {
	namePattern := `^[A-Za-z\s]{3,50}$` // Name should be 3-50 characters long and contain only alphabets and spaces
	matched, _ := regexp.MatchString(namePattern, name)
	if !matched {
		return errors.New("invalid name: must be between 3-50 characters and contain only alphabets")
	}
	return nil
}

// ValidatePassword ensures the password has at least 8 characters.
func ValidatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}

// ValidateUser validates the Worker or Employer registration details.
func ValidateUser(name, mobile, email, password string) error {
	if err := ValidateName(name); err != nil {
		return err
	}
	if err := ValidateMobileNumber(mobile); err != nil {
		return err
	}
	if err := ValidateEmail(email); err != nil {
		return err
	}
	if err := ValidatePassword(password); err != nil {
		return err
	}
	return nil
}

func ValidateUpdateUser(name, mobile, email string) error {
	if err := ValidateName(name); err != nil {
		return err
	}
	if err := ValidateMobileNumber(mobile); err != nil {
		return err
	}
	if err := ValidateEmail(email); err != nil {
		return err
	}
	return nil
}
