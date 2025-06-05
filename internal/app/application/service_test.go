package application

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ApplicationServiceTestSuite struct {
	suite.Suite
	service         Service
	applicationRepo mocks.ApplicationStorer
}

func (suite *ApplicationServiceTestSuite) SetupTest() {
	suite.applicationRepo = mocks.ApplicationStorer{}
	suite.service = NewService(&suite.applicationRepo)
}

func (suite *ApplicationServiceTestSuite) TearDownTest() {
	suite.applicationRepo.AssertExpectations(suite.T())
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationServiceTestSuite))
}

func (suite *ApplicationServiceTestSuite) TestFetchApplicationById() {
	type testCase struct {
		name           string
		setup          func()
		applicationId  int
		expectedOutput Application
		expectedError  bool
	}

	testCases := []testCase{
		{
			name:          "success",
			applicationId: 1,
			setup: func() {
				suite.applicationRepo.On("FetchApplicationByID", mock.Anything, 1).Return(repo.Application{
					ID:             1,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
				}, nil)
			},
			expectedOutput: Application{
				ID:            1,
				JobID:         3,
				WorkerID:      12,
				Status:        "Pending",
				ExpectedWage:  1200,
				ModeOfArrival: "Pick-Up",
				PickUpLocation: Address{
					ID:      5,
					Details: "location details",
					Street:  "location street",
					City:    "location city",
					State:   "location state",
					Pincode: 411025,
				},
				WorkerComment: "some random comments by worker",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			expectedError: false,
		},
		{
			name:          "application not found",
			applicationId: 1,
			setup: func() {
				suite.applicationRepo.On("FetchApplicationByID", mock.Anything, 1).Return(repo.Application{}, apperrors.ErrNoApplicationExists)
			},
			expectedOutput: Application{},
			expectedError:  true,
		},
		{
			name:          "application not found",
			applicationId: 1,
			setup: func() {
				suite.applicationRepo.On("FetchApplicationByID", mock.Anything, 1).Return(repo.Application{}, sql.ErrNoRows)
			},
			expectedOutput: Application{},
			expectedError:  true,
		},
		{
			name:          "internal error",
			applicationId: 1,
			setup: func() {
				suite.applicationRepo.On("FetchApplicationByID", mock.Anything, 1).Return(repo.Application{}, apperrors.ErrInternalServerError)
			},
			expectedOutput: Application{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			application, err := suite.service.FetchApplicationById(context.Background(), test.applicationId)
			suite.Equal(test.expectedOutput, application)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *ApplicationServiceTestSuite) TestFetchAllApplications() {
	type testCase struct {
		name            string
		setup           func()
		expectedOutput  []ApplicationComplete
		isExpectedError bool
	}

	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.applicationRepo.On("FetchAllApplications", mock.Anything).Return([]repo.ApplicationComplete{
					{
						ID:             1,
						JobID:          3,
						WorkerID:       12,
						Status:         "Pending",
						ExpectedWage:   1200,
						ModeOfArrival:  "Pick-Up",
						PickUpLocation: 5,
						WorkerComment:  "some random comments by worker",
						AppliedAt:      time.Time{},
						UpdatedAt:      time.Time{},
						Details:        "location details",
						Street:         "location street",
						City:           "location city",
						State:          "location state",
						Pincode:        411025,
						JobTitle:       "Full Stack Developer",
						Description:    "some random description",
						SkillsRequired: "Frontend, Backend",
						JobSectors:     "IT, Technology, Computer",
						JobWage:        1022,
						Vacancy:        5,
						JobDate:        "2025-02-12",
						EmployerName:   "Employer XYZ",
						ContactNumber:  "9067691363",
						EmployerEmail:  "employer@gmail.com",
						EmployerType:   "Organization",
					},
					{
						ID:             2,
						JobID:          3,
						WorkerID:       12,
						Status:         "Pending",
						ExpectedWage:   1200,
						ModeOfArrival:  "Pick-Up",
						PickUpLocation: 5,
						WorkerComment:  "some random comments by worker",
						AppliedAt:      time.Time{},
						UpdatedAt:      time.Time{},
						Details:        "location details",
						Street:         "location street",
						City:           "location city",
						State:          "location state",
						Pincode:        411025,
						JobTitle:       "Full Stack Developer",
						Description:    "some random description",
						SkillsRequired: "Frontend, Backend",
						JobSectors:     "IT, Technology, Computer",
						JobWage:        1022,
						Vacancy:        5,
						JobDate:        "2025-02-12",
						EmployerName:   "Employer XYZ",
						ContactNumber:  "9067691363",
						EmployerEmail:  "employer@gmail.com",
						EmployerType:   "Organization",
					},
				}, nil)
			},
			expectedOutput: []ApplicationComplete{
				{
					ID:             1,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
					JobTitle:       "Full Stack Developer",
					Description:    "some random description",
					SkillsRequired: "Frontend, Backend",
					JobSectors:     "IT, Technology, Computer",
					JobWage:        1022,
					Vacancy:        5,
					JobDate:        "2025-02-12",
					EmployerName:   "Employer XYZ",
					ContactNumber:  "9067691363",
					EmployerEmail:  "employer@gmail.com",
					EmployerType:   "Organization",
				},
				{
					ID:             2,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
					JobTitle:       "Full Stack Developer",
					Description:    "some random description",
					SkillsRequired: "Frontend, Backend",
					JobSectors:     "IT, Technology, Computer",
					JobWage:        1022,
					Vacancy:        5,
					JobDate:        "2025-02-12",
					EmployerName:   "Employer XYZ",
					ContactNumber:  "9067691363",
					EmployerEmail:  "employer@gmail.com",
					EmployerType:   "Organization",
				},
			},
			isExpectedError: false,
		},
		{
			name: "failed",
			setup: func() {
				suite.applicationRepo.On("FetchAllApplications", mock.Anything).Return([]repo.ApplicationComplete{}, apperrors.ErrInternalServerError)
			},
			expectedOutput:  []ApplicationComplete{},
			isExpectedError: true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			applications, err := suite.service.FetchAllApplications(context.Background())
			suite.Equal(test.expectedOutput, applications)
			suite.Equal(test.isExpectedError, err != nil)
		})
	}

}

