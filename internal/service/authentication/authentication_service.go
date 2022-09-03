package authentication

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/mrizkimaulidan/storial/internal/database"
	"github.com/mrizkimaulidan/storial/internal/entity"
	model "github.com/mrizkimaulidan/storial/internal/model/authentication"
	"github.com/mrizkimaulidan/storial/internal/repository/authentication"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/authentication"
	"github.com/mrizkimaulidan/storial/pkg/jwt"
	"github.com/mrizkimaulidan/storial/pkg/password"
	"github.com/mrizkimaulidan/storial/pkg/time"
)

type authenticationService struct {
	authenticationRepository authentication.AuthenticationRepository
	db                       *sql.DB
}

func (as *authenticationService) Register(ctx context.Context, r model.RegisterRequest) (*model.RegisterResponse, error) {
	tx, err := as.db.Begin()
	if err != nil {
		return nil, err
	}
	defer database.CommitOrRollback(tx)

	err = r.Validate()
	if err != nil {
		return nil, err
	}

	ok, err := as.authenticationRepository.CheckIfEmailExists(ctx, tx, r.Email)
	if err != nil {
		return nil, err
	}

	if *ok {
		return nil, exception.ErrEmailAlreadyExists
	}

	ok, err = as.authenticationRepository.CheckIfUsernameExists(ctx, tx, r.Username)
	if err != nil {
		return nil, err
	}

	if *ok {
		return nil, exception.ErrUsernameAlreadyExists
	}

	password, err := password.HashPassword(r.Password)
	if err != nil {
		return nil, err
	}

	sex, err := strconv.Atoi(r.Sex)
	if err != nil {
		return nil, err
	}

	s := entity.User{
		Name:      r.Name,
		Username:  r.Username,
		Email:     r.Email,
		Password:  password,
		Sex:       uint8(sex),
		CreatedAt: time.CurrentTimeToUnixTimestamp(),
	}

	s.Id = uint64(s.GenerateID())

	registeredUser, err := as.authenticationRepository.Register(ctx, tx, s)
	if err != nil {
		return nil, err
	}

	token, err := jwt.GenerateToken(*registeredUser)
	if err != nil {
		return nil, err
	}

	return &model.RegisterResponse{
		Id:       registeredUser.Id,
		Name:     registeredUser.Name,
		Username: registeredUser.Username,
		Email:    registeredUser.Email,
		Token:    token,
		Sex:      registeredUser.GetGenderName(),
	}, nil
}

func (as *authenticationService) Login(ctx context.Context, r model.LoginRequest) (*model.LoginResponse, error) {
	tx, err := as.db.Begin()
	if err != nil {
		return nil, err
	}

	err = r.Validate()
	if err != nil {
		return nil, err
	}

	u := entity.User{
		Email:    r.Email,
		Password: r.Password,
	}

	user, err := as.authenticationRepository.Login(ctx, tx, u)
	if err != nil {
		return nil, err
	}

	ok := password.CheckPassword(user.Password, u.Password)
	if !ok {
		return nil, exception.ErrPasswordAreWrong
	}

	token, err := jwt.GenerateToken(*user)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		Id:       user.Id,
		Name:     user.Name,
		Email:    user.Email,
		Username: user.Username,
		Token:    token,
		Sex:      user.GetGenderName(),
	}, nil
}

func NewService(authenticationRepository authentication.AuthenticationRepository, db *sql.DB) AuthenticationService {
	return &authenticationService{
		authenticationRepository: authenticationRepository,
		db:                       db,
	}
}
