package worker

import (
	"context"
	"testing"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo/mocks"
	"github.com/stretchr/testify/suite"
)

type WorkerServiceTestSuite struct {
	suite.Suite
	service    Service
	workerRepo mocks.WorkerStorer
}

func (suite *WorkerServiceTestSuite) SetupTest() {
	suite.service = NewService(&mocks.WorkerStorer{})
	suite.workerRepo = mocks.WorkerStorer{}
}

func (suite *WorkerServiceTestSuite) TearDownTest() {
	suite.workerRepo.AssertExpectations(suite.T())
}

func TestWorkerServiceTestSuite(t *testing.T) {
	suite.Run(t, new(WorkerServiceTestSuite))
}

func (suite *WorkerServiceTestSuite) TestFetchWorkerById() {
	type testCase struct {
		name           string
		workeId        int
		setup          func()
		expectedOutput Worker
		expectedError  bool
	}

	testCases := []testCase{
		// {
		// 	name:    "success",
		// 	workeId: 1,
		// 	setup: func() {
		// 		suite.workerRepo.On("FetchWorkerByID", mock.Anything, 1).Return(repo.Worker{
		// 			ID:            1,
		// 			Name:          "Harsh",
		// 			ContactNumber: "9067691363",
		// 			Email:         "harsh@gmail.com",
		// 			Gender:        "Male",
		// 			Sectors:       "IT, Technology",
		// 			Password:      "",
		// 			Skills:        "React, Golang",
		// 			Location:      2,
		// 			CreatedAt:     time.Time{},
		// 			UpdatedAt:     time.Time{},
		// 			Language:      "English",
		// 			Details:       "details",
		// 			Street:        "street",
		// 			City:          "city",
		// 			State:         "state",
		// 			Pincode:       411025,
		// 		}, nil)
		// 	},
		// 	expectedOutput: Worker{
		// 		ID:            1,
		// 		Name:          "Harsh",
		// 		ContactNumber: "9067691363",
		// 		Email:         "harsh@gmail.com",
		// 		Gender:        "Male",
		// 		Sectors:       "IT, Technology",
		// 		Skills:        "React, Golang",
		// 		Location: Address{
		// 			Details: "details",
		// 			Street:  "street",
		// 			City:    "city",
		// 			State:   "state",
		// 			Pincode: 411025,
		// 		},
		// 		IsAvailable:     true,
		// 		Rating:          0,
		// 		TotalJobsWorked: 0,
		// 		CreatedAt:       time.Time{},
		// 		UpdatedAt:       time.Time{},
		// 		Language:        "English",
		// 	},
		// 	expectedError: false,
		// },
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()
			worker, err := suite.service.FetchWorkerByID(context.Background(), test.workeId)
			suite.Equal(test.expectedOutput, worker)
			suite.Equal(test.expectedError, err != nil)
		})
		suite.TearDownTest()
	}
}

// func (suite WorkerServiceTestSuite) TestDeleteWorkerByID () {
// 	type testCase struct {
// 		name           string
// 		workeId        int
// 		setup          func()
// 		expectedOutput Worker
// 		expectedError  bool
// 	}

// }
