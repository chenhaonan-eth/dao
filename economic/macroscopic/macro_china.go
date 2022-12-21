// 宏观数据-中国
package macroscopic

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/chenhaonan-eth/dao/dal/model"
	"github.com/chenhaonan-eth/dao/pkg/utils"
	"github.com/go-resty/resty/v2"
	"github.com/robertkrimen/otto"
	"github.com/shopspring/decimal"
	"github.com/tidwall/gjson"
	"golang.org/x/text/encoding/simplifiedchinese"
)

var Client = resty.New()

const SCRIPT = `
function e(n) {
    var e, t, r = "", o = -1, f;
    if (n && n.length) {
        f = n.length;
        while ((o += 1) < f) {
            e = n.charCodeAt(o);
            t = o + 1 < f ? n.charCodeAt(o + 1) : 0;
            if (55296 <= e && e <= 56319 && 56320 <= t && t <= 57343) {
                e = 65536 + ((e & 1023) << 10) + (t & 1023);
                o += 1
            }
            if (e <= 127) {
                r += String.fromCharCode(e)
            } else if (e <= 2047) {
                r += String.fromCharCode(192 | e >>> 6 & 31, 128 | e & 63)
            } else if (e <= 65535) {
                r += String.fromCharCode(224 | e >>> 12 & 15, 128 | e >>> 6 & 63, 128 | e & 63)
            } else if (e <= 2097151) {
                r += String.fromCharCode(240 | e >>> 18 & 7, 128 | e >>> 12 & 63, 128 | e >>> 6 & 63, 128 | e & 63)
            }
        }
    }
    return r
}

function t(n) {
    var e, t, r, o, f, i = [], h;
    e = t = r = o = f = 0;
    if (n && n.length) {
        h = n.length;
        n += "";
        while (e < h) {
            r = n.charCodeAt(e);
            t += 1;
            if (r < 128) {
                i[t] = String.fromCharCode(r);
                e += 1
            } else if (r > 191 && r < 224) {
                o = n.charCodeAt(e + 1);
                i[t] = String.fromCharCode((r & 31) << 6 | o & 63);
                e += 2
            } else {
                o = n.charCodeAt(e + 1);
                f = n.charCodeAt(e + 2);
                i[t] = String.fromCharCode((r & 15) << 12 | (o & 63) << 6 | f & 63);
                e += 3
            }
        }
    }
    return i.join("")
}

function r(n, e) {
    var t = (n & 65535) + (e & 65535)
        , r = (n >> 16) + (e >> 16) + (t >> 16);
    return r << 16 | t & 65535
}

function o(n, e) {
    return n << e | n >>> 32 - e
}

function f(n, e) {
    var t = e ? "0123456789ABCDEF" : "0123456789abcdef", r = "", o, f = 0, i = n.length;
    for (; f < i; f += 1) {
        o = n.charCodeAt(f);
        r += t.charAt(o >>> 4 & 15) + t.charAt(o & 15)
    }
    return r
}

function i(n) {
    var e, t = n.length, r = "";
    for (e = 0; e < t; e += 1) {
        r += String.fromCharCode(n.charCodeAt(e) & 255, n.charCodeAt(e) >>> 8 & 255)
    }
    return r
}

function h(n) {
    var e, t = n.length, r = "";
    for (e = 0; e < t; e += 1) {
        r += String.fromCharCode(n.charCodeAt(e) >>> 8 & 255, n.charCodeAt(e) & 255)
    }
    return r
}

function u(n) {
    var e, t = n.length * 32, r = "";
    for (e = 0; e < t; e += 8) {
        r += String.fromCharCode(n[e >> 5] >>> 24 - e % 32 & 255)
    }
    return r
}

function a(n) {
    var e, t = n.length * 32, r = "";
    for (e = 0; e < t; e += 8) {
        r += String.fromCharCode(n[e >> 5] >>> e % 32 & 255)
    }
    return r
}

function c(n) {
    var e, t = n.length * 8, r = Array(n.length >> 2), o = r.length;
    for (e = 0; e < o; e += 1) {
        r[e] = 0
    }
    for (e = 0; e < t; e += 8) {
        r[e >> 5] |= (n.charCodeAt(e / 8) & 255) << e % 32
    }
    return r
}

function l(n) {
    var e, t = n.length * 8, r = Array(n.length >> 2), o = r.length;
    for (e = 0; e < o; e += 1) {
        r[e] = 0
    }
    for (e = 0; e < t; e += 8) {
        r[e >> 5] |= (n.charCodeAt(e / 8) & 255) << 24 - e % 32
    }
    return r
}

function D(n, e) {
    var t = e.length, r = Array(), o, f, i, h, u, a, c, l;
    a = Array(Math.ceil(n.length / 2));
    h = a.length;
    for (o = 0; o < h; o += 1) {
        a[o] = n.charCodeAt(o * 2) << 8 | n.charCodeAt(o * 2 + 1)
    }
    while (a.length > 0) {
        u = Array();
        i = 0;
        for (o = 0; o < a.length; o += 1) {
            i = (i << 16) + a[o];
            f = Math.floor(i / t);
            i -= f * t;
            if (u.length > 0 || f > 0) {
                u[u.length] = f
            }
        }
        r[r.length] = i;
        a = u
    }
    c = "";
    for (o = r.length - 1; o >= 0; o--) {
        c += e.charAt(r[o])
    }
    l = Math.ceil(n.length * 8 / (Math.log(e.length) / Math.log(2)));
    for (o = c.length; o < l; o += 1) {
        c = e[0] + c
    }
    return c
}

function B(n, e) {
    var t = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/", r = "", o = n.length, f, i, h;
    e = e || "=";
    for (f = 0; f < o; f += 3) {
        h = n.charCodeAt(f) << 16 | (f + 1 < o ? n.charCodeAt(f + 1) << 8 : 0) | (f + 2 < o ? n.charCodeAt(f + 2) : 0);
        for (i = 0; i < 4; i += 1) {
            if (f * 8 + i * 6 > n.length * 8) {
                r += e
            } else {
                r += t.charAt(h >>> 6 * (3 - i) & 63)
            }
        }
    }
    return r
}

function hex(n) {
    return f(u(n, h), t)
}

function u(n) {
    n = h ? e(n) : n;
    return a(C(c(n), n.length * 8))
}

function l(n, t) {
    var r, o, f, i, u;
    n = h ? e(n) : n;
    t = h ? e(t) : t;
    r = c(n);
    if (r.length > 16) {
        r = C(r, n.length * 8)
    }
    o = Array(16),
        f = Array(16);
    for (u = 0; u < 16; u += 1) {
        o[u] = r[u] ^ 909522486;
        f[u] = r[u] ^ 1549556828
    }
    i = C(o.concat(c(t)), 512 + t.length * 8);
    return a(C(f.concat(i), 512 + 128))
}

function C(n, e) {
    var t, o, f, i, h, u = 1732584193, a = -271733879, c = -1732584194, l = 271733878;
    n[e >> 5] |= 128 << e % 32;
    n[(e + 64 >>> 9 << 4) + 14] = e;
    for (t = 0; t < n.length; t += 16) {
        o = u;
        f = a;
        i = c;
        h = l;
        u = s(u, a, c, l, n[t + 0], 7, -680876936);
        l = s(l, u, a, c, n[t + 1], 12, -389564586);
        c = s(c, l, u, a, n[t + 2], 17, 606105819);
        a = s(a, c, l, u, n[t + 3], 22, -1044525330);
        u = s(u, a, c, l, n[t + 4], 7, -176418897);
        l = s(l, u, a, c, n[t + 5], 12, 1200080426);
        c = s(c, l, u, a, n[t + 6], 17, -1473231341);
        a = s(a, c, l, u, n[t + 7], 22, -45705983);
        u = s(u, a, c, l, n[t + 8], 7, 1770035416);
        l = s(l, u, a, c, n[t + 9], 12, -1958414417);
        c = s(c, l, u, a, n[t + 10], 17, -42063);
        a = s(a, c, l, u, n[t + 11], 22, -1990404162);
        u = s(u, a, c, l, n[t + 12], 7, 1804603682);
        l = s(l, u, a, c, n[t + 13], 12, -40341101);
        c = s(c, l, u, a, n[t + 14], 17, -1502002290);
        a = s(a, c, l, u, n[t + 15], 22, 1236535329);
        u = w(u, a, c, l, n[t + 1], 5, -165796510);
        l = w(l, u, a, c, n[t + 6], 9, -1069501632);
        c = w(c, l, u, a, n[t + 11], 14, 643717713);
        a = w(a, c, l, u, n[t + 0], 20, -373897302);
        u = w(u, a, c, l, n[t + 5], 5, -701558691);
        l = w(l, u, a, c, n[t + 10], 9, 38016083);
        c = w(c, l, u, a, n[t + 15], 14, -660478335);
        a = w(a, c, l, u, n[t + 4], 20, -405537848);
        u = w(u, a, c, l, n[t + 9], 5, 568446438);
        l = w(l, u, a, c, n[t + 14], 9, -1019803690);
        c = w(c, l, u, a, n[t + 3], 14, -187363961);
        a = w(a, c, l, u, n[t + 8], 20, 1163531501);
        u = w(u, a, c, l, n[t + 13], 5, -1444681467);
        l = w(l, u, a, c, n[t + 2], 9, -51403784);
        c = w(c, l, u, a, n[t + 7], 14, 1735328473);
        a = w(a, c, l, u, n[t + 12], 20, -1926607734);
        u = F(u, a, c, l, n[t + 5], 4, -378558);
        l = F(l, u, a, c, n[t + 8], 11, -2022574463);
        c = F(c, l, u, a, n[t + 11], 16, 1839030562);
        a = F(a, c, l, u, n[t + 14], 23, -35309556);
        u = F(u, a, c, l, n[t + 1], 4, -1530992060);
        l = F(l, u, a, c, n[t + 4], 11, 1272893353);
        c = F(c, l, u, a, n[t + 7], 16, -155497632);
        a = F(a, c, l, u, n[t + 10], 23, -1094730640);
        u = F(u, a, c, l, n[t + 13], 4, 681279174);
        l = F(l, u, a, c, n[t + 0], 11, -358537222);
        c = F(c, l, u, a, n[t + 3], 16, -722521979);
        a = F(a, c, l, u, n[t + 6], 23, 76029189);
        u = F(u, a, c, l, n[t + 9], 4, -640364487);
        l = F(l, u, a, c, n[t + 12], 11, -421815835);
        c = F(c, l, u, a, n[t + 15], 16, 530742520);
        a = F(a, c, l, u, n[t + 2], 23, -995338651);
        u = E(u, a, c, l, n[t + 0], 6, -198630844);
        l = E(l, u, a, c, n[t + 7], 10, 1126891415);
        c = E(c, l, u, a, n[t + 14], 15, -1416354905);
        a = E(a, c, l, u, n[t + 5], 21, -57434055);
        u = E(u, a, c, l, n[t + 12], 6, 1700485571);
        l = E(l, u, a, c, n[t + 3], 10, -1894986606);
        c = E(c, l, u, a, n[t + 10], 15, -1051523);
        a = E(a, c, l, u, n[t + 1], 21, -2054922799);
        u = E(u, a, c, l, n[t + 8], 6, 1873313359);
        l = E(l, u, a, c, n[t + 15], 10, -30611744);
        c = E(c, l, u, a, n[t + 6], 15, -1560198380);
        a = E(a, c, l, u, n[t + 13], 21, 1309151649);
        u = E(u, a, c, l, n[t + 4], 6, -145523070);
        l = E(l, u, a, c, n[t + 11], 10, -1120210379);
        c = E(c, l, u, a, n[t + 2], 15, 718787259);
        a = E(a, c, l, u, n[t + 9], 21, -343485551);
        u = r(u, o);
        a = r(a, f);
        c = r(c, i);
        l = r(l, h)
    }
    return Array(u, a, c, l)
}

function A(n, e, t, f, i, h) {
    return r(o(r(r(e, n), r(f, h)), i), t)
}

function s(n, e, t, r, o, f, i) {
    return A(e & t | ~e & r, n, e, o, f, i)
}

function w(n, e, t, r, o, f, i) {
    return A(e & r | t & ~r, n, e, o, f, i)
}

function F(n, e, t, r, o, f, i) {
    return A(e ^ t ^ r, n, e, o, f, i)
}

function E(n, e, t, r, o, f, i) {
    return A(t ^ (e | ~r), n, e, o, f, i)
}
`

