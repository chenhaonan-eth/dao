package main

// 生成DB代码 https://github.com/go-gorm/gen/tree/master/examples
import (
	"fmt"

	"github.com/chenhaonan-eth/dao/dal/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

var (
	genOutPath string = "../../dal/query"
	splDsn     string = "../../guide_sqlite.db"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: genOutPath,
		Mode:    gen.WithDefaultQuery,
	})
	// Initialize a *gorm.DB instance
	db, err := gorm.Open(sqlite.Open(splDsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	g.UseDB(db)

	for _, tabledb := range model.OpArgsMap {
		// init db table
		// db.AutoMigrate(tabledb)

		// generate from struct in project
		g.ApplyBasic(tabledb)
	}

	g.Execute()
}
