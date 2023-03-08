package configuration

import (
	"github.com/go-redis/redis"
	"log"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       int
}

var RedisClient *redis.Client

func InitRedis() error {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     AllConfig.Redis.Host + ":" + AllConfig.Redis.Port,
		Password: AllConfig.Redis.Password, // no password set
		DB:       AllConfig.Redis.DB,       // use default DB
	})
	_, err := RedisClient.Ping().Result()
	if err != nil {
		log.Printf("Redis connect ping failed, err: %v", err)
		return err
	}
	log.Printf("Redis connect success")
	return nil
}