/*
商务数据中心-国内贸易-社会融资规模增量统计

http://data.mofcom.gov.cn/gnmy/shrzgm.shtml
https://www.chinabond.com.cn/Info/20007290 社会融资规模存量各构成指标的定义及统计方法
*/
func MacroChinaShrzgm() ([]*model.SocialFinancingFlow, error) {
	var m = []*model.SocialFinancingFlow{}
	url := "http://data.mofcom.gov.cn/datamofcom/front/gnmy/shrzgmQuery"
	_, err := Client.R().SetResult(&m).Post(url)
	if err != nil {
		return m, err
	}
	return m, nil
}

/*
东方财富-货币供应量 M0 M1 M2
M表示monetary
M0:被称为基础货币，它指的是在银行体系外流通的现金，也就是大家没有存在银行而是拿在自己手上的钱，是货币构成中流动性最强的部分
M1:是在M0的基础上，加上企业活期存款，代表了货币构成中流动性较强的那部分
M2:则是在M1的基础上，再加上企业定期存款、个人储蓄存款和其它存款，因为它包括了M1里的所有货币，所以范围比M1、M0都要大，代表了货币构成中流动性较弱的部分
M0:流通中的现金
M1:M0+企业活期存款+机关团体部队存款+农村存款+个人持有的信用卡类存款
M2:M1+城乡居民储蓄存款+企业存款中具有定期性质的存款+外币存款+信托类存款

http://data.eastmoney.com/cjsj/hbgyl.html
*/

