package middleware

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v4"
	jwtpkg "github.com/mrizkimaulidan/storial/pkg/jwt"
	"github.com/mrizkimaulidan/storial/pkg/response"
)

type middleware struct {
	response *response.Response
}

// JWT Authorization middleware. If no authorization token on
// header, it will return unathorized status.
func (m *middleware) JWTAuthorization(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		t := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.ParseWithClaims(t, &jwtpkg.CustomClaims{}, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenUnverifiable
			}

			return jwtpkg.SECRET_KEY, nil
		})

		if !strings.Contains(authorizationHeader, "Bearer ") {
			m.response.SetCode(http.StatusUnauthorized).SetMessage(jwt.ErrInvalidKey.Error()).SetData(nil).JSON(w)
			return
		}

		// error checking
		if err != nil {
			v, ok := err.(*jwt.ValidationError)
			if !ok {
				log.Println("[ERROR]", err)
				log.Println("v failed type assertion")
				m.response.SetCode(http.StatusInternalServerError).SetMessage("INTERNAL SERVER ERROR").SetData(nil).JSON(w)
				return
			}

			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				m.response.SetCode(http.StatusUnauthorized).SetMessage(jwt.ErrSignatureInvalid.Error()).SetData(nil).JSON(w)
				return
			case jwt.ValidationErrorExpired:
				m.response.SetCode(http.StatusUnauthorized).SetMessage(jwt.ErrTokenExpired.Error()).SetData(nil).JSON(w)
				return
			case jwt.ValidationErrorMalformed:
				m.response.SetCode(http.StatusUnauthorized).SetMessage(jwt.ErrTokenMalformed.Error()).SetData(nil).JSON(w)
				return
			}
		}

		claims, ok := token.Claims.(*jwtpkg.CustomClaims)
		if !ok {
			log.Println("[ERROR]", err)
			log.Println("claims failed type assertion")
			m.response.SetCode(http.StatusInternalServerError).SetMessage("INTERNAL SERVER ERROR").SetData(nil).JSON(w)
			return
		}

		if token.Valid {
			ctx := context.WithValue(r.Context(), jwtpkg.CtxKeyUserInformation, claims)
			r = r.WithContext(ctx)

			n.ServeHTTP(w, r)
		}
	})
}

// Logging incoming request to terminal.
func (m *middleware) LoggingMiddleware(n http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.RemoteAddr, r.URL.Path)

		n.ServeHTTP(w, r)
	})
}

func New() *middleware {
	return &middleware{
		response: new(response.Response),
	}
}
