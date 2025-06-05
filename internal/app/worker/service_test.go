package worker

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/application"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WorkerServiceTestSuite struct {
	suite.Suite
	service    Service
	workerRepo mocks.WorkerStorer
}

func (suite *WorkerServiceTestSuite) SetupTest() {
	suite.workerRepo = mocks.WorkerStorer{}
	suite.service = NewService(&suite.workerRepo)
}

func (suite *WorkerServiceTestSuite) TearDownTest() {
	suite.workerRepo.AssertExpectations(suite.T())
}

func TestWorkerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerServiceTestSuite))
}
	
func (suite *WorkerServiceTestSuite) TestDeleteWorkerById() {
	type testCase struct {
		name           string
		workerId       int
		setup          func()
		expectedOutput int
		expectedError  bool
	}
	testCases := []testCase{
		{
			name:     "success",
			workerId: 1,
			setup: func() {
				suite.workerRepo.On("FindWorkerById", mock.Anything, 1).Return(true)
				suite.workerRepo.On("DeleteWorkerByID", mock.Anything, 1).Return(1, nil)
			},
			expectedOutput: 1,
			expectedError:  false,
		},
		{
			name:     "db error",
			workerId: 1,
			setup: func() {
				suite.workerRepo.On("FindWorkerById", mock.Anything, 1).Return(true)
				suite.workerRepo.On("DeleteWorkerByID", mock.Anything, 1).Return(-1, errors.New("db error while delete worker"))
			},
			expectedOutput: -1,
			expectedError:  true,
		},
		{
			name:     "worker with id doesn't exists",
			workerId: 1,
			setup: func() {
				suite.workerRepo.On("FindWorkerById", mock.Anything, 1).Return(false)
			},
			expectedOutput: -1,
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			id, err := suite.service.DeleteWorkerByID(context.Background(), test.workerId)
			suite.Equal(test.expectedOutput, id)
			suite.Equal(test.expectedError, err != nil)
		})
	}
}

func (suite *WorkerServiceTestSuite) TestFetchWorkerById() {
	type testCase struct {
		name           string
		workerId       int
		setup          func()
		expectedOutput Worker
		expectedError  bool
	}

	testCases := []testCase{
		{
			name:     "success",
			workerId: 1,
			setup: func() {
				suite.workerRepo.On("FetchWorkerByID", mock.Anything, 1).Return(repo.Worker{
					ID:              1,
					Name:            "John",
					ContactNumber:   "9067691363",
					Email:           "john@gmail.com",
					Gender:          "male",
					Sectors:         "something",
					Skills:          "some skills",
					Location:        1,
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "English, Hindi",
					Details:         "details",
					Street:          "Street",
					City:            "city",
					State:           "state",
					Pincode:         411057,
				}, nil)
			},
			expectedOutput: Worker{
				ID:            1,
				Name:          "John",
				ContactNumber: "9067691363",
				Email:         "john@gmail.com",
				Gender:        "male",
				Sectors:       "something",
				Skills:        "some skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "Street",
					City:    "city",
					State:   "state",
					Pincode: 411057,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "",
			},
			expectedError: false,
		},
		{
			name:     "db error",
			workerId: 1,
			setup: func() {
				suite.workerRepo.On("FetchWorkerByID", mock.Anything, 1).Return(repo.Worker{}, errors.New("failed to fetch worker"))
			},
			expectedOutput: Worker{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			worker, err := suite.service.FetchWorkerByID(context.Background(), test.workerId)
			suite.Equal(test.expectedOutput, worker)
			suite.Equal(test.expectedError, err != nil)
		})
		suite.TearDownTest()
	}
}