func MacroChinaMoneySupply() ([]*model.MacroChinaMoneySupply, error) {
	result := []*model.MacroChinaMoneySupply{}
	b, err := GetEastmoney("REPORT_DATE,BASIC_CURRENCY,BASIC_CURRENCY_SAME,BASIC_CURRENCY_SEQUENTIAL,CURRENCY,CURRENCY_SAME,CURRENCY_SEQUENTIAL,FREE_CASH,FREE_CASH_SAME,FREE_CASH_SEQUENTIAL", "2000", "RPT_ECONOMY_CURRENCY_SUPPLY")
	if err != nil {
		return result, err
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	return result, err
}

/*
新浪财经-中国宏观经济数据-社会消费品零售总额
http://data.eastmoney.com/cjsj/xfp.html
*/
func MacroChinaConsumerGoodsRetail() ([]*model.MacroChinaConsumerGoodsRetail, error) {
	result := []*model.MacroChinaConsumerGoodsRetail{}
	b, err := GetEastmoney("REPORT_DATE,RETAIL_TOTAL,RETAIL_TOTAL_SAME,RETAIL_TOTAL_SEQUENTIAL,RETAIL_TOTAL_ACCUMULATE,RETAIL_ACCUMULATE_SAME",
		"1000", "RPT_ECONOMY_TOTAL_RETAIL")
	if err != nil {
		return result, err
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	return result, nil
}

// 金十数据中心-中国 GDP 年率报告, 数据区间从 20110120-至今
// https://datacenter.jin10.com/reportType/dc_chinese_gdp_yoy
func MacroChinaGdpYearlyJin10() (data map[string]string, err error) {
	data, err = getJin10Yearly("https://cdn.jin10.com/dc/reports/dc_chinese_gdp_yoy_all.js",
		"中国GDP年率报告", "57")
	if err != nil {
		return data, err
	}
	return data, nil
}

// 金十数据中心-经济指标-中国-国民经济运行状况-物价水平-中国CPI月率报告
// https://datacenter.jin10.com/reportType/dc_chinese_cpi_yoy
func MacroChinaCpiYearlyJin10() (data map[string]string, err error) {
	data, err = getJin10Yearly("https://cdn.jin10.com/dc/reports/dc_chinese_cpi_yoy_all.js",
		"中国CPI年率报告", "56")
	if err != nil {
		return data, err
	}
	return data, nil
}

// 金十数据中心-经济指标-中国-国民经济运行状况-物价水平-中国PPI年率报告
// 中国年度 PPI 数据, 数据区间从 19950801-至今
// https://datacenter.jin10.com/reportType/dc_chinese_ppi_yoy
func MacroChinaPpiYearlyJin10() (data map[string]string, err error) {
	data, err = getJin10Yearly("https://cdn.jin10.com/dc/reports/dc_chinese_ppi_yoy_all.js",
		"中国PPI年率报告", "60")
	if err != nil {
		return data, err
	}
	return data, nil
}

// 金十数据中心-经济指标-中国-产业指标-官方制造业PMI
// https://zhuanlan.zhihu.com/p/100723005 解读
// https://datacenter.jin10.com/reportType/dc_chinese_manufacturing_pmi
// https://cdn.jin10.com/dc/reports/dc_chinese_manufacturing_pmi_all.js?v=1578817858
func MacroChinaPmiYearlyJin10() (data map[string]string, err error) {
	data, err = getJin10Yearly("https://cdn.jin10.com/dc/reports/dc_chinese_manufacturing_pmi_all.js",
		"中国官方制造业PMI报告", "65")
	if err != nil {
		return data, err
	}
	return data, nil
}

func getJin10Yearly(args ...string) (map[string]string, error) {
	url := args[0]
	result := make(map[string]string)
	t := time.Now().Unix()
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"v": strconv.FormatInt(t, 10),
			"_": strconv.FormatInt(t+90, 10),
		}).
		Get(url)
	if err != nil {
		return result, err
	}
	b := resp.Body()
	by := b[bytes.IndexAny(b, "{"):]
	value := gjson.GetBytes(by, "list").Array()
	for _, v := range value {
		rdate := v.Get("date")
		datas := v.Get(fmt.Sprintf("datas.%s", args[1]))
		restime, _ := time.ParseInLocation("20060102", rdate.String(), time.Local)
		result[restime.Format("2006-01-02")] = datas.Array()[1].String()
	}
	/*************************************************************/
	resp, err = Client.R().
		SetQueryParams(map[string]string{
			"max_date": "",
			"category": "ec",
			"attr_id":  args[2],
			"_":        strconv.FormatInt(t, 10),
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
		return result, err
	}
	b = resp.Body()
	value = gjson.GetBytes(b, "data.values").Array()
	for _, v := range value {
		vl := v.Array()
		result[vl[0].String()] = vl[1].String()
	}
	return result, nil
}

// 乐咕乐股-全A 股市盈率
// https://www.legulegu.com/stockdata/market_pe
// :param symbol: choice of {"sh", "sz", "cy", "zx", "000016.XSHG" ...}
func StockAPe() (map[string]interface{}, error) {
	result := make(map[string]interface{})
	vm := otto.New()
	vm.Run(SCRIPT)
	t := time.Now().Format("2006-01-02")

	token, err := vm.Call("hex", nil, t)
	if err != nil {
		panic(err)
	}
	_, err = Client.R().
		SetResult(&result).
		SetQueryParams(map[string]string{
			"token": strings.ToLower(token.String()),
		}).
		Get("https://www.legulegu.com/stockdata/market_pe/getmarket_pe")
	if err != nil {
		return result, err
	}
	return result, nil
}

// 沪深300 市盈率
// https://legulegu.com/stockdata/hs300-ttm-lyr
func SH300PE() ([]*model.SH300PE, error) {
	result := []*model.SH300PE{}
	vm := otto.New()
	vm.Run(SCRIPT)
	t := time.Now().Format("2006-01-02")

	token, err := vm.Call("hex", nil, t)
	if err != nil {
		panic(err)
	}
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"indexCode": "000300.SH",
			"token":     strings.ToLower(token.String()),
		}).
		Get("https://legulegu.com/api/stockdata/index-basic-pe")
	if err != nil {
		return result, err
	}
	b := resp.Body()
	value := gjson.GetBytes(b, "data")
	json.Unmarshal([]byte(value.String()), &result)
	for _, v := range result {
		t := time.UnixMilli(int64(v.Date))
		ts := t.Format("2006-01-02 15:04:05")
		v.Time = ts
	}
	return result, nil
}

