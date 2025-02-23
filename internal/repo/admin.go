package repo

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type adminRepo struct {
	DB *sqlx.DB
}

type AdminStorer interface {
	RegisterAdmin(ctx context.Context, adminData Admin) (Admin, error)
	FindAdminByEmail(ctx context.Context, email string) bool
}

func NewAdminRepo(db *sqlx.DB) AdminStorer {
	return &adminRepo{
		DB: db,
	}
}

func (admR *adminRepo) RegisterAdmin(ctx context.Context, adminData Admin) (Admin, error) {
	query := `INSERT INTO admins (name, contact_no, email, password, created_at, updated_at) VALUES (:name, :contact_no, :email, :password, NOW(), NOW()) RETURNING *;`

	var createdAdmin Admin

	rows, err := admR.DB.NamedQuery(query, adminData)
	if err != nil {
		return Admin{}, err
	}

	defer rows.Close()

	if rows.Next() {
		err = rows.StructScan(&createdAdmin)
		if err != nil {
			return Admin{}, err
		}
	}

	return createdAdmin, nil
}

func (admR *adminRepo) FindAdminByEmail(ctx context.Context, email string) bool {
	var ID int

	query := `SELECT id from admins where email = $1`
	err := admR.DB.QueryRow(query, email).Scan(&ID)
	return err == nil
}
