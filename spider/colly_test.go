package spider

import (
	"context"
	"testing"

	"github.com/chenhaonan-eth/dao/config"
	"github.com/chenhaonan-eth/dao/dal"
	"github.com/chenhaonan-eth/dao/dal/model"
	"github.com/chenhaonan-eth/dao/dal/query"
)

func TestCollySH300PE(t *testing.T) {
	dal.DB = dal.ConnectDB("test_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	// dal.DB.Migrator().DropTable(&model.SH300PE{})
	dal.DB.AutoMigrate(&model.SH300PE{})
	CollySH300PE()
}

func TestCollyBondZhUsRate(t *testing.T) {
	config.Init()
	dal.DB = dal.ConnectDB("test_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	// dal.DB.Migrator().DropTable(&model.BondZhUsRate{})
	// dal.DB.AutoMigrate(&model.BondZhUsRate{})
	ti := q.BondZhUsRate
	do := ti.WithContext(context.Background())
	do.Create(&model.BondZhUsRate{Date: "2022-12-04 00:00:00", CN2Years: 1, CN5Years: 1, CN10Years: 1, CN30Years: 1, CN10_2Years: 1, US2Years: 1, US5Years: 1, US10Years: 1, US30Years: 1, US10_2Years: 1})
	CollyBondZhUsRate()
}

func TestCollyChinaMoneySupply(t *testing.T) {
	dal.DB = dal.ConnectDB("test_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	// dal.DB.Migrator().DropTable(&model.MacroChinaMoneySupply{})
	// dal.DB.AutoMigrate(&model.MacroChinaMoneySupply{})
	CollyChinaMoneySupply()
}

func TestCollyCNGDP(t *testing.T) {
	config.Init()
	dal.DB = dal.ConnectDB("test_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	// dal.DB.Migrator().DropTable(&model.MacroChinaMoneySupply{})
	// dal.DB.AutoMigrate(&model.MacroGDP{})

	// ti := q.MacroGDP
	// do := ti.WithContext(context.Background())
	// do.Create(&model.MacroGDP{Date: "2022-10-14", Country: "cn", Gdp: "2.3"})
	CollyCNGDP()
}
func TestCollyCNCPI(t *testing.T) {
	config.Init()
	dal.DB = dal.ConnectDB("test_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	// dal.DB.Migrator().DropTable(&model.MacroChinaMoneySupply{})
	dal.DB.AutoMigrate(&model.ChinaCPI{})

	// ti := q.MacroGDP
	// do := ti.WithContext(context.Background())
	// do.Create(&model.MacroGDP{Date: "2022-10-14", Country: "cn", Gdp: "2.3"})
	CollyCNCPI()
}

func TestCollyCNSocialFinancingFlow(t *testing.T) {
	config.Init()
	// dal.DB = dal.ConnectDB("test_sqlite.db").Debug()
	// query.SetDefault(dal.DB)
	CollyCNSocialFinancingFlow()
}

func TestCollyCADFuturesForeignHist(t *testing.T) {
	config.Init()
	dal.DB = dal.ConnectDB("../dao_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	CollyCADFuturesForeignHist()
}