// 东方财富网-数据中心-经济数据-中美国债收益率
// http://data.eastmoney.com/cjsj/zmgzsyl.html
// [{2022-10-27 00:00:00 0 2.4636 2.6953 3.1142 0 0 0 0 0 0},...]
func BondZhUsRate() ([]*model.BondZhUsRate, error) {
	result := []*model.BondZhUsRate{}
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"type":    "RPTA_WEB_TREASURYYIELD",
			"sty":     "ALL",
			"st":      "SOLAR_DATE",
			"sr":      "-1",
			"token":   "894050c76af8597a853f5b408b759f5d",
			"p":       "1",
			"ps":      "500",
			"pageNo":  "1",
			"pageNum": "1",
			"_":       "1615791534490",
		}).
		Get("http://datacenter.eastmoney.com/api/data/get")
	if err != nil {
		return result, err
	}
	b := resp.Body()
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	total_page := gjson.GetBytes(b, "result.pages").Int()
	for i := 2; i < int(total_page); i++ {
		r := []*model.BondZhUsRate{}
		resp, err := Client.R().
			SetQueryParams(map[string]string{
				"type":    "RPTA_WEB_TREASURYYIELD",
				"sty":     "ALL",
				"st":      "SOLAR_DATE",
				"sr":      "-1",
				"token":   "894050c76af8597a853f5b408b759f5d",
				"p":       strconv.Itoa(i),
				"ps":      "500",
				"pageNo":  strconv.Itoa(i),
				"pageNum": strconv.Itoa(i),
				"_":       "1615791534490",
			}).
			Get("http://datacenter.eastmoney.com/api/data/get")
		if err != nil {
			return result, err
		}
		b := resp.Body()
		v := gjson.GetBytes(b, "result.data")
		json.Unmarshal([]byte(v.String()), &r)
		result = append(result, r...)
	}
	return result, nil
}

