package article

type Article struct {
	Id              int     `gorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Timestamp       int64   `gorm:"int(10)" json:"timestamp"`
	Author          string  `gorm:"varchar(20)" json:"author"`
	Reviewer        string  `gorm:"varchar(20)" json:"reviewer"`
	Title           string  `gorm:"varchar(20)" json:"title"`
	Contentshort    string  `gorm:"varchar(20)" json:"content_short"`
	Content         string  `gorm:"varchar(200)" json:"content"`
	Forecast        float64 `gorm:"float(20)" json:"forecast"`
	Importance      int     `gorm:"int(10)" json:"importance"`
	Type            string  `gorm:"varchar(20)" json:"type"`
	Status          string  `gorm:"varchar(20)" json:"status"`
	Displaytime     string  `gorm:"varchar(20)" json:"display_time"`
	Commentdisabled bool    `json:"comment_disabled"`
	Pageviews       int     `gorm:"int(10)" json:"pageviews"`
	Imageuri        string  `gorm:"varchar(20)" json:"image_uri"`
}
