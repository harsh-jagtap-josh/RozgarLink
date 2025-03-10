package employer

import (
	"context"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
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
		{
			name:       "db error",
			employerId: 1,
			setup: func() {
				suite.employerRepo.On("FetchEmployerByID", mock.Anything, 1).Return(repo.Employer{}, errors.New("db error"))
			},
			expectedOutput: Employer{},
			expectedError:  errors.New("db error"),
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			emp, err := suite.service.FetchEmployerByID(context.Background(), tc.employerId)
			suite.Equal(tc.expectedOutput, emp)
			suite.Equal(tc.expectedError, err)
		})
		suite.TearDownTest()
	}
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
		{
			name: "success",
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
			setup: func() {
				suite.employerRepo.On("UpdateEmployerById", mock.Anything, repo.Employer{
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
					Details:      "location details",
					Street:       "Street",
					City:         "City",
					State:        "State",
					Pincode:      412544,
				}).Return(repo.Employer{
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
			},
			expectedError: nil,
		},
		{
			name: "db error",
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
			setup: func() {
				suite.employerRepo.On("UpdateEmployerById", mock.Anything, repo.Employer{
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
					Details:      "location details",
					Street:       "Street",
					City:         "City",
					State:        "State",
					Pincode:      412544,
				}).Return(repo.Employer{}, errors.New("db error while update employer"))
			},
			expectedOutput: Employer{},
			expectedError:  errors.New("db error while update employer"),
		},
		{
			name: "validation fail in employer mobile",
			employerData: Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "90676913",
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
			setup: func() {
			},
			expectedOutput: Employer{},
			expectedError:  fmt.Errorf("%w: %w", apperrors.ErrInvalidUserDetails, errors.New("invalid mobile number: must be 10 digits and start with 6-9")),
		},
		{
			name: "validation fail in employer email",
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
			setup: func() {
			},
			expectedOutput: Employer{},
			expectedError:  fmt.Errorf("%w: %w", apperrors.ErrInvalidUserDetails, errors.New("invalid email address format")),
		},
		{
			name: "validation fail in employer name",
			employerData: Employer{
				ID:        1,
				Name:      "John Doe1",
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
			setup: func() {
			},
			expectedOutput: Employer{},
			expectedError:  fmt.Errorf("%w: %w", apperrors.ErrInvalidUserDetails, errors.New("invalid name: must be between 3-50 characters and contain only alphabets")),
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			emp, err := suite.service.UpdateEmployerById(context.Background(), tc.employerData)
			suite.Equal(tc.expectedOutput, emp)
			suite.Equal(tc.expectedError, err)
		})
		suite.TearDownTest()
	}
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
		{
			name: "db error",
			setup: func() {
				suite.employerRepo.On("FetchAllEmployers", mock.Anything).Return([]repo.Employer{}, errors.New("db error while list employers"))
			},
			expectedOutput: []Employer{},
			expectedError:  errors.New("db error while list employers"),
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			emp, err := suite.service.FetchAllEmployers(context.Background())
			suite.Equal(tc.expectedOutput, emp)
			suite.Equal(tc.expectedError, err)
		})
		suite.TearDownTest()
	}
}