// 搜狐证券
// https://q.stock.sohu.com/cn/zs.shtml
// symbol:zs_000300
func StockSohuCom(symbol string) error {

	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"code":                "zs_000300",
			"start":               "20221030",
			"end":                 "20221101",
			"stat":                "1",
			"order":               "D",
			"period":              "d",
			"callback":            "historySearchHandler",
			"rt":                  "jsonp",
			"r":                   "0.9563785074743043",
			"0.44466337614464213": "",
		}).
		Get("https://q.stock.sohu.com/hisHq")
	if err != nil {
		return err
	}
	byBody := resp.Body()
	bydata := byBody[bytes.IndexAny(byBody, "[")+1 : bytes.LastIndexAny(byBody, "]")]
	fmt.Println(string(bydata))
	return nil
}

func ChinaGDP() ([]*model.ChinaGDP, error) {
	result := []*model.ChinaGDP{}
	b, err := GetEastmoney("ALL", "1000", "RPT_ECONOMY_GDP")
	if err != nil {
		return result, err
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	return result, nil
}

func ChinaCPI() ([]*model.ChinaCPI, error) {
	result := []*model.ChinaCPI{}
	b, err := GetEastmoney("ALL", "1000", "RPT_ECONOMY_CPI")
	if err != nil {
		return result, err
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	return result, nil
}

func ChinaPPI() ([]*model.ChinaPPI, error) {
	result := []*model.ChinaPPI{}
	b, err := GetEastmoney("ALL", "1000", "RPT_ECONOMY_PPI")
	if err != nil {
		return result, err
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	return result, nil
}

func ChinaPMI() ([]*model.ChinaPMI, error) {
	result := []*model.ChinaPMI{}
	b, err := GetEastmoney("ALL", "1000", "RPT_ECONOMY_PMI")
	if err != nil {
		return result, err
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	return result, nil
}

// 东方财富网
func GetEastmoney(columns, pageSize, reportName string) ([]byte, error) {
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"columns":     columns,
			"pageNumber":  "1",
			"pageSize":    pageSize,
			"sortColumns": "REPORT_DATE",
			"sortTypes":   "-1",
			"source":      "WEB",
			"client":      "WEB",
			"reportName":  reportName,
			"p":           "1",
			"pageNo":      "1",
			"pageNum":     "1",
			"_":           "1660718498421",
		}).
		Get("https://datacenter-web.eastmoney.com/api/data/v1/get")
	if err != nil {
		return nil, err
	}
	b := resp.Body()
	return b, nil
}

// 财新PMI https://s.ccxe.com.cn/indices/pmi
func PmiCx() ([]*model.PmiCx, error) {
	result := []*model.PmiCx{}
	// list := []string{"com", "man", "ser"}
	resp, err := Client.R().SetQueryParams(map[string]string{"type": "com"}).
		Get("https://s.ccxe.com.cn/api/index/pro/cxIndexTrendInfo")
	if err != nil {
		return result, err
	}
	b := resp.Body()
	v := gjson.GetBytes(b, "data")
	manufacture := []model.CxPmi{}
	m := make(map[float64]*model.PmiCx)
	json.Unmarshal([]byte(v.String()), &manufacture)
	for _, v := range manufacture {
		changeRate, _ := decimal.NewFromFloat(v.ChangeRate).Round(2).Float64()
		m[v.Time] = &model.PmiCx{
			Date:                    time.UnixMilli(int64(v.Time)).Format("2006-01-02 15:04:05"),
			Time:                    v.Time,
			Manufacture:             0,
			ManufactureYearOverYear: 0,
			Service:                 0,
			ServiceYearOverYear:     0,
			Synthesis:               v.Data,
			SynthesisYearOverYear:   changeRate,
		}
	}
	resp, err = Client.R().SetQueryParams(map[string]string{"type": "man"}).
		Get("https://s.ccxe.com.cn/api/index/pro/cxIndexTrendInfo")
	if err != nil {
		return result, err
	}
	b = resp.Body()
	v = gjson.GetBytes(b, "data")
	json.Unmarshal([]byte(v.String()), &manufacture)
	for _, v := range manufacture {
		changeRate, _ := decimal.NewFromFloat(v.ChangeRate).Round(2).Float64()
		m[v.Time].Manufacture = v.Data
		m[v.Time].ManufactureYearOverYear = changeRate
	}

	resp, err = Client.R().SetQueryParams(map[string]string{"type": "ser"}).
		Get("https://s.ccxe.com.cn/api/index/pro/cxIndexTrendInfo")
	if err != nil {
		return result, err
	}
	b = resp.Body()
	v = gjson.GetBytes(b, "data")
	json.Unmarshal([]byte(v.String()), &manufacture)
	for _, v := range manufacture {
		changeRate, _ := decimal.NewFromFloat(v.ChangeRate).Round(2).Float64()
		m[v.Time].Service = v.Data
		m[v.Time].ServiceYearOverYear = changeRate
		result = append(result, m[v.Time])
	}
	// 排序 -- 升序 由小到大
	sort.SliceStable(result, func(i int, j int) bool {
		return result[i].Time > result[j].Time
	})
	return result, nil
}

// 工业生产增加值
func ValueAddedOfIndustrialProduction() ([]*model.ValueAddedOfIndustrialProduction, error) {
	result := []*model.ValueAddedOfIndustrialProduction{}
	b, err := GetEastmoney("ALL", "1000", "RPT_ECONOMY_INDUS_GROW")
	if err != nil {
		return result, err
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	return result, nil
}

func GetSocialElectricityConsumption(num, from string) ([]byte, error) {
	t := time.Now().UnixMilli()
	st := strconv.FormatInt(t, 10)
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"cate":      "industry",
			"event":     "6",
			"from":      from,
			"num":       num,
			"condition": "",
			"_":         st,
		}).
		Get("https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1671102964958/MacPage_Service.get_pagedata")
	if err != nil {
		return resp.Body(), err
	}
	byBody := resp.Body()
	return byBody, nil
}

// 全社会用电
func SocialElectricityConsumption() ([]*model.SocialElectricityConsumption, error) {
	listVal := []*model.SocialElectricityConsumption{}
	byBody, err := GetSocialElectricityConsumption("1", "0")
	if err != nil {
		return listVal, err
	}
	b := byBody[bytes.LastIndexAny(byBody, "c")+7 : bytes.LastIndexAny(byBody, "da")-5]
	aByteToInt, err := strconv.Atoi(string(b))
	if err != nil {
		return listVal, err
	}
	// 获取数据总数量
	for i := 0; i <= aByteToInt/31; i++ {
		form := fmt.Sprintf("%d", i*31)
		byBody, err = GetSocialElectricityConsumption("31", form)
		if err != nil {
			return listVal, err
		}
		b = byBody[bytes.LastIndexAny(byBody, "data:")+1 : bytes.LastIndexAny(byBody, "}")]
		strList := make([][]string, 0)
		err = json.Unmarshal([]byte(b), &strList)
		if err != nil {
			return listVal, nil
		}
		for _, v := range strList {
			m := model.SocialElectricityConsumption{}
			ivalue := reflect.ValueOf(&m).Elem()
			for i, v := range v {
				if i == 0 {
					elem := ivalue.Field(0)
					var ti time.Time
					if len(v) == 7 {
						ti, err = time.Parse("2006.01", v)
						if err != nil {
							log.Println(err)
						}
					} else {
						ti, err = time.Parse("2006.1", v)
						if err != nil {
							log.Println(err)
						}
					}
					elem.SetString(ti.Format(utils.DATE_FORMAT))
					continue
				}
				elem := ivalue.Field(i)
				elem.SetString(strings.TrimSpace(v))
			}
			listVal = append(listVal, &m)
		}
	}

	return listVal, nil
}

func GetPassengerAndFreightTraffic(num, from string) ([]byte, error) {
	t := time.Now().UnixMilli()
	st := strconv.FormatInt(t, 10)
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"cate":      "industry",
			"event":     "10",
			"from":      from,
			"num":       num,
			"condition": "",
			"_":         st,
		}).
		Get("https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1671270562484/MacPage_Service.get_pagedata")
	if err != nil {
		return resp.Body(), err
	}
	byBody := resp.Body()
	return byBody, nil
}