func (suite *WorkerServiceTestSuite) TestUpdateWorkerById() {
	type testCase struct {
		name           string
		workerData     Worker
		setup          func()
		expectedOutput Worker
		expectedError  bool
	}

	testCases := []testCase{
		{
			name: "sucessfully updated",
			workerData: Worker{
				ID:            1,
				Name:          "John",
				ContactNumber: "9067691363",
				Email:         "john@gmail.com",
				Gender:        "male",
				Sectors:       "something",
				Skills:        "some skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "Street",
					City:    "city",
					State:   "state",
					Pincode: 411057,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "",
			},
			setup: func() {
				suite.workerRepo.On("UpdateWorkerByID", mock.Anything, repo.Worker{
					ID:              1,
					Name:            "John",
					ContactNumber:   "9067691363",
					Email:           "john@gmail.com",
					Gender:          "male",
					Sectors:         "something",
					Skills:          "some skills",
					Location:        1,
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "",
					Details:         "details",
					Street:          "Street",
					City:            "city",
					State:           "state",
					Pincode:         411057,
				}).Return(repo.Worker{
					ID:              1,
					Name:            "John",
					ContactNumber:   "9067691363",
					Email:           "john@gmail.com",
					Gender:          "male",
					Sectors:         "something",
					Skills:          "some skills",
					Location:        1,
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "",
					Details:         "details",
					Street:          "Street",
					City:            "city",
					State:           "state",
					Pincode:         411057,
				}, nil)
			},
			expectedOutput: Worker{
				ID:            1,
				Name:          "John",
				ContactNumber: "9067691363",
				Email:         "john@gmail.com",
				Gender:        "male",
				Sectors:       "something",
				Skills:        "some skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "Street",
					City:    "city",
					State:   "state",
					Pincode: 411057,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "",
			},
			expectedError: false,
		},
		{
			name: "db error",
			workerData: Worker{
				ID:            1,
				Name:          "John",
				ContactNumber: "9067691363",
				Email:         "john@gmail.com",
				Gender:        "male",
				Sectors:       "something",
				Skills:        "some skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "Street",
					City:    "city",
					State:   "state",
					Pincode: 411057,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "",
			},
			setup: func() {
				suite.workerRepo.On("UpdateWorkerByID", mock.Anything, repo.Worker{
					ID:              1,
					Name:            "John",
					ContactNumber:   "9067691363",
					Email:           "john@gmail.com",
					Gender:          "male",
					Sectors:         "something",
					Skills:          "some skills",
					Location:        1,
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "",
					Details:         "details",
					Street:          "Street",
					City:            "city",
					State:           "state",
					Pincode:         411057,
				}).Return(repo.Worker{}, errors.New("db error while update worker by id"))
			},
			expectedOutput: Worker{},
			expectedError:  true,
		},
		{
			name: "validation error in name",
			workerData: Worker{
				ID:            1,
				Name:          "John1",
				ContactNumber: "9067691363",
				Email:         "john@gmail.com",
				Gender:        "male",
				Sectors:       "something",
				Skills:        "some skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "Street",
					City:    "city",
					State:   "state",
					Pincode: 411057,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "",
			},
			setup: func() {

			},
			expectedOutput: Worker{},
			expectedError:  true,
		},
		{
			name: "validation error in email",
			workerData: Worker{
				ID:            1,
				Name:          "John",
				ContactNumber: "9067691363",
				Email:         "johngmail.com",
				Gender:        "male",
				Sectors:       "something",
				Skills:        "some skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "Street",
					City:    "city",
					State:   "state",
					Pincode: 411057,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "",
			},
			setup: func() {

			},
			expectedOutput: Worker{},
			expectedError:  true,
		},
		{
			name: "validation error in mobile",
			workerData: Worker{
				ID:            1,
				Name:          "John1",
				ContactNumber: "90676913",
				Email:         "john@gmail.com",
				Gender:        "male",
				Sectors:       "something",
				Skills:        "some skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "Street",
					City:    "city",
					State:   "state",
					Pincode: 411057,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "",
			},
			setup: func() {

			},
			expectedOutput: Worker{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			worker, err := suite.service.UpdateWorkerByID(context.Background(), test.workerData)
			suite.Equal(test.expectedOutput, worker)
			suite.Equal(test.expectedError, err != nil)
		})
		suite.TearDownTest()
	}
}

