package spider

import (
	"testing"

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