// 全社会客货运输量
func PassengerAndFreightTraffic() ([]*model.PassengerAndFreightTraffic, error) {
	listVal := []*model.PassengerAndFreightTraffic{}
	// 获取数据总数量
	for i := 0; i <= 1; i++ {
		form := fmt.Sprintf("%d", i*1000)
		byBody, err := GetPassengerAndFreightTraffic("1000", form)
		if err != nil {
			return listVal, err
		}
		// fmt.Println("isGBK:", utils.IsGBK(byBody))
		utf8Data, _ := simplifiedchinese.GBK.NewDecoder().Bytes(byBody)
		strBody := string(utf8Data)
		b := strBody[strings.LastIndex(strBody, "'非累计':")+12 : strings.LastIndex(strBody, ",'累计")]
		strList := make([][]string, 0)
		err = json.Unmarshal([]byte(b), &strList)
		if err != nil {
			return listVal, nil
		}
		for _, v := range strList {
			m := model.PassengerAndFreightTraffic{}
			ivalue := reflect.ValueOf(&m).Elem()
			for i, v := range v {
				elem := ivalue.Field(i)
				if i == 0 {
					// elem := ivalue.Field(0)
					var ti time.Time
					if len(v) == 7 {
						ti, err = time.Parse("2006.01", v)
						if err != nil {
							log.Println(err)
						}
					} else {
						ti, err = time.Parse("2006.1", v)
						if err != nil {
							log.Println(err)
						}
					}
					elem.SetString(ti.Format(utils.DATE_FORMAT))
					continue
				} else if i == 1 {
					// elem := ivalue.Field(i)
					switch v {
					case "国际航线":
						elem.SetInt(int64(model.CategoryOfTraffic_InternationalService))
						continue
					case "港澳地区航线":
						elem.SetInt(int64(model.CategoryOfTraffic_HongKongMacaoRegionalRoute))
						continue
					case "国内航线":
						elem.SetInt(int64(model.CategoryOfTraffic_DomesticAirline))
						continue
					case "民航":
						elem.SetInt(int64(model.CategoryOfTraffic_CivilAviation))
						continue
					case "水运":
						elem.SetInt(int64(model.CategoryOfTraffic_WaterTransport))
						continue
					case "公路":
						elem.SetInt(int64(model.CategoryOfTraffic_Highway))
						continue
					case "铁路":
						elem.SetInt(int64(model.CategoryOfTraffic_Railway))
						continue
					case "合计":
						elem.SetInt(int64(model.CategoryOfTraffic_Total))
						continue
					default:
						continue
					}
				}
				elem.SetString(strings.TrimSpace(v))
			}
			listVal = append(listVal, &m)
		}
	}
	return listVal, nil
}

