// 业务层 只处理业务数据
package biz

import (
	"context"
	"log"

	"github.com/chenhaonan-eth/dao/dal/query"
	"github.com/chenhaonan-eth/dao/pkg/utils"
)

var q = query.Q

func Query(ctx context.Context) {
	t := q.BondZhUsRate
	do := t.WithContext(context.Background())

	data, err := do.Take()
	utils.CatchError("Take", err)
	log.Printf("got %+v\n", data)

	dataArray, err := do.Find()
	utils.CatchError("Find", err)
	log.Printf("got %+v\n", dataArray)
}
