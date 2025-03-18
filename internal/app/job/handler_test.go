package job_test

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
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/job/mocks"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/logger"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type JobHandlerTestSuite struct {
	suite.Suite
	jobService mocks.Service
	router     mux.Router
}

func (suite *JobHandlerTestSuite) SetupTest() {
	suite.jobService = mocks.Service{}
	suite.router = *mux.NewRouter()
}

func (suite *JobHandlerTestSuite) TearDownTest() {
	suite.jobService.AssertExpectations(suite.T())
}

func TestJobHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(JobHandlerTestSuite))
}

func (suite *JobHandlerTestSuite) TestFetchJobByID() {
	t := suite.T()
	type testCase struct {
		name               string
		job_id             interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:   "success",
			job_id: 1,
			setup: func() {
				suite.jobService.On("FetchJobByID", mock.Anything, 1).Return(job.Job{
					ID:              1,
					EmployerID:      1,
					Title:           "Random Job Title",
					RequiredGender:  "Male",
					Description:     "some desccription",
					DurationInHours: 10,
					SkillsRequired:  "random",
					Sectors:         "",
					Wage:            1200,
					Vacancy:         5,
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:   "invalid job id",
			job_id: "abc",
			setup: func() {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:   "no job found",
			job_id: 1,
			setup: func() {
				suite.jobService.On("FetchJobByID", mock.Anything, 1).Return(job.Job{}, apperrors.ErrNoJobExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:   "error from db",
			job_id: 1,
			setup: func() {
				suite.jobService.On("FetchJobByID", mock.Anything, 1).Return(job.Job{}, errors.New("error while fetch job by id"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/job/{job_id}", job.FetchJobByID(&suite.jobService)).Methods(http.MethodGet)
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/job/%v", test.job_id), bytes.NewBuffer([]byte(``)))
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

func (suite *JobHandlerTestSuite) TestFetchApplicationsByJobId() {
	t := suite.T()
	type testCase struct {
		name               string
		job_id             interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:   "success",
			job_id: 1,
			setup: func() {
				suite.jobService.On("FetchApplicationsByJobId", mock.Anything, 1).Return([]application.ApplicationCompleteEmp{
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
						Pincode:        411025,
						JobTitle:       "Random Title",
						Description:    "description",
						SkillsRequired: "skills",
						JobSectors:     "sectors",
						JobWage:        1500,
						Vacancy:        5,
						JobDate:        "",
						WorkerName:     "John Doe",
						ContactNumber:  "9067691363",
						WorkerEmail:    "john@gmail.com",
						WorkerGender:   "Male",
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
						Pincode:        411025,
						JobTitle:       "Random Title",
						Description:    "description",
						SkillsRequired: "skills",
						JobSectors:     "sectors",
						JobWage:        1500,
						Vacancy:        5,
						JobDate:        "",
						WorkerName:     "John Doe",
						ContactNumber:  "9067691363",
						WorkerEmail:    "john@gmail.com",
						WorkerGender:   "Male",
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:   "no applications exist",
			job_id: 1,
			setup: func() {
				suite.jobService.On("FetchApplicationsByJobId", mock.Anything, 1).Return([]application.ApplicationCompleteEmp{}, apperrors.ErrNoJobExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:   "db error",
			job_id: 1,
			setup: func() {
				suite.jobService.On("FetchApplicationsByJobId", mock.Anything, 1).Return([]application.ApplicationCompleteEmp{}, errors.New("db error while fetch applications"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:   "invalid job id",
			job_id: "abc",
			setup: func() {
			},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/job/{job_id}/applications", job.FetchApplicationsByJobId(&suite.jobService))
			req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/job/%v/applications", test.job_id), bytes.NewBuffer([]byte(``)))
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

func (suite *JobHandlerTestSuite) TestFetchAllJobs() {
	t := suite.T()
	type testCase struct {
		name               string
		setup              func()
		urlParams          string
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.jobService.On("FetchAllJobs", mock.Anything, job.JobFilters{}).Return([]job.Job{
					{
						ID:              1,
						EmployerID:      1,
						Title:           "Random Job Title",
						RequiredGender:  "Male",
						Description:     "some desccription",
						DurationInHours: 10,
						SkillsRequired:  "random",
						Sectors:         "",
						Wage:            1200,
						Vacancy:         5,
						Location: worker.Address{
							ID:      1,
							Details: "details",
							Street:  "street",
							City:    "city",
							State:   "state",
							Pincode: 411051,
						},
					},
					{
						ID:              1,
						EmployerID:      1,
						Title:           "Random Job Title",
						RequiredGender:  "Male",
						Description:     "some desccription",
						DurationInHours: 10,
						SkillsRequired:  "random",
						Sectors:         "",
						Wage:            1200,
						Vacancy:         5,
						Location: worker.Address{
							ID:      1,
							Details: "details",
							Street:  "street",
							City:    "city",
							State:   "state",
							Pincode: 411051,
						},
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
			urlParams:          "",
		},
		{
			name:      "filters",
			urlParams: "?title=Construction&sector=IT&wage_min=1200&wage_max=1500&start_date=2024-10-09&end_date=2024-10-09&city=Pune&required_gender=Male",
			setup: func() {
				startParsed, _ := time.Parse("2006-01-02", "2024-10-09")
				endParsed, _ := time.Parse("2006-01-02", "2024-10-09")
				suite.jobService.On("FetchAllJobs", mock.Anything, job.JobFilters{
					Title:     "Construction",
					Sector:    "IT",
					WageMin:   1200,
					WageMax:   1500,
					StartDate: startParsed,
					EndDate:   endParsed,
					City:      "Pune",
					Gender:    "Male",
				}).Return([]job.Job{
					{
						ID:              1,
						EmployerID:      1,
						Title:           "Random Job Title",
						RequiredGender:  "Male",
						Description:     "some desccription",
						DurationInHours: 10,
						SkillsRequired:  "random",
						Sectors:         "",
						Wage:            1200,
						Vacancy:         5,
						Location: worker.Address{
							ID:      1,
							Details: "details",
							Street:  "street",
							City:    "city",
							State:   "state",
							Pincode: 411051,
						},
					},
					{
						ID:              1,
						EmployerID:      1,
						Title:           "Random Job Title",
						RequiredGender:  "Male",
						Description:     "some desccription",
						DurationInHours: 10,
						SkillsRequired:  "random",
						Sectors:         "",
						Wage:            1200,
						Vacancy:         5,
						Location: worker.Address{
							ID:      1,
							Details: "details",
							Street:  "street",
							City:    "city",
							State:   "state",
							Pincode: 411051,
						},
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:      "db error",
			urlParams: "",
			setup: func() {
				suite.jobService.On("FetchAllJobs", mock.Anything, job.JobFilters{}).Return([]job.Job{}, errors.New("db error while fetch all jobs"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/job/all", job.FetchAllJobs(&suite.jobService)).Methods(http.MethodGet)
			req, err := http.NewRequest(http.MethodGet, "/job/all"+test.urlParams, bytes.NewBuffer([]byte(``)))
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

func (suite *JobHandlerTestSuite) TestDeleteJobByID() {
	t := suite.T()
	type testCase struct {
		name               string
		job_id             interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:   "success",
			job_id: 1,
			setup: func() {
				suite.jobService.On("DeleteJobByID", mock.Anything, 1).Return(1, nil)
			},
			expectedStatusCode: http.StatusNoContent,
		},
		{
			name:   "job doesn't exist",
			job_id: 1,
			setup: func() {
				suite.jobService.On("DeleteJobByID", mock.Anything, 1).Return(-1, apperrors.ErrNoJobExists)
			},
			expectedStatusCode: http.StatusNotFound,
		},
		{
			name:   "job doesn't exist",
			job_id: 1,
			setup: func() {
				suite.jobService.On("DeleteJobByID", mock.Anything, 1).Return(-1, errors.New("db error while delete job"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "invalid job id",
			job_id:             "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/job/{job_id}", job.DeleteJobByID(&suite.jobService)).Methods(http.MethodDelete)
			req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/job/%v", test.job_id), bytes.NewBuffer([]byte(``)))
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

func (suite *JobHandlerTestSuite) TestUpdateJob() {
	t := suite.T()
	type testCase struct {
		name               string
		jobData            interface{}
		jobId              interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name:  "success",
			jobId: 1,
			jobData: job.Job{
				ID:              1,
				EmployerID:      1,
				Title:           "Random Job Title",
				RequiredGender:  "Male",
				Description:     "some desccription",
				DurationInHours: 10,
				SkillsRequired:  "random",
				Sectors:         "",
				Wage:            1200,
				Vacancy:         5,
				Location: worker.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
			},
			setup: func() {
				suite.jobService.On("UpdateJobByID", mock.Anything, job.Job{
					ID:              1,
					EmployerID:      1,
					Title:           "Random Job Title",
					RequiredGender:  "Male",
					Description:     "some desccription",
					DurationInHours: 10,
					SkillsRequired:  "random",
					Sectors:         "",
					Wage:            1200,
					Vacancy:         5,
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
				}).Return(job.Job{
					ID:              1,
					EmployerID:      1,
					Title:           "Random Job Title",
					RequiredGender:  "Male",
					Description:     "some desccription",
					DurationInHours: 10,
					SkillsRequired:  "random",
					Sectors:         "",
					Wage:            1200,
					Vacancy:         5,
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
				}, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
		{
			name:               "invalid job id provided",
			jobData:            job.Job{},
			jobId:              "a",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
		{
			name:    "db error",
			jobData: job.Job{},
			jobId:   1,
			setup: func() {
				suite.jobService.On("UpdateJobByID", mock.Anything, job.Job{}).Return(job.Job{}, errors.New("error faced while update job by id"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "invalid request body",
			jobData:            "job data",
			jobId:              1,
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/job/{job_id}", job.UpdateJobById(&suite.jobService)).Methods(http.MethodPut)
			reqBody, err := json.Marshal(test.jobData)
			if err != nil {
				logger.Errorw(context.Background(), "facing error while json marshal input data")
			}

			req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/job/%v", test.jobId), bytes.NewBuffer(reqBody))
			if err != nil {
				t.Errorf("error faced while making an http request, error: %v", err.Error())
			}
			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
		suite.TearDownTest()
	}

}

func (suite *JobHandlerTestSuite) TestCreateJob() {
	t := suite.T()
	type testCase struct {
		name               string
		jobData            interface{}
		setup              func()
		expectedStatusCode int
	}

	testCases := []testCase{
		{
			name: "success",
			jobData: job.Job{
				ID:              1,
				EmployerID:      1,
				Title:           "Random Job Title",
				RequiredGender:  "Male",
				Description:     "some desccription",
				DurationInHours: 10,
				SkillsRequired:  "random",
				Sectors:         "",
				Wage:            1200,
				Vacancy:         5,
				Location: worker.Address{
					ID:      1,
					Details: "details",
					Street:  "street",
					City:    "city",
					State:   "state",
					Pincode: 411051,
				},
			},
			setup: func() {
				suite.jobService.On("CreateJob", mock.Anything, job.Job{
					ID:              1,
					EmployerID:      1,
					Title:           "Random Job Title",
					RequiredGender:  "Male",
					Description:     "some desccription",
					DurationInHours: 10,
					SkillsRequired:  "random",
					Sectors:         "",
					Wage:            1200,
					Vacancy:         5,
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
				}).Return(job.Job{
					ID:              1,
					EmployerID:      1,
					Title:           "Random Job Title",
					RequiredGender:  "Male",
					Description:     "some desccription",
					DurationInHours: 10,
					SkillsRequired:  "random",
					Sectors:         "",
					Wage:            1200,
					Vacancy:         5,
					Location: worker.Address{
						ID:      1,
						Details: "details",
						Street:  "street",
						City:    "city",
						State:   "state",
						Pincode: 411051,
					},
				}, nil)
			},
			expectedStatusCode: http.StatusCreated,
		},
		{
			name:    "db error",
			jobData: job.Job{},
			setup: func() {
				suite.jobService.On("CreateJob", mock.Anything, job.Job{}).Return(job.Job{}, errors.New("error faced while update job by id"))
			},
			expectedStatusCode: http.StatusInternalServerError,
		},
		{
			name:               "invalid request body",
			jobData:            "job data",
			setup:              func() {},
			expectedStatusCode: http.StatusBadRequest,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			suite.router.HandleFunc("/job/create", job.CreateJob(&suite.jobService)).Methods(http.MethodPost)
			reqBody, err := json.Marshal(test.jobData)
			if err != nil {
				logger.Errorw(context.Background(), "facing error while json marshal input data")
			}

			req, err := http.NewRequest(http.MethodPost, "/job/create", bytes.NewBuffer(reqBody))
			if err != nil {
				t.Errorf("error faced while making an http request, error: %v", err.Error())
			}
			recorder := httptest.NewRecorder()
			suite.router.ServeHTTP(recorder, req)

			suite.Equal(test.expectedStatusCode, recorder.Code)
		})
		suite.TearDownTest()
	}

}