func (suite *EmployerServiceTestSuite) TestRegisterEmployer() {
	type testCase struct {
		name           string
		setup          func()
		employerData   Employer
		expectedOutput Employer
		expectedError  bool
	}

	testCases := []testCase{
		{
			name: "employer email already exists",
			employerData: Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9067691363",
				Email:     "employer@gmail.com",
				Password:  "Employer@123",
				Type:      "Employer",
				Language:  "English",
				Sectors:   "IT",
				Location: Address{
					ID:      1,
					Details: "location details",
					Street:  "Street",
					City:    "City",
					State:   "State",
					Pincode: 412544,
				},
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			setup: func() {
				suite.employerRepo.On("FindEmployerByEmail", mock.Anything, "employer@gmail.com").Return(true)
			},
			expectedOutput: Employer{},
			expectedError:  true,
		},
		{
			name: "validation error password not provided",
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
			setup:          func() {},
			expectedOutput: Employer{},
			expectedError:  true,
		},
		{
			name: "validation error invalid email address",
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
			setup:          func() {},
			expectedOutput: Employer{},
			expectedError:  true,
		},
		{
			name: "validation error invalid phone number",
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
			setup:          func() {},
			expectedOutput: Employer{},
			expectedError:  true,
		},
		{
			name: "validation error invalid employer name",
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
			setup:          func() {},
			expectedOutput: Employer{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			output, err := suite.service.RegisterEmployer(context.Background(), test.employerData)
			suite.Equal(test.expectedOutput, output)
			suite.Equal(test.expectedError, err != nil)
		})
		suite.TearDownTest()
	}
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
		{
			name:       "success",
			employerId: 1,
			setup: func() {
				suite.employerRepo.On("FindEmployerById", mock.Anything, 1).Return(true)
				suite.employerRepo.On("DeleteEmployerByID", mock.Anything, 1).Return(1, nil)
			},
			expectedOutput: 1,
			expectedError:  nil,
		},
		{
			name:       "db error",
			employerId: 1,
			setup: func() {
				suite.employerRepo.On("FindEmployerById", mock.Anything, 1).Return(true)
				suite.employerRepo.On("DeleteEmployerByID", mock.Anything, 1).Return(-1, errors.New("db error while delete employer"))
			},
			expectedOutput: -1,
			expectedError:  errors.New("db error while delete employer"),
		},
		{
			name:       "employer with id not found",
			employerId: 1,
			setup: func() {
				suite.employerRepo.On("FindEmployerById", mock.Anything, 1).Return(false)
			},
			expectedOutput: -1,
			expectedError:  apperrors.ErrNoEmployerExists,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			output, err := suite.service.DeleteEmployerById(context.Background(), test.employerId)
			suite.Equal(test.expectedOutput, output)
			suite.Equal(test.expectedError, err)
		})
		suite.TearDownTest()
	}
}

func (suite *EmployerServiceTestSuite) TestFetchJobsByEmployerId() {
	type testCase struct {
		name           string
		employerId     int
		setup          func()
		expectedOutput []job.Job
		expectedError  bool
	}
	testCases := []testCase{
		{
			name:       "success",
			employerId: 1,
			setup: func() {
				suite.employerRepo.On("FindEmployerById", mock.Anything, 1).Return(true)
				suite.employerRepo.On("FindJobByEmployerId", mock.Anything, 1).Return([]repo.Job{
					{
						ID:              1,
						EmployerID:      1,
						Title:           "title",
						RequiredGender:  "Male",
						Description:     "some random description",
						DurationInHours: 10,
						SkillsRequired:  "some skills",
						Sectors:         "random skills",
						Wage:            1200,
						Vacancy:         5,
						Location:        1,
						Date:            "2025-12-03",
						StartHour:       "",
						EndHour:         "",
						CreatedAt:       time.Time{},
						UpdatedAt:       time.Time{},
						Details:         "details",
						Street:          "street",
						City:            "city",
						State:           "state",
						Pincode:         41205,
					},
				}, nil)
			},
			expectedOutput: []job.Job{
				{
					ID:              1,
					EmployerID:      1,
					Title:           "title",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 10,
					SkillsRequired:  "some skills",
					Sectors:         "random skills",
					Wage:            1200,
					Vacancy:         5,
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 41205,
					},
					Date:      "2025-12-03",
					StartHour: "",
					EndHour:   "",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
			},
			expectedError: false,
		},
		{
			name:       "db error",
			employerId: 1,
			setup: func() {
				suite.employerRepo.On("FindEmployerById", mock.Anything, 1).Return(true)
				suite.employerRepo.On("FindJobByEmployerId", mock.Anything, 1).Return([]repo.Job{}, errors.New("db error while fetch jobs by employer id"))
			},
			expectedOutput: []job.Job{},
			expectedError:  true,
		},
		{
			name:       "employer with id not found",
			employerId: 1,
			setup: func() {
				suite.employerRepo.On("FindEmployerById", mock.Anything, 1).Return(false)
			},
			expectedOutput: []job.Job{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			jobs, err := suite.service.FetchJobsByEmployerId(context.Background(), test.employerId)
			suite.Equal(test.expectedOutput, jobs)
			suite.Equal(test.expectedError, err != nil)
		})
		suite.TearDownTest()
	}
}
