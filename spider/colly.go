package spider

import (
	"context"
	"log"
	"strings"

	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chenhaonan-eth/dao/dal/initialize"
	"github.com/chromedp/chromedp"

	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
)

var (
	Client = resty.New()
	do     = initialize.QueryCtxDo
	q      = initialize.Q
)

// 爬取CN社融ppi，从每月9号开始
func CollyCNPPI() {
	// 计算本月时间
	t := time.Now().Format("2006-01")
	log.Println(t)
	// 查询数据库是否存在最近一个月数据
	m, err := do.MacroPpi.Order(q.MacroPpi.Date.Desc()).First()
	if err != nil {
		return
	}
	if strings.Contains(m.Date, t) {
		return
	}
	// 无，获取最新数据
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"max_date": "",
			"category": "ec",
			"attr_id":  "60",
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
		return
	}
	b := resp.Body()
	value := gjson.GetBytes(b, "data.values").Array()
	for _, v := range value {
		vl := v.Array()
		if find := strings.Contains(vl[0].String(), t); find {
			m.Date = vl[0].String()
			m.Ppi = vl[1].String()
			do.MacroPpi.Create(m)
			return
		}
	}
}

// 社会融资存量
func CollyCNSocialFinancingStock() {
	// 计算本月时间
	t := time.Now().Format("2006.01")
	log.Println(t)
	// 查询数据库是否存在最近一个月数据
	m, err := do.SocialFinancingStock.Order(q.MacroPpi.Date.Desc()).First()
	if err != nil {
		return
	}
	if strings.Contains(m.Date, t) {
		return
	}
	// 无，爬最新数据
	htmlContent := ""
	url := `http://www.pbc.gov.cn/diaochatongjisi/116219/116319/index.html`
	// // create context
	chromeCtx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	err = chromedp.Run(chromeCtx, chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#l_con`),
		chromedp.OuterHTML(`.qhkuang2`, &htmlContent, chromedp.ByQuery),
		chromedp.Stop(),
	})
	if err != nil {
		log.Fatalf("Run err : %v\n", err)
		return
	}

	// htmlContent, err := utils.ReaderText("./out.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }
	// dom选择器
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
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
				log.Fatalf("Run err : %v\n", err)
				return
			}
			dom, err = goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
			if err != nil {
				log.Fatal(err)
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
							log.Fatalf("Run err : %v\n", err)
							return
						}
						datalist, err := initialize.ProcessedSocialFinancingStockTable(htmlContent)
						if err != nil {
							log.Fatal(err)
							return
						}
						log.Println(datalist)
						last := datalist[len(datalist)-1]
						if last.Date != m.Date {
							do.SocialFinancingStock.Create(last)
						}
					}
				}
			})
		}
	})
}

// 伦铜 每日
func CollyCADFuturesForeignHist() {
	// TODO
	// 计算本日时间
	// 查询数据库是否存在最近三日数据
	// 无，爬最新数据

}
