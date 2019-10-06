package modes

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sumeragibi_blog/log_init"
	"time"
)

var log = log_init.LogInit()

type User struct {
	gorm.Model
	UserName string `grom:"size:255"`
	Email    string `gorm:"type:varchar(100);unique_index"`
	PassWord string `gorm:"size:255"`
}

type Article struct {
	gorm.Model
	Title     string `grom:"type:varchar(64);index"`
	Content   string `grom:"type:text;index"`
	CreatTime *time.Time
	Labels    [] *Label `gorm:"many2many:article_labels;"`
	Reply     [] *Reply `gorm:"many2many:article_reply;"`
}

type Label struct {
	gorm.Model
	LabelName string      `grom:"type:varchar(100);index"`
	Articles  [] *Article `gorm:"many2many:article_labels;"`
}

type Reply struct {
	gorm.Model
	Email     string      `gorm:"type:varchar(100);"`
	UserName  string      `gorm:"size:255"`
	Text      string      `gorm:"type:text;"`
	Articles  [] *Article `gorm:"many2many:article_reply;"`
	CreatTime *time.Time
}

func CloseDataBase(db *gorm.DB) {
	err := db.Close()
	if err != nil {
		log.Error(err.Error())
		panic(fmt.Errorf("Close database error: %s \n", err))
	}
}

func DataBaseCline(host string, port string, user string, dbname string, password string) *gorm.DB {
	args := "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"
	args = fmt.Sprintf(args, host, port, user, dbname, password)
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Error(err.Error())
		panic(fmt.Errorf("Cline Database Error: %s \n", err))
	}
	db.AutoMigrate(&User{}, &Article{}, &Label{}, &Reply{})
	return db
}
