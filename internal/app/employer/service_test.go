package employer

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EmployerServiceTestSuite struct {
	suite.Suite
	service      Service
	employerRepo mocks.EmployerStorer
}

func (suite *EmployerServiceTestSuite) SetupTest() {
	suite.employerRepo = mocks.EmployerStorer{}
	suite.service = NewService(&suite.employerRepo)
}

func (suite *EmployerServiceTestSuite) TearDownTest() {
	suite.employerRepo.AssertExpectations(suite.T())
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(EmployerServiceTestSuite))
}

func (suite *EmployerServiceTestSuite) TestFetchEmployerByID() {
	type testCase struct {
		name           string
		setup          func()
		employerId     int
		expectedOutput Employer
		expectedError  error
	}

	testCases := []testCase{
		{
			name:       "success",
			employerId: 1,
			setup: func() {
				suite.employerRepo.On("FetchEmployerByID", mock.Anything, 1).Return(repo.Employer{
					ID:           1,
					Name:         "John Doe",
					ContactNo:    "9067691363",
					Email:        "employer@gmail.com",
					Type:         "Employer",
					Sectors:      "IT",
					Location:     1,
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
					Details:      "location details",
					Street:       "Street",
					City:         "City",
					State:        "State",
					Pincode:      412544,
				}, nil)
			},
			expectedOutput: Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9067691363",
				Email:     "employer@gmail.com",
				Type:      "Employer",
				Sectors:   "IT",
				Location: Address{
					ID:      1,
					Details: "location details",
					Street:  "Street",
					City:    "City",
					State:   "State",
					Pincode: 412544,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			expectedError: nil,
		},
	}

	for _, test := range testCases {
		// fmt.Println(test)
		suite.Run(test.name, func() {
			test.setup()
			output, err := suite.service.FetchEmployerByID(context.Background(), test.employerId)
			suite.Equal(test.expectedOutput, output)
			fmt.Println(output)
			suite.Equal(test.expectedError, err)
		})
	}
	suite.TearDownTest()
}

func (suite *EmployerServiceTestSuite) TestUpdateEmployerById() {
	type testCase struct {
		name           string
		setup          func()
		employerData   Employer
		expectedOutput Employer
		expectedError  error
	}

	testCases := []testCase{
		// {
		// 	name: "success",
		// 	employerData: Employer{
		// 		ID:        1,
		// 		Name:      "John Doe",
		// 		ContactNo: "9067691363",
		// 		Email:     "employer@gmail.com",
		// 		Type:      "Employer",
		// 		Sectors:   "IT",
		// 		Location: Address{
		// 			ID:      1,
		// 			Details: "location details",
		// 			Street:  "Street",
		// 			City:    "City",
		// 			State:   "State",
		// 			Pincode: 412544,
		// 		},
		// 		IsVerified:   true,
		// 		Rating:       0,
		// 		WorkersHired: 0,
		// 		CreatedAt:    time.Time{},
		// 		UpdatedAt:    time.Time{},
		// 	},
		// 	setup: func() {
		// 		suite.employerRepo.On("UpdateEmployerByID", mock.Anything, repo.Employer{
		// 			ID:           1,
		// 			Name:         "John Doe",
		// 			ContactNo:    "9067691363",
		// 			Email:        "employer@gmail.com",
		// 			Type:         "Employer",
		// 			Sectors:      "IT",
		// 			Location:     1,
		// 			IsVerified:   true,
		// 			Rating:       0,
		// 			WorkersHired: 0,
		// 			CreatedAt:    time.Time{},
		// 			UpdatedAt:    time.Time{},
		// 			Details:      "location details",
		// 			Street:       "Street",
		// 			City:         "City",
		// 			State:        "State",
		// 			Pincode:      412544,
		// 		}).Return(repo.Employer{
		// 			ID:           1,
		// 			Name:         "John Doe",
		// 			ContactNo:    "9067691363",
		// 			Email:        "employer@gmail.com",
		// 			Type:         "Employer",
		// 			Sectors:      "IT",
		// 			Location:     1,
		// 			IsVerified:   true,
		// 			Rating:       0,
		// 			WorkersHired: 0,
		// 			CreatedAt:    time.Time{},
		// 			UpdatedAt:    time.Time{},
		// 			Details:      "location details",
		// 			Street:       "Street",
		// 			City:         "City",
		// 			State:        "State",
		// 			Pincode:      412544,
		// 		}, nil)
		// 	},
		// 	expectedOutput: Employer{
		// 		ID:        1,
		// 		Name:      "John Doe",
		// 		ContactNo: "9067691363",
		// 		Email:     "employer@gmail.com",
		// 		Type:      "Employer",
		// 		Sectors:   "IT",
		// 		Location: Address{
		// 			ID:      1,
		// 			Details: "location details",
		// 			Street:  "Street",
		// 			City:    "City",
		// 			State:   "State",
		// 			Pincode: 412544,
		// 		},
		// 		IsVerified:   true,
		// 		Rating:       0,
		// 		WorkersHired: 0,
		// 		CreatedAt:    time.Time{},
		// 		UpdatedAt:    time.Time{},
		// 	},
		// 	expectedError: nil,
		// },
	}

	for _, test := range testCases {
		suite.Run(test.name, func() {
			test.setup()
			output, err := suite.service.UpdateEmployerById(context.Background(), test.employerData)
			suite.Equal(test.expectedOutput, output)
			suite.Equal(test.expectedError, err)
		})
	}
	suite.TearDownTest()
}

func (suite *EmployerServiceTestSuite) TestFetchAllEmployers() {
	type testCase struct {
		name           string
		setup          func()
		expectedOutput []Employer
		expectedError  error
	}

	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.employerRepo.On("FetchAllEmployers", mock.Anything).Return([]repo.Employer{
					{
						ID:           1,
						Name:         "John Doe",
						ContactNo:    "9067691363",
						Email:        "employer@gmail.com",
						Type:         "Employer",
						Sectors:      "IT",
						Location:     1,
						IsVerified:   true,
						Rating:       0,
						WorkersHired: 0,
						CreatedAt:    time.Time{},
						UpdatedAt:    time.Time{},
						Details:      "",
						Street:       "",
						City:         "",
						State:        "",
						Pincode:      0,
					},
				}, nil)
			},
			expectedOutput: []Employer{
				{
					ID:        1,
					Name:      "John Doe",
					ContactNo: "9067691363",
					Email:     "employer@gmail.com",
					Type:      "Employer",
					Sectors:   "IT",
					Location: Address{
						ID:      1,
						Details: "",
						Street:  "",
						City:    "",
						State:   "",
						Pincode: 0,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
				},
			},
			expectedError: nil,
		},
	}

	for _, test := range testCases {
		suite.Run(test.name, func() {
			test.setup()
			output, err := suite.service.FetchAllEmployers(context.Background())
			suite.Equal(test.expectedOutput, output)
			suite.Equal(test.expectedError, err)
		})
	}
	suite.TearDownTest()
}

