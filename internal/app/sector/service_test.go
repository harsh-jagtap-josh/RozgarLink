package sector

import (
	"context"
	"errors"
	"testing"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/pkg/apperrors"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ServiceTestSuite struct {
	suite.Suite
	service    Service
	sectorRepo mocks.SectoreStorer
}

// this function executes before the test suite begins execution
func (suite *ServiceTestSuite) SetupTest() {
	suite.service = NewService(&suite.sectorRepo)
	suite.sectorRepo = mocks.SectoreStorer{}
}

// this function executes after all tests executed
func (suite *ServiceTestSuite) TearDownTest() {
	suite.sectorRepo.AssertExpectations(suite.T())
}

func (suite *ServiceTestSuite) TestGetAllSectors() {
	type testCase struct {
		name           string
		setup          func()
		expectedOutput []Sector
		expectedError  bool
	}
	testCases := []testCase{
		{
			name: "success",
			setup: func() {
				suite.sectorRepo.On("FetchAllSectors", mock.Anything, mock.Anything).Return([]repo.Sector{
					{
						ID:          1,
						Name:        "IT",
						Description: "Information Technology",
					},
					{
						ID:          2,
						Name:        "Healthcare",
						Description: "Healthcare",
					},
				}, nil)
			},
			expectedOutput: []Sector{
				{
					ID:          1,
					Name:        "IT",
					Description: "Information Technology",
				},
				{
					ID:          2,
					Name:        "Healthcare",
					Description: "Healthcare",
				},
			},
			expectedError: false,
		}, {
			name: "error",
			setup: func() {
				suite.sectorRepo.On("FetchAllSectors", mock.Anything, mock.Anything).Return([]repo.Sector{}, errors.New("some db error"))
			},
			expectedOutput: []Sector{},
			expectedError:  true,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			sectors, err := suite.service.FetchAllSectors(context.Background())
			if test.expectedError {
				suite.Error(err)
			} else {
				suite.NoError(err)
				suite.Equal(test.expectedOutput, sectors)
			}
		})
		suite.TearDownTest()
	}
}

func (suite *ServiceTestSuite) TestFetchSectorById() {
	type testCase struct {
		name           string
		orderId        int
		setup          func()
		expectedOutput Sector
		expectedError  error
	}
	testCases := []testCase{
		{
			name:    "fetch sector by id success",
			orderId: 1,
			setup: func() {
				suite.sectorRepo.On("FetchSectorById", mock.Anything, 1).Return(repo.Sector{
					ID:          1,
					Name:        "IT",
					Description: "Information Technology",
				}, nil)
			},
			expectedOutput: Sector{
				ID:          1,
				Name:        "IT",
				Description: "Information Technology",
			},
			expectedError: nil,
		}, {
			name:    "fetch sector by id error",
			orderId: 1,
			setup: func() {
				suite.sectorRepo.On("FetchSectorById", mock.Anything, 1).Return(repo.Sector{}, errors.New("db fetch error"))
			},
			expectedOutput: Sector{},
			expectedError:  errors.New("db fetch error"),
		}, {
			name:    "sector by id not found",
			orderId: 1,
			setup: func() {
				suite.sectorRepo.On("FetchSectorById", mock.Anything, 1).Return(repo.Sector{}, apperrors.ErrNoSectorExists)
			},
			expectedOutput: Sector{},
			expectedError:  apperrors.ErrNoSectorExists,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			sector, err := suite.service.FetchSectorById(context.Background(), test.orderId)
			suite.Equal(test.expectedOutput, sector)
			suite.Equal(test.expectedError, err)
		})
		suite.TearDownTest()
	}
}

