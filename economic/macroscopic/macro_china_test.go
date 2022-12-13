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
	for _, v := range v {
		t.Logf("%v", *v)
	}
}

func TestMacroChinaPmiYearly(t *testing.T) {
	v, err := ChinaPMI()
	if err != nil {
		t.Error(err)
	}
	for _, v := range v {
		t.Logf("%v", *v)
	}
}

func TestMacroChinaConsumerGoodsRetail(t *testing.T) {
	v, err := MacroChinaConsumerGoodsRetail()
	if err != nil {
		t.Error(err)
	}
	for _, v := range v {
		t.Logf("%v", *v)
	}
}

func TestChinaGDP(t *testing.T) {
	v, err := ChinaGDP()
	if err != nil {
		t.Error(err)
	}
	for _, v := range v {
		t.Logf("%v", *v)
	}
}

func TestMacroChinaCpiMonthly(t *testing.T) {
	v, err := ChinaCPI()
	if err != nil {
		t.Error(err)
	}
	for _, v := range v {
		t.Logf("%v", *v)
	}
}

func TestMacroChinaPpiYearly(t *testing.T) {
	v, err := ChinaPPI()
	if err != nil {
		t.Error(err)
	}
	for _, v := range v {
		t.Logf("%v", *v)
	}
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

func TestGetEastmoney(t *testing.T) {
	b, err := GetEastmoney("ALL", "1000", "RPT_ECONOMY_CPI")
	if err != nil {
		t.Log(err)
	}
	t.Log(string(b))
}

func TestPmiManCx(t *testing.T) {
	b, err := PmiCx()
	if err != nil {
		t.Log(err)
	}
	for _, v := range b {
		t.Logf("%v", *v)
	}
}
