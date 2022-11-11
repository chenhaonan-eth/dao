package init

import (
	"github.com/chenhaonan-eth/dao/dal"
	"github.com/chenhaonan-eth/dao/dal/model"
)

func InitDB() {
	// 初始化数据库
	for _, tabledb := range model.OpArgsMap {
		// init db table 它会自动检测是否存在
		dal.DB.AutoMigrate(tabledb)
	}
	//TODO step1 检查数据库是否有数据

	//TODO step2 无数据，加载数据

	//TODO step3 有数据，检查是否完整

}
