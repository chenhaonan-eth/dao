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
	"github.com/chenhaonan-eth/dao/pkg/utils"
	"github.com/chromedp/chromedp"
	"go.uber.org/zap"
	"gorm.io/gorm"

	browser "github.com/EDDYCJY/fake-useragent"
	"github.com/go-resty/resty/v2"
	"github.com/noaway/dateparse"
	"github.com/tidwall/gjson"
)

var (
	Client = resty.New()
	q      = query.Q
)

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
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"columns":     "REPORT_DATE,BASIC_CURRENCY,BASIC_CURRENCY_SAME,BASIC_CURRENCY_SEQUENTIAL,CURRENCY,CURRENCY_SAME,CURRENCY_SEQUENTIAL,FREE_CASH,FREE_CASH_SAME,FREE_CASH_SEQUENTIAL",
			"pageNumber":  "1",
			"pageSize":    "1",
			"sortColumns": "REPORT_DATE",
			"sortTypes":   "-1",
			"source":      "WEB",
			"client":      "WEB",
			"reportName":  "RPT_ECONOMY_CURRENCY_SUPPLY",
			"p":           "1",
			"pageNo":      "1",
			"pageNum":     "1",
			"_":           "1669047266881",
		}).
		Get("https://datacenter-web.eastmoney.com/api/data/v1/get")
	if err != nil {
		config.G_LOG.Error(err.Error())
		return
	}
	b := resp.Body()
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &res)

	if m.Date == res[0].Date {
		config.G_LOG.Error("db is exist !", zap.Any("date", m.Date))
		return
	}
	do.Create(res[0])
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
		"User-Agent":   browser.Random(),
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
		t, err := dateparse.ParseAny(str)
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

// 爬取CN GDP
func CollyCNGDP() {
	// TODO
	// config.G_LOG.Debug("Start CollyCNGDP ")
	// t := q.MacroGDP
	// do := t.WithContext(context.Background())
	// // 计算上月时间
	// t1 := time.Now().Format("2006-01")
	// // 查询数据库是否存在最近一个月数据
	// m, err := do.Order(t.Date.Desc()).First()
	// if err != nil {
	// 	config.G_LOG.Error("find db", zap.Error(err))
	// 	return
	// }
	// if strings.Contains(m.Date, t1) && m.Gdp != "" {
	// 	config.G_LOG.Debug("CollyCNGDP It already exists in the database", zap.Any("data", m))
	// 	return
	// }
	// v, err := jin10("CollyCNGDP", "57")
	// if err != nil {
	// 	config.G_LOG.Error("CollyCNGDP ", zap.Error(err))
	// 	return
	// }
	// for _, v := range v {
	// 	vl := v.Array()
	// 	if find := strings.Contains(vl[0].String(), t1); find {
	// 		if strings.Contains(m.Date, t1) && m.Gdp == "" {
	// 			do.Where(t.Date.Eq(vl[0].String())).Update(t.Gdp, vl[1].String())
	// 			config.G_LOG.Debug("CollyCNGDP Update", zap.Any("date", vl[0].String()), zap.Any("GDP", vl[1].String()))
	// 		}
	// 		if !strings.Contains(m.Date, t1) {
	// 			m.Date = vl[0].String()
	// 			m.Gdp = vl[1].String()
	// 			do.Create(m)
	// 			config.G_LOG.Debug("CollyCNGDP Create", zap.Any("data", *m))
	// 		}
	// 		break
	// 	}
	// }
	// config.G_LOG.Debug("End CollyCNGDP ")
}

// 爬取CN CPI
func CollyCNCPI() {
	// TODO
	// config.G_LOG.Debug("Start CollyCNCPI ")
	// t := q.MacroCpi
	// do := t.WithContext(context.Background())
	// // 计算本月时间
	// t1 := time.Now().Format("2006-01")
	// // 查询数据库是否存在最近一个月数据
	// m, err := do.Order(t.Date.Desc()).First()
	// if err != nil {
	// 	config.G_LOG.Error("find db", zap.Error(err))
	// 	return
	// }
	// if strings.Contains(m.Date, t1) && m.Cpi != "" {
	// 	config.G_LOG.Debug("CollyCNCPI It already exists in the database", zap.Any("data", m))
	// 	return
	// }
	// v, err := jin10("CollyCNPPI", "56")
	// if err != nil {
	// 	config.G_LOG.Error("CollyCNPPI ", zap.Error(err))
	// 	return
	// }
	// for _, v := range v {
	// 	vl := v.Array()
	// 	if find := strings.Contains(vl[0].String(), t1); find {
	// 		if strings.Contains(m.Date, t1) && m.Cpi == "" {
	// 			do.Where(t.Date.Eq(vl[0].String())).Update(t.Cpi, vl[1].String())
	// 			config.G_LOG.Debug("CollyCNCPI Update", zap.Any("date", vl[0].String()), zap.Any("PPI", vl[1].String()))
	// 		}
	// 		if !strings.Contains(m.Date, t1) {
	// 			m.Date = vl[0].String()
	// 			m.Cpi = vl[1].String()
	// 			do.Create(m)
	// 			config.G_LOG.Debug("CollyCNCPI Create", zap.Any("data", *m))
	// 		}
	// 		break
	// 	}
	// }

	// config.G_LOG.Debug("End CollyCNCPI ")
}

