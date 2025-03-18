package sector_test

import (
	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/sector/mocks"
	"github.com/stretchr/testify/suite"
)

type SectorHandlerTestSuite struct {
	suite.Suite
	sectorService *mocks.Service
	router        mux.Router
}

