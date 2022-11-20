// 期货
package ftures

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/chenhaonan-eth/dao/dal/model"
	"github.com/go-resty/resty/v2"
)

var Client = resty.New()

// 描述: 新浪财经-外盘商品期货品种代码表数据
// https://finance.sina.com.cn/money/future/hf.html
// return :{"LME镍3个月":"NID","LME铜3个月":"CAD"}
var FuturesMap = map[string]string{
	"NYBOT-棉花": "CT",
	"LME镍3个月":  "NID",
	"LME铅3个月":  "PBD",
	"LME锡3个月":  "SND",
	"LME锌3个月":  "ZSD",
	"LME铝3个月":  "AHD",
	"LME铜3个月":  "CAD",
	"CBOT-黄豆":  "S",
	"CBOT-小麦":  "W",
	"CBOT-玉米":  "C",
	"CBOT-黄豆油": "BO",
	"CBOT-黄豆粉": "SM",
	"日本橡胶":     "TRB",
	"COMEX铜":   "HG",
	"NYMEX天然气": "NG",
	"NYMEX原油":  "CL",
	"COMEX白银":  "SI",
	"COMEX黄金":  "GC",
	"CME-瘦肉猪":  "LHC",
	"布伦特原油":    "OIL",
	"伦敦金":      "XAU",
	"伦敦银":      "XAG",
	"伦敦铂金":     "XPT",
	"伦敦钯金":     "XPD",
	"马棕油":      "FCPO",
	"欧洲碳排放":    "EUA",
}

// 获取所有外盘期货
func ALLFuturesForeignHist() ([]*model.FturesFoewign, error) {
	result := []*model.FturesFoewign{}
	for _, v := range FuturesMap {
		r, err := FuturesForeignHist(v)
		if err != nil {
			log.Printf("%s ->err:%s", v, err.Error())
			return result, err
		}
		result = append(result, r...)
	}
	return result, nil
}

// 外盘期货-历史行情数据-日频率
// http://finance.sina.com.cn/money/future/hf.html
// symbol: NID、CAD
func FuturesForeignHist(symbol string) ([]*model.FturesFoewign, error) {
	result := []*model.FturesFoewign{}
	// 2022_11_14
	t := time.Now().Format("2006_01_02")
	url := fmt.Sprintf("https://stock2.finance.sina.com.cn/futures/api/jsonp.php/var%%20_S%s=/GlobalFuturesService.getGlobalFuturesDailyKLine", t)
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"symbol": symbol,
			"_":      t,
			"source": "web",
		}).
		Get(url)
	if err != nil {
		return result, err
	}
	byBody := resp.Body()
	bydata := byBody[bytes.IndexAny(byBody, "[") : len(byBody)-2]
	json.Unmarshal(bydata, &result)
	return result, nil
}
