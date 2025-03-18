package employer_test

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
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/employer"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/employer/mocks"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type EmployerHandlerTestSuite struct {
	suite.Suite
	empService mocks.Service
	router     mux.Router
}

func (suite *EmployerHandlerTestSuite) SetupTest() {
	suite.empService = mocks.Service{}
	suite.router = *mux.NewRouter()
}

func (suite *EmployerHandlerTestSuite) TearDownTest() {
	suite.empService.AssertExpectations(suite.T())
}

func TestEmployerHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(EmployerHandlerTestSuite))
}

func (suite *EmployerHandlerTestSuite) TestFetchEmployerById() {
	t := suite.T()
	type testCase struct {
		name               string
		employer_id        interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:        "success",
			employer_id: 1,
			setup: func() {
				suite.empService.On("FetchEmployerByID", mock.Anything, 1).Return(employer.Employer{
					ID:        1,
					Name:      "John Doe",
					ContactNo: "9037691363",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid employer id",
			employer_id:        "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "db error",
			employer_id: 1,
			setup: func() {
				suite.empService.On("FetchEmployerByID", mock.Anything, 1).Return(employer.Employer{}, errors.New("some random error while fetch employer from db"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "no employer exists with id",
			employer_id: 1,
			setup: func() {
				suite.empService.On("FetchEmployerByID", mock.Anything, 1).Return(employer.Employer{}, apperrors.ErrNoEmployerExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/employer/{employer_id}", employer.FetchEmployerByID(&suite.empService)).Methods(http.MethodGet)
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/employer/%v", test.employer_id), bytes.NewBuffer([]byte(``)))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
	}
}

func (suite *EmployerHandlerTestSuite) TestDeleteEmployerByID() {
	type testCase struct {
		name               string
		employer_id        interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:        "success",
			employer_id: 1,
			setup: func() {
				suite.empService.On("DeleteEmployerById", mock.Anything, 1).Return(1, nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:               "invalid employer id",
			employer_id:        "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "db error",
			employer_id: 1,
			setup: func() {
				suite.empService.On("DeleteEmployerById", mock.Anything, 1).Return(-1, errors.New("some random error while fetch employer from db"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:        "no employer exists with id",
			employer_id: 1,
			setup: func() {
				suite.empService.On("DeleteEmployerById", mock.Anything, 1).Return(-1, apperrors.ErrNoEmployerExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/employer/{employer_id}", employer.DeleteEmployerByID(&suite.empService)).Methods(http.MethodDelete)
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/employer/%v", test.employer_id), bytes.NewBuffer([]byte(``)))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
	}
}

func (suite *EmployerHandlerTestSuite) TestFetchJobsByEmployerId() {
	type testCase struct {
		name               string
		employer_id        interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:        "success",
			employer_id: 1,
			setup: func() {
				suite.empService.On("FetchJobsByEmployerId", mock.Anything, 1).Return([]job.Job{
					{
						ID:              1,
						EmployerID:      1,
						Title:           "Some Random Title",
						RequiredGender:  "Male",
						Description:     "some random description",
						DurationInHours: 10,
						SkillsRequired:  "some random skills",
						Sectors:         "sectors",
						Wage:            2000,
						Vacancy:         4,
						Location: worker.Address{
							ID:      1,
							Details: "Details",
							Street:  "Street",
							City:    "city",
							State:   "state",
							Pincode: 411025,
						},
						Date:      "",
						StartHour: "",
						EndHour:   "",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					},
					{
						ID:              2,
						EmployerID:      1,
						Title:           "Some Random Title",
						RequiredGender:  "Male",
						Description:     "some random description",
						DurationInHours: 10,
						SkillsRequired:  "some random skills",
						Sectors:         "sectors",
						Wage:            2000,
						Vacancy:         4,
						Location: worker.Address{
							ID:      1,
							Details: "Details",
							Street:  "Street",
							City:    "city",
							State:   "state",
							Pincode: 411025,
						},
						Date:      "",
						StartHour: "",
						EndHour:   "",
						CreatedAt: time.Time{},
						UpdatedAt: time.Time{},
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:        "employer doesn't exist",
			employer_id: 1,
			setup: func() {
				suite.empService.On("FetchJobsByEmployerId", mock.Anything, 1).Return([]job.Job{}, apperrors.ErrNoEmployerExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:        "invalid employer id",
			employer_id: "a",
			setup: func() {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:        "db error",
			employer_id: 1,
			setup: func() {
				suite.empService.On("FetchJobsByEmployerId", mock.Anything, 1).Return([]job.Job{}, errors.New("error occured while fetch jobs"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/employer/{employer_id}/jobs", employer.FetchJobsByEmployerId(&suite.empService)).Methods(http.MethodGet)
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/employer/%v/jobs", test.employer_id), bytes.NewBuffer([]byte(``)))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
	}
}

func (suite *EmployerHandlerTestSuite) TestFetchAllEmployers() {
	type testCase struct {
		name               string
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.empService.On("FetchAllEmployers", mock.Anything).Return([]employer.Employer{
					{
						ID:        1,
						Name:      "John Doe",
						ContactNo: "9037691363",
						Email:     "john@gmail.com",
						Type:      "Employer",
						Sectors:   "sectors",
						Location: employer.Address{
							ID:      1,
							Details: "details",
							Street:  "street",
							City:    "city",
							State:   "state",
							Pincode: 411051,
						},
						IsVerified:   true,
						Rating:       0,
						WorkersHired: 0,
						CreatedAt:    time.Time{},
						UpdatedAt:    time.Time{},
						Language:     "English",
					},
					{
						ID:        1,
						Name:      "John Doe",
						ContactNo: "9037691363",
						Email:     "john@gmail.com",
						Type:      "Employer",
						Sectors:   "sectors",
						Location: employer.Address{
							ID:      1,
							Details: "details",
							Street:  "street",
							City:    "city",
							State:   "state",
							Pincode: 411051,
						},
						IsVerified:   true,
						Rating:       0,
						WorkersHired: 0,
						CreatedAt:    time.Time{},
						UpdatedAt:    time.Time{},
						Language:     "English",
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name: "db error",
			setup: func() {
				suite.empService.On("FetchAllEmployers", mock.Anything).Return([]employer.Employer{}, errors.New("error while fetch all employers"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/employers", employer.FetchAllEmployers(&suite.empService)).Methods(http.MethodGet)
			req, err := http.NewRequest(http.MethodGet, "/employers", bytes.NewBuffer([]byte(``)))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
	}
}

func (suite *EmployerHandlerTestSuite) TestUpdateEmployerById() {
	type testCase struct {
		name               string
		employerData       interface{}
		employer_id        interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			employerData: employer.Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9037691363",
				Email:     "john@gmail.com",
				Type:      "Employer",
				Sectors:   "sectors",
				Location: employer.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			employer_id: 1,
			setup: func() {
				suite.empService.On("UpdateEmployerById", mock.Anything, employer.Employer{
					ID:        1,
					Name:      "John Doe",
					ContactNo: "9037691363",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}).Return(employer.Employer{
					ID:        1,
					Name:      "John Doe",
					ContactNo: "9037691363",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid request body",
			employerData:       "employer data",
			employer_id:        1,
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "invalid employer id",
			employerData: employer.Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9037691363",
				Email:     "john@gmail.com",
				Type:      "Employer",
				Sectors:   "sectors",
				Location: employer.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			employer_id:        "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "user validation failed",
			employerData: employer.Employer{
				ID:        1,
				Name:      "John Doe1",
				ContactNo: "90376913",
				Email:     "john@gmail.com",
				Type:      "Employer",
				Sectors:   "sectors",
				Location: employer.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			employer_id: 1,
			setup: func() {
				suite.empService.On("UpdateEmployerById", mock.Anything, employer.Employer{
					ID:        1,
					Name:      "John Doe1",
					ContactNo: "90376913",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}).Return(employer.Employer{}, apperrors.ErrInvalidUserDetails)
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "db error",
			employerData: employer.Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9037691363",
				Email:     "john@gmail.com",
				Type:      "Employer",
				Sectors:   "sectors",
				Location: employer.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			employer_id: 1,
			setup: func() {
				suite.empService.On("UpdateEmployerById", mock.Anything, employer.Employer{
					ID:        1,
					Name:      "John Doe",
					ContactNo: "9037691363",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}).Return(employer.Employer{}, errors.New("db error while update employer"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "invalid request body",
			employerData:       "employerData",
			employer_id:        1,
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/employer/{employer_id}", employer.UpdateEmployerById(&suite.empService)).Methods(http.MethodPut)
			reqBody, err := json.Marshal(test.employerData)
			if err != nil {
				logger.Errorw(context.Background(), "error faced while json marshal request body")
			}

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/employer/%v", test.employer_id), bytes.NewBuffer(reqBody))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
	}
}

func (suite *EmployerHandlerTestSuite) TestRegisterEmployer() {
	type testCase struct {
		name               string
		employerData       interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			employerData: employer.Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9037691363",
				Email:     "john@gmail.com",
				Type:      "Employer",
				Sectors:   "sectors",
				Location: employer.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			setup: func() {
				suite.empService.On("RegisterEmployer", mock.Anything, employer.Employer{
					ID:        1,
					Name:      "John Doe",
					ContactNo: "9037691363",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}).Return(employer.Employer{
					ID:        1,
					Name:      "John Doe",
					ContactNo: "9037691363",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}, nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:               "invalid request body",
			employerData:       "employer data",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "employer with email already exists",
			employerData: employer.Employer{
				ID:        1,
				Name:      "John Doe1",
				ContactNo: "90376913",
				Email:     "john@gmail.com",
				Type:      "Employer",
				Sectors:   "sectors",
				Location: employer.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			setup: func() {
				suite.empService.On("RegisterEmployer", mock.Anything, employer.Employer{
					ID:        1,
					Name:      "John Doe1",
					ContactNo: "90376913",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}).Return(employer.Employer{}, apperrors.ErrEmployerAlreadyExists)
			},
			expectedStatusCode: http.StatusConflict,
		},
		{
			name: "user validation failed",
			employerData: employer.Employer{
				ID:        1,
				Name:      "John Doe1",
				ContactNo: "90376913",
				Email:     "john@gmail.com",
				Type:      "Employer",
				Sectors:   "sectors",
				Location: employer.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			setup: func() {
				suite.empService.On("RegisterEmployer", mock.Anything, employer.Employer{
					ID:        1,
					Name:      "John Doe1",
					ContactNo: "90376913",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}).Return(employer.Employer{}, apperrors.ErrInvalidUserDetails)
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name: "db error",
			employerData: employer.Employer{
				ID:        1,
				Name:      "John Doe",
				ContactNo: "9037691363",
				Email:     "john@gmail.com",
				Type:      "Employer",
				Sectors:   "sectors",
				Location: employer.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
				IsVerified:   true,
				Rating:       0,
				WorkersHired: 0,
				CreatedAt:    time.Time{},
				UpdatedAt:    time.Time{},
				Language:     "English",
			},
			setup: func() {
				suite.empService.On("RegisterEmployer", mock.Anything, employer.Employer{
					ID:        1,
					Name:      "John Doe",
					ContactNo: "9037691363",
					Email:     "john@gmail.com",
					Type:      "Employer",
					Sectors:   "sectors",
					Location: employer.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
					IsVerified:   true,
					Rating:       0,
					WorkersHired: 0,
					CreatedAt:    time.Time{},
					UpdatedAt:    time.Time{},
					Language:     "English",
				}).Return(employer.Employer{}, errors.New("db error while update employer"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "invalid request body",
			employerData:       "employerData",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	t := suite.T()

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/register/employer", employer.RegisterEmployer(&suite.empService)).Methods(http.MethodPost)
			reqBody, err := json.Marshal(test.employerData)
			if err != nil {
				logger.Errorw(context.Background(), "error faced while json marshal request body")
			}

			req, err := http.NewRequest(http.MethodPost, "/register/employer", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Errorf("error occured while making http request, error : %v", err.Error())
			}

			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
	}
}
