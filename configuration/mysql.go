package configuration

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

type MysqlConfig struct {
	Username string
	Password string
	Database string
	Url      string
	Port     string
}

func InitMySQL() error {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", AllConfig.Mysql.Username, AllConfig.Mysql.Password, AllConfig.Mysql.Url, AllConfig.Mysql.Port, AllConfig.Mysql.Database)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
