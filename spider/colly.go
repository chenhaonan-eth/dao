package spider

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"reflect"
	"strings"

	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chenhaonan-eth/dao/config"
	"github.com/chenhaonan-eth/dao/dal/initialize"
	"github.com/chenhaonan-eth/dao/dal/model"
	"github.com/chenhaonan-eth/dao/dal/query"
	"github.com/chenhaonan-eth/dao/economic/macroscopic"
	"github.com/chenhaonan-eth/dao/pkg/utils"
	"github.com/chromedp/chromedp"
	"go.uber.org/zap"
	"gorm.io/gorm"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

var (
	Client = resty.New()
	q      = query.Q
)

//中美国债收益率
func CollyBondZhUsRate() {
	config.G_LOG.Debug("Start CollyBondZhUsRate ")
	res := []*model.BondZhUsRate{}
	// 计算如果上月时间对比数据如已入库，则取消请求
	t := q.BondZhUsRate
	do := t.WithContext(context.Background())
	// 查询数据库是否存在最近一天数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error(err.Error())
		return
	}

	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"type":    "RPTA_WEB_TREASURYYIELD",
			"sty":     "ALL",
			"st":      "SOLAR_DATE",
			"sr":      "-1",
			"token":   "894050c76af8597a853f5b408b759f5d",
			"p":       "1",
			"ps":      "1",
			"pageNo":  "1",
			"pageNum": "1",
			"_":       "1615791534490",
		}).
		Get("http://datacenter.eastmoney.com/api/data/get")
	if err != nil {
		config.G_LOG.Error("CollyBondZhUsRate HTTP get ", zap.Error(err))
		return
	}
	b := resp.Body()
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &res)

	if m.Date == res[0].Date {
		config.G_LOG.Error("CollyBondZhUsRate The latest data is the same as the database ", zap.Any("date", *res[0]))
		return
	}
	if err := do.Create(res[0]); err != nil {
		config.G_LOG.Error("CollyBondZhUsRate Create ", zap.Error(err))
	}
	config.G_LOG.Debug("End CollyBondZhUsRate ")
}

//社会消费品零售总额
func CollyMacroChinaConsumerGoodsRetail() {
	config.G_LOG.Debug("Start CollyMacroChinaConsumerGoodsRetail ")
	// 计算如果上月时间对比数据如已入库，则取消请求
	t := q.MacroChinaConsumerGoodsRetail
	do := t.WithContext(context.Background())
	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error(err.Error())
		return
	}
	if m.Date == utils.FirstDayOfLastMonth() {
		config.G_LOG.Error("db is exist", zap.Any("date", m.Date))
		return
	}
	res := []*model.MacroChinaConsumerGoodsRetail{}
	b, err := macroscopic.GetEastmoney(
		"REPORT_DATE,RETAIL_TOTAL,RETAIL_TOTAL_SAME,RETAIL_TOTAL_SEQUENTIAL,RETAIL_TOTAL_ACCUMULATE,RETAIL_ACCUMULATE_SAME",
		"1", "RPT_ECONOMY_TOTAL_RETAIL")
	if err != nil {
		config.G_LOG.Error("CollyMacroChinaConsumerGoodsRetail HTTP get ", zap.Error(err))
		return
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &res)

	if m.Date == res[0].Date {
		config.G_LOG.Error("CollyMacroChinaConsumerGoodsRetail The latest data is the same as the database ", zap.Any("date", *res[0]))
		return
	}
	if err := do.Create(res[0]); err != nil {
		config.G_LOG.Error("CollyMacroChinaConsumerGoodsRetail Create ", zap.Error(err))
	}
	config.G_LOG.Debug("End CollyMacroChinaConsumerGoodsRetail ")

}