// 爬取CN社融ppi，从每月9号开始
func CollyCNPPI() {
	// TODO
	// config.G_LOG.Debug("start CollyCNPPI ")
	// t := q.MacroPpi
	// do := t.WithContext(context.Background())

	// // 计算本月时间
	// t1 := time.Now().Format("2006-01")
	// log.Println(t)
	// // 查询数据库是否存在最近一个月数据
	// m, err := do.Order(t.Date.Desc()).First()
	// if err != nil {
	// 	config.G_LOG.Error("CollyCNPPI ", zap.Error(err))
	// 	return
	// }
	// if strings.Contains(m.Date, t1) {
	// 	config.G_LOG.Debug("CollyCNPPI It already exists in the database", zap.Any("data", m))
	// 	return
	// }
	// // 无，获取最新数据
	// value, err := jin10("CollyCNPPI", "60")
	// if err != nil {
	// 	config.G_LOG.Error("CollyCNPPI ", zap.Error(err))
	// 	return
	// }
	// for _, v := range value {
	// 	vl := v.Array()
	// 	if find := strings.Contains(vl[0].String(), t1); find {
	// 		m.Date = vl[0].String()
	// 		m.Ppi = vl[1].String()
	// 		err := do.Create(m)
	// 		if err != nil {
	// 			config.G_LOG.Error("CollyCNPPI ", zap.Error(err))
	// 		}
	// 		config.G_LOG.Debug("CollyCNPPI Create", zap.Any("data", *m))
	// 		return
	// 	}
	// }
	// config.G_LOG.Debug("End CollyCNPPI ")
}

func jin10(name, id string) ([]gjson.Result, error) {
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"max_date": "",
			"category": "ec",
			"attr_id":  id,
			"_":        strconv.FormatInt(time.Now().Unix(), 10),
		}).SetHeaders(map[string]string{
		"accept":          "*/*",
		"accept-encoding": "gzip, deflate, br",
		"accept-language": "zh-CN,zh;q=0.9,en;q=0.8",
		"cache-control":   "no-cache",
		"origin":          "https://datacenter.jin10.com",
		"pragma":          "no-cache",
		"referer":         "https://datacenter.jin10.com/reportType/dc_usa_michigan_consumer_sentiment",
		"sec-fetch-dest":  "empty",
		"sec-fetch-mode":  "cors",
		"sec-fetch-site":  "same-site",
		"user-agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36",
		"x-app-id":        "rU6QIu7JHe2gOUeR",
		"x-csrf-token":    "",
		"x-version":       "1.0.0",
	}).
		Get("https://datacenter-api.jin10.com/reports/list_v2")
	if err != nil {
		config.G_LOG.Error("get jin10 ", zap.Any("name", name), zap.Error(err))
		return nil, err
	}
	b := resp.Body()
	value := gjson.GetBytes(b, "data.values").Array()
	return value, nil
}

// 社会融资存量
func CollyCNSocialFinancingStock() {
	config.G_LOG.Debug("start CollyCNSocialFinancingStock ")
	t := q.SocialFinancingStock
	do := t.WithContext(context.Background())

	// 计算本月时间
	tf := time.Now().Format("2006.01")
	// 查询数据库是否存在最近一个月数据
	m, err := do.Order(t.Date.Desc()).First()
	if err != nil {
		config.G_LOG.Error("find db", zap.Error(err))
		return
	}
	if strings.Contains(m.Date, tf) {
		config.G_LOG.Debug("CollyCNSocialFinancingStock It already exists in the database", zap.Any("data", m))
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
	// log.Println(htmlContent)
	dom.Find(`table > tbody > tr > td > table > tbody > tr:nth-child(5) a`).Each(func(i int, s *goquery.Selection) {
		t := s.Text()
		log.Println(t)
		href, e := s.Attr("href")
		if e {
			log.Println(href)
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
						log.Println(datalist)
						last := datalist[len(datalist)-1]
						if last.Date != m.Date {
							do.Create(last)
							config.G_LOG.Debug("CollyCNSocialFinancingStock Create successful ", zap.Any("data", *last))
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
