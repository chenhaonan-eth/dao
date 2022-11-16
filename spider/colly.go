package spider

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	// "github.com/chenhaonan-eth/dao/core"

	"github.com/PuerkitoBio/goquery"
	"github.com/chenhaonan-eth/dao/dal/model"
	"github.com/chenhaonan-eth/dao/dal/query"
	"github.com/chenhaonan-eth/dao/economic"
	"github.com/chromedp/chromedp"
	"github.com/go-resty/resty/v2"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/robertkrimen/otto"
	"github.com/tidwall/gjson"
)

var (
	client = resty.New()
	q      = query.Q
)

// 获取所有股票市盈率|市净率|股息率|市销率|总市值
func CollyleguleguPePSPb() {
	t := q.PePbPsDvTotalmv
	do := t.WithContext(context.Background())

	c := colly.NewCollector(
		colly.AllowedDomains("legulegu.com", "www.legulegu.com"),
	)

	extensions.RandomUserAgent(c)

	//限速
	err := c.Limit(&colly.LimitRule{
		DomainRegexp: `legulegu\.com`,
		RandomDelay:  5000 * time.Millisecond,
		Parallelism:  10,
	})
	if err != nil {
		log.Fatal(err)
	}

	// On every a element which has href attribute call callback
	c.OnHTML("div.col-md-2 > a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")
		// Print link
		log.Printf("Link found: %q -> %s\n", e.Text, link)
		r, err := Process(e.Text, link)
		if err == nil {
			do.CreateInBatches(r, 5000)
		}
	})

	// Before making a request print "Visiting ..."
	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
	})

	// Set error handler
	c.OnError(func(r *colly.Response, err error) {
		log.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	// Start scraping on https://legulegu.com/stocklist
	c.Visit("https://legulegu.com/stocklist")
}

// 存储  冠龙节能(301151) /s/301151
func Process(name, link string) ([]*model.PePbPsDvTotalmv, error) {
	symbolName := string([]byte(name)[:strings.Index(name, "(")])
	result := []*model.PePbPsDvTotalmv{}
	symbol := strings.Split(link, "/")[2]

	token, err := getToken()
	if err != nil {
		panic(err)
	}
	resp, err := client.R().
		SetQueryParams(map[string]string{
			"token": token,
		}).
		Get(fmt.Sprintf("https://www.legulegu.com/api/s/base-info/%s", symbol))
	if err != nil {
		return result, err
	}
	b := resp.Body()
	v := gjson.GetBytes(b, "data.items")
	arrdata := v.Array()
	for _, v := range arrdata {
		v1 := v.Array()
		m := model.PePbPsDvTotalmv{}
		ivalue := reflect.ValueOf(&m).Elem()
		for i, v := range v1 {
			elem := ivalue.Field(i)
			elem.SetString(v.String())
		}
		m.Name = symbolName
		m.Symbol = symbol
		result = append(result, &m)
	}
	return result, nil
}

func getToken() (string, error) {
	vm := otto.New()
	vm.Run(economic.SCRIPT)
	t := time.Now().Format("2006-01-02")

	token, err := vm.Call("hex", nil, t)
	if err != nil {
		return "", err
	}
	return strings.ToLower(token.String()), nil
}

// 获取 M0 M1 M2存储
func CollyMacroChinaMoneySupply() error {
	t := q.MacroChinaMoneySupply
	do := t.WithContext(context.Background())
	v, err := economic.MacroChinaMoneySupply()
	if err != nil {
		// core.G_LOG.Error("find Colly Macro China Money Supply err", zap.Any("err", err))
		return err
	}
	err = do.CreateInBatches(v, 5000)
	if err != nil {
		// core.G_LOG.Info("find CollyMacroChinaMoneySupply err", zap.Any("err", err))
		return err
	}
	return nil
}

// pmi 存储
func CollyPMI() error {
	t := q.MacroPMI
	do := t.WithContext(context.Background())
	v, err := economic.MacroChinaPmiYearly()
	if err != nil {
		// core.G_LOG.Error("store Macro China Pmi Yearly err", zap.Any("err", err))
		return err
	}
	for k, v := range v {
		pmi := &model.MacroPMI{}
		pmi.Date = k
		pmi.Pmi = v
		pmi.Country = "cn"
		do.Create(pmi)
	}
	return nil
}

// 社会融资总额
func CollyTotalSocialFinancing() error {
	t := q.MacroChinaShrzgm
	do := t.WithContext(context.Background())
	m, err := economic.MacroChinaShrzgm()
	if err != nil {
		return err
	}
	do.CreateInBatches(m, 5000)
	return nil
}

// GDP
func CollyGDP() error {
	t := q.MacroGDP
	do := t.WithContext(context.Background())
	v, err := economic.MacroChinaGdpYearly()
	if err != nil {
		return err
	}
	for k, v := range v {
		gdp := &model.MacroGDP{}
		gdp.Date = k
		gdp.Gdp = v
		gdp.Country = "cn"
		do.Create(gdp)
	}
	return nil
}

// 社会消费品零售总额
func CollyMacroChinaConsumerGoodsRetail() error {
	v, err := economic.MacroChinaConsumerGoodsRetail()
	if err != nil {
		return err
	}
	t := q.MacroChinaConsumerGoodsRetail
	do := t.WithContext(context.Background())
	do.CreateInBatches(v, 5000)
	return nil
}

// cpi
func CollyMacroChinaCpi() error {
	v, err := economic.MacroChinaCpiYearly()
	if err != nil {
		return err
	}
	t := q.MacroCpi
	do := t.WithContext(context.Background())
	for k, v := range v {
		cpi := &model.MacroCpi{}
		cpi.Date = k
		cpi.Cpi = v
		cpi.Country = "cn"
		do.Create(cpi)
	}
	return nil
}

// ppi
func CollyMacroPpi() error {
	v, err := economic.MacroChinaPpiYearly()
	if err != nil {
		return err
	}
	t := q.MacroPpi
	do := t.WithContext(context.Background())
	for k, v := range v {
		ppi := &model.MacroPpi{}
		ppi.Date = k
		ppi.Ppi = v
		ppi.Country = "cn"
		do.Create(ppi)
	}
	return nil
}

// 社会融资存量
func CollySocialFinancingStock() {
	urlmap, err := getSocialFinancingStockUrlMap()
	if err != nil {
		log.Println(err)
	}
	for k, v := range urlmap {
		log.Printf("%s -> %t\n", k, v)
	}
}

// 获取社会融资url map
func getSocialFinancingStockUrlMap() (map[string]bool, error) {
	htmlContent := ""
	urlMap := make(map[string]bool, 11)
	url := `http://www.pbc.gov.cn/diaochatongjisi/116219/116319/index.html`
	err := GetHttpHtmlContent(getSocialFinancingStockHTML(url, &htmlContent))
	if err != nil {
		log.Fatal(err)
		return urlMap, err
	}
	// dom选择器
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return urlMap, err
	}

	dom.Find(`.guidlevel02_style2`).Each(func(i int, s *goquery.Selection) {
		href, e := s.Attr("href")
		if e {
			content := strings.TrimSpace(s.Text())
			if content == "社会融资规模" {
				href = strings.TrimSpace(href)
				urlMap[href] = false
			}
		}
	})
	return urlMap, nil
}

// 动态加载js
func GetHttpHtmlContent(tasks chromedp.Tasks) error {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("blink-settings", "imagesEnabled=false"),
	}
	//初始化参数，先传一个空的数据
	options = append(chromedp.DefaultExecAllocatorOptions[:], options...)

	c, _ := chromedp.NewExecAllocator(context.Background(), options...)

	// create context
	chromeCtx, cancel := chromedp.NewContext(c, chromedp.WithLogf(log.Printf))
	defer cancel()

	// 执行一个空task, 用提前创建Chrome实例
	chromedp.Run(chromeCtx, make([]chromedp.Action, 0, 1)...)

	//创建一个上下文，超时时间为20s
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 20*time.Second)
	defer cancel()

	// var htmlContent string
	err := chromedp.Run(timeoutCtx, tasks)
	if err != nil {
		log.Fatalf("Run err : %v\n", err)
		return err
	}
	return nil
}

// 获取社融数据html
func getSocialFinancingStockHTML(url string, htmlContent *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(`#l_con`),
		chromedp.OuterHTML(`.qhkuang2 > tbody`, htmlContent, chromedp.ByQuery),
		chromedp.Stop(),
	}
}
