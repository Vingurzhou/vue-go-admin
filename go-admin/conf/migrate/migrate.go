package main

import (
	"fmt"
	"zwzProject/model/article"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:Z00a0319@tcp(139.196.137.117:4306)/zwzProject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4").AutoMigrate(&article.Article{})
	// for i := 1; i < 100; i++ {
	// 	article := article.Article{
	// 		Id:              i,
	// 		Timestamp:       588550467007,
	// 		Author:          "Patricia",
	// 		Reviewer:        "Cynthia",
	// 		Title:           "Bso Olwfid Ioq Njkleodh Bqvmabo",
	// 		Contentshort:    "mock data",
	// 		Content:         "<p>I am testing data, I am testing data.</p><p><img src=\"https://wpimg.wallstcn.com/4c69009c-0fd4-4153-b112-6cb53d1cf943\"></p>",
	// 		Forecast:        83.82,
	// 		Importance:      2,
	// 		Type:            "JP",
	// 		Status:          "published",
	// 		Displaytime:     "1983-06-29 20:12:03",
	// 		Commentdisabled: true,
	// 		Pageviews:       634,
	// 		Imageuri:        "https://wpimg.wallstcn.com/e4558086-631c-425c-9430-56ffb46e70b3",
	// 	}
	// 	db.Create(&article)
	// }
	fmt.Println("init data base ok")
}
