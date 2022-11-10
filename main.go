package main

// import (
// 	_ "github.com/chenhaonan-eth/guide-sisyphean/core"
// 	"github.com/chenhaonan-eth/guide-sisyphean/router"
// )

// // go:generate go env -w GO111MODULE=on
// // go:generate go env -w GOPROXY=https://goproxy.cn,direct
// // go:generate go mod tidy
// // go:generate go mod download

// //TODO: 启动时候检查配置文件是否正确
// func main() {
// 	router.Run()
// }

// import (
// 	"fmt"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// type Product struct {
// 	gorm.Model
// 	Code  string
// 	Price uint
// }

// func main() {
// 	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// Migrate the schema
// 	db.AutoMigrate(&Product{})

// 	// Create
// 	// var r = []Product{{Code: "D42", Price: 100}, {Code: "D111", Price: 2323}}
// 	// result := db.Create(&r)
// 	// Read
// 	var product Product
// 	var products []Product
// 	// db.First(&product, 1) // find product with integer primary key
// 	db.First(&product, "code = ?", "D42") // find product with code D42

// 	// // Update - update product's price to 200
// 	// db.Model(&product).Update("Price", 200)
// 	// // Update - update multiple fields
// 	db.Model(&product).Updates(Product{Code: "D42", Price: 200}) // non-zero fields
// 	// db.Model(&product).Updates(map[string]interface{}{"price": 200, "code": "D42"})
// 	// Get all records
// 	db.Find(&products)
// 	fmt.Printf("%+v", products)
// 	// Delete - delete product
// 	// db.Delete(&product, 1)
// }

import (
	"context"
	"fmt"

	"github.com/chenhaonan-eth/guide-sisyphean/biz"
	"github.com/chenhaonan-eth/guide-sisyphean/dal"
	"github.com/chenhaonan-eth/guide-sisyphean/dal/query"
)

func init() {
	dal.DB = dal.ConnectDB("guide_sqlite.db").Debug()
}
func main() {
	// start your project here
	fmt.Println("hello world")
	defer fmt.Println("bye~")

	query.SetDefault(dal.DB)
	biz.Query(context.Background())
}
