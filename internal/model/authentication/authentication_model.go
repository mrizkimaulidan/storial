package authentication

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type RegisterRequest struct {
	Name     string
	Username string
	Email    string
	Password string
	Sex      string
}

func (rr *RegisterRequest) Validate() error {
	return validation.ValidateStruct(rr,
		validation.Field(&rr.Name, validation.Required, validation.Length(5, 255)),
		validation.Field(&rr.Username, validation.Required, validation.Length(5, 255)),
		validation.Field(&rr.Email, validation.Required, is.Email, validation.Length(5, 255)),
		validation.Field(&rr.Password, validation.Required, validation.Length(5, 255)),
		validation.Field(&rr.Sex, validation.Required, validation.In("0", "1", "2", "9")),
	)
}

type RegisterResponse struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Sex      string `json:"sex"`
}

type LoginRequest struct {
	Email    string
	Password string
}

func (lr *LoginRequest) Validate() error {
	return validation.ValidateStruct(lr,
		validation.Field(&lr.Email, validation.Required, is.Email, validation.Length(5, 255)),
		validation.Field(&lr.Password, validation.Required, validation.Length(5, 255)),
	)
}

type LoginResponse struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	Sex      string `json:"sex"`
}
