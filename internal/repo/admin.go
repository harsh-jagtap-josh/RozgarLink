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
	DeleteAdmin(ctx context.Context, adminId int) error
	FindAdminByEmail(ctx context.Context, email string) bool
	FindAdminById(ctx context.Context, adminId int) bool
}

func NewAdminRepo(db *sqlx.DB) AdminStorer {
	return &adminRepo{
		DB: db,
	}
}

// PostgreSQL Queries
const (
	registerAdminQuery    = `INSERT INTO admins (name, contact_no, email, password, created_at, updated_at) VALUES (:name, :contact_no, :email, :password, NOW(), NOW()) RETURNING *;`
	deleteAdminQuery      = `DELETE FROM admins WHERE id=$1 RETURNING id;`
	findAdminByEmailQuery = `SELECT id from admins where email = $1;`
	findAdminByIdQuery    = `SELECT id from admins where id = $1;`
)

func (admR *adminRepo) RegisterAdmin(ctx context.Context, adminData Admin) (Admin, error) {

	var createdAdmin Admin

	rows, err := admR.DB.NamedQuery(registerAdminQuery, adminData)
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

func (adminR *adminRepo) DeleteAdmin(ctx context.Context, adminId int) error {
	var ID int
	err := adminR.DB.Get(&ID, deleteAdminQuery, adminId)

	if err != nil {
		return err
	}
	return nil
}

func (adminR *adminRepo) FindAdminById(ctx context.Context, adminId int) bool {
	var ID int
	err := adminR.DB.QueryRow(findAdminByIdQuery, ID).Scan(&ID)

	return err == nil
}

func (admR *adminRepo) FindAdminByEmail(ctx context.Context, email string) bool {
	var ID int

	err := admR.DB.QueryRow(findAdminByEmailQuery, email).Scan(&ID)
	return err == nil
}
