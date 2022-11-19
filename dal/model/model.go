// 定义的数据模型，使用dorm gen可生成model的CURD
// 已定义好执行脚本 generate_gorm_gen_code.sh
package model

import (
	"fmt"
	"time"
)

// 中国社会融资规模数据对象
// 单位：万亿元人民币
type MacroChinaShrzgm struct {
	Date                            string `json:"date"`                    //时间
	Tyep                            string `json:"tyep"`                    // Flow or Stock 总量或存量
	Ndbab                           string `json:"ndbab"`                   //其中-未贴现银行承兑汇票
	Entrustloan                     string `json:"entrustloan"`             //其中-委托贷款
	Forcloan                        string `json:"forcloan"`                //其中-委托贷款外币(折合人民币)
	Rmblaon                         string `json:"rmblaon"`                 //其中-人民币贷款
	Bibae                           string `json:"bibae"`                   //其中-企业债券
	Tiosfs                          string `json:"tiosfs"`                  //社会融资规模增量(亿元)/存量
	Sfinfe                          string `json:"sfinfe"`                  //其中-非金融企业境内股票融资
	Trustloan                       string `json:"trustloan"`               //其中-信托贷款
	AssetBackedSecurities           string `json:"asset_backed_securities"` //存款类金融机构资产支持证券
	LoansWrittenOff                 string `json:"loans_written_off"`       //贷款核销
	GovernmentBonds                 string `json:"government_bonds"`        //政府债券
	NdbabGrowthRate                 string `json:"ndbabgrowthrate"`         //增速（%）
	EntrustloanGrowthRate           string `json:"entrustloangrowthrate"`
	ForcloanGrowthRate              string `json:"forcloangrowthrate"`
	RmblaonGrowthRate               string `json:"rmblaongrowthrate"`
	BibaeGrowthRate                 string `json:"bibaegrowthrate"`
	TiosfsGrowthRate                string `json:"tiosfsgrowthrate"`
	SfinfeGrowthRate                string `json:"sfinfegrowthrate"`
	TrustloanGrowthRate             string `json:"trustloangrowthrate"`
	AssetBackedSecuritiesGrowthRate string `json:"asset_backed_securitiesgrowthrate"`
	LoansWrittenOffGrowthRate       string `json:"loans_written_offgrowthrate"`
	GovernmentBondsGrowthRate       string `json:"government_bondsgrowthrate"`
}

func (m *MacroChinaShrzgm) String() string {
	return fmt.Sprintf(
		"{date: %v, tyep: %v, ndbab: %v,forcloan: %v, rmblaon: %v, bibae: %v,tiosfs: %v, sfinfe: %v, trustloan: %v,asset_backed_securities: %v, loans_written_off: %v, ndbabgrowthrate: %v,entrustloangrowthrate: %v, forcloangrowthrate: %v, rmblaongrowthrate: %v,bibaegrowthrate: %v,tiosfsgrowthrate: %v, sfinfegrowthrate: %v, trustloangrowthrate: %v,asset_backed_securitiesgrowthrate: %v,loans_written_offgrowthrate: %v,government_bonds: %v,government_bondsgrowthrate:%v}",
		m.Date, m.Tyep, m.Ndbab, m.Forcloan, m.Rmblaon, m.Bibae, m.Tiosfs, m.Sfinfe, m.Trustloan, m.AssetBackedSecurities, m.LoansWrittenOff, m.NdbabGrowthRate, m.EntrustloanGrowthRate, m.ForcloanGrowthRate, m.RmblaonGrowthRate, m.BibaeGrowthRate, m.TiosfsGrowthRate, m.SfinfeGrowthRate, m.TrustloanGrowthRate, m.AssetBackedSecuritiesGrowthRate, m.LoansWrittenOffGrowthRate, m.GovernmentBonds, m.GovernmentBondsGrowthRate)
}

type MacroChinaMoneySupply struct {
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

type MacroChinaConsumerGoodsRetail struct {
	Date                        string `gorm:"index"`
	TotalRetailSales            string // 社会消费品零售总额(亿元)
	YearOnYear                  string // 同比增长
	YearOverYear                string // 环比增长
	TotalAccumulation           string // 累计
	TotalAccumulationYearOnYear string // 累计同比增长
}

// 中美国债模型
type BondZhUsRate struct {
	Date        string  `json:"SOLAR_DATE"`
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

// 沪深300市盈率
type SH300PE struct {
	Date        float64 `json:"date"`        // 时间
	MiddleLyrPe float64 `json:"middleLyrPe"` //沪深300静态市盈率中位数
	LyrPe       float64 `json:"lyrPe"`       //沪深300静态市盈率等权平均
	AddLyrPe    float64 `json:"addLyrPe"`    //沪深300静态市盈率
	MiddleTtmPe float64 `json:"middleTtmPe"` //沪深300滚动市盈率(TTM)中位数
	TtmPe       float64 `json:"ttmPe"`       //沪深300滚动市盈率(TTM)等权平均
	AddTtmPe    float64 `json:"addTtmPe"`    //沪深300滚动市盈率(TTM)
}

// 全A 市盈率|市净率|股息率|市销率|总市值
type PePbPsDvTotalmv struct {
	Date     string `json:"date"`     // 时间
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

// PMI
type MacroPMI struct {
	Date    string `json:"date" gorm:"index"`
	Country string `json:"country"`
	Pmi     string `json:"pmi"`
}

// GDP
type MacroGDP struct {
	Date    string `json:"date" gorm:"index"`
	Country string `json:"country"`
	Gdp     string `json:"gdp"`
}

// cpi
type MacroCpi struct {
	Date    string `json:"date" gorm:"index"`
	Country string `json:"country"`
	Cpi     string `json:"cpi"`
}

// ppi
type MacroPpi struct {
	Date    string `json:"date" gorm:"index"`
	Country string `json:"country"`
	Ppi     string `json:"Ppi"`
}

// 外盘期货交易
type FturesFoewign struct {
	Date   string `json:"date"`   //交易日
	Open   string `json:"open"`   //开盘价
	High   string `json:"high"`   //最高价
	Low    string `json:"low"`    //最低价
	Close  string `json:"close"`  //收盘价
	Volume string `json:"volume"` //成交量
}

func (f *FturesFoewign) String() string {
	return fmt.Sprintf("{date: %v, open: %v, high: %v,low: %v, close: %v, volume: %v}", f.Date, f.Open, f.High, f.Low, f.Close, f.Volume)
}
