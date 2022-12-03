package spider

import (
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
