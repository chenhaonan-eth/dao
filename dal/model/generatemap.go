package model

var (
	OpArgsMap = make(map[string]interface{})
)

// 注册所有数据库
func init() {
	OpArgsMap["bond_zh_us_rate_models"] = &BondZhUsRateModel{}
	OpArgsMap["sH300_pe_models"] = &SH300PEModel{}
	OpArgsMap["PePbPsDvTotalmv"] = &PePbPsDvTotalmv{}
	OpArgsMap["MacroChinaMoneySupplyModel"] = &MacroChinaMoneySupplyModel{}
}
