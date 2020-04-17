package users

import (
	"strings"
	"github.com/james077/bookstore_utils-go/rest_errors"
)
const (
	StatusActive = "active"
)

// User repesenta la entidad usuario
type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
	Password    string `json:"password"`
} 
/*
func (user *User) Validate() *errors.RestErr{
	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return errors.NewBadRequestError("Direccion de correo electrónico inválida")
	}
	return nil
}*/

type Users []User

func (user *User) Validate() rest_errors.RestErr {
	user.FirstName = strings.TrimSpace(user.FirstName)
	user.LastName = strings.TrimSpace(user.LastName)

	user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	if user.Email == "" {
		return rest_errors.NewBadRequestError("Direccion de correo electrónico inválida")
	}

	user.Password = strings.TrimSpace(user.Password)
	if user.Password == "" {
		return rest_errors.NewBadRequestError("Contraseña invalida")
	}
	return nil
}