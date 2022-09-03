package authentication

import (
	"net/http"
)

type AuthenticationHandler interface {
	Register() http.Handler
	Login() http.Handler
}