func (suite *WorkerServiceTestSuite) TestApplicationsByWorkerId() {
	type testCase struct {
		name           string
		workerId       int
		setup          func()
		expectedOutput []application.ApplicationComplete
		expectedError  bool
	}

	testCases := []testCase{
		{
			name:     "success",
			workerId: 1,
			setup: func() {
				suite.workerRepo.On("FindWorkerById", mock.Anything, 1).Return(true)
				suite.workerRepo.On("FetchApplicationsByWorkerId", mock.Anything, 1).Return([]repo.ApplicationComplete{
					{
						ID:             1,
						JobID:          1,
						WorkerID:       1,
						Status:         "Pending",
						ExpectedWage:   1200,
						ModeOfArrival:  "Personal",
						PickUpLocation: 1,
						WorkerComment:  "some random comments",
						AppliedAt:      time.Time{},
						UpdatedAt:      time.Time{},
						Details:        "details",
						Street:         "street",
						City:           "city",
						State:          "state",
						Pincode:        411057,
						JobTitle:       "job title",
						Description:    "some description",
						SkillsRequired: "some random skills",
						JobSectors:     "some random sectors",
						JobWage:        1500,
						Vacancy:        5,
						JobDate:        "2025-02-02",
						EmployerName:   "John Doe",
						ContactNumber:  "9067691363",
						EmployerEmail:  "emp@gmail.com",
						EmployerType:   "Organization",
					},
					{
						ID:             1,
						JobID:          1,
						WorkerID:       1,
						Status:         "Pending",
						ExpectedWage:   1200,
						ModeOfArrival:  "Personal",
						PickUpLocation: 1,
						WorkerComment:  "some random comments",
						AppliedAt:      time.Time{},
						UpdatedAt:      time.Time{},
						Details:        "details",
						Street:         "street",
						City:           "city",
						State:          "state",
						Pincode:        411057,
						JobTitle:       "job title",
						Description:    "some description",
						SkillsRequired: "some random skills",
						JobSectors:     "some random sectors",
						JobWage:        1500,
						Vacancy:        5,
						JobDate:        "2025-02-02",
						EmployerName:   "John Doe",
						ContactNumber:  "9067691363",
						EmployerEmail:  "emp@gmail.com",
						EmployerType:   "Organization",
					},
				}, nil)
			},
			expectedOutput: []application.ApplicationComplete{
				{
					ID:             1,
					JobID:          1,
					WorkerID:       1,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Personal",
					PickUpLocation: 1,
					WorkerComment:  "some random comments",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
					Details:        "details",
					Street:         "street",
					City:           "city",
					State:          "state",
					Pincode:        411057,
					JobTitle:       "job title",
					Description:    "some description",
					SkillsRequired: "some random skills",
					JobSectors:     "some random sectors",
					JobWage:        1500,
					Vacancy:        5,
					JobDate:        "2025-02-02",
					EmployerName:   "John Doe",
					ContactNumber:  "9067691363",
					EmployerEmail:  "emp@gmail.com",
					EmployerType:   "Organization",
				},
				{
					ID:             1,
					JobID:          1,
					WorkerID:       1,
					Status:         "Pending",
					ExpectedWage:   1200,
					ModeOfArrival:  "Personal",
					PickUpLocation: 1,
					WorkerComment:  "some random comments",
					AppliedAt:      time.Time{},
					UpdatedAt:      time.Time{},
					Details:        "details",
					Street:         "street",
					City:           "city",
					State:          "state",
					Pincode:        411057,
					JobTitle:       "job title",
					Description:    "some description",
					SkillsRequired: "some random skills",
					JobSectors:     "some random sectors",
					JobWage:        1500,
					Vacancy:        5,
					JobDate:        "2025-02-02",
					EmployerName:   "John Doe",
					ContactNumber:  "9067691363",
					EmployerEmail:  "emp@gmail.com",
					EmployerType:   "Organization",
				},
			},
			expectedError: false,
		},
		{
			name:     "db error",
			workerId: 1,
			setup: func() {
				suite.workerRepo.On("FindWorkerById", mock.Anything, 1).Return(true)
				suite.workerRepo.On("FetchApplicationsByWorkerId", mock.Anything, 1).Return([]repo.ApplicationComplete{}, errors.New("db error while fetch applications by worker id"))
			},
			expectedOutput: []application.ApplicationComplete{},
			expectedError:  true,
		},
		{
			name:     "worker by id not found",
			workerId: 1,
			setup: func() {
				suite.workerRepo.On("FindWorkerById", mock.Anything, 1).Return(false)
			},
			expectedOutput: []application.ApplicationComplete{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			worker, err := suite.service.FetchApplicationsByWorkerId(context.Background(), test.workerId)
			suite.Equal(test.expectedOutput, worker)
			suite.Equal(test.expectedError, err != nil)
		})
		suite.TearDownTest()
	}
}

