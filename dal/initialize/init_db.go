package initialize

import (
	"context"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chenhaonan-eth/dao/core"
	"github.com/chenhaonan-eth/dao/dal"
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
	client     = resty.New()
	q          = query.Q
	queryCtxDo *query.QueryCtx
)

func InitDB() {
	// 创建本地数据库
	dal.DB = dal.ConnectDB(core.G_Config.System.Dsn).Debug()
	query.SetDefault(dal.DB)
	queryCtxDo = q.WithContext(context.Background())
	// 初始化数据库
	for _, tabledb := range model.OpArgsMap {
		// init db table 自动检测是否存在
		dal.DB.AutoMigrate(tabledb)
	}
	// 检查数据库是否有数据 没有的话加载数据
	if err := initTotalSocialFinancing(); err != nil {
		log.Println(err)
	}
	// TODO 东财接口失效 等待重写代码
	// if err := initMacroChinaMoneySupply(); err != nil {
	// 	log.Println(err)
	// }
	// if err := initMacroChinaConsumerGoodsRetail(); err != nil {
	// 	log.Println(err)
	// }
	if err := initPMI(); err != nil {
		log.Println(err)
	}
	if err := initGDP(); err != nil {
		log.Println(err)
	}

	if err := initMacroChinaCpi(); err != nil {
		log.Println(err)
	}
	if err := initMacroPpi(); err != nil {
		log.Println(err)
	}
	if err := initSocialFinancingStock(); err != nil {
		log.Println(err)
	}
	if err := initSH300PE(); err != nil {
		log.Println(err)
	}
	if err := initBondZhUsRate(); err != nil {
		log.Println(err)
	}
}

// 国债收益率
func initBondZhUsRate() error {
	do := queryCtxDo.BondZhUsRate
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
	v, err := economic.BondZhUsRate()
	if err != nil {
		// core.G_LOG.Error("find Colly Macro China Money Supply err", zap.Any("err", err))
		return err
	}
	err = do.CreateInBatches(v, 2048)
	if err != nil {
		// core.G_LOG.Info("find CollyMacroChinaMoneySupply err", zap.Any("err", err))
		return err
	}
	return nil
}

// 沪深300市盈率
func initSH300PE() error {
	do := queryCtxDo.SH300PE
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
	v, err := economic.SH300PE()
	if err != nil {
		// core.G_LOG.Error("find Colly Macro China Money Supply err", zap.Any("err", err))
		return err
	}
	err = do.CreateInBatches(v, 1024)
	if err != nil {
		// core.G_LOG.Info("find CollyMacroChinaMoneySupply err", zap.Any("err", err))
		return err
	}
	return nil
}

