package dal

import (
	"testing"

	"github.com/chenhaonan-eth/dao/dal/model"
	"gorm.io/gen/examples/dal"
)

func TestCreateTable(t *testing.T) {
	dal.DB = ConnectDB("guide_sqlite.db").Debug()
	// if err := dal.DB.Migrator().CreateTable(&model.PePbPsDvTotalmv{}); err != nil {
	// 	t.Error(err)
	// }
}

func TestAutoMigrate(t *testing.T) {
	dal.DB = ConnectDB("guide_sqlite.db").Debug()
	if err := dal.DB.AutoMigrate(&model.PePbPsDvTotalmv{}); err != nil {
		t.Error(err)
	}
}
