package admin

import (
	"context"
	"fmt"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/utils"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	adminRepo repo.AdminStorer
}

type AdminService interface {
	RegisterAdmin(ctx context.Context, adminData Admin) (Admin, error)
	DeleteAdmin(ctx context.Context, adminId int) error
}

func NewAdminService(adminRepo repo.AdminStorer) AdminService {
	return &service{
		adminRepo: adminRepo,
	}
}

func (adminS *service) RegisterAdmin(ctx context.Context, adminData Admin) (Admin, error) {

	err := utils.ValidateUser(adminData.Name, adminData.ContactNo, adminData.Email, adminData.Password)
	if err != nil {
		return Admin{}, fmt.Errorf("%w, %w", apperrors.ErrInvalidUserDetails, err)
	}

	alreadyExists := adminS.adminRepo.FindAdminByEmail(ctx, adminData.Email)
	if alreadyExists {
		return Admin{}, apperrors.ErrAdminExists
	}

	hashed_password, err := utils.HashPassword(adminData.Password)
	if err != nil {
		return Admin{}, fmt.Errorf("%w: %w", apperrors.ErrEncrPassword, err)
	}
	adminData.Password = hashed_password

	createdAdmin, err := adminS.adminRepo.RegisterAdmin(ctx, repo.Admin(adminData))
	if err != nil {
		return Admin{}, err
	}

	return Admin(createdAdmin), nil
}

func (adminS *service) DeleteAdmin(ctx context.Context, adminId int) error {
	exists := adminS.adminRepo.FindAdminById(ctx, adminId)
	if !exists {
		return apperrors.ErrNoAdminExists
	}
	err := adminS.adminRepo.DeleteAdmin(ctx, adminId)
	if err != nil {
		return err
	}
	return nil
}
