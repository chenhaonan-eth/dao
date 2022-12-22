package model

var OpArgsMap = make(map[string]interface{})

// 注册所有数据库
func init() {
	OpArgsMap["BondZhUsRate"] = &BondZhUsRate{}
	OpArgsMap["SH300PE"] = &SH300PE{}
	OpArgsMap["PePbPsDvTotalmv"] = &PePbPsDvTotalmv{}
	OpArgsMap["MacroChinaMoneySupply"] = &MacroChinaMoneySupply{}
	OpArgsMap["MacroPMI"] = &ChinaPMI{}
	OpArgsMap["SocialFinancingFlow"] = &SocialFinancingFlow{}
	OpArgsMap["SocialFinancingStock"] = &SocialFinancingStock{}
	OpArgsMap["MacroGDP"] = &ChinaGDP{}
	OpArgsMap["MacroChinaConsumerGoodsRetail"] = &MacroChinaConsumerGoodsRetail{}
	OpArgsMap["MacroCpi"] = &ChinaCPI{}
	OpArgsMap["MacroPpi"] = &ChinaPPI{}
	OpArgsMap["FturesFoewign"] = &FturesFoewign{}
	OpArgsMap["CxPmi"] = &PmiCx{}
	OpArgsMap["ValueAddedOfIndustrialProduction"] = &ValueAddedOfIndustrialProduction{}
	OpArgsMap["SocialElectricityConsumption"] = &SocialElectricityConsumption{}
	OpArgsMap["PassengerAndFreightTraffic"] = &PassengerAndFreightTraffic{}
	OpArgsMap["NewFinancialCredit"] = &NewFinancialCredit{}
	OpArgsMap["ForeignReserveAndGold"] = &ForeignReserveAndGold{}
	OpArgsMap["InvestmentInFixedAssets"] = &InvestmentInFixedAssets{}
	OpArgsMap["CentralBankMonetaryAuthorityAssetsAndLiabilities"] = &CentralBankMonetaryAuthorityAssetsAndLiabilities{}
	OpArgsMap["ManufacturingPmiParticulars"] = &ManufacturingPmiParticulars{}
}
