package driver

import (
	"fmt"
	"os"
	"sync"

	. "tasks_list/config"
	"tasks_list/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	once sync.Once
	db   *gorm.DB
)

func InitGorm() *gorm.DB {

	var err error
	host := os.Getenv("DB_HOST")
	fmt.Println(host)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		GlobalConfig.MySQLConfig.Username,
		GlobalConfig.MySQLConfig.Password,
		GlobalConfig.MySQLConfig.DbHost,
		GlobalConfig.MySQLConfig.DbPORT,
		GlobalConfig.MySQLConfig.DbName,
	)
	once.Do(func() {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
		if err != nil {
			panic(err)
		}
		err = db.AutoMigrate(models.Tasks{})
		if err != nil {
			panic(err)
		}
	})
	return db

}