// 货币供应 M0 M1 M2
func CollyChinaMoneySupply() {
	config.G_LOG.Debug("Start CollyChinaMoneySupply ")
	// 计算如果上月时间对比数据如已入库，则取消请求
	t := q.MacroChinaMoneySupply
	do := t.WithContext(context.Background())
	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error(err.Error())
		return
	}
	if m.Date == utils.FirstDayOfLastMonth() {
		config.G_LOG.Error("db is exist", zap.Any("date", m.Date))
		return
	}
	res := []*model.MacroChinaMoneySupply{}
	b, err := macroscopic.GetEastmoney(
		"REPORT_DATE,BASIC_CURRENCY,BASIC_CURRENCY_SAME,BASIC_CURRENCY_SEQUENTIAL,CURRENCY,CURRENCY_SAME,CURRENCY_SEQUENTIAL,FREE_CASH,FREE_CASH_SAME,FREE_CASH_SEQUENTIAL",
		"1", "RPT_ECONOMY_CURRENCY_SUPPLY")
	if err != nil {
		config.G_LOG.Error("CollyChinaMoneySupply HTTP get ", zap.Error(err))
		return
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &res)

	if m.Date == res[0].Date {
		config.G_LOG.Error("CollyChinaMoneySupply The latest data is the same as the database ", zap.Any("date", *res[0]))
		return
	}
	if err := do.Create(res[0]); err != nil {
		config.G_LOG.Error("CollyChinaMoneySupply Create ", zap.Error(err))
	}
	config.G_LOG.Debug("End CollyChinaMoneySupply ")
}

// 沪深300市盈率
func CollySH300PE() {
	config.G_LOG.Debug("Start CollySH300PE. ")
	t := q.SH300PE
	do := t.WithContext(context.Background())
	url := "https://legulegu.com/stockdata/hs300-ttm-lyr"
	resp, err := Client.R().SetHeaders(map[string]string{
		"Content-Type": "application/json",
		"User-Agent":   utils.Random(),
	}).Get(url)
	if err != nil {
		return
	}
	b := resp.Body()
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(b)))
	if err != nil {
		log.Fatal(err)
		return
	}

	result := model.SH300PE{}
	dom.Find("#data-description .index-data-update-date").Each(func(i int, s *goquery.Selection) {
		str := strings.TrimSpace(s.Text())
		//2022年11月30日
		t, err := utils.ParseTime(str)
		if err != nil {
			config.G_LOG.Error(err.Error())
			return
		}
		result.Time = t.Format("2006-01-02 15:04:05")
		result.Date = float64(t.UnixMilli())
	})

	_, err = do.Where(t.Date.Eq(result.Date)).First()
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if err != nil {
			config.G_LOG.Error(err.Error())
		} else {
			config.G_LOG.Debug("CollySH300PE. is exist", zap.Any("Time", result.Time))
		}
		return
	}

	ivalue := reflect.ValueOf(&result).Elem()
	dom.Find("#data-description tbody > tr td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		str := strings.TrimSpace(s.Text())
		elem := ivalue.Field(i + 2)
		float, _ := strconv.ParseFloat(str, 64)
		elem.SetFloat(float)
	})

	do.Create(&result)
	config.G_LOG.Debug("End  CollySH300PE. Create", zap.Any("result", result))
}

// PMI
func CollyCNPMI() {
	config.G_LOG.Debug("Start CollyCNPMI ")
	// 计算如果上月时间对比数据如已入库，则取消请求
	t := q.ChinaGDP
	do := t.WithContext(context.Background())
	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error(err.Error())
		return
	}
	if m.Date == utils.FirstDayOfLastMonth() {
		config.G_LOG.Error("db is exist", zap.Any("date", m.Date))
		return
	}
	res := []*model.ChinaGDP{}
	b, err := macroscopic.GetEastmoney("ALL", "1", "RPT_ECONOMY_PMI")
	if err != nil {
		config.G_LOG.Error("CollyCNPMI HTTP get ", zap.Error(err))
		return
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &res)

	if m.Date == res[0].Date {
		config.G_LOG.Error("CollyCNPMI The latest data is the same as the database ", zap.Any("date", *res[0]))
		return
	}
	if err := do.Create(res[0]); err != nil {
		config.G_LOG.Error("CollyCNPMI Create ", zap.Error(err))
	}
	config.G_LOG.Debug("End CollyCNPMI ")
}

