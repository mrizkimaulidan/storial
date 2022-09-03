package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mrizkimaulidan/storial/internal/config"
	"github.com/mrizkimaulidan/storial/internal/entity"
)

// Context key type for user information.
type ContextKeyUserInformation string

var c = config.New().GetConfig()

var (
	SECRET_KEY                                      = []byte(c.JWT_SECRET_KEY)
	CtxKeyUserInformation ContextKeyUserInformation = "userInformation"
)

// Custom claims for JWT.
type CustomClaims struct {
	jwt.RegisteredClaims
	Id    uint64
	Name  string
	Email string
}

// Generate JSON Web Token.
func GenerateToken(u entity.User) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)),
		},
		Id:    u.Id,
		Name:  u.Name,
		Email: u.Email,
	}

	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := t.SignedString(SECRET_KEY)
	if err != nil {
		return "", err
	}

	return token, nil
}