func (suite *WorkerServiceTestSuite) TestFetchAllWorkers() {
	type testCase struct {
		name           string
		setup          func()
		expectedOutput []Worker
		expectedError  bool
	}

	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.workerRepo.On("FetchAllWorkers", mock.Anything).Return([]repo.Worker{
					{
						ID:              1,
						Name:            "John",
						ContactNumber:   "9067691363",
						Email:           "john@gmail.com",
						Gender:          "male",
						Sectors:         "something",
						Skills:          "some skills",
						Location:        1,
						IsAvailable:     true,
						Rating:          0,
						TotalJobsWorked: 0,
						CreatedAt:       time.Time{},
						UpdatedAt:       time.Time{},
						Language:        "",
						Details:         "details",
						Street:          "Street",
						City:            "city",
						State:           "state",
						Pincode:         411057,
					},
					{
						ID:              1,
						Name:            "John",
						ContactNumber:   "9067691363",
						Email:           "john@gmail.com",
						Gender:          "male",
						Sectors:         "something",
						Skills:          "some skills",
						Location:        1,
						IsAvailable:     true,
						Rating:          0,
						TotalJobsWorked: 0,
						CreatedAt:       time.Time{},
						UpdatedAt:       time.Time{},
						Language:        "",
						Details:         "details",
						Street:          "Street",
						City:            "city",
						State:           "state",
						Pincode:         411057,
					},
				}, nil)
			},
			expectedOutput: []Worker{
				{
					ID:            1,
					Name:          "John",
					ContactNumber: "9067691363",
					Email:         "john@gmail.com",
					Gender:        "male",
					Sectors:       "something",
					Skills:        "some skills",
					Location: Address{
						ID:      1,
						Details: "details",
						Street:  "Street",
						City:    "city",
						State:   "state",
						Pincode: 411057,
					},
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "",
				},
				{
					ID:            1,
					Name:          "John",
					ContactNumber: "9067691363",
					Email:         "john@gmail.com",
					Gender:        "male",
					Sectors:       "something",
					Skills:        "some skills",
					Location: Address{
						ID:      1,
						Details: "details",
						Street:  "Street",
						City:    "city",
						State:   "state",
						Pincode: 411057,
					},
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "",
				},
			},
			expectedError: false,
		},
		{
			name: "db error",
			setup: func() {
				suite.workerRepo.On("FetchAllWorkers", mock.Anything).Return([]repo.Worker{}, errors.New("db error while list all workers"))
			},
			expectedOutput: []Worker{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			worker, err := suite.service.FetchAllWorkers(context.Background())
			suite.Equal(test.expectedOutput, worker)
			suite.Equal(test.expectedError, err != nil)
		})
		suite.TearDownTest()
	}

}

func (suite *WorkerServiceTestSuite) TestCreateWorker() {
	type testCase struct {
		name           string
		workerData     Worker
		setup          func()
		expectedOutput Worker
		expectedError  bool
	}

	testCases := []testCase{
		{
			name: "invalid name",
			workerData: Worker{
				ID:            1,
				Name:          "John Doe1",
				ContactNumber: "9067691363",
				Email:         "john@gmail.com",
				Gender:        "Male",
				Password:      "Something@123",
				Sectors:       "some random sectors",
				Skills:        "some random skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411054,
				},
				Language: "English, Hindi",
			},
			setup:          func() {},
			expectedOutput: Worker{},
			expectedError:  true,
		},
		{
			name: "invalid email",
			workerData: Worker{
				ID:            1,
				Name:          "John Doe",
				ContactNumber: "9067691363",
				Email:         "johngmail.com",
				Gender:        "Male",
				Password:      "Something@123",
				Sectors:       "some random sectors",
				Skills:        "some random skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411054,
				},
				Language: "English, Hindi",
			},
			setup:          func() {},
			expectedOutput: Worker{},
			expectedError:  true,
		},
		{
			name: "invalid password",
			workerData: Worker{
				ID:            1,
				Name:          "John Doe",
				ContactNumber: "9067691363",
				Email:         "john@gmail.com",
				Gender:        "Male",
				Password:      "",
				Sectors:       "some random sectors",
				Skills:        "some random skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411054,
				},
				Language: "English, Hindi",
			},
			setup:          func() {},
			expectedOutput: Worker{},
			expectedError:  true,
		},
		{
			name: "invalid mobile number",
			workerData: Worker{
				ID:            1,
				Name:          "John Doe",
				ContactNumber: "906769136",
				Email:         "john@gmail.com",
				Gender:        "Male",
				Password:      "Something@123",
				Sectors:       "some random sectors",
				Skills:        "some random skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411054,
				},
				Language: "English, Hindi",
			},
			setup:          func() {},
			expectedOutput: Worker{},
			expectedError:  true,
		},
		{
			name: "worker with email already exists",
			workerData: Worker{
				ID:            1,
				Name:          "John Doe",
				ContactNumber: "9067691363",
				Email:         "john@gmail.com",
				Gender:        "Male",
				Password:      "Something@123",
				Sectors:       "some random sectors",
				Skills:        "some random skills",
				Location: Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411054,
				},
				Language: "English, Hindi",
			},
			setup: func() {
				suite.workerRepo.On("FindWorkerByEmail", mock.Anything, "john@gmail.com").Return(true)
			},
			expectedOutput: Worker{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			worker, err := suite.service.CreateWorker(context.Background(), test.workerData)
			suite.Equal(test.expectedOutput, worker)
			suite.Equal(test.expectedError, err != nil)
		})
		suite.TearDownTest()
	}

}