// 爬取CN GDP
func CollyCNGDP() {
	config.G_LOG.Debug("Start CollyCNGDP ")
	// 计算如果上月时间对比数据如已入库，则取消请求
	t := q.ChinaGDP
	do := t.WithContext(context.Background())
	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error(err.Error())
		return
	}
	if m.Date == utils.FirstDayOfLastMonth() {
		config.G_LOG.Error("db is exist", zap.Any("date", m.Date))
		return
	}
	res := []*model.ChinaGDP{}
	b, err := macroscopic.GetEastmoney("ALL", "1", "RPT_ECONOMY_GDP")
	if err != nil {
		config.G_LOG.Error("CollyCNGDP HTTP get ", zap.Error(err))
		return
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &res)

	if m.Date == res[0].Date {
		config.G_LOG.Error("CollyCNGDP The latest data is the same as the database ", zap.Any("date", *res[0]))
		return
	}
	if err := do.Create(res[0]); err != nil {
		config.G_LOG.Error("CollyCNGDP Create ", zap.Error(err))
	}
	config.G_LOG.Debug("End CollyCNGDP ")
}

// 爬取CN CPI
func CollyCNCPI() {
	config.G_LOG.Debug("Start CollyCNCPI ")
	// 计算如果上月时间对比数据如已入库，则取消请求
	t := q.ChinaCPI
	do := t.WithContext(context.Background())
	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error(err.Error())
		return
	}
	if m.Date == utils.FirstDayOfLastMonth() {
		config.G_LOG.Error("db is exist", zap.Any("date", m.Date))
		return
	}
	res := []*model.ChinaCPI{}
	b, err := macroscopic.GetEastmoney("ALL", "1", "RPT_ECONOMY_CPI")
	if err != nil {
		config.G_LOG.Error("CollyCNCPI HTTP get ", zap.Error(err))
		return
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &res)

	if m.Date == res[0].Date {
		config.G_LOG.Error("CollyCNCPI The latest data is the same as the database ", zap.Any("date", *res[0]))
		return
	}
	if err := do.Create(res[0]); err != nil {
		config.G_LOG.Error("CollyCNCPI Create ", zap.Error(err))
	}
	config.G_LOG.Debug("End CollyCNCPI ")
}

// 爬取CN社融ppi，从每月9号开始
func CollyCNPPI() {
	config.G_LOG.Debug("Start CollyCNPPI ")
	// 计算如果上月时间对比数据如已入库，则取消请求
	t := q.ChinaCPI
	do := t.WithContext(context.Background())
	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error(err.Error())
		return
	}
	if m.Date == utils.FirstDayOfLastMonth() {
		config.G_LOG.Error("db is exist", zap.Any("date", m.Date))
		return
	}
	res := []*model.ChinaCPI{}
	b, err := macroscopic.GetEastmoney("ALL", "1", "RPT_ECONOMY_PPI")
	if err != nil {
		config.G_LOG.Error("CollyCNPPI HTTP get ", zap.Error(err))
		return
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &res)

	if m.Date == res[0].Date {
		config.G_LOG.Error("CollyCNPPI The latest data is the same as the database ", zap.Any("date", *res[0]))
		return
	}
	if err := do.Create(res[0]); err != nil {
		config.G_LOG.Error("CollyCNPPI Create ", zap.Error(err))
	}
	config.G_LOG.Debug("End CollyCNPPI ")
}