func (suite *ServiceTestSuite) TestCreateNewSector() {
	type testCase struct {
		name           string
		sectorData     Sector
		setup          func()
		expectedOutput Sector
		expectedError  error
	}
	testCases := []testCase{
		{
			name: "create new sector success",
			sectorData: Sector{
				Name:        "IT",
				Description: "Information Technology",
			},
			setup: func() {
				suite.sectorRepo.On("CreateNewSector", mock.Anything, mock.Anything).Return(repo.Sector{
					ID:          1,
					Name:        "IT",
					Description: "Information Technology",
				}, nil)
			},
			expectedOutput: Sector{
				ID:          1,
				Name:        "IT",
				Description: "Information Technology",
			},
			expectedError: nil,
		}, {
			name: "create new sector error",
			sectorData: Sector{
				Name:        "IT",
				Description: "Information Technology",
			},
			setup: func() {
				suite.sectorRepo.On("CreateNewSector", mock.Anything, mock.Anything).Return(repo.Sector{}, errors.New("db create error"))
			},
			expectedOutput: Sector{},
			expectedError:  errors.New("db create error"),
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			sector, err := suite.service.CreateNewSector(context.Background(), test.sectorData)
			suite.Equal(test.expectedOutput, sector)
			suite.Equal(test.expectedError, err)
		})
		suite.TearDownTest()
	}
}

func (suite *ServiceTestSuite) TestUpdateSectorById() {
	type testCase struct {
		name           string
		sectorData     Sector
		setup          func()
		expectedOutput Sector
		expectedError  error
	}
	testCases := []testCase{
		{
			name: "update sector success",
			sectorData: Sector{
				ID:          1,
				Name:        "IT",
				Description: "Information Technology",
			},
			setup: func() {
				suite.sectorRepo.On("UpdateSectorById", mock.Anything, mock.Anything).Return(repo.Sector{
					ID:          1,
					Name:        "IT",
					Description: "Information Technology",
				}, nil)
			},
			expectedOutput: Sector{
				ID:          1,
				Name:        "IT",
				Description: "Information Technology",
			},
			expectedError: nil,
		}, {
			name: "update sector error",
			sectorData: Sector{
				ID:          1,
				Name:        "IT",
				Description: "Information Technology",
			},
			setup: func() {
				suite.sectorRepo.On("UpdateSectorById", mock.Anything, mock.Anything).Return(repo.Sector{}, errors.New("db update error"))
			},
			expectedOutput: Sector{},
			expectedError:  errors.New("db update error"),
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			sector, err := suite.service.UpdateSectorById(context.Background(), test.sectorData)
			suite.Equal(test.expectedOutput, sector)
			suite.Equal(test.expectedError, err)
		})
		suite.TearDownTest()
	}
}

func (suite *ServiceTestSuite) TestDeleteSectorById() {
	type testCase struct {
		name           string
		sectorId       int
		setup          func()
		expectedOutput int
		expectedError  error
	}
	testCases := []testCase{
		{
			name:     "delete sector success",
			sectorId: 1,
			setup: func() {
				suite.sectorRepo.On("DeleteSectorById", mock.Anything, 1).Return(1, nil)
			},
			expectedOutput: 1,
			expectedError:  nil,
		}, {
			name:     "delete sector error",
			sectorId: 1,
			setup: func() {
				suite.sectorRepo.On("DeleteSectorById", mock.Anything, 1).Return(-1, errors.New("db delete error"))
			},
			expectedOutput: -1,
			expectedError:  errors.New("db delete error"),
		}, {
			name:     "sector not found",
			sectorId: 1,
			setup: func() {
				suite.sectorRepo.On("DeleteSectorById", mock.Anything, 1).Return(-1, apperrors.ErrNoSectorExists)
			},
			expectedOutput: -1,
			expectedError:  apperrors.ErrNoSectorExists,
		},
	}

	for _, test := range testCases {
		suite.SetupTest()
		suite.Run(test.name, func() {
			test.setup()

			sector, err := suite.service.DeleteSectorById(context.Background(), test.sectorId)
			suite.Equal(test.expectedOutput, sector)
			suite.Equal(test.expectedError, err)
		})
		suite.TearDownTest()
	}
}

func TestSectorServciceTestSuite(t *testing.T) {
	suite.Run(t, new(ServiceTestSuite))
}
