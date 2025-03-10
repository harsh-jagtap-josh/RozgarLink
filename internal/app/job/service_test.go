package job

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/application"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"

	// "github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JobServiceTestSuite struct {
	suite.Suite
	service Service
	jobRepo mocks.JobStorer
}

func (suite *JobServiceTestSuite) SetupTest() {
	suite.jobRepo = mocks.JobStorer{}
	suite.service = NewService(&suite.jobRepo)
}

func (suite *JobServiceTestSuite) TearDownTest() {
	suite.jobRepo.AssertExpectations(suite.T())
}

func (suite *JobServiceTestSuite) TestFetchAllJobs() {
	type testCase struct {
		name           string
		input          JobFilters
		setup          func()
		expectedOutput []Job
		expectedError  bool
	}
	testCases := []testCase{
		{
			name:  "success",
			input: JobFilters{},
			setup: func() {
				suite.jobRepo.On("FetchAllJobs", mock.Anything, repo.JobFilters{}).Return([]repo.Job{
					{
						ID:              1,
						EmployerID:      3,
						Title:           "Software Developer",
						RequiredGender:  "Male",
						Description:     "some random description",
						DurationInHours: 12,
						SkillsRequired:  "Frontend, Backend",
						Sectors:         "IT, Technology, Computers",
						Wage:            2500,
						Vacancy:         3,
						Location:        1,
						Date:            "2025-12-12",
						StartHour:       "",
						EndHour:         "",
						CreatedAt:       time.Time{},
						UpdatedAt:       time.Time{},
						Details:         "Steet 123, Near ABC",
						Street:          "Street 123",
						City:            "Pune",
						State:           "Maharastra",
						Pincode:         411057,
					}, {
						ID:              2,
						EmployerID:      4,
						Title:           "Construction Worker",
						RequiredGender:  "Female",
						Description:     "some random description",
						DurationInHours: 12,
						SkillsRequired:  "Construction",
						Sectors:         "Construction",
						Wage:            1000,
						Vacancy:         5,
						Location:        2,
						Date:            "2025-12-12",
						StartHour:       "",
						EndHour:         "",
						CreatedAt:       time.Time{},
						UpdatedAt:       time.Time{},
						Details:         "Steet 123, Near ABC",
						Street:          "Street 123",
						City:            "Pune",
						State:           "Maharastra",
						Pincode:         411057,
					},
				}, nil)
			},
			expectedOutput: []Job{
				{
					ID:              1,
					EmployerID:      3,
					Title:           "Software Developer",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Frontend, Backend",
					Sectors:         "IT, Technology, Computers",
					Wage:            2500,
					Vacancy:         3,
					Location: worker.Address{
						ID:      1,
						Details: "Steet 123, Near ABC",
						Street:  "Street 123",
						City:    "Pune",
						State:   "Maharastra",
						Pincode: 411057,
					},
					Date:      "2025-12-12",
					StartHour: "",
					EndHour:   "",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				}, {
					ID:              2,
					EmployerID:      4,
					Title:           "Construction Worker",
					RequiredGender:  "Female",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Construction",
					Sectors:         "Construction",
					Wage:            1000,
					Vacancy:         5,
					Location: worker.Address{
						ID:      2,
						Details: "Steet 123, Near ABC",
						Street:  "Street 123",
						City:    "Pune",
						State:   "Maharastra",
						Pincode: 411057,
					},
					Date:      "2025-12-12",
					StartHour: "",
					EndHour:   "",
					CreatedAt: time.Time{},
					UpdatedAt: time.Time{},
				},
			},
			expectedError: false,
		},
		{
			name:  "db error",
			input: JobFilters{},
			setup: func() {
				suite.jobRepo.On("FetchAllJobs", mock.Anything, repo.JobFilters{}).Return([]repo.Job{}, errors.New("db error while list jobs"))
			},
			expectedOutput: []Job{},
			expectedError:  true,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			job, err := suite.service.FetchAllJobs(context.Background(), tc.input)
			if tc.expectedError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expectedOutput, job)
			}
		})
		suite.TearDownTest()
	}
}

