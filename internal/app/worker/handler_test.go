package worker_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/application"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker/mocks"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type WorkerHandlerTestSuite struct {
	suite.Suite
	workerService *mocks.Service
	router        mux.Router
}

func (suite *WorkerHandlerTestSuite) SetupTest() {
	suite.workerService = &mocks.Service{}
	suite.router = *mux.NewRouter()
}

func (suite *WorkerHandlerTestSuite) TearDownTest() {
	suite.workerService.AssertExpectations(suite.T())
}

func TestApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerHandlerTestSuite))
}

func (suite *WorkerHandlerTestSuite) TestFetchWorkerByID() {
	type testCase struct {
		name               string
		worker_id          interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:      "success",
			worker_id: 1,
			setup: func() {
				suite.workerService.On("FetchWorkerByID", mock.Anything, 1).Return(worker.Worker{
					ID:            1,
					Name:          "John Doe",
					ContactNumber: "9067691363",
					Email:         "john@gmail.com",
					Gender:        "Male",
					Sectors:       "sectors",
					Skills:        "skills",
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411057,
					},
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "English",
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid worker id",
			worker_id:          "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "worker doesn't exist",
			worker_id: 1,
			setup: func() {
				suite.workerService.On("FetchWorkerByID", mock.Anything, 1).Return(worker.Worker{}, apperrors.ErrNoWorkerExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:      "db error",
			worker_id: 1,
			setup: func() {
				suite.workerService.On("FetchWorkerByID", mock.Anything, 1).Return(worker.Worker{}, errors.New("error while fetch worker"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/worker/{worker_id}", worker.FetchWorkerByID(suite.workerService)).Methods(http.MethodGet)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/worker/%v", test.worker_id), bytes.NewBuffer([]byte(``)))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
		suite.TearDownTest()
	}
}

func (suite *WorkerHandlerTestSuite) TestDeleteWorkerByID() {
	type testCase struct {
		name               string
		setup              func()
		workerId           interface{}
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.workerService.On("DeleteWorkerByID", mock.Anything, 1).Return(1, nil)
			},
			workerId:           1,
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name: "db error",
			setup: func() {
				suite.workerService.On("DeleteWorkerByID", mock.Anything, 1).Return(-1, errors.New("error while delete worker"))
			},
			workerId:           1,
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "worker not found",
			setup: func() {
				suite.workerService.On("DeleteWorkerByID", mock.Anything, 1).Return(-1, apperrors.ErrNoWorkerExists)
			},
			workerId:           1,
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:               "invalid worker id",
			setup:              func() {},
			workerId:           "a",
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/worker/{worker_id}", worker.DeleteWorkerByID(suite.workerService)).Methods(http.MethodDelete)

			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/worker/%v", test.workerId), bytes.NewBuffer([]byte(``)))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
		suite.TearDownTest()
	}
}

func (suite *WorkerHandlerTestSuite) TestFetchApplicationsByWorkerId() {
	type testCase struct {
		name               string
		worker_id          interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:      "success",
			worker_id: 1,
			setup: func() {
				suite.workerService.On("FetchApplicationsByWorkerId", mock.Anything, 1).Return([]application.ApplicationComplete{
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
						Pincode:        541205,
						JobTitle:       "Job Title",
						Description:    "random description",
						SkillsRequired: "skills",
						JobSectors:     "sectors",
						JobWage:        1502,
						Vacancy:        5,
						JobDate:        "",
						EmployerName:   "John Doe",
						ContactNumber:  "9067691363",
						EmployerEmail:  "john@gmail.com",
						EmployerType:   "Employer",
					},
					{
						ID:             2,
						JobID:          3,
						WorkerID:       4,
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
						Pincode:        541205,
						JobTitle:       "Job Title",
						Description:    "random description",
						SkillsRequired: "skills",
						JobSectors:     "sectors",
						JobWage:        1502,
						Vacancy:        5,
						JobDate:        "",
						EmployerName:   "John Doe",
						ContactNumber:  "9067691363",
						EmployerEmail:  "john@gmail.com",
						EmployerType:   "Employer",
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid worker id",
			worker_id:          "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:      "worker with id not found",
			worker_id: 1,
			setup: func() {
				suite.workerService.On("FetchApplicationsByWorkerId", mock.Anything, 1).Return([]application.ApplicationComplete{}, apperrors.ErrNoWorkerExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:      "internal error",
			worker_id: 1,
			setup: func() {
				suite.workerService.On("FetchApplicationsByWorkerId", mock.Anything, 1).Return([]application.ApplicationComplete{}, errors.New("failed to fetch from db"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/worker/{worker_id}/applications", worker.FetchApplicationsByWorkerId(suite.workerService)).Methods(http.MethodGet)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/worker/%v/applications", test.worker_id), bytes.NewBuffer([]byte(``)))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
		suite.TearDownTest()
	}
}

func (suite *WorkerHandlerTestSuite) TestFetchAllWorkers() {
	type testCase struct {
		name               string
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.workerService.On("FetchAllWorkers", mock.Anything).Return([]worker.Worker{
					{
						ID:            1,
						Name:          "John Doe",
						ContactNumber: "9067691363",
						Email:         "harsh@gmail.com",
						Gender:        "Male",
						Sectors:       "sectors",
						Skills:        "skills",
						Location: worker.Address{
							ID:      1,
							Details: "details",
							Street:  "street",
							City:    "city",
							State:   "state",
							Pincode: 541254,
						},
						IsAvailable:     true,
						Rating:          0,
						TotalJobsWorked: 0,
						CreatedAt:       time.Time{},
						UpdatedAt:       time.Time{},
						Language:        "English",
					},
					{
						ID:            2,
						Name:          "Harsh Jagtap",
						ContactNumber: "9067691363",
						Email:         "harsh@gmail.com",
						Gender:        "Male",
						Sectors:       "sectors",
						Skills:        "skills",
						Location: worker.Address{
							ID:      1,
							Details: "details",
							Street:  "street",
							City:    "city",
							State:   "state",
							Pincode: 541254,
						},
						IsAvailable:     true,
						Rating:          0,
						TotalJobsWorked: 0,
						CreatedAt:       time.Time{},
						UpdatedAt:       time.Time{},
						Language:        "English",
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "internal error",
			setup: func() {
				suite.workerService.On("FetchAllWorkers", mock.Anything).Return([]worker.Worker{}, errors.New("internal error while fetch all workers"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/workers", worker.FetchAllWorkers(suite.workerService)).Methods(http.MethodGet)

			req, err := http.NewRequest(http.MethodGet, "/workers", bytes.NewBuffer([]byte(``)))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
		suite.TearDownTest()
	}
}

func (suite *WorkerHandlerTestSuite) TestUpdateWorkerByID() {
	type testCase struct {
		name               string
		input              interface{}
		worker_id          interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:      "success",
			worker_id: 2,
			input: worker.Worker{
				ID:            2,
				Name:          "Harsh Jagtap",
				ContactNumber: "9067691363",
				Email:         "harsh@gmail.com",
				Gender:        "Male",
				Sectors:       "sectors",
				Skills:        "skills",
				Location: worker.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 541254,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "English",
			},
			setup: func() {
				suite.workerService.On("UpdateWorkerByID", mock.Anything, worker.Worker{
					ID:            2,
					Name:          "Harsh Jagtap",
					ContactNumber: "9067691363",
					Email:         "harsh@gmail.com",
					Gender:        "Male",
					Sectors:       "sectors",
					Skills:        "skills",
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 541254,
					},
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "English",
				}).Return(worker.Worker{
					ID:            2,
					Name:          "Harsh Jagtap",
					ContactNumber: "9067691363",
					Email:         "harsh@gmail.com",
					Gender:        "Male",
					Sectors:       "sectors",
					Skills:        "skills",
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 541254,
					},
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "English",
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "invalid worker id",
			input: worker.Worker{
				ID:            2,
				Name:          "Harsh Jagtap",
				ContactNumber: "9067691363",
				Email:         "harsh@gmail.com",
				Gender:        "Male",
				Sectors:       "sectors",
				Skills:        "skills",
				Location: worker.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 541254,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "English",
			},
			worker_id:          "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:               "invalid request data",
			input:              "worker data",
			worker_id:          1,
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid user details",
			input: worker.Worker{
				ID:            2,
				Name:          "Harsh Jagtap1",
				ContactNumber: "9067691363",
				Email:         "harsh@gmail.com",
				Gender:        "Male",
				Sectors:       "sectors",
				Skills:        "skills",
				Location: worker.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 541254,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "English",
			},
			worker_id: 2,
			setup: func() {
				suite.workerService.On("UpdateWorkerByID", mock.Anything, worker.Worker{
					ID:            2,
					Name:          "Harsh Jagtap1",
					ContactNumber: "9067691363",
					Email:         "harsh@gmail.com",
					Gender:        "Male",
					Sectors:       "sectors",
					Skills:        "skills",
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 541254,
					},
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "English",
				}).Return(worker.Worker{}, apperrors.ErrInvalidUserDetails)
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "internal error",
			input: worker.Worker{
				ID:            2,
				Name:          "Harsh Jagtap1",
				ContactNumber: "9067691363",
				Email:         "harsh@gmail.com",
				Gender:        "Male",
				Sectors:       "sectors",
				Skills:        "skills",
				Location: worker.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 541254,
				},
				IsAvailable:     true,
				Rating:          0,
				TotalJobsWorked: 0,
				CreatedAt:       time.Time{},
				UpdatedAt:       time.Time{},
				Language:        "English",
			},
			worker_id: 2,
			setup: func() {
				suite.workerService.On("UpdateWorkerByID", mock.Anything, worker.Worker{
					ID:            2,
					Name:          "Harsh Jagtap1",
					ContactNumber: "9067691363",
					Email:         "harsh@gmail.com",
					Gender:        "Male",
					Sectors:       "sectors",
					Skills:        "skills",
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 541254,
					},
					IsAvailable:     true,
					Rating:          0,
					TotalJobsWorked: 0,
					CreatedAt:       time.Time{},
					UpdatedAt:       time.Time{},
					Language:        "English",
				}).Return(worker.Worker{}, errors.New("internal error while update user with id"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	t := suite.T()
	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/worker/{worker_id}", worker.UpdateWorkerByID(suite.workerService)).Methods(http.MethodPut)

			reqBody, err := json.Marshal(test.input)
			if err != nil {
				logger.Errorw(context.Background(), "error occured while marshal req body")
			}

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/worker/%v", test.worker_id), bytes.NewBuffer(reqBody))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
		suite.TearDownTest()
	}
}
