package authentication

import (
	"errors"
	"log"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	model "github.com/mrizkimaulidan/storial/internal/model/authentication"
	"github.com/mrizkimaulidan/storial/internal/service/authentication"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/authentication"
	"github.com/mrizkimaulidan/storial/pkg/response"
)

type authenticationHandler struct {
	authenticationService authentication.AuthenticationService
	response              *response.Response
}

func (ah *authenticationHandler) Register() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := model.RegisterRequest{
			Name:     r.PostFormValue("name"),
			Username: r.PostFormValue("username"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
			Sex:      r.PostFormValue("sex"),
		}

		registeredResponse, err := ah.authenticationService.Register(r.Context(), request)
		if err != nil {
			ah.handleErr(err).JSON(w)
			return
		}

		ah.response.SetCode(http.StatusCreated).SetMessage("OK").SetData(registeredResponse).JSON(w)
	})
}

func (ah *authenticationHandler) Login() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		request := model.LoginRequest{
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}

		loginResponse, err := ah.authenticationService.Login(r.Context(), request)
		if err != nil {
			ah.handleErr(err).JSON(w)
			return
		}

		ah.response.SetCode(http.StatusOK).SetMessage("OK").SetData(loginResponse).JSON(w)
	})
}

func (ah *authenticationHandler) handleErr(err error) *response.Response {
	switch {
	case errors.As(err, &validation.Errors{}):
		return ah.response.Error(err).SetCode(http.StatusBadRequest)
	case errors.Is(err, exception.ErrEmailAlreadyExists):
		return ah.response.Error(err).SetCode(http.StatusBadRequest)
	case errors.Is(err, exception.ErrEmailNotFound):
		return ah.response.Error(err).SetCode(http.StatusNotFound)
	case errors.Is(err, exception.ErrUsernameAlreadyExists):
		return ah.response.Error(err).SetCode(http.StatusBadRequest)
	case errors.Is(err, exception.ErrPasswordAreWrong):
		return ah.response.Error(err).SetCode(http.StatusUnauthorized)
	}

	log.Println("[ERROR]", err)
	return ah.response.Error(err).SetCode(http.StatusInternalServerError).SetMessage("internal server error")
}

func NewHandler(authenticationService authentication.AuthenticationService) AuthenticationHandler {
	return &authenticationHandler{
		authenticationService: authenticationService,
		response:              new(response.Response),
	}
}
