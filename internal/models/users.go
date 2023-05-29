package models

import (
	"fmt"
	"regexp"
)

type User struct {
	Base

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	Email     string `json:"email"`
}

var (
	uppercaseRegex   = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex   = regexp.MustCompile(`[a-z]`)
	numberRegex      = regexp.MustCompile(`[0-9]`)
	specialCharRegex = regexp.MustCompile(`[!@#%&*_+=-]`)
	emailRegex       = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func (u *User) Validate() error {
	if err := valideNames(u); err != nil {
		return err
	}

	if err := checkPass(u.Password); err != nil {
		return err
	}

	if !emailRegex.MatchString(u.Email) {
		return fmt.Errorf("the email is invalid")
	}

	return nil
}

func valideNames(u *User) error {
	if u.FirstName == "" {
		return fmt.Errorf("first name cannot be empty")
	}

	if u.LastName == "" {
		return fmt.Errorf("last name cannot be empty")
	}

	return nil
}

func checkPass(p string) error {
	if len(p) < 8 {
		return fmt.Errorf("password must be at least 8 characters long")
	}

	if !uppercaseRegex.MatchString(p) {
		return fmt.Errorf("password must contain at least one uppercase letter")
	}

	if !lowercaseRegex.MatchString(p) {
		return fmt.Errorf("password must contain at least one lowercase letter")
	}

	if !numberRegex.MatchString(p) {
		return fmt.Errorf("password must contain at least one number")
	}

	if !specialCharRegex.MatchString(p) {
		return fmt.Errorf("password must contain at least one special character")
	}
	return nil
}
