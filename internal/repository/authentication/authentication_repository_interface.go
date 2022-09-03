package authentication

import (
	"context"
	"database/sql"

	"github.com/mrizkimaulidan/storial/internal/entity"
)

type AuthenticationRepository interface {
	Register(ctx context.Context, tx *sql.Tx, u entity.User) (*entity.User, error)
	CheckIfEmailExists(ctx context.Context, tx *sql.Tx, e string) (*bool, error)
	CheckIfUsernameExists(ctx context.Context, tx *sql.Tx, u string) (*bool, error)
	Login(ctx context.Context, tx *sql.Tx, u entity.User) (*entity.User, error)
}
