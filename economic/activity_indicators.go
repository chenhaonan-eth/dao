// 周期经济指标
package economic

import (
	"math"
	"strings"
	"time"
)

//股债利差来源自FED估值模型（美联储估值模型），也叫风险溢价模型
//股债利差=1/市盈率/十年期国债收益率
//股债利差就代表股市的“投资收益率”高于债市“无风险利率”的部分叫“风险溢价”，也即为风险付出的溢价
//利差越大，说明股市性价比越高，更具有投资价值，反之，则债市更值得配置。
func EquityBbondYieldSpreads() (map[string]float64, error) {
	// maplist := make(map[string][]float64)
	maplist := make(map[string]float64)
	band, err := BondZhUsRate()
	if err != nil {
		return maplist, err
	}
	pe, err := SH300PE()
	if err != nil {
		return maplist, err
	}
	for _, p := range pe {
		t := time.Unix(int64(p.Date/1000), 0)
		maplist[t.Format("2006-01-02")] = p.AddTtmPe / 100
	}
	// 对齐时间
	for _, b := range band {
		t := strings.Split(b.Date, " ")[0]
		_, ok := maplist[t]
		if ok {
			pe := maplist[t]
			spreads := 1 / pe / b.CN10Years
			//四舍五入
			n10 := math.Pow10(2)
			v := math.Trunc((spreads+0.5/n10)*n10) / n10
			maplist[t] = v
		} else {
			delete(maplist, t)
		}
	}
	return maplist, nil
}
