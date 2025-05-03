//go:build integration
// +build integration

package sql

import (
	"context"
	"testing"

	"github.com/cryptellation/candlesticks/configs"
	"github.com/cryptellation/candlesticks/configs/sql/down"
	"github.com/cryptellation/candlesticks/configs/sql/up"
	"github.com/cryptellation/candlesticks/svc/db"
	"github.com/cryptellation/dbmigrator"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/suite"
)

func TestCandlesticksSuite(t *testing.T) {
	suite.Run(t, new(CandlesticksSuite))
}

type CandlesticksSuite struct {
	db.CandlesticksSuite
}

func (suite *CandlesticksSuite) SetupSuite() {
	act, err := New(context.Background(), viper.GetString(configs.EnvSQLDSN))
	suite.Require().NoError(err)

	mig, err := dbmigrator.NewMigrator(context.Background(), act.db, up.Migrations, down.Migrations, nil)
	suite.Require().NoError(err)
	suite.Require().NoError(mig.MigrateToLatest(context.Background()))

	suite.DB = act
}

func (suite *CandlesticksSuite) SetupTest() {
	db := suite.DB.(*Activities)
	suite.Require().NoError(db.Reset(context.Background()))
}
