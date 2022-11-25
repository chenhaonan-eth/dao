package main

// 生成DB代码 https://github.com/go-gorm/gen/tree/master/examples
import (
	"github.com/chenhaonan-eth/dao/dal/model"
	"gorm.io/gen"
)

var (
	genOutPath string = "./dal/query"
	// splDsn     string = "../../dao.db"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: genOutPath,
		Mode:    gen.WithDefaultQuery,
	})
	// Initialize a *gorm.DB instance
	// db, err := gorm.Open(sqlite.Open(splDsn), &gorm.Config{})
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// g.UseDB(db)

	for _, tabledb := range model.OpArgsMap {
		// init db table
		// db.AutoMigrate(tabledb)

		// generate from struct in project
		g.ApplyBasic(tabledb)
	}

	g.Execute()
}
