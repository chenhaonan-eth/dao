// 业务层 只处理业务数据
package biz

import (
	"context"
	"fmt"

	"github.com/chenhaonan-eth/guide-sisyphean/dal/model"
	"github.com/chenhaonan-eth/guide-sisyphean/dal/query"
)

var q = query.Q

func Query(ctx context.Context) {
	t := q.BondZhUsRateModel
	do := t.WithContext(context.Background())
	b := model.BondZhUsRateModel{Data: "2022-02-01", CN2Years: 213123, CN5Years: 142341, CN10Years: 15215, CN30Years: 325125, CN10_2Years: 21342, US2Years: 21355, US5Years: 66, US10Years: 345, US30Years: 34, US10_2Years: 324}

	err := do.Create(&b)
	catchError("Create", err)

	data, err := do.Take()
	catchError("Take", err)
	fmt.Printf("got %+v\n", data)

	dataArray, err := do.Find()
	catchError("Find", err)
	fmt.Printf("got %+v\n", dataArray)
}

func catchError(detail string, err error) {
	if err != nil {
		fmt.Printf("%s: %v\n", detail, err)
	}
}