// 社会融资总量
func CollyCNSocialFinancingFlow() {
	config.G_LOG.Debug("start CollyCNSocialFinancingFlow ")
	t := q.SocialFinancingFlow
	do := t.WithContext(context.Background())

	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error("find db", zap.Error(err))
		return
	}
	if m.Date == utils.FirstDayOfLastMonth() {
		config.G_LOG.Error("db is exist", zap.Any("date", m.Date))
		return
	}
	// 无，爬最新数据
	htmlContent := ""
	url := `http://www.pbc.gov.cn/diaochatongjisi/116219/116319/index.html`
	chromeCtx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	err = chromedp.Run(chromeCtx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#l_con`),
		chromedp.OuterHTML(`.qhkuang2`, &htmlContent, chromedp.ByQuery),
		chromedp.Stop(),
	})
	if err != nil {
		config.G_LOG.Error("CollyCNSocialFinancingFlow Run err", zap.Error(err))
		return
	}

	// dom选择器
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		config.G_LOG.Error("CollyCNSocialFinancingFlow NewDocumentFromReader", zap.Error(err))
		return
	}
	dom.Find(`table > tbody > tr > td > table > tbody > tr:nth-child(5) a`).Each(func(i int, s *goquery.Selection) {
		href, e := s.Attr("href")
		if e {
			href = strings.TrimSpace(href)
			err := chromedp.Run(chromeCtx, chromedp.Tasks{
				chromedp.Navigate(`http://www.pbc.gov.cn` + href),
				chromedp.WaitVisible(`tbody`),
				chromedp.OuterHTML(`.border_nr > tbody`, &htmlContent, chromedp.ByQuery),
				chromedp.Stop(),
			})
			if err != nil {
				config.G_LOG.Error("CollyCNSocialFinancingFlow chromedp run", zap.Error(err))
				return
			}
			dom, err = goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
			if err != nil {
				config.G_LOG.Error("CollyCNSocialFinancingFlow NewDocumentFromReader", zap.Error(err))
				return
			}
			dom.Find("table:nth-child(1) a").Each(func(i int, s *goquery.Selection) {
				content := strings.TrimSpace(s.Text())
				href, e := s.Attr("href")
				if e && content == "htm" {
					if h := strings.TrimSpace(href); h != "" {
						err := chromedp.Run(chromeCtx, chromedp.Tasks{
							chromedp.Navigate(`http://www.pbc.gov.cn` + h),
							chromedp.WaitVisible(`tbody`),
							chromedp.OuterHTML(`table`, &htmlContent, chromedp.ByQuery),
							chromedp.Stop(),
						})
						if err != nil {
							config.G_LOG.Error("CollyCNSocialFinancingFlow Run", zap.Error(err))
							return
						}
						datalist, err := initialize.ProcessedSocialFinancingFlowTable(htmlContent)
						if err != nil {
							config.G_LOG.Error("CollyCNSocialFinancingFlow ProcessedSocialFinancingStockTable", zap.Error(err))
							return
						}
						for _, v := range datalist {
							_, err := do.Where(t.Date.Eq(v.Date)).First()
							if errors.Is(err, gorm.ErrRecordNotFound) {
								do.Create(v)
								config.G_LOG.Debug("CollyCNSocialFinancingFlow Create successful ", zap.Any("data", *v))
							}
						}
					}
				}
			})
		}
	})
	config.G_LOG.Debug("End CollyCNSocialFinancingFlow ")
}