func (suite *ApplicationServiceTestSuite) TestCreateNewApplication() {
	type testCase struct {
		name            string
		input           Application
		setup           func()
		expectedOutput  Application
		isExpectedError bool
	}

	testCases := []testCase{
		{
			name: "create success",
			input: Application{
				ID:            1,
				JobID:         3,
				WorkerID:      12,
				Status:        "Pending",
				ExpectedWage:  1200,
				ModeOfArrival: "Pick-Up",
				PickUpLocation: Address{
					ID:      5,
					Details: "location details",
					Street:  "location street",
					City:    "location city",
					State:   "location state",
					Pincode: 411025,
				},
				WorkerComment: "some random comments by worker",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			setup: func() {
				suite.applicationRepo.On("CreateNewApplication", mock.Anything, repo.Application{
					ID:             1,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
				}).Return(repo.Application{
					ID:             1,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
				}, nil)
			},
			expectedOutput: Application{
				ID:            1,
				JobID:         3,
				WorkerID:      12,
				Status:        "Pending",
				ExpectedWage:  1200,
				ModeOfArrival: "Pick-Up",
				PickUpLocation: Address{
					ID:      5,
					Details: "location details",
					Street:  "location street",
					City:    "location city",
					State:   "location state",
					Pincode: 411025,
				},
				WorkerComment: "some random comments by worker",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			isExpectedError: false,
		},
		{
			name: "create fail",
			input: Application{
				ID:            1,
				JobID:         3,
				WorkerID:      12,
				Status:        "Pending",
				ExpectedWage:  1200,
				ModeOfArrival: "Pick-Up",
				PickUpLocation: Address{
					ID:      5,
					Details: "location details",
					Street:  "location street",
					City:    "location city",
					State:   "location state",
					Pincode: 411025,
				},
				WorkerComment: "some random comments by worker",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			setup: func() {
				suite.applicationRepo.On("CreateNewApplication", mock.Anything, repo.Application{
					ID:             1,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
				}).Return(repo.Application{}, apperrors.ErrCreateApplication)
			},
			expectedOutput:  Application{},
			isExpectedError: true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			application, err := suite.service.CreateNewApplication(context.Background(), test.input)
			suite.Equal(test.expectedOutput, application)
			suite.Equal(test.isExpectedError, err != nil)
		})
	}
}

func (suite *ApplicationServiceTestSuite) TestUpdateApplicationById() {
	type testCase struct {
		name            string
		input           Application
		setup           func()
		expectedOutput  Application
		isExpectedError bool
	}

	testCases := []testCase{
		{
			name: "update success",
			input: Application{
				ID:            1,
				JobID:         3,
				WorkerID:      12,
				Status:        "Pending",
				ExpectedWage:  1200,
				ModeOfArrival: "Pick-Up",
				PickUpLocation: Address{
					ID:      5,
					Details: "location details",
					Street:  "location street",
					City:    "location city",
					State:   "location state",
					Pincode: 411025,
				},
				WorkerComment: "some random comments by worker",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			setup: func() {
				suite.applicationRepo.On("UpdateApplicationByID", mock.Anything, repo.Application{
					ID:             1,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
				}).Return(repo.Application{
					ID:             1,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
				}, nil)
			},
			expectedOutput: Application{
				ID:            1,
				JobID:         3,
				WorkerID:      12,
				Status:        "Pending",
				ExpectedWage:  1200,
				ModeOfArrival: "Pick-Up",
				PickUpLocation: Address{
					ID:      5,
					Details: "location details",
					Street:  "location street",
					City:    "location city",
					State:   "location state",
					Pincode: 411025,
				},
				WorkerComment: "some random comments by worker",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			isExpectedError: false,
		},
		{
			name: "update fail",
			input: Application{
				ID:            1,
				JobID:         3,
				WorkerID:      12,
				Status:        "Pending",
				ExpectedWage:  1200,
				ModeOfArrival: "Pick-Up",
				PickUpLocation: Address{
					ID:      5,
					Details: "location details",
					Street:  "location street",
					City:    "location city",
					State:   "location state",
					Pincode: 411025,
				},
				WorkerComment: "some random comments by worker",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			setup: func() {
				suite.applicationRepo.On("UpdateApplicationByID", mock.Anything, repo.Application{
					ID:             1,
					JobID:          3,
					WorkerID:       12,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Pick-Up",
					PickUpLocation: 5,
					Details:        "location details",
					Street:         "location street",
					City:           "location city",
					State:          "location state",
					Pincode:        411025,
					WorkerComment:  "some random comments by worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
				}).Return(repo.Application{}, apperrors.ErrCreateApplication)
			},
			expectedOutput:  Application{},
			isExpectedError: true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			application, err := suite.service.UpdateApplicationById(context.Background(), test.input)
			suite.Equal(test.expectedOutput, application)
			suite.Equal(test.isExpectedError, err != nil)
		})
	}
}

func (suite *ApplicationServiceTestSuite) TestDeleteApplicationById() {
	type testCase struct {
		name            string
		application_id  int
		setup           func()
		expectedOutput  int
		isExpectedError bool
	}

	testCases := []testCase{
		{
			name:           "success",
			application_id: 1,
			setup: func() {
				suite.applicationRepo.On("FindApplicationById", mock.Anything, 1).Return(true)
				suite.applicationRepo.On("DeleteApplicationByID", mock.Anything, 1).Return(1, nil)
			},
			expectedOutput:  1,
			isExpectedError: false,
		},
		{
			name:           "db error",
			application_id: 1,
			setup: func() {
				suite.applicationRepo.On("FindApplicationById", mock.Anything, 1).Return(true)
				suite.applicationRepo.On("DeleteApplicationByID", mock.Anything, 1).Return(-1, errors.New("db error while delete application"))
			},
			expectedOutput:  -1,
			isExpectedError: true,
		},
		{
			name:           "application with id not found",
			application_id: 1,
			setup: func() {
				suite.applicationRepo.On("FindApplicationById", mock.Anything, 1).Return(false)
			},
			expectedOutput:  -1,
			isExpectedError: true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			id, err := suite.service.DeleteApplicationById(context.Background(), test.application_id)
			suite.Equal(test.expectedOutput, id)
			suite.Equal(test.isExpectedError, err != nil)
		})
	}
}
