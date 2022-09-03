package authentication

import (
	"context"

	"github.com/mrizkimaulidan/storial/internal/model/authentication"
)

type AuthenticationService interface {
	Register(ctx context.Context, r authentication.RegisterRequest) (*authentication.RegisterResponse, error)
	Login(ctx context.Context, r authentication.LoginRequest) (*authentication.LoginResponse, error)
}
