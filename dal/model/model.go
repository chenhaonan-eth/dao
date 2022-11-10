// 定义的数据模型，使用dorm gen可生成model的CURD
// 已定义好执行脚本 generate_gorm_gen_code.sh
package model

import "time"

// 中国社会融资规模数据对象
type MacroChinaShrzgmModel struct {
	Date        string  `json:"date"`        //时间
	Ndbab       float32 `json:"ndbab"`       //其中-未贴现银行承兑汇票
	Entrustloan float32 `json:"entrustloan"` //其中-委托贷款
	Forcloan    float32 `json:"forcloan"`    //其中-委托贷款外币(折合人民币)
	Rmblaon     float32 `json:"rmblaon"`     //其中-人民币贷款
	Bibae       float32 `json:"bibae"`       //其中-企业债券
	Tiosfs      float32 `json:"tiosfs"`      //社会融资规模增量(亿元)
	Sfinfe      float32 `json:"sfinfe"`      //其中-非金融企业境内股票融资
	Trustloan   float32 `json:"trustloan"`   //其中-信托贷款
}

type MacroChinaMoneySupplyModel struct {
	Date           time.Time //月份
	M2             float64   //货币和准货币(M2)数量(亿元)
	M2YearOnYear   float64   //货币和准货币(M2)同比增长
	M2YearOverYear float64   //货币和准货币(M2)环比增长
	M1             float64   //货币(M1)数量(亿元)
	M1YearOnYear   float64   //货币(M1)同比增长
	M1YearOverYear float64   //货币(M1)环比增长
	M0             float64   //流通中的现金(M0)数量(亿元)
	M0YearOnYear   float64   //流通中的现金(M0)同比增长
	M0YearOverYear float64   //流通中的现金(M0)环比增长
}

type MacroChinaConsumerGoodsRetailMode struct {
	Data                        string
	TotalRetailSales            string // 社会消费品零售总额(亿元)
	YearOnYear                  string // 同比增长
	YearOverYear                string // 环比增长
	TotalAccumulation           string // 累计
	TotalAccumulationYearOnYear string // 累计同比增长
}

// 中美国债模型
type BondZhUsRateModel struct {
	Data        string  `json:"SOLAR_DATE"`
	CN2Years    float64 `json:"EMM00588704"` //中国国债收益率2年
	CN5Years    float64 `json:"EMM00166462"` //中国国债收益率5年
	CN10Years   float64 `json:"EMM00166466"` //中国国债收益率10年
	CN30Years   float64 `json:"EMM00166469"` //中国国债收益率30年
	CN10_2Years float64 `json:"EMM01276014"` //中国国债收益率10年-2年
	US2Years    float64 `json:"EMG00001306"` //美国国债收益率2年
	US5Years    float64 `json:"EMG00001308"` //美国国债收益率5年
	US10Years   float64 `json:"EMG00001310"` //美国国债收益率10年
	US30Years   float64 `json:"EMG00001312"` //美国国债收益率30年
	US10_2Years float64 `json:"EMG01339436"` //美国国债收益率10年-2年
}

//沪深300市盈率
type SH300PEModel struct {
	Data        float64 `json:"date"`        // 时间
	MiddleLyrPe float64 `json:"middleLyrPe"` //沪深300静态市盈率中位数
	LyrPe       float64 `json:"lyrPe"`       //沪深300静态市盈率等权平均
	AddLyrPe    float64 `json:"addLyrPe"`    //沪深300静态市盈率
	MiddleTtmPe float64 `json:"middleTtmPe"` //沪深300滚动市盈率(TTM)中位数
	TtmPe       float64 `json:"ttmPe"`       //沪深300滚动市盈率(TTM)等权平均
	AddTtmPe    float64 `json:"addTtmPe"`    //沪深300滚动市盈率(TTM)
}

// 全A 市盈率|市净率|股息率|市销率|总市值
type PePbPsDvTotalmv struct {
	Data     string `json:"date"`     // 时间
	Pe       string `json:"pe"`       // 市盈率
	PeTtm    string `json:"pe_ttm"`   // 市盈率TTM
	Pb       string `json:"pb"`       // 市净率
	Ps       string `json:"ps"`       // 市销率
	PsTtm    string `json:"ps_ttm"`   // 市销率TTM
	Dv_ratio string `json:"dv_ratio"` // 股息率 （%）
	Dv_ttm   string `json:"dv_ttm"`   // 股息率（TTM）（%）
	Total_mv string `json:"total_mv"` // 总市值 万为单位
	Name     string `json:"name"`     // 名字
	Symbol   string `json:"symbol"`   // 股票代码
}
