package auth

import (
	"context"
	"testing"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo/mocks"
	"github.com/stretchr/testify/suite"
)

type AuthServiceTestSuite struct {
	suite.Suite
	authSerivce Service
	authRepo    mocks.AuthStorer
}

func (suite *AuthServiceTestSuite) SetupTest() {
	suite.authSerivce = NewService(&mocks.AuthStorer{})
	suite.authRepo = mocks.AuthStorer{}
}

func (suite *AuthServiceTestSuite) TearDownTest() {
	suite.authRepo.AssertExpectations(suite.T())
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(AuthServiceTestSuite))
}

func (suite *AuthServiceTestSuite) TestLogin() {
	type testCase struct {
		name           string
		input          LoginRequest
		setup          func()
		expectedOutput LoginResponse
		iExpectedError bool
	}

	testCases := []testCase{
		// {
		// 	name: "success",
		// 	input: LoginRequest{
		// 		Email:    "harsh@gmail.com",
		// 		Password: "Harsh@123",
		// 	},
		// 	setup: func() {
		// 		suite.authRepo.On("Login", mock.Anything, repo.LoginRequest{
		// 			Email:    "harsh@gmail.com",
		// 			Password: "Harsh@123",
		// 		}).Return(repo.LoginUserData{
		// 			ID:       1,
		// 			Name:     "Harsh",
		// 			Email:    "harsh@gmail.com",
		// 			Password: "Harsh@123",
		// 			Role:     "worker",
		// 		}, nil)
		// 	},
		// 	expectedOutput: LoginResponse{
		// 		Token: "somerandomtokengeneratedafterlogin",
		// 		User: LoginUserData{
		// 			ID:    1,
		// 			Name:  "Harsh",
		// 			Email: "harsh@gmail.com",
		// 			Role:  "worker",
		// 		},
		// 	},
		// 	iExpectedError: false,
		// },

	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			resp, err := suite.authSerivce.Login(context.Background(), test.input)
			suite.Equal(test.expectedOutput, resp)
			suite.Equal(test.expectedOutput, err != nil)
		})
	}
}