// 获取所有股票市盈率|市净率|股息率|市销率|总市值
// 写入sqlite数据库时间较长
func initleguleguPePSPb() {
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
func initMacroChinaMoneySupply() error {
	do := queryCtxDo.MacroChinaMoneySupply

	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
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
func initPMI() error {
	do := queryCtxDo.MacroPMI
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
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
func initTotalSocialFinancing() error {
	do := queryCtxDo.SocialFinancingFlow
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m == nil {
		m, err := economic.MacroChinaShrzgm()
		if err != nil {
			return err
		}
		err = do.CreateInBatches(m, 5000)
		if err != nil {
			return err
		}
	}
	return nil
}

// GDP
func initGDP() error {
	do := queryCtxDo.MacroGDP
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
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
func initMacroChinaConsumerGoodsRetail() error {
	do := queryCtxDo.MacroChinaConsumerGoodsRetail
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
	v, err := economic.MacroChinaConsumerGoodsRetail()
	if err != nil {
		return err
	}
	do.CreateInBatches(v, 512)
	return nil
}

// cpi
func initMacroChinaCpi() error {
	do := queryCtxDo.MacroCpi
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
	v, err := economic.MacroChinaCpiYearly()
	if err != nil {
		return err
	}
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
func initMacroPpi() error {
	do := queryCtxDo.MacroPpi
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
	v, err := economic.MacroChinaPpiYearly()
	if err != nil {
		return err
	}
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
func initSocialFinancingStock() error {
	do := queryCtxDo.SocialFinancingStock
	m, err := do.First()
	if err != nil {
		log.Println(err)
	}
	if m != nil {
		return nil
	}
	var htmlContent string
	// macroChinaShrzgmList := []*model.SocialFinancingStock{}
	urlSli, err := getSocialFinancingStockUrlMap()
	if err != nil {
		log.Println(err)
	}
	for _, v := range urlSli {
		// fmt.Printf("%d -> %s\n", i, v)
		url := `http://www.pbc.gov.cn` + v
		err = GetHttpHtmlContent(getSocialFinancingStockHTML(url, `tbody`, `table`, &htmlContent))
		if err != nil {
			log.Fatal(err)
		}
		datalist, err := processedSocialFinancingStockTable(htmlContent)
		if err != nil {
			log.Fatal(err)
			return err
		}
		do.CreateInBatches(datalist, 512)
		// macroChinaShrzgmList = append(macroChinaShrzgmList, datalist...)
	}
	return nil
}

// 获取社会融资url map
func getSocialFinancingStockUrlMap() ([]string, error) {
	htmlContent := ""
	urlSli := make([]string, 0)
	url := `http://www.pbc.gov.cn/diaochatongjisi/116219/116319/index.html`
	err := GetHttpHtmlContent(getSocialFinancingStockHTML(url, `#l_con`, `.qhkuang2 > tbody`, &htmlContent))
	if err != nil {
		log.Fatal(err)
		return urlSli, err
	}
	// dom选择器
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
		return urlSli, err
	}
	dom.Find(`.guidlevel02_style2`).Each(func(i int, s *goquery.Selection) {
		var html string
		href, e := s.Attr("href")
		if e {
			content := strings.TrimSpace(s.Text())
			if content == "社会融资规模" {
				href = strings.TrimSpace(href)
				// 进一步获取 所有talbe 的url
				url := `http://www.pbc.gov.cn` + href
				err := GetHttpHtmlContent(getSocialFinancingStockHTML(url, `tbody`, `.border_nr > tbody`, &html))
				if err != nil {
					log.Fatal(err)
				}
				dom, err = goquery.NewDocumentFromReader(strings.NewReader(html))
				if err != nil {
					log.Fatal(err)
				}
				dom.Find("table:nth-child(2) a").Each(func(i int, s *goquery.Selection) {
					content := strings.TrimSpace(s.Text())
					href, e := s.Attr("href")
					if e && content == "htm" {
						if h := strings.TrimSpace(href); h != "" {
							urlSli = append(urlSli, h)
						}
					}
				})
			}
		}
	})
	return urlSli, nil
}

// 动态加载js
func GetHttpHtmlContent(tasks chromedp.Tasks) error {
	options := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
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
	timeoutCtx, cancel := context.WithTimeout(chromeCtx, 40*time.Second)
	defer cancel()

	// var htmlContent string
	err := chromedp.Run(timeoutCtx, tasks)
	if err != nil {
		log.Fatalf("Run err : %v\n", err)
		return err
	}
	return nil
}

// 获取html
func getSocialFinancingStockHTML(url, waitVisible, sel string, htmlContent *string) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(url),
		chromedp.WaitVisible(waitVisible),
		chromedp.OuterHTML(sel, htmlContent, chromedp.ByQuery),
		chromedp.Stop(),
	}
}

// 处理 社会融资规模存量 HTML table
// return 二维数组
func processedSocialFinancingStockTable(htmlDom string) ([]*model.SocialFinancingStock, error) {
	dataList := make([][]string, 0)
	macroChinaShrzgmList := []*model.SocialFinancingStock{}
	dom, err := goquery.NewDocumentFromReader(strings.NewReader(htmlDom))
	if err != nil {
		log.Fatal(err)
		return macroChinaShrzgmList, err
	}

	// 获取日期
	list := make([]string, 0)
	dom.Find("tbody > tr:nth-child(5) td").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			return
		}
		str := strings.TrimSpace(s.Text())
		if s.Text() != "　" && len(str) != 0 {
			list = append(list, str)
		}
	})
	dataList = append(dataList, list)
	log.Println(list[0])
	time := strings.Split(list[0], ".")
	// 22、21、20年 9~19 td
	// 19年需要特殊处理
	// 18年多了一项地方政府专项债券
	// 17、16年 9~16 td
	// 15年时间需要特殊处理
	if time[0] == "2018" {
		for i := 9; i < 18; i++ {
			list := make([]string, 0)
			dom.Find(fmt.Sprintf("tbody > tr:nth-child(%d) td", i)).Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					return
				}
				str := strings.TrimSpace(s.Text())
				if s.Text() != "　" && len(str) != 0 {
					list = append(list, str)
				}
			})
			dataList = append(dataList, list)
		}
		socialFinancingStockConvertToObject1(&macroChinaShrzgmList, dataList, true)
	} else if time[0] == "2015" {
		dataList[0] = []string{"2015.3", "2015.6", "2015.9", "2015.12"}
		for i := 9; i < 17; i++ {
			list := make([]string, 0)
			dom.Find(fmt.Sprintf("tbody > tr:nth-child(%d) td", i)).Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					return
				}
				str := strings.TrimSpace(s.Text())
				if s.Text() != "　" && len(str) != 0 {
					list = append(list, str)
				}
			})
			dataList = append(dataList, list)
		}
		socialFinancingStockConvertToObject1(&macroChinaShrzgmList, dataList, false)
	} else if time[0] == "2017" || time[0] == "2016" {
		for i := 9; i < 17; i++ {
			list := make([]string, 0)
			dom.Find(fmt.Sprintf("tbody > tr:nth-child(%d) td", i)).Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					return
				}
				str := strings.TrimSpace(s.Text())
				if s.Text() != "　" && len(str) != 0 {
					list = append(list, str)
				}
			})
			dataList = append(dataList, list)
		}
		// 转换为对象
		socialFinancingStockConvertToObject1(&macroChinaShrzgmList, dataList, false)
	} else if time[0] == "2019" {
		for i := 9; i < 30; i += 2 {
			list := make([]string, 0)
			dom.Find(fmt.Sprintf("tbody > tr:nth-child(%d) td", i)).Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					return
				}
				str := strings.TrimSpace(s.Text())
				if s.Text() != "　" && len(str) != 0 {
					list = append(list, str)
				}
			})
			dataList = append(dataList, list)
		}
		// 转换为对象
		socialFinancingStockConvertToObject(&macroChinaShrzgmList, dataList)
	} else {
		for i := 9; i < 20; i++ {
			list := make([]string, 0)
			dom.Find(fmt.Sprintf("tbody > tr:nth-child(%d) td", i)).Each(func(i int, s *goquery.Selection) {
				if i == 0 {
					return
				}
				// fmt.Println("--->" + s.Text() + "<---")
				if s.Text() != "　" {
					list = append(list, s.Text())
				}
			})
			dataList = append(dataList, list)
		}
		// 转换为对象
		socialFinancingStockConvertToObject(&macroChinaShrzgmList, dataList)
	}
	return macroChinaShrzgmList, err
}