func (suite *JobServiceTestSuite) TestFetchJobByID() {
	type testCase struct {
		name           string
		setup          func()
		input          int
		expectedOutput Job
		expectedError  bool
	}
	testCases := []testCase{
		{
			name:  "success",
			input: 1,
			setup: func() {
				suite.jobRepo.On("FetchJobById", mock.Anything, 1).Return(repo.Job{
					ID:              1,
					EmployerID:      3,
					Title:           "Software Developer",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Frontend, Backend",
					Sectors:         "IT, Technology, Computers",
					Wage:            2500,
					Vacancy:         3,
					Location:        1,
					Date:            "2025-12-12",
					StartHour:       "",
					EndHour:         "",
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Details:         "Steet 123, Near ABC",
					Street:          "Street 123",
					City:            "Pune",
					State:           "Maharastra",
					Pincode:         411057,
				}, nil)
			},
			expectedOutput: Job{
				ID:              1,
				EmployerID:      3,
				Title:           "Software Developer",
				RequiredGender:  "Male",
				Description:     "some random description",
				DurationInHours: 12,
				SkillsRequired:  "Frontend, Backend",
				Sectors:         "IT, Technology, Computers",
				Wage:            2500,
				Vacancy:         3,
				Location: worker.Address{
					ID:      1,
					Details: "Steet 123, Near ABC",
					Street:  "Street 123",
					City:    "Pune",
					State:   "Maharastra",
					Pincode: 411057,
				},
				Date:      "2025-12-12",
				StartHour: "",
				EndHour:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			expectedError: false,
		},
		{
			name:  "db error",
			input: 1,
			setup: func() {
				suite.jobRepo.On("FetchJobById", mock.Anything, 1).Return(repo.Job{}, errors.New("db error while fetch job details"))
			},
			expectedOutput: Job{},
			expectedError:  true,
		},
	}
	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			job, err := suite.service.FetchJobByID(context.Background(), tc.input)
			if tc.expectedError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expectedOutput, job)
			}
		})
		suite.TearDownTest()
	}
}

