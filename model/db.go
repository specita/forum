package model

import (
	logrus "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func InitDB(dsn string, logModel bool) {
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("db connect failed" + err.Error())
	}

	DB.SingularTable(true)
	DB.LogMode(logModel)

	logrus.Info("db connect successed!")
}

func MigrateDB() {
	err := DB.AutoMigrate(&Article{},
		&ArticleImages{},
		&ArticleBlacklist{},
		&ArticleNotes{},
		&DeletedArticle{},
	).Error

	if err != nil {
		panic("migrate db failed" + err.Error())
	}
}
