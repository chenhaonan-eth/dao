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
	b, err := GetEastmoney("ALL", "1000", "RPT_ECONOMY_INDUS_GROW")
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

func TestValueAddedOfIndustrialProduction(t *testing.T) {
	b, err := ValueAddedOfIndustrialProduction()
	if err != nil {
		t.Log(err)
	}
	for _, v := range b {
		t.Logf("%v", *v)
	}
}

func TestSocialElectricityConsumption(t *testing.T) {
	// byBody := `/*<script>location.href='//sina.com';</script>*/
	// SINAREMOTECALLCALLBACK1671102964958(({config:{all:[],defaultItems:[1,2],querylist:[{title:"time",data:[["2022.10","2022.10",true],["2022.9","2022.9",true],["2022.8","2022.8",true],["2022.7","2022.7",true],["2022.6","2022.6",true],["2022.5","2022.5",true],["2022.4","2022.4",true],["2022.3","2022.3",true],["2022.2","2022.2",true],["2021.12","2021.12",true],["2021.11","2021.11",true],["2021.10","2021.10",true],["2021.9","2021.9",false],["2021.8","2021.8",false],["2021.7","2021.7",false],["2021.6","2021.6",false],["2021.5","2021.5",false],["2021.4","2021.4",false],["2021.3","2021.3",false],["2021.2","2021.2",false],["2020.12","2020.12",false],["2020.11","2020.11",false],["2020.10","2020.10",false],["2020.9","2020.9",false],["2020.8","2020.8",false],["2020.7","2020.7",false],["2020.6","2020.6",false],["2020.5","2020.5",false],["2020.4","2020.4",false],["2020.3","2020.3",false],["2020.2","2020.2",false],["2019.12","2019.12",false],["2019.11","2019.11",false],["2019.10","2019.10",false],["2019.9","2019.9",false],["2019.8","2019.8",false],["2019.7","2019.7",false],["2019.6","2019.6",false],["2019.5","2019.5",false],["2019.4","2019.4",false],["2019.3","2019.3",false],["2019.2","2019.2",false],["2018.12","2018.12",false],["2018.11","2018.11",false],["2018.10","2018.10",false],["2018.9","2018.9",false],["2018.8","2018.8",false],["2018.7","2018.7",false],["2018.6","2018.6",false],["2018.5","2018.5",false],["2018.4","2018.4",false],["2018.3","2018.3",false],["2018.2","2018.2",false],["2017.12","2017.12",false],["2017.11","2017.11",false],["2017.10","2017.10",false],["2017.9","2017.9",false],["2017.8","2017.8",false],["2017.7","2017.7",false],["2017.6","2017.6",false],["2017.5","2017.5",false],["2017.4","2017.4",false],["2017.3","2017.3",false],["2017.2","2017.2",false],["2016.12","2016.12",false],["2016.11","2016.11",false],["2016.10","2016.10",false],["2016.9","2016.9",false],["2016.8","2016.8",false],["2016.7","2016.7",false],["2016.6","2016.6",false],["2016.5","2016.5",false],["2016.4","2016.4",false],["2016.3","2016.3",false],["2016.2","2016.2",false],["2015.12","2015.12",false],["2015.11","2015.11",false],["2015.10","2015.10",false],["2015.9","2015.9",false],["2015.8","2015.8",false],["2015.7","2015.7",false],["2015.6","2015.6",false],["2015.5","2015.5",false],["2015.4","2015.4",false],["2015.3","2015.3",false],["2015.2","2015.2",false],["2014.12","2014.12",false],["2014.11","2014.11",false],["2014.10","2014.10",false],["2014.9","2014.9",false],["2014.8","2014.8",false],["2014.7","2014.7",false],["2014.6","2014.6",false],["2014.5","2014.5",false],["2014.4","2014.4",false],["2014.3","2014.3",false],["2014.2","2014.2",false],["2013.12","2013.12",false],["2013.11","2013.11",false],["2013.10","2013.10",false],["2013.9","2013.9",false],["2013.8","2013.8",false],["2013.7","2013.7",false],["2013.6","2013.6",false],["2013.5","2013.5",false],["2013.4","2013.4",false],["2013.3","2013.3",false],["2013.2","2013.2",false],["2012.12","2012.12",false],["2012.11","2012.11",false],["2012.10","2012.10",false],["2012.9","2012.9",false],["2012.8","2012.8",false],["2012.7","2012.7",false],["2012.6","2012.6",false],["2012.5","2012.5",false],["2012.4","2012.4",false],["2012.3","2012.3",false],["2012.2","2012.2",false],["2011.12","2011.12",false],["2011.11","2011.11",false],["2011.10","2011.10",false],["2011.9","2011.9",false],["2011.8","2011.8",false],["2011.7","2011.7",false],["2011.6","2011.6",false],["2011.5","2011.5",false],["2011.4","2011.4",false],["2011.3","2011.3",false],["2011.2","2011.2",false],["2010.12","2010.12",false],["2010.11","2010.11",false],["2010.10","2010.10",false],["2010.9","2010.9",false],["2010.8","2010.8",false],["2010.7","2010.7",false],["2010.6","2010.6",false],["2010.5","2010.5",false],["2010.4","2010.4",false],["2010.3","2010.3",false],["2010.2","2010.2",false],["2010.1","2010.1",false],["2009.12","2009.12",false],["2009.11","2009.11",false],["2009.10","2009.10",false],["2009.9","2009.9",false],["2009.8","2009.8",false],["2009.7","2009.7",false],["2009.6","2009.6",false],["2009.5","2009.5",false],["2009.4","2009.4",false],["2009.3","2009.3",false],["2009.2","2009.2",false],["2008.12","2008.12",false],["2008.11","2008.11",false],["2008.10","2008.10",false],["2008.9","2008.9",false],["2008.8","2008.8",false],["2008.7","2008.7",false],["2008.6","2008.6",false],["2008.5","2008.5",false],["2008.4","2008.4",false],["2008.3","2008.3",false],["2008.2","2008.2",false],["2007.12","2007.12",false],["2007.11","2007.11",false],["2007.10","2007.10",false],["2007.9","2007.9",false],["2007.8","2007.8",false],["2007.7","2007.7",false],["2007.6","2007.6",false],["2007.5","2007.5",false],["2007.4","2007.4",false],["2007.3","2007.3",false],["2007.2","2007.2",false],["2006.12","2006.12",false],["2006.11","2006.11",false],["2006.10","2006.10",false],["2006.9","2006.9",false],["2006.8","2006.8",false],["2006.7","2006.7",false],["2006.6","2006.6",false],["2006.5","2006.5",false],["2006.4","2006.4",false],["2006.3","2006.3",false],["2006.2","2006.2",false],["2005.11","2005.11",false],["2005.10","2005.10",false],["2005.9","2005.9",false],["2005.8","2005.8",false],["2005.6","2005.6",false],["2005.5","2005.5",false],["2005.4","2005.4",false],["2005.3","2005.3",false],["2005.2","2005.2",false],["2004.11","2004.11",false],["2004.10","2004.10",false],["2004.9","2004.9",false],["2004.3","2004.3",false],["2003.12","2003.12",false]]}],index:1},count:"200",data:[["2022.10","717600000.00","3.80",null,null,"9480000.00","9.90","470860000.00","1.70","124790000.00","4.20","112470000.00","12.60",null,null,null,null],["2022.9","649310000.00","4.00",null,null,"8570000.00","8.40","423640000.00","1.60","113790000.00","4.90","103310000.00","13.50",null,null,null,null]]}));`

	// b := byBody[bytes.LastIndexAny([]byte(byBody), "c")+7 : bytes.LastIndexAny([]byte(byBody), "da")-5]
	b, err := SocialElectricityConsumption()
	if err != nil {
		t.Log(err)
	}
	for _, v := range b {
		t.Logf("%v \n", *v)
	}
}

func TestMacroChinaGdpYearlyJin10(t *testing.T) {
	MacroChinaGdpYearlyJin10()
}
