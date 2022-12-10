package initialize

import (
	"fmt"
	"testing"

	"github.com/chenhaonan-eth/dao/dal"
	"github.com/chenhaonan-eth/dao/dal/model"
	"github.com/chenhaonan-eth/dao/dal/query"
	"github.com/chenhaonan-eth/dao/pkg/utils"
)

func TestInitlegulegu(t *testing.T) {
	dal.DB = dal.ConnectDB("guide_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	initleguleguPePSPb()
}

func TestProcess(t *testing.T) {
	r, err := Process("冠龙节能(301151)", "/s/301151")
	if err != nil {
		t.Error(err)
	}
	for _, v := range r {
		fmt.Printf("%+v", *v)
	}
}

func TestInitPMI(t *testing.T) {
	dal.DB = dal.ConnectDB("guide_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	dal.DB.AutoMigrate(&model.ChinaPMI{})
	err := initPMI()
	if err != nil {
		t.Error(err)
	}
}

func TestInitTotalSocialFinancing(t *testing.T) {
	dal.DB = dal.ConnectDB("guide_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	dal.DB.AutoMigrate(&model.SocialFinancingFlow{})
	err := initTotalSocialFinancing()
	if err != nil {
		t.Error(err)
	}
}

func TestInitGDP(t *testing.T) {
	dal.DB = dal.ConnectDB("guide_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	dal.DB.AutoMigrate(&model.ChinaGDP{})
	err := initGDP()
	if err != nil {
		t.Error(err)
	}
}

func TestInitMacroChinaConsumerGoodsRetail(t *testing.T) {
	dal.DB = dal.ConnectDB("guide_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	dal.DB.AutoMigrate(&model.MacroChinaConsumerGoodsRetail{})
	if err := initMacroChinaConsumerGoodsRetail(); err != nil {
		t.Error(err)
	}
}

func TestInitMacroCpi(t *testing.T) {
	dal.DB = dal.ConnectDB("../../dao_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	dal.DB.AutoMigrate(&model.ChinaCPI{})
	// dal.DB.Migrator().DropTable(&model.ChinaCPI{})
	if err := initMacroChinaCpi(); err != nil {
		t.Error(err)
	}
}
func TestInitMacroChinaPpi(t *testing.T) {
	dal.DB = dal.ConnectDB("guide_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	dal.DB.AutoMigrate(&model.ChinaPPI{})
	if err := initMacroPpi(); err != nil {
		t.Error(err)
	}
}

func TestInitSocialFinancingStock(t *testing.T) {
	err := initSocialFinancingStock()
	if err != nil {
		t.Error(err)
	}
}
func TestProcessedSocialFinancingFlowTable(t *testing.T) {
	strhtml, err := utils.ReaderText("./out.txt")
	if err != nil {
		t.Error(err)
	}
	ProcessedSocialFinancingFlowTable(strhtml)
}
