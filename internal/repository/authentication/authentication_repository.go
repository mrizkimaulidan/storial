package authentication

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
	exception "github.com/mrizkimaulidan/storial/pkg/exception/authentication"
)

type authenticationRepository struct {
	//
}

// Checking email if already exists on database.
// Returning boolean true if exists, false if not exists.
func (ar *authenticationRepository) CheckIfEmailExists(ctx context.Context, tx *sql.Tx, e string) (*bool, error) {
	query := `
		SELECT
		EXISTS(
		SELECT
			email
		FROM
			users
		WHERE
			email = ?
	)
	`

	row := tx.QueryRowContext(ctx, query, e)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return nil, err
	}

	return &exists, nil
}

// Checking username if already exists on database.
// Returning boolean true if exists, false if not exists.
func (ar *authenticationRepository) CheckIfUsernameExists(ctx context.Context, tx *sql.Tx, u string) (*bool, error) {
	query := `
		SELECT
		EXISTS(
		SELECT
			username
		FROM
			users
		WHERE
			username = ?
	)
	`

	row := tx.QueryRowContext(ctx, query, u)

	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		return nil, err
	}

	return &exists, nil
}

// Function that handle login. If no rows returned, there must be no email
// registered with provided email.
func (ar *authenticationRepository) Login(ctx context.Context, tx *sql.Tx, s entity.User) (*entity.User, error) {
	query := `
		SELECT
		id,
		name,
		username,
		email,
		password,
		sex
	FROM
		users
	WHERE
		email = ?
	`

	row := tx.QueryRowContext(ctx, query, s.Email)

	var user entity.User
	err := row.Scan(&user.Id, &user.Name, &user.Username, &user.Email, &user.Password, &user.Sex)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, exception.ErrEmailNotFound
		}
		return nil, err
	}

	return &user, nil
}

// Function that handle register.
func (ar *authenticationRepository) Register(ctx context.Context, tx *sql.Tx, s entity.User) (*entity.User, error) {
	query := `
		INSERT INTO users(
			id,
			name,
			username,
			email,
			password,
			sex,
			created_at
		)
		VALUES(?, ?, ?, ?, ?, ?, ?)
	`
	_, err := tx.ExecContext(ctx, query, s.Id, s.Name, s.Username, s.Email, s.Password, s.Sex, s.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &s, nil
}

func NewRepository() AuthenticationRepository {
	return &authenticationRepository{}
}