// 社会融资存量
func CollyCNSocialFinancingStock() {
	config.G_LOG.Debug("start CollyCNSocialFinancingStock ")
	t := q.SocialFinancingStock
	do := t.WithContext(context.Background())

	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error("find db", zap.Error(err))
		return
	}
	if m.Date == utils.FirstDayOfLastMonth() {
		config.G_LOG.Error("db is exist", zap.Any("date", m.Date))
		return
	}
	// 无，爬最新数据
	htmlContent := ""
	url := `http://www.pbc.gov.cn/diaochatongjisi/116219/116319/index.html`
	chromeCtx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	err = chromedp.Run(chromeCtx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#l_con`),
		chromedp.OuterHTML(`.qhkuang2`, &htmlContent, chromedp.ByQuery),
		chromedp.Stop(),
	})
	if err != nil {
		config.G_LOG.Error("CollyCNSocialFinancingStock Run err", zap.Error(err))
		return
	}

	// dom选择器
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		config.G_LOG.Error("CollyCNSocialFinancingStock NewDocumentFromReader", zap.Error(err))
		return
	}
	dom.Find(`table > tbody > tr > td > table > tbody > tr:nth-child(5) a`).Each(func(i int, s *goquery.Selection) {
		href, e := s.Attr("href")
		if e {
			href = strings.TrimSpace(href)
			err := chromedp.Run(chromeCtx, chromedp.Tasks{
				chromedp.Navigate(`http://www.pbc.gov.cn` + href),
				chromedp.WaitVisible(`tbody`),
				chromedp.OuterHTML(`.border_nr > tbody`, &htmlContent, chromedp.ByQuery),
				chromedp.Stop(),
			})
			if err != nil {
				config.G_LOG.Error("CollyCNSocialFinancingStock chromedp run", zap.Error(err))
				return
			}
			dom, err = goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
			if err != nil {
				config.G_LOG.Error("CollyCNSocialFinancingStock NewDocumentFromReader", zap.Error(err))
				return
			}
			dom.Find("table:nth-child(2) a").Each(func(i int, s *goquery.Selection) {
				content := strings.TrimSpace(s.Text())
				href, e := s.Attr("href")
				if e && content == "htm" {
					if h := strings.TrimSpace(href); h != "" {
						err := chromedp.Run(chromeCtx, chromedp.Tasks{
							chromedp.Navigate(`http://www.pbc.gov.cn` + h),
							chromedp.WaitVisible(`tbody`),
							chromedp.OuterHTML(`table`, &htmlContent, chromedp.ByQuery),
							chromedp.Stop(),
						})
						if err != nil {
							config.G_LOG.Error("CollyCNSocialFinancingStock Run", zap.Error(err))
							return
						}
						datalist, err := initialize.ProcessedSocialFinancingStockTable(htmlContent)
						if err != nil {
							config.G_LOG.Error("CollyCNSocialFinancingStock ProcessedSocialFinancingStockTable", zap.Error(err))
							return
						}
						for _, v := range datalist {
							_, err := do.Where(t.Date.Eq(v.Date)).First()
							if errors.Is(err, gorm.ErrRecordNotFound) {
								do.Create(v)
								config.G_LOG.Debug("CollyCNSocialFinancingFlow Create successful ", zap.Any("data", *v))
							}
						}
					}
				}
			})
		}
	})
	config.G_LOG.Debug("End CollyCNSocialFinancingStock ")
}

// 伦铜 每日9点
func CollyCADFuturesForeignHist() {
	config.G_LOG.Debug("Start CollyCADFuturesForeignHist ")
	t := q.FturesFoewign
	do := t.WithContext(context.Background())

	// 计算本日时间
	tf := time.Now()
	yesterday := tf.AddDate(0, 0, -1).Format("2006-01-02")
	today := tf.Format("2006-01-02")
	url := `https://vip.stock.finance.sina.com.cn/q/view/vFutures_History.php`
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"jys":   "LME",
			"pz":    "CAD",
			"hy":    "",
			"breed": "CAD",
			"type":  "global",
			"start": yesterday,
			"end":   today,
		}).
		Get(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(string(resp.Body())))
	if err != nil {
		log.Fatal(err)
		return
	}
	list := make([]string, 0)
	dom.Find("body > div.wrap > div.middleSection > div.centerColumn > div.historyList > table:nth-child(1) > tbody > tr.tr_2 td").Each(func(i int, s *goquery.Selection) {

		str := strings.TrimSpace(s.Text())
		list = append(list, str)
	})
	f := &model.FturesFoewign{
		Date:   list[0],
		Symbol: "CAD",
		Close:  list[1],
		Open:   list[2],
		High:   list[3],
		Low:    list[4],
		Volume: list[5],
	}
	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error("CollyCADFuturesForeignHist find db", zap.Error(err))
		return
	}
	if strings.Contains(m.Date, f.Date) {
		config.G_LOG.Debug("CollyCADFuturesForeignHist It already exists in the database", zap.Any("data", m))
		return
	}
	do.Create(f)
	config.G_LOG.Debug("CollyCADFuturesForeignHist Create successful ", zap.Any("data", *f))
	config.G_LOG.Debug("End CollyCADFuturesForeignHist ")
}
