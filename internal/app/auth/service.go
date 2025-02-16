package auth

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/middleware"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/utils"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type service struct {
	authRepo repo.AuthStorer
}

type Service interface {
	Login(ctx context.Context, loginData LoginRequest) (LoginResponse, error, error)
}

func NewService(authRepo repo.AuthStorer) Service {
	return &service{
		authRepo: authRepo,
	}
}

func (authS *service) Login(ctx context.Context, loginData LoginRequest) (LoginResponse, error, error) {

	var resp LoginResponse

	// Check in Workers table
	user, err := authS.authRepo.Login(ctx, repo.LoginRequest(loginData))
	if err != nil {
		return LoginResponse{}, err, apperrors.ErrInvalidLoginCredentials
	}
	if user.Email == "" {
		return LoginResponse{}, err, apperrors.ErrInvalidLoginCredentials
	}

	resp.User = LoginUserData(user)

	// check bcrypt password match
	match := utils.CheckPasswordHash(loginData.Password, user.Password)
	if !match {
		return LoginResponse{}, err, apperrors.ErrInvalidLoginCredentials
	}

	// create jwt
	token, err := middleware.GenerateToken(user.ID, user.Role)
	if err != nil {
		return LoginResponse{}, err, apperrors.ErrCreateToken
	}

	resp.Token = token
	return resp, nil, nil
}
