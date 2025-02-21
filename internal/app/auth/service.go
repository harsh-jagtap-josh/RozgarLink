package auth

import (
	"context"
	"fmt"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/middleware"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/utils"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	authRepo repo.AuthStorer
}

type Service interface {
	Login(ctx context.Context, loginData LoginRequest) (LoginResponse, error)
}

func NewService(authRepo repo.AuthStorer) Service {
	return &service{
		authRepo: authRepo,
	}
}

func (authS *service) Login(ctx context.Context, loginData LoginRequest) (LoginResponse, error) {

	var resp LoginResponse

	user, err := authS.authRepo.Login(ctx, repo.LoginRequest(loginData))
	if err != nil {
		return LoginResponse{}, fmt.Errorf("%w: %w", apperrors.ErrInvalidLoginCredentials, err)
	}

	resp.User = LoginUserData(user)

	// check bcrypt password match
	match := utils.CheckPasswordHash(loginData.Password, user.Password)
	if !match {
		return LoginResponse{}, fmt.Errorf("%w: %w", apperrors.ErrFailedLogin, apperrors.ErrIncorrectLoginData)
	}

	// create jwt
	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("%w: %w", apperrors.ErrCreateToken, err)
	}

	resp.Token = token
	return resp, nil
}
