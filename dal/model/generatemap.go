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
	OpArgsMap["MacroChinaShrzgm"] = &MacroChinaShrzgm{}
	OpArgsMap["MacroGDP"] = &MacroGDP{}
}
