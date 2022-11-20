package model

var (
	OpArgsMap = make(map[string]interface{})
)

// 注册所有数据库
func init() {
	OpArgsMap["BondZhUsRate"] = &BondZhUsRate{}
	OpArgsMap["SH300PE"] = &SH300PE{}
	OpArgsMap["PePbPsDvTotalmv"] = &PePbPsDvTotalmv{}
	OpArgsMap["MacroChinaMoneySupply"] = &MacroChinaMoneySupply{}
	OpArgsMap["MacroPMI"] = &MacroPMI{}
	OpArgsMap["SocialFinancingFlow"] = &SocialFinancingFlow{}
	OpArgsMap["SocialFinancingStock"] = &SocialFinancingStock{}
	OpArgsMap["MacroGDP"] = &MacroGDP{}
	OpArgsMap["MacroChinaConsumerGoodsRetail"] = &MacroChinaConsumerGoodsRetail{}
	OpArgsMap["MacroCpi"] = &MacroCpi{}
	OpArgsMap["MacroPpi"] = &MacroPpi{}
	OpArgsMap["FturesFoewign"] = &FturesFoewign{}
}
