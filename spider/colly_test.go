package spider

import (
	"fmt"
	"testing"

	"github.com/chenhaonan-eth/dao/dal"
	"github.com/chenhaonan-eth/dao/dal/query"
)

func TestCollylegulegu(t *testing.T) {
	dal.DB = dal.ConnectDB("guide_sqlite.db").Debug()
	query.SetDefault(dal.DB)
	Collylegulegu()
}

func TestProcess(t *testing.T) {
	r, err := Process("冠龙节能(301151)", "/s/301151")
	if err != nil {
		t.Error(err)
	}
	for _, v := range r {
		fmt.Printf("%+v", *v)
	}
}