// 中国新增信贷数据
func ChinaNewFinancialCredit() ([]*model.NewFinancialCredit, error) {
	result := []*model.NewFinancialCredit{}
	b, err := GetEastmoney("ALL", "1000", "RPT_ECONOMY_RMB_LOAN")
	if err != nil {
		return result, err
	}
	v := gjson.GetBytes(b, "result.data")
	json.Unmarshal([]byte(v.String()), &result)
	return result, nil
}

// 外汇储备与黄金
func ForeignReserveAndGold(num string) ([]*model.ForeignReserveAndGold, error) {
	result := []*model.ForeignReserveAndGold{}
	t := time.Now().UnixMilli()
	st := strconv.FormatInt(t, 10)
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"cate":      "fininfo",
			"event":     "5",
			"from":      "0",
			"num":       num,
			"condition": "",
			"_":         st,
		}).
		Get("https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1671456338405/MacPage_Service.get_pagedata")
	if err != nil {
		return result, err
	}
	byBody := resp.Body()
	b := byBody[bytes.LastIndexAny(byBody, "data:")+1 : bytes.LastIndexAny(byBody, "}")]
	strList := make([][]string, 0)
	err = json.Unmarshal([]byte(b), &strList)
	if err != nil {
		return result, nil
	}
	for _, v := range strList {
		m := model.ForeignReserveAndGold{}
		ivalue := reflect.ValueOf(&m).Elem()
		for i, v := range v {
			if i == 0 {
				elem := ivalue.Field(0)
				var ti time.Time
				if len(v) == 7 {
					ti, err = time.Parse("2006.01", v)
					if err != nil {
						log.Println(err)
					}
				} else {
					ti, err = time.Parse("2006.1", v)
					if err != nil {
						log.Println(err)
					}
				}
				elem.SetString(ti.Format(utils.DATE_FORMAT))
				continue
			}
			elem := ivalue.Field(i)
			elem.SetString(strings.TrimSpace(v))
		}
		result = append(result, &m)
	}
	return result, nil
}