func (suite *JobServiceTestSuite) TestCreateJob() {
	type testCase struct {
		name           string
		setup          func()
		input          Job
		expectedOutput Job
		expectedError  bool
	}
	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.jobRepo.On("CreateJob", mock.Anything, repo.Job{
					EmployerID:      3,
					Title:           "Software Developer",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Frontend, Backend",
					Sectors:         "IT, Technology, Computers",
					Wage:            2500,
					Vacancy:         3,
					Location:        1,
					Date:            "2025-12-12",
					StartHour:       "",
					EndHour:         "",
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Details:         "Steet 123, Near ABC",
					Street:          "Street 123",
					City:            "Pune",
					State:           "Maharastra",
					Pincode:         411057,
				}).Return(repo.Job{
					ID:              1,
					EmployerID:      3,
					Title:           "Software Developer",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Frontend, Backend",
					Sectors:         "IT, Technology, Computers",
					Wage:            2500,
					Vacancy:         3,
					Location:        1,
					Date:            "2025-12-12",
					StartHour:       "",
					EndHour:         "",
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Details:         "Steet 123, Near ABC",
					Street:          "Street 123",
					City:            "Pune",
					State:           "Maharastra",
					Pincode:         411057,
				}, nil)
			},
			input: Job{
				EmployerID:      3,
				Title:           "Software Developer",
				RequiredGender:  "Male",
				Description:     "some random description",
				DurationInHours: 12,
				SkillsRequired:  "Frontend, Backend",
				Sectors:         "IT, Technology, Computers",
				Wage:            2500,
				Vacancy:         3,
				Location: worker.Address{
					ID:      1,
					Details: "Steet 123, Near ABC",
					Street:  "Street 123",
					City:    "Pune",
					State:   "Maharastra",
					Pincode: 411057,
				},
				Date:      "2025-12-12",
				StartHour: "",
				EndHour:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			expectedOutput: Job{
				ID:              1,
				EmployerID:      3,
				Title:           "Software Developer",
				RequiredGender:  "Male",
				Description:     "some random description",
				DurationInHours: 12,
				SkillsRequired:  "Frontend, Backend",
				Sectors:         "IT, Technology, Computers",
				Wage:            2500,
				Vacancy:         3,
				Location: worker.Address{
					ID:      1,
					Details: "Steet 123, Near ABC",
					Street:  "Street 123",
					City:    "Pune",
					State:   "Maharastra",
					Pincode: 411057,
				},
				Date:      "2025-12-12",
				StartHour: "",
				EndHour:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			expectedError: false,
		},
		{
			name: "db error",
			setup: func() {
				suite.jobRepo.On("CreateJob", mock.Anything, repo.Job{
					EmployerID:      3,
					Title:           "Software Developer",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Frontend, Backend",
					Sectors:         "IT, Technology, Computers",
					Wage:            2500,
					Vacancy:         3,
					Location:        1,
					Date:            "2025-12-12",
					StartHour:       "",
					EndHour:         "",
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Details:         "Steet 123, Near ABC",
					Street:          "Street 123",
					City:            "Pune",
					State:           "Maharastra",
					Pincode:         411057,
				}).Return(repo.Job{}, errors.New("db error while create job"))
			},
			input: Job{
				EmployerID:      3,
				Title:           "Software Developer",
				RequiredGender:  "Male",
				Description:     "some random description",
				DurationInHours: 12,
				SkillsRequired:  "Frontend, Backend",
				Sectors:         "IT, Technology, Computers",
				Wage:            2500,
				Vacancy:         3,
				Location: worker.Address{
					ID:      1,
					Details: "Steet 123, Near ABC",
					Street:  "Street 123",
					City:    "Pune",
					State:   "Maharastra",
					Pincode: 411057,
				},
				Date:      "2025-12-12",
				StartHour: "",
				EndHour:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			expectedOutput: Job{},
			expectedError:  true,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			job, err := suite.service.CreateJob(context.Background(), tc.input)
			if tc.expectedError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expectedOutput, job)
			}
		})
		suite.TearDownTest()
	}
}

