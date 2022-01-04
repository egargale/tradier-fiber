package redis

import (
	"fmt"
	"log"
	"github.com/go-redis/redis"
	"tradier-fiber/internals/util"
)

func TestRedis() error {
	fmt.Println("Go Redis Tutorial")

	client := redis.NewClient(&redis.Options{
		Addr: util.MyConfig.RedisHost + ":" + util.MyConfig.RedisPort,
		Password: "",
		DB: 0,
	})

	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Redis is not ready. Error:", err)
	}
	return err

}