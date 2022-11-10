// 业务层 只处理业务数据
package biz

import (
	"context"
	"fmt"

	"github.com/chenhaonan-eth/dao/dal/query"
)

var q = query.Q

func Query(ctx context.Context) {
	t := q.BondZhUsRateModel
	do := t.WithContext(context.Background())

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
