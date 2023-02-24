package configuration

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
)

var (
	DB *gorm.DB
)

type Configuration struct {
	Mysql MysqlConfiguration
}

type MysqlConfiguration struct {
	Username string
	Password string
	Database string
	Url      string
	Port     string
}

func InitMySQL() error {
	data, err := ioutil.ReadFile("./configuration/configuration.yaml")
	if err != nil {
		log.Printf("resolve yaml file failed, error: %v\n", err)
		return err
	}
	log.Printf("data: %v", string(data))
	//2解析文件
	var y Configuration
	err = yaml.Unmarshal(data, &y)
	if err != nil {
		log.Printf("yaml unmarshal failed, error: %v", err)
		return err
	}

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", y.Mysql.Username, y.Mysql.Password, y.Mysql.Url, y.Mysql.Port, y.Mysql.Database)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return nil
}