// 将表格转换为对象
func socialFinancingStockConvertToObject(macroChinaShrzgmList *[]*model.SocialFinancingStock, dataList [][]string) {
	// slice需要扩容的情况下，需要传入slice
	count := len(dataList[1]) / 2
	for i := 0; i < len(dataList[0]); i++ {
		if i == count {
			return
		}
		data := &model.SocialFinancingStock{
			Date:                            dataList[0][i],
			Tiosfs:                          dataList[1][i*2],
			TiosfsGrowthRate:                dataList[1][i*2+1],
			Rmblaon:                         dataList[2][i*2],
			RmblaonGrowthRate:               dataList[2][i*2+1],
			Forcloan:                        dataList[3][i*2],
			ForcloanGrowthRate:              dataList[3][i*2+1],
			Entrustloan:                     dataList[4][i*2],
			EntrustloanGrowthRate:           dataList[4][i*2+1],
			Trustloan:                       dataList[5][i*2],
			TrustloanGrowthRate:             dataList[5][i*2+1],
			Ndbab:                           dataList[6][i*2],
			NdbabGrowthRate:                 dataList[6][i*2+1],
			Bibae:                           dataList[7][i*2],
			BibaeGrowthRate:                 dataList[7][i*2+1],
			GovernmentBonds:                 dataList[8][i*2],
			GovernmentBondsGrowthRate:       dataList[8][i*2+1],
			Sfinfe:                          dataList[9][i*2],
			SfinfeGrowthRate:                dataList[9][i*2+1],
			AssetBackedSecurities:           dataList[10][i*2],
			AssetBackedSecuritiesGrowthRate: dataList[10][i*2+1],
			LoansWrittenOff:                 dataList[11][i*2],
			LoansWrittenOffGrowthRate:       dataList[11][i*2+1],
		}
		*macroChinaShrzgmList = append(*macroChinaShrzgmList, data)
	}
}

func socialFinancingStockConvertToObject1(macroChinaShrzgmList *[]*model.SocialFinancingStock, dataList [][]string, is2018 bool) {
	index := 8
	if is2018 {
		index = 9 // 除去地方政府专项债券
	}
	for i := 0; i < len(dataList[0]); i++ {
		data := &model.SocialFinancingStock{
			Date:                  dataList[0][i],
			Tiosfs:                dataList[1][i*2],
			TiosfsGrowthRate:      dataList[1][i*2+1],
			Rmblaon:               dataList[2][i*2],
			RmblaonGrowthRate:     dataList[2][i*2+1],
			Forcloan:              dataList[3][i*2],
			ForcloanGrowthRate:    dataList[3][i*2+1],
			Entrustloan:           dataList[4][i*2],
			EntrustloanGrowthRate: dataList[4][i*2+1],
			Trustloan:             dataList[5][i*2],
			TrustloanGrowthRate:   dataList[5][i*2+1],
			Ndbab:                 dataList[6][i*2],
			NdbabGrowthRate:       dataList[6][i*2+1],
			Bibae:                 dataList[7][i*2],
			BibaeGrowthRate:       dataList[7][i*2+1],
			Sfinfe:                dataList[index][i*2],
			SfinfeGrowthRate:      dataList[index][i*2+1],
		}
		*macroChinaShrzgmList = append(*macroChinaShrzgmList, data)
	}
}
