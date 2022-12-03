package article

import (
	"net/http"

	"zwzProject/model/article"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func List(c *gin.Context) {
	// page := c.PostForm("page")
	// limit := c.PostForm("limit")

	dsn := "root:Z00a0319@tcp(139.196.137.117:4306)/zwzProject?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	articles := make([]article.Article, 0)
	db.Raw("select *from articles").Scan(&articles)
	data := struct {
		Total int               `json:"total"`
		Items []article.Article `json:"items"`
	}{Total: len(articles),
		Items: articles}

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 20000,
			"data": data,
		})
	}
}
