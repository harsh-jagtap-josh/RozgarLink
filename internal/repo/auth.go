package repo

import (
	"context"
	"database/sql"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/jmoiron/sqlx"
)

type authStore struct {
	BaseRepository
}

type AuthStorer interface {
	Login(ctx context.Context, loginData LoginRequest) (LoginUserData, error)
}

func NewAuthRepo(db *sql.DB) AuthStorer {
	return &authStore{
		BaseRepository: BaseRepository{db},
	}
}

const FetchWorkerByEmailQuery = "SELECT id, name, email, password FROM workers WHERE email=:email;"
const FetchEmployerByEmailQuery = "SELECT id, name, email, password FROM employers WHERE email=:email;"

func (authS *authStore) Login(ctx context.Context, loginData LoginRequest) (LoginUserData, error) {
	var user LoginUserData
	db := sqlx.NewDb(authS.DB, "postgres")

	rows, err := db.NamedQuery(FetchWorkerByEmailQuery, loginData)
	if err != nil {
		return LoginUserData{}, err
	} else {
		defer rows.Close()
		if rows.Next() {
			err = rows.StructScan(&user)
			if err != nil {
				return LoginUserData{}, err
			} else {
				user.Role = "worker"
			}
		}
	}
	if len(user.Email) == 0 {
		rows, err = db.NamedQuery(FetchEmployerByEmailQuery, loginData)
		if err != nil {
			return LoginUserData{}, nil
		} else {
			defer rows.Close()
			if rows.Next() {
				err = rows.StructScan(&user)
				if err != nil {
					return LoginUserData{}, apperrors.ErrInvalidLoginCredentials
				} else {
					user.Role = "employer"
				}
			}
		}
	}

	return user, nil
}