// 固定资产投资
func InvestmentInFixedAssets(num string) ([]*model.InvestmentInFixedAssets, error) {
	result := []*model.InvestmentInFixedAssets{}
	t := time.Now().UnixMilli()
	st := strconv.FormatInt(t, 10)
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"cate":      "fixed",
			"event":     "1",
			"from":      "0",
			"num":       num,
			"condition": "",
			"_":         st,
		}).
		Get("https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1671538840147/MacPage_Service.get_pagedata")
	if err != nil {
		return result, err
	}
	byBody := resp.Body()
	b := byBody[bytes.LastIndexAny(byBody, "data:")+1 : bytes.LastIndexAny(byBody, "}")]
	strList := make([][]string, 0)
	err = json.Unmarshal([]byte(b), &strList)
	if err != nil {
		return result, nil
	}
	for _, v := range strList {
		var ti time.Time
		if len(v[0]) == 7 {
			ti, err = time.Parse("2006.01", v[0])
			if err != nil {
				log.Println(err)
			}
		} else {
			ti, err = time.Parse("2006.1", v[0])
			if err != nil {
				log.Println(err)
			}
		}
		date := ti.Format(utils.DATE_FORMAT)
		m := &model.InvestmentInFixedAssets{
			Date:                                date,
			InvestmentCompletedAmount:           v[1],  //投资完成额(亿元)
			InvestmentCompletedAmountYearOnYear: v[2],  //投资完成额同比增长(%)
			PrimaryIndustry:                     v[7],  //第一产业(亿元)
			PrimaryIndustryYearOnYear:           v[8],  //第一产业用同比%
			SecondaryIndustry:                   v[9],  //第二产业
			SecondaryIndustryYearOnYear:         v[10], //第二产业同比%
			TertiaryIndustry:                    v[11], //第三产业
			TertiaryIndustryYearOnYear:          v[12], //第三产业同比%
			RealEstateDevelopment:               v[45], //房地产开发(亿元)
			RealEstateDevelopmentYearOnYear:     v[46], //房地产开发同比增长(%)
		}

		result = append(result, m)
	}
	return result, nil
}

// 央行货币当局资产负债
func CentralBankMonetaryAuthorityAssetsAndLiabilities(num string) ([]*model.CentralBankMonetaryAuthorityAssetsAndLiabilities, error) {
	result := []*model.CentralBankMonetaryAuthorityAssetsAndLiabilities{}
	t := time.Now().UnixMilli()
	st := strconv.FormatInt(t, 10)
	resp, err := Client.R().
		SetQueryParams(map[string]string{
			"cate":      "fininfo",
			"event":     "8",
			"from":      "0",
			"num":       num,
			"condition": "",
			"_":         st,
		}).
		Get("https://quotes.sina.cn/mac/api/jsonp_v3.php/SINAREMOTECALLCALLBACK1671456338405/MacPage_Service.get_pagedata")
	if err != nil {
		return result, err
	}
	byBody := resp.Body()
	b := byBody[bytes.LastIndexAny(byBody, "data:")+1 : bytes.LastIndexAny(byBody, "}")]
	strList := make([][]string, 0)
	err = json.Unmarshal([]byte(b), &strList)
	if err != nil {
		return result, nil
	}
	for _, v := range strList {
		m := model.CentralBankMonetaryAuthorityAssetsAndLiabilities{}
		ivalue := reflect.ValueOf(&m).Elem()
		for i, v := range v {
			if i == 0 {
				elem := ivalue.Field(0)
				var ti time.Time
				if len(v) == 7 {
					ti, err = time.Parse("2006.01", v)
					if err != nil {
						log.Println(err)
					}
				} else {
					ti, err = time.Parse("2006.1", v)
					if err != nil {
						log.Println(err)
					}
				}
				elem.SetString(ti.Format(utils.DATE_FORMAT))
				continue
			}
			elem := ivalue.Field(i)
			elem.SetString(strings.TrimSpace(v))
		}
		result = append(result, &m)
	}
	return result, nil
}
