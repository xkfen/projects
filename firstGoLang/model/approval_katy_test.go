package model

import (
	"testing"
	"gcoresys/common/logger"
	"github.com/stretchr/testify/suite"
	"firstGoLang/db/config"
)

type iHttpReqSuite struct {
	suite.Suite
}

func (s *iHttpReqSuite) SetupTest() {
}

func (s *iHttpReqSuite) TearDownTest() {
	config.ClearAllData()
}

func TestRun(t *testing.T){
	logger.InitLogger(logger.LvlDebug, nil)
	config.GetApprovalDbConfig("test")
	//config.GetDb().LogMode(true)
	//config.GetApprovalDbConfig("test")
	//config.GetDb().LogMode(false)
	suite.Run(t, &iHttpReqSuite{})
}

func (a *iHttpReqSuite) TestO() {

}

