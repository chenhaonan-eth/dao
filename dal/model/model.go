// 定义的数据模型，使用dorm gen可生成model的CURD
// 已定义好执行脚本 generate_gorm_gen_code.sh
package model

import (
	"fmt"
)

// 中国社会融资规模增量数据对象
type SocialFinancingFlow struct {
	Date        string  `json:"date"`        //时间
	Tiosfs      float32 `json:"tiosfs"`      //社会融资规模增量(亿元)
	Rmblaon     float32 `json:"rmblaon"`     //其中-人民币贷款
	Forcloan    float32 `json:"forcloan"`    //其中-委托贷款外币(折合人民币)
	Entrustloan float32 `json:"entrustloan"` //其中-委托贷款
	Trustloan   float32 `json:"trustloan"`   //其中-信托贷款
	Ndbab       float32 `json:"ndbab"`       //其中-未贴现银行承兑汇票
	Bibae       float32 `json:"bibae"`       //其中-企业债券
	Sfinfe      float32 `json:"sfinfe"`      //其中-非金融企业境内股票融资
}

// 中国社会融资规模存量数据
// 单位：万亿元人民币
type SocialFinancingStock struct {
	Date                            string `json:"date"`                    //时间
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

func (m *SocialFinancingStock) String() string {
	return fmt.Sprintf(
		"{date: %v, ndbab: %v,forcloan: %v, rmblaon: %v, bibae: %v,tiosfs: %v, sfinfe: %v, trustloan: %v,asset_backed_securities: %v, loans_written_off: %v, ndbabgrowthrate: %v,entrustloangrowthrate: %v, forcloangrowthrate: %v, rmblaongrowthrate: %v,bibaegrowthrate: %v,tiosfsgrowthrate: %v, sfinfegrowthrate: %v, trustloangrowthrate: %v,asset_backed_securitiesgrowthrate: %v,loans_written_offgrowthrate: %v,government_bonds: %v,government_bondsgrowthrate:%v}",
		m.Date, m.Ndbab, m.Forcloan, m.Rmblaon, m.Bibae, m.Tiosfs, m.Sfinfe, m.Trustloan, m.AssetBackedSecurities, m.LoansWrittenOff, m.NdbabGrowthRate, m.EntrustloanGrowthRate, m.ForcloanGrowthRate, m.RmblaonGrowthRate, m.BibaeGrowthRate, m.TiosfsGrowthRate, m.SfinfeGrowthRate, m.TrustloanGrowthRate, m.AssetBackedSecuritiesGrowthRate, m.LoansWrittenOffGrowthRate, m.GovernmentBonds, m.GovernmentBondsGrowthRate)
}

type MacroChinaMoneySupply struct {
	Date           string  `json:"REPORT_DATE"`               //月份
	M2             float64 `json:"BASIC_CURRENCY"`            //货币和准货币(M2)数量(亿元)
	M2YearOnYear   float64 `json:"BASIC_CURRENCY_SAME"`       //货币和准货币(M2)同比增长 / %
	M2YearOverYear float64 `json:"BASIC_CURRENCY_SEQUENTIAL"` //货币和准货币(M2)环比增长
	M1             float64 `json:"CURRENCY"`                  //货币(M1)数量(亿元)
	M1YearOnYear   float64 `json:"CURRENCY_SAME"`             //货币(M1)同比增长
	M1YearOverYear float64 `json:"CURRENCY_SEQUENTIAL"`       //货币(M1)环比增长
	M0             float64 `json:"FREE_CASH"`                 //流通中的现金(M0)数量(亿元)
	M0YearOnYear   float64 `json:"FREE_CASH_SAME"`            //流通中的现金(M0)同比增长
	M0YearOverYear float64 `json:"FREE_CASH_SEQUENTIAL"`      //流通中的现金(M0)环比增长
}

type MacroChinaConsumerGoodsRetail struct {
	Date                        string  `json:"REPORT_DATE" gorm:"index"`
	TotalRetailSales            float64 `json:"RETAIL_TOTAL"`            // 社会消费品零售总额(亿元)
	YearOnYear                  float64 `json:"RETAIL_TOTAL_SAME"`       // 同比增长
	YearOverYear                float64 `json:"RETAIL_TOTAL_SEQUENTIAL"` // 环比增长
	TotalAccumulation           float64 `json:"RETAIL_TOTAL_ACCUMULATE"` // 累计
	TotalAccumulationYearOnYear float64 `json:"RETAIL_ACCUMULATE_SAME"`  // 累计同比增长
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
	Time        string  `json:"time"`
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
type ChinaPMI struct {
	Date                       string  `json:"REPORT_DATE" gorm:"index"` //2022-10-01 00:00:00
	Time                       string  `json:"TIME"`                     //2022年10月份
	Manufacturing              float64 `json:"MAKE_INDEX"`               // 制造业
	ManufacturingYearOnYear    float64 `json:"MAKE_SAME"`                // 同比增加(%)
	NonManufacturing           float64 `json:"NMAKE_INDEX"`              // 非制造业
	NonManufacturingYearOnYear float64 `json:"NMAKE_SAME"`               // 同比增加(%)
}

// GDP
type ChinaGDP struct {
	Date                          string  `json:"REPORT_DATE" gorm:"index"` //2022-09-01 00:00:00
	Time                          string  `json:"TIME"`                     //"2022年第1-3季度"
	GDP                           float64 `json:"DOMESTICL_PRODUCT_BASE"`   // 国内生产总值(亿元)
	GDPYearOnYear                 float64 `json:"SUM_SAME"`                 // 国内生产总值季度累计同比增长（%）
	PrimaryIndustry               float64 `json:"FIRST_PRODUCT_BASE"`       // 第一产业增加值（亿元）
	PrimaryIndustryYearOnYear     float64 `json:"FIRST_SAME"`               // 第一产业增加值季度累计同比增长（%）
	SecondaryIndustries           float64 `json:"SECOND_PRODUCT_BASE"`      // 第二产业增加值（亿元）
	SecondaryIndustriesYearOnYear float64 `json:"SECOND_SAME"`              // 第二产业增加值季度累计同比增长（%）
	TertiaryIndustry              float64 `json:"THIRD_PRODUCT_BASE"`       // 第三产业增加值（亿元）
	TertiaryIndustryYearOnYear    float64 `json:"THIRD_SAME"`               // 第三产业增加值季度累计同比增长（%）
}

// cpi
type ChinaCPI struct {
	Date                 string  `json:"REPORT_DATE" gorm:"index"` //2022-10-01 00:00:00
	Time                 string  `json:"TIME"`                     //2022年10月份
	National             float64 `json:"NATIONAL_BASE"`            // 全国当月
	NationalYearOnYear   float64 `json:"NATIONAL_SAME"`            // 全国同比增加(%)
	NationalYearOverYear float64 `json:"NATIONAL_SEQUENTIAL"`      // 全国环比增长(%)
	NationalAccumulative float64 `json:"NATIONAL_ACCUMULATE"`      // 全国累计
	City                 float64 `json:"CITY_BASE"`                // 城市当月
	CityYearOnYear       float64 `json:"CITY_SAME"`                // 城市同比增加(%)
	CityYearOverYear     float64 `json:"CITY_SEQUENTIAL"`          // 城市环比增长(%)
	CityAccumulative     float64 `json:"CITY_ACCUMULATE"`          // 城市累计
	Rural                float64 `json:"RURAL_BASE"`               // 农村当月
	RuralYearOnYear      float64 `json:"RURAL_SAME"`               // 农村同比增加(%)
	RuralYearOverYear    float64 `json:"RURAL_SEQUENTIAL"`         // 农村环比增长(%)
	RuralAccumulative    float64 `json:"RURAL_ACCUMULATE"`         // 农村累计
}

// ppi
type ChinaPPI struct {
	Date         string  `json:"REPORT_DATE" gorm:"index"` //2022-10-01 00:00:00
	Time         string  `json:"TIME"`                     //2022年10月份
	PPI          float64 `json:"BASE"`                     // 工业品出厂价格指数
	YearOnYear   float64 `json:"BASE_SAME"`                // 当月同比增长(%)
	Accumulative float64 `json:"BASE_ACCUMULATE"`          // 累计
}

// 外盘期货交易
type FturesFoewign struct {
	Date   string `json:"date"`   //交易日
	Open   string `json:"open"`   //开盘价
	High   string `json:"high"`   //最高价
	Low    string `json:"low"`    //最低价
	Close  string `json:"close"`  //收盘价
	Volume string `json:"volume"` //成交量
	Symbol string `json:"symbol"` //品种
}

func (f *FturesFoewign) String() string {
	return fmt.Sprintf("{date: %v, open: %v, high: %v,low: %v, close: %v, volume: %v}", f.Date, f.Open, f.High, f.Low, f.Close, f.Volume)
}

type CxPmi struct {
	Time       float64 `json:"time"`
	Data       float64 `json:"data"`
	ChangeRate float64 `json:"changeRate"`
}

type PmiCx struct {
	Date                    string `gorm:"index"`
	Time                    float64
	Manufacture             float64 //制造业
	ManufactureYearOverYear float64 //环比
	Service                 float64 //服务业
	ServiceYearOverYear     float64 //服务业环比
	Synthesis               float64 //综合
	SynthesisYearOverYear   float64 //综合环比
}

type ValueAddedOfIndustrialProduction struct {
	Date             string  `json:"REPORT_DATE" gorm:"index"` //2022-10-01 00:00:00
	Time             string  `json:"TIME"`                     //2022年10月份
	YearOnYear       float64 `json:"BASE_SAME"`                // 同比增长
	CumulativeGrowth float64 `json:"BASE_ACCUMULATE"`          // 累计增长
}

type SocialElectricityConsumption struct {
	Date                        string `gorm:"index"`
	WholeSociety                string //全社会用电量/万千瓦时
	WholeSocietyYearOnYear      string //全社会用电量同比 %
	AllIndustries               string //各行业用电量合计/万千瓦时
	AllIndustriesYearOnYear     string //各行业用电量合计同比 %
	PrimaryIndustry             string //第一产业用电量 万千瓦时
	PrimaryIndustryYearOnYear   string //第一产业用电量同比%
	SecondaryIndustry           string //第二产业用电量 万千瓦时
	SecondaryIndustryYearOnYear string //第二产业用电量同比%
	TertiaryIndustry            string //第三产业用电量万千瓦时
	TertiaryIndustryYearOnYear  string //第三产业用电量同比%
	CitiesAndVillages           string //城乡居民生活用电量合计/ 万千瓦时
	CitiesAndVillagesYearOnYear string //城乡居民生活用电量合计同比 %
	Cities                      string //城镇居民用电量 万千瓦时
	CitiesYearOnYear            string //城镇居民用电量同比%
	Villages                    string //乡村居民用电量 万千瓦时
	VillagesYearOnYear          string //乡村居民用电量同比 %
}

const (
	InternationalService       = iota //国际航线
	HongKongMacaoRegionalRoute        //港澳地区航线
	DomesticAirline                   //国内航线
	CivilAviation                     //民航
	WaterTransport                    //水运
	Highway                           //公路
	Railway                           //铁路
	Total                             //合计

)

type CategoryOfTraffic int32

const (
	CategoryOfTraffic_InternationalService       CategoryOfTraffic = iota //国际航线
	CategoryOfTraffic_HongKongMacaoRegionalRoute                          //港澳地区航线
	CategoryOfTraffic_DomesticAirline                                     //国内航线
	CategoryOfTraffic_CivilAviation                                       //民航
	CategoryOfTraffic_WaterTransport                                      //水运
	CategoryOfTraffic_Highway                                             //公路
	CategoryOfTraffic_Railway                                             //铁路
	CategoryOfTraffic_Total                                               //合计
)

type PassengerAndFreightTraffic struct {
	Date                                         string
	Class                                        CategoryOfTraffic //运输种类
	FreightVolume                                string            //货运量/亿吨
	FreightVolumeYearOnYear                      string            //货运量同比增长 %
	FreightTurnover                              string            //货物周转量/亿
	FreightTurnoverYearOnYear                    string            //公里货物周转量同比增长 %
	PassengerCapacity                            string            //客运量/亿人
	PassengerCapacityYearOnYear                  string            //客运量同比增长 %
	PassengerTurnover                            string            //旅客周转量/亿
	PassengerTurnoverYearOnYear                  string            //公里旅客周转量同比增长/%
	CargoThroughputOfMajorCoastalPorts           string            //沿海主要港口货物吞吐量/亿吨
	CargoThroughputOfMajorCoastalPortsYearOnYear string            //沿海主要港口货物吞吐量同比增长 %
	ForeignTradeCargoThroughput                  string            //其中:外贸货物吞吐量 /亿吨
	ForeignTradeCargoThroughputYearOnYear        string            //其中:外贸货物吞吐量同比增长 %
	TotalTurnoverOfCivilAviation                 string            //民航总周转量 /亿
	KmTotalTurnoverOfCivilAviation               string            //公里民航总周转/%
}

// 中国新增信贷数据
type NewFinancialCredit struct {
	Date               string  `json:"REPORT_DATE" gorm:"index"` //2022-10-01 00:00:00
	Time               string  `json:"TIME"`                     //2022年10月份
	Loany              float64 `json:"RMB_LOAN"`                 // 当月(亿元)
	LoanYearOnYear     float64 `json:"RMB_LOAN_SAME"`            // 同比增长(%)
	LoanyYearOverYear  float64 `json:"RMB_LOAN_SEQUENTIAL"`      // 环比增长
	LoanyAcc           float64 `json:"RMB_LOAN_ACCUMULATE"`      // 累计(亿元)
	LoanyAccYearOnYear float64 `json:"LOAN_ACCUMULATE_SAME"`     // 同比增长(%)
}

// 外汇储备与黄金
type ForeignReserveAndGold struct {
	Date string `gorm:"index"` //2022-10-01 00:00:00
	Gold string //黄金(万盎司)
	// GoldYearOnYear           float64 //黄金同比%
	ForeignReserve string //外汇储备(亿美元)
	// ForeignReserveYearOnYear float64 //外汇储备同比%
}

// 固定资产投资
type InvestmentInFixedAssets struct {
	Date                                string `gorm:"index"` //2022-10-01 00:00:00
	InvestmentCompletedAmount           string //投资完成额(亿元)
	InvestmentCompletedAmountYearOnYear string //投资完成额同比增长(%)
	PrimaryIndustry                     string //第一产业(亿元)
	PrimaryIndustryYearOnYear           string //第一产业用同比%
	SecondaryIndustry                   string //第二产业
	SecondaryIndustryYearOnYear         string //第二产业同比%
	TertiaryIndustry                    string //第三产业
	TertiaryIndustryYearOnYear          string //第三产业同比%
	RealEstateDevelopment               string //房地产开发(亿元)
	RealEstateDevelopmentYearOnYear     string //房地产开发同比增长(%)
}

// 央行货币当局资产负债
type CentralBankMonetaryAuthorityAssetsAndLiabilities struct {
	Date                                     string `gorm:"index"` //2022-10-01 00:00:00
	ExternalAssets                           string //国外资产 亿元
	Forex                                    string //外汇 亿元
	MonetaryGold                             string //货币黄金 亿元
	OtherForeignAssets                       string //其他国外资产 亿元
	ClaimOnTheGovernment                     string // 对政府债权 亿元
	TheCentralGovernment                     string //其中:中央政府 亿元
	ClaimsOnOtherDepositoryCompanies         string //对其他存款性公司债权 亿元
	ClaimsOnOtherFinancialCompanies          string //对其他金融性公司债权 亿元
	ClaimsOnNonMonetaryFinancialInstitutions string //对非货币金融机构债权 亿元
	ClaimsOnNonFinancialCorporations         string //对非金融性公司债权 亿元
	AllOtherAssets                           string //其他资产 亿元
	TotalAssets                              string //总资产 亿元
	ReserveCurrency                          string //储备货币 亿元
	IssueCurrency                            string //发行货币 亿元
	FinancialCompanyDeposits                 string //金融性公司存款 亿元
	OtherDepositoryCompanies                 string //其他存款性公司 亿元
	OtherFinancialCompanies                  string //其他金融性公司 亿元
	LiabilitiesToFinancialInstitutions       string //对金融机构负债 亿元
	ReserveDeposit                           string //准备金存款 亿元
	NonFinancialCorporateDeposits            string //非金融性公司存款 亿元
	CurrentDeposit                           string //活期存款 亿元
	Bond                                     string //债券 亿元
	ForeignLiabilities                       string //国外负债 亿元
	PublicDeposit                            string //政府存款 亿元
	EquityFund                               string //自有资金 亿元
	OtherLiabilities                         string //其他负债 亿元
	GrossLiability                           string //总负债 亿元
}
