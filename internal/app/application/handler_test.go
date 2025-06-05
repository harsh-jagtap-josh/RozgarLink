package application_test

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
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/application/mocks"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ApplicationHandlerTestSuite struct {
	suite.Suite
	appService mocks.Service
	router     mux.Router
}

func (suite *ApplicationHandlerTestSuite) SetupTest() {
	suite.appService = mocks.Service{}
	suite.router = *mux.NewRouter()
}

func (suite *ApplicationHandlerTestSuite) TearDownTest() {
	suite.appService.AssertExpectations(suite.T())
}

func TestApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(ApplicationHandlerTestSuite))
}

func (suite *ApplicationHandlerTestSuite) TestCreateNewApplication() {
	type testCase struct {
		name               string
		input              interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			input: application.Application{
				ID:            1,
				JobID:         1,
				WorkerID:      1,
				Status:        "Pending",
				ExpectedWage:  1500,
				ModeOfArrival: "Personal",
				PickUpLocation: application.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411052,
				},
				WorkerComment: "some random comments",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			setup: func() {
				suite.appService.On("CreateNewApplication", mock.Anything, application.Application{
					ID:            1,
					JobID:         1,
					WorkerID:      1,
					Status:        "Pending",
					ExpectedWage:  1500,
					ModeOfArrival: "Personal",
					PickUpLocation: application.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411052,
					},
					WorkerComment: "some random comments",
					AppliedAt:     time.Time{},
					UpdatedAt:     time.Time{},
				}).Return(application.Application{
					ID:            1,
					JobID:         1,
					WorkerID:      1,
					Status:        "Pending",
					ExpectedWage:  1500,
					ModeOfArrival: "Personal",
					PickUpLocation: application.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411052,
					},
					WorkerComment: "some random comments",
					AppliedAt:     time.Time{},
					UpdatedAt:     time.Time{},
				}, nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "invalid req body",
			input:              "application Data",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "db error",
			input: application.Application{
				ID:            1,
				JobID:         1,
				WorkerID:      1,
				Status:        "Pending",
				ExpectedWage:  1500,
				ModeOfArrival: "Personal",
				PickUpLocation: application.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411052,
				},
				WorkerComment: "some random comments",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			setup: func() {
				suite.appService.On("CreateNewApplication", mock.Anything, application.Application{
					ID:            1,
					JobID:         1,
					WorkerID:      1,
					Status:        "Pending",
					ExpectedWage:  1500,
					ModeOfArrival: "Personal",
					PickUpLocation: application.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411052,
					},
					WorkerComment: "some random comments",
					AppliedAt:     time.Time{},
					UpdatedAt:     time.Time{},
				}).Return(application.Application{}, errors.New("db error while create application"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/application/create", application.CreateNewApplication(&suite.appService)).Methods(http.MethodPost)

			reqBody, err := json.Marshal(test.input)
			if err != nil {
				logger.Errorw(context.Background(), "error while json marshal test data, error : "+err.Error())
			}

			req, err := http.NewRequest(http.MethodPost, "/application/create", bytes.NewBuffer(reqBody))
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

func (suite *ApplicationHandlerTestSuite) TestUpdateApplicationByID() {
	type testCase struct {
		name               string
		input              interface{}
		applicationId      interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			input: application.Application{
				ID:            1,
				JobID:         1,
				WorkerID:      1,
				Status:        "Pending",
				ExpectedWage:  1500,
				ModeOfArrival: "Personal",
				PickUpLocation: application.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411052,
				},
				WorkerComment: "some random comments",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			applicationId: 1,
			setup: func() {
				suite.appService.On("UpdateApplicationById", mock.Anything, application.Application{
					ID:            1,
					JobID:         1,
					WorkerID:      1,
					Status:        "Pending",
					ExpectedWage:  1500,
					ModeOfArrival: "Personal",
					PickUpLocation: application.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411052,
					},
					WorkerComment: "some random comments",
					AppliedAt:     time.Time{},
					UpdatedAt:     time.Time{},
				}).Return(application.Application{
					ID:            1,
					JobID:         1,
					WorkerID:      1,
					Status:        "Pending",
					ExpectedWage:  1500,
					ModeOfArrival: "Personal",
					PickUpLocation: application.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411052,
					},
					WorkerComment: "some random comments",
					AppliedAt:     time.Time{},
					UpdatedAt:     time.Time{},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid req body",
			input:              "application Data",
			setup:              func() {},
			applicationId:      1,
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:          "db error",
			applicationId: 1,
			input: application.Application{
				ID:            1,
				JobID:         1,
				WorkerID:      1,
				Status:        "Pending",
				ExpectedWage:  1500,
				ModeOfArrival: "Personal",
				PickUpLocation: application.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411052,
				},
				WorkerComment: "some random comments",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			setup: func() {
				suite.appService.On("UpdateApplicationById", mock.Anything, application.Application{
					ID:            1,
					JobID:         1,
					WorkerID:      1,
					Status:        "Pending",
					ExpectedWage:  1500,
					ModeOfArrival: "Personal",
					PickUpLocation: application.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411052,
					},
					WorkerComment: "some random comments",
					AppliedAt:     time.Time{},
					UpdatedAt:     time.Time{},
				}).Return(application.Application{}, errors.New("db error while create application"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name: "invalid application id",
			input: application.Application{
				ID:            1,
				JobID:         1,
				WorkerID:      1,
				Status:        "Pending",
				ExpectedWage:  1500,
				ModeOfArrival: "Personal",
				PickUpLocation: application.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411052,
				},
				WorkerComment: "some random comments",
				AppliedAt:     time.Time{},
				UpdatedAt:     time.Time{},
			},
			applicationId:      "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/application/{application_id}", application.UpdateApplicationByID(&suite.appService)).Methods(http.MethodPut)

			reqBody, err := json.Marshal(test.input)
			if err != nil {
				logger.Errorw(context.Background(), "error while json marshal test data, error : "+err.Error())
			}

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/application/%v", test.applicationId), bytes.NewBuffer(reqBody))
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

func (suite *ApplicationHandlerTestSuite) TestFetchApplicationByID() {
	type testCase struct {
		name               string
		applicationId      interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:          "success",
			applicationId: 1,
			setup: func() {
				suite.appService.On("FetchApplicationById", mock.Anything, 1).Return(application.Application{
					ID:            1,
					JobID:         1,
					WorkerID:      1,
					Status:        "Pending",
					ExpectedWage:  1500,
					ModeOfArrival: "Personal",
					PickUpLocation: application.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411052,
					},
					WorkerComment: "some random comments",
					AppliedAt:     time.Time{},
					UpdatedAt:     time.Time{},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid application id",
			applicationId:      "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:          "no application exists",
			applicationId: 1,
			setup: func() {
				suite.appService.On("FetchApplicationById", mock.Anything, 1).Return(application.Application{}, apperrors.ErrNoApplicationExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:          "db error",
			applicationId: 1,
			setup: func() {
				suite.appService.On("FetchApplicationById", mock.Anything, 1).Return(application.Application{}, errors.New("db error while fetch application"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}
	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/application/{application_id}", application.FetchApplicationByID(&suite.appService)).Methods(http.MethodGet)

			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/application/%v", test.applicationId), bytes.NewBuffer([]byte(``)))
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

func (suite *ApplicationHandlerTestSuite) TestDeleteApplicationByID() {
	type testCase struct {
		name               string
		applicationId      interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:          "success",
			applicationId: 1,
			setup: func() {
				suite.appService.On("DeleteApplicationById", mock.Anything, 1).Return(1, nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:          "no application exists",
			applicationId: 1,
			setup: func() {
				suite.appService.On("DeleteApplicationById", mock.Anything, 1).Return(-1, apperrors.ErrNoApplicationExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:          "db error",
			applicationId: 1,
			setup: func() {
				suite.appService.On("DeleteApplicationById", mock.Anything, 1).Return(-1, errors.New("db error while delete application"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "invalid application id",
			applicationId:      "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	t := suite.T()
	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/application/{application_id}", application.DeleteApplicationByID(&suite.appService)).Methods(http.MethodDelete)

			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/application/%v", test.applicationId), bytes.NewBuffer([]byte(``)))
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

func (suite *ApplicationHandlerTestSuite) TestFetchAllApplications() {
	type testCase struct {
		name               string
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.appService.On("FetchAllApplications", mock.Anything).Return([]application.ApplicationComplete{
					{},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "db error",
			setup: func() {
				suite.appService.On("FetchAllApplications", mock.Anything).Return([]application.ApplicationComplete{}, errors.New("db error while fetch all applications"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	t := suite.T()
	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/applications", application.FetchAllApplications(&suite.appService)).Methods(http.MethodGet)

			req, err := http.NewRequest(http.MethodGet, "/applications", bytes.NewBuffer([]byte(``)))
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