func (suite *EmployerServiceTestSuite) TestCreateEmployer() {
	type testCase struct {
		name           string
		setup          func()
		repoRequired   bool
		employerData   Employer
		expectedOutput Employer
		expectedError  bool
	}

	testCases := []testCase{
		{
			name:         "password not provided",
			repoRequired: false,
			employerData: Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9067691363",
				Email:     "employer@gmail.com",
				Type:      "Employer",
				Sectors:   "IT",
				Location: Address{
					ID:      1,
					Details: "location details",
					Street:  "Street",
					City:    "City",
					State:   "State",
					Pincode: 412544,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
			},
			expectedOutput: Employer{},
			expectedError:  true,
		},
		{
			name:         "invalid email address",
			repoRequired: false,
			employerData: Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9067691363",
				Email:     "employer",
				Type:      "Employer",
				Sectors:   "IT",
				Location: Address{
					ID:      1,
					Details: "location details",
					Street:  "Street",
					City:    "City",
					State:   "State",
					Pincode: 412544,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
			},
			expectedOutput: Employer{},
			expectedError:  true,
		},
		{
			name:         "invalid phone number",
			repoRequired: false,
			employerData: Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "906769136",
				Email:     "employer",
				Type:      "Employer",
				Sectors:   "IT",
				Location: Address{
					ID:      1,
					Details: "location details",
					Street:  "Street",
					City:    "City",
					State:   "State",
					Pincode: 412544,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
			},
			expectedOutput: Employer{},
			expectedError:  true,
		},
		{
			name:         "invalid employer name",
			repoRequired: false,
			employerData: Employer{
				ID:        1,
				Name:      "John Doe1",
				ContactNo: "906769136",
				Email:     "employer",
				Type:      "Employer",
				Sectors:   "IT",
				Location: Address{
					ID:      1,
					Details: "location details",
					Street:  "Street",
					City:    "City",
					State:   "State",
					Pincode: 412544,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
			},
			expectedOutput: Employer{},
			expectedError:  true,
		},
		// {
		// 	name:         "success",
		// 	repoRequired: true,
		// 	employerData: Employer{
		// 		ID:        1,
		// 		Name:      "John Doe",
		// 		ContactNo: "9067691363",
		// 		Email:     "employer@gmail.com",
		// 		Password:  "Employer@123",
		// 		Type:      "Employer",
		// 		Language:  "English",
		// 		Sectors:   "IT",
		// 		Location: Address{
		// 			ID:      1,
		// 			Details: "location details",
		// 			Street:  "Street",
		// 			City:    "City",
		// 			State:   "State",
		// 			Pincode: 412544,
		// 		},
		// 		CreatedAt: time.Time{},
		// 		UpdatedAt: time.Time{},
		// 	},
		// 	setup: func() {
		// 		suite.employerRepo.On("CreateEmployer", mock.Anything, repo.Employer{
		// 			ID:        1,
		// 			Name:      "John Doe",
		// 			ContactNo: "9067691363",
		// 			Email:     "employer@gmail.com",
		// 			Type:      "Employer",
		// 			Sectors:   "IT",
		// 			Location:  1,
		// 			CreatedAt: time.Time{},
		// 			UpdatedAt: time.Time{},
		// 			Details:   "location details",
		// 			Street:    "Street",
		// 			City:      "City",
		// 			State:     "State",
		// 			Pincode:   412544,
		// 			Language:  "English",
		// 		}).Return(repo.Employer{
		// 			ID:           1,
		// 			Name:         "John Doe",
		// 			ContactNo:    "9067691363",
		// 			Email:        "employer@gmail.com",
		// 			Type:         "Employer",
		// 			Sectors:      "IT",
		// 			Location:     1,
		// 			IsVerified:   true,
		// 			Rating:       0,
		// 			WorkersHired: 0,
		// 			CreatedAt:    time.Time{},
		// 			UpdatedAt:    time.Time{},
		// 			Details:      "location details",
		// 			Street:       "Street",
		// 			City:         "City",
		// 			State:        "State",
		// 			Pincode:      412544,
		// 			Language:     "English",
		// 		}, nil)
		// 	},
		// 	expectedOutput: Employer{
		// 		ID:        1,
		// 		Name:      "John Doe",
		// 		ContactNo: "9067691363",
		// 		Email:     "employer@gmail.com",
		// 		Type:      "Employer",
		// 		Sectors:   "IT",
		// 		Location: Address{
		// 			ID:      1,
		// 			Details: "location details",
		// 			Street:  "Street",
		// 			City:    "City",
		// 			State:   "State",
		// 			Pincode: 412544,
		// 		},
		// 		CreatedAt:    time.Time{},
		// 		UpdatedAt:    time.Time{},
		// 		Language:     "English",
		// 		IsVerified:   true,
		// 		Rating:       0,
		// 		WorkersHired: 0,
		// 	},
		// 	expectedError: false,
		// },
	}

	for _, test := range testCases {
		suite.Run(test.name, func() {
			if test.repoRequired {
				test.setup()
			}
			output, err := suite.service.RegisterEmployer(context.Background(), test.employerData)
			suite.Equal(test.expectedOutput, output)
			suite.Equal(test.expectedError, err != nil)
		})
	}
	suite.TearDownTest()
}

func (suite *EmployerServiceTestSuite) TestDeleteEmployerByID() {
	type testCase struct {
		name           string
		setup          func()
		employerId     int
		expectedOutput int
		expectedError  error
	}

	testCases := []testCase{
		// {
		// 	name:       "success",
		// 	employerId: 1,
		// 	setup: func() {
		// 		suite.employerRepo.On("DeleteEmployerByID", mock.Anything, 1).Return(1, nil)
		// 	},
		// 	expectedOutput: 1,
		// 	expectedError:  nil,
		// },
	}

	for _, test := range testCases {
		suite.Run(test.name, func() {
			test.setup()
			output, err := suite.service.DeleteEmployerById(context.Background(), test.employerId)
			suite.Equal(test.expectedOutput, output)
			suite.Equal(test.expectedError, err)
		})
	}
	suite.TearDownTest()
}