package model

import (
	"errors"
	"github.com/asaskevich/govalidator"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id" db:"id"`
	Login        string `json:"login" binding:"required,min=4,max=15"`
	Email        string `json:"email" binding:"required,email"`
	Password     string `json:"password,omitempty" binding:"required,min=6,max=25"`
	RePassword   string `json:"rePassword,omitempty" binding:"required,min=6,max=25"`
	PasswordHash string `json:"-"`
	Role         int    `json:"role"`
	AccessToken  string `json:"-"`
	RefreshToken string `json:"-"`
}

func (u *User) Validate() bool {
	return govalidator.IsEmail(u.Email) &&
		govalidator.StringLength(u.Login, "5", "15") &&
		govalidator.StringLength(u.Password, "5", "30")
}

func (u *User) BeforeCreate() error {
	if len(u.Password) >= 5 {
		enc, err := encryptPassword(u.Password)
		if err != nil {
			return err
		}
		u.PasswordHash = enc
		return nil
	}
	return errors.New("пароль слишком короткий")
}

func (u *User) Sanitize() {
	u.Password = ""
	u.RePassword = ""
}

func (u *User) CheckUserPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func encryptPassword(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
