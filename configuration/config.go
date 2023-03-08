package configuration

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

var AllConfig Config = Config{}

type Config struct {
	Mysql MysqlConfig
	Redis RedisConfig
}

func init() {
	//1 读取yaml文件
	data, err := os.ReadFile("./configuration/configuration.yaml")
	if err != nil {
		log.Printf("resolve yaml file failed, error: %v\n", err)
	}
	//log.Printf("data: %v", string(data))
	//2 解析文件
	err = yaml.Unmarshal(data, &AllConfig)
	if err != nil {
		log.Printf("yaml unmarshal failed, error: %v", err)
	}
}
