package macroscopic

import (
	"testing"
)

func TestMacro_china_shrzgm(t *testing.T) {
	t.Log("start TestMacro_china_shrzgm ")

	m, err := MacroChinaShrzgm()
	if err != nil {
		t.Error(err)
	}
	t.Logf("rsp: %+v", m)

	t.Log("end TestMacro_china_shrzgm ")
}

func TestMacroChinaMoneySupply(t *testing.T) {
	v, err := MacroChinaMoneySupply()
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}

func TestMacroChinaPmiYearly(t *testing.T) {
	v, err := MacroChinaPmiYearly()
	if err != nil {
		t.Error(err)
	}
	t.Log(v)
}

func TestMacroChinaConsumerGoodsRetail(t *testing.T) {
	v, err := MacroChinaConsumerGoodsRetail()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v", v)
}

func TestMacroChinaGdpYearly(t *testing.T) {
	v, err := MacroChinaGdpYearly()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", v)
}

func TestMacroChinaCpiMonthly(t *testing.T) {
	v, err := MacroChinaCpiYearly()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", v)
}
func TestMacroChinaPpiYearly(t *testing.T) {
	v, err := MacroChinaPpiYearly()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", v)
}
func TestStockAPe(t *testing.T) {
	v, err := StockAPe()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", v)
}

func TestSH300PE(t *testing.T) {
	v, err := SH300PE()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", v)
}

func TestBondZhUsRate(t *testing.T) {
	v, err := BondZhUsRate()
	if err != nil {
		t.Error(err)
	}
	t.Logf("%v", v)
}

func TestStockSohuCom(t *testing.T) {
	err := StockSohuCom("")
	if err != nil {
		t.Error(err)
	}
}