func (suite *JobServiceTestSuite) TestUpdateJob() {
	type testCase struct {
		name           string
		setup          func()
		input          Job
		expectedOutput Job
		expectedError  bool
	}
	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.jobRepo.On("UpdateJobById", mock.Anything, repo.Job{
					ID:              1,
					EmployerID:      3,
					Title:           "Software Developer",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Frontend, Backend",
					Sectors:         "IT, Technology, Computers",
					Wage:            2500,
					Vacancy:         3,
					Location:        1,
					Date:            "2025-12-12",
					StartHour:       "",
					EndHour:         "",
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Details:         "Steet 123, Near ABC",
					Street:          "Street 123",
					City:            "Pune",
					State:           "Maharastra",
					Pincode:         411057,
				}).Return(repo.Job{
					ID:              1,
					EmployerID:      3,
					Title:           "Software Developer",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Frontend, Backend",
					Sectors:         "IT, Technology, Computers",
					Wage:            2500,
					Vacancy:         3,
					Location:        1,
					Date:            "2025-12-12",
					StartHour:       "",
					EndHour:         "",
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Details:         "Steet 123, Near ABC",
					Street:          "Street 123",
					City:            "Pune",
					State:           "Maharastra",
					Pincode:         411057,
				}, nil)
			},
			input: Job{
				ID:              1,
				EmployerID:      3,
				Title:           "Software Developer",
				RequiredGender:  "Male",
				Description:     "some random description",
				DurationInHours: 12,
				SkillsRequired:  "Frontend, Backend",
				Sectors:         "IT, Technology, Computers",
				Wage:            2500,
				Vacancy:         3,
				Location: worker.Address{
					ID:      1,
					Details: "Steet 123, Near ABC",
					Street:  "Street 123",
					City:    "Pune",
					State:   "Maharastra",
					Pincode: 411057,
				},
				Date:      "2025-12-12",
				StartHour: "",
				EndHour:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			expectedOutput: Job{
				ID:              1,
				EmployerID:      3,
				Title:           "Software Developer",
				RequiredGender:  "Male",
				Description:     "some random description",
				DurationInHours: 12,
				SkillsRequired:  "Frontend, Backend",
				Sectors:         "IT, Technology, Computers",
				Wage:            2500,
				Vacancy:         3,
				Location: worker.Address{
					ID:      1,
					Details: "Steet 123, Near ABC",
					Street:  "Street 123",
					City:    "Pune",
					State:   "Maharastra",
					Pincode: 411057,
				},
				Date:      "2025-12-12",
				StartHour: "",
				EndHour:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			expectedError: false,
		},
		{
			name: "db error",
			setup: func() {
				suite.jobRepo.On("UpdateJobById", mock.Anything, repo.Job{
					EmployerID:      3,
					Title:           "Software Developer",
					RequiredGender:  "Male",
					Description:     "some random description",
					DurationInHours: 12,
					SkillsRequired:  "Frontend, Backend",
					Sectors:         "IT, Technology, Computers",
					Wage:            2500,
					Vacancy:         3,
					Location:        1,
					Date:            "2025-12-12",
					StartHour:       "",
					EndHour:         "",
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Details:         "Steet 123, Near ABC",
					Street:          "Street 123",
					City:            "Pune",
					State:           "Maharastra",
					Pincode:         411057,
				}).Return(repo.Job{}, errors.New("db error while create job"))
			},
			input: Job{
				EmployerID:      3,
				Title:           "Software Developer",
				RequiredGender:  "Male",
				Description:     "some random description",
				DurationInHours: 12,
				SkillsRequired:  "Frontend, Backend",
				Sectors:         "IT, Technology, Computers",
				Wage:            2500,
				Vacancy:         3,
				Location: worker.Address{
					ID:      1,
					Details: "Steet 123, Near ABC",
					Street:  "Street 123",
					City:    "Pune",
					State:   "Maharastra",
					Pincode: 411057,
				},
				Date:      "2025-12-12",
				StartHour: "",
				EndHour:   "",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
			},
			expectedOutput: Job{},
			expectedError:  true,
		},
	}
	for _, tc := range testCases {
		suite.T().Run(tc.name, func(t *testing.T) {
			tc.setup()
			job, err := suite.service.UpdateJobByID(context.Background(), tc.input)
			if tc.expectedError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expectedOutput, job)
			}
		})
		suite.TearDownTest()
	}
}

func (suite *JobServiceTestSuite) TestDeleteJob() {
	type testCase struct {
		name           string
		setup          func()
		input          int
		expectedOutput int
		expectedError  bool
	}
	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.jobRepo.On("FindJobById", mock.Anything, 1).Return(true)
				suite.jobRepo.On("DeleteJobById", mock.Anything, 1).Return(1, nil)
			},
			input:          1,
			expectedOutput: 1,
			expectedError:  false,
		},
		{
			name: "db error while delete job",
			setup: func() {
				suite.jobRepo.On("FindJobById", mock.Anything, 1).Return(true)
				suite.jobRepo.On("DeleteJobById", mock.Anything, 1).Return(-1, errors.New("db error while delete job"))
			},
			input:          1,
			expectedOutput: -1,
			expectedError:  true,
		},
		{
			name:  "job with Id not found",
			input: 1,
			setup: func() {
				suite.jobRepo.On("FindJobById", mock.Anything, 1).Return(false)
			},
			expectedOutput: -1,
			expectedError:  true,
		},
	}
	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			job, err := suite.service.DeleteJobByID(context.Background(), tc.input)
			if tc.expectedError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expectedOutput, job)
			}
		})
		suite.TearDownTest()
	}
}

