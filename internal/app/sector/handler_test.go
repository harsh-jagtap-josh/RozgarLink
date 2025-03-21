package sector_test

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/sector/mocks"
	"github.com/stretchr/testify/suite"
)

type SectorHandlerTestSuite struct {
	suite.Suite
	sectorService *mocks.Service
	router        mux.Router
}

func (suite *SectorHandlerTestSuite) SetupTest() {
	suite.sectorService = &mocks.Service{}
	suite.router = *mux.NewRouter()
}

func (suite *SectorHandlerTestSuite) TearDownTest() {
	suite.sectorService.AssertExpectations(suite.T())
}

func TestApplicationTestSuite(t *testing.T) {
	suite.Run(t, new(SectorHandlerTestSuite))
}

// func (suite *SectorHandlerTestSuite) TestCreateSector() {
// 	type testCase struct {
// 		name               string
// 		input              interface{}
// 		setup              func()
// 		expectedStatusCode int
// 	}

// 	testCases := []testCase{
// 		{
// 			name: "success",
// 			input: sector.Sector{
// 				Name:        "Sector Name",
// 				Description: "Sector Description",
// 			},
// 			setup: func() {
// 				suite.sectorService.On("CreateNewSector", mock.Anything, sector.Sector{
// 					Name:        "Sector Name",
// 					Description: "Sector Description",
// 				}).Return(sector.Sector{
// 					Name:        "Sector Name",
// 					Description: "Sector Description",
// 				}, nil)
// 			},
// 			expectedStatusCode: http.StatusCreated,
// 		},
// 	}

// 	t := suite.T()

// 	for _, test := range testCases {
// 		suite.SetupTest()
// 		suite.Run(test.name, func() {
// 			test.setup()

// 			suite.router.HandleFunc("/sector/create", sector.CreateSector(suite.sectorService)).Methods(http.MethodPost)

// 			reqBody, err := json.Marshal(test.input)
// 			if err != nil {
// 				logger.Errorw(context.Background(), "error while json marshal test data, error : "+err.Error())
// 			}

// 			req, err := http.NewRequest(http.MethodPost, "/sector/create", bytes.NewBuffer(reqBody))
// 			if err != nil {
// 				t.Errorf("error occured while making http request, error : %v", err.Error())
// 			}

// 			recorder := httptest.NewRecorder()
// 			suite.router.ServeHTTP(recorder, req)

// 			suite.Equal(test.expectedStatusCode, recorder.Code)
// 		})
// 		suite.TearDownTest()
// 	}
// }