func (suite *JobServiceTestSuite) TestFetchApplicationsByJobId() {
	type testCase struct {
		name           string
		jobId          int
		setup          func()
		expectedOutput []application.ApplicationCompleteEmp
		expectedError  bool
	}
	testCases := []testCase{
		{
			name:  "success",
			jobId: 1,
			setup: func() {
				suite.jobRepo.On("FindJobById", mock.Anything, 1).Return(true)
				suite.jobRepo.On("FetchApplicationsByJobId", mock.Anything, 1).Return([]repo.ApplicationCompleteEmp{
					{
						ID:             1,
						JobID:          2,
						WorkerID:       2,
						Status:         "Pending",
						ExpectedWage:   1000,
						ModeOfArrival:  "Personal",
						PickUpLocation: 5,
						WorkerComment:  "some random commetns from worker",
						AppliedAt:      time.Time{},
						UpdatedAt:      time.Time{},
						Details:        "details",
						Street:         "street",
						City:           "city",
						State:          "state",
						Pincode:        4125,
						JobTitle:       "some random title",
						Description:    "description",
						SkillsRequired: "random skills",
						JobSectors:     "sectors",
						JobWage:        1500,
						Vacancy:        5,
						JobDate:        "2025-12-3",
						WorkerName:     "Jogn Doe",
						ContactNumber:  "9067691363",
						WorkerEmail:    "harsh@gmail.com",
						WorkerGender:   "Male",
					},
				}, nil)
			},
			expectedOutput: []application.ApplicationCompleteEmp{
				{
					ID:             1,
					JobID:          2,
					WorkerID:       2,
					Status:         "Pending",
					ExpectedWage:   1000,
					ModeOfArrival:  "Personal",
					PickUpLocation: 5,
					WorkerComment:  "some random commetns from worker",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
					Details:        "details",
					Street:         "street",
					City:           "city",
					State:          "state",
					Pincode:        4125,
					JobTitle:       "some random title",
					Description:    "description",
					SkillsRequired: "random skills",
					JobSectors:     "sectors",
					JobWage:        1500,
					Vacancy:        5,
					JobDate:        "2025-12-3",
					WorkerName:     "Jogn Doe",
					ContactNumber:  "9067691363",
					WorkerEmail:    "harsh@gmail.com",
					WorkerGender:   "Male",
				},
			},
			expectedError: false,
		},
		{
			name:  "db error",
			jobId: 1,
			setup: func() {
				suite.jobRepo.On("FindJobById", mock.Anything, 1).Return(true)
				suite.jobRepo.On("FetchApplicationsByJobId", mock.Anything, 1).Return([]repo.ApplicationCompleteEmp{}, errors.New("db error while fetch applications by Job ID"))
			},
			expectedOutput: []application.ApplicationCompleteEmp{},
			expectedError:  true,
		},
		{
			name:  "job with id not found",
			jobId: 1,
			setup: func() {
				suite.jobRepo.On("FindJobById", mock.Anything, 1).Return(false)
			},
			expectedOutput: []application.ApplicationCompleteEmp{},
			expectedError:  true,
		},
	}

	for _, tc := range testCases {
		suite.SetupTest()
		suite.Run(tc.name, func() {
			tc.setup()
			job, err := suite.service.FetchApplicationsByJobId(context.Background(), tc.jobId)
			if tc.expectedError {
				suite.Require().Error(err)
			} else {
				suite.Require().NoError(err)
				suite.Require().Equal(tc.expectedOutput, job)
			}
		})
		suite.TearDownTest()
	}
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(JobServiceTestSuite))
}
