package waiting

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func NewRedisCheck(waitRedisKey, redisAddr, redisPassword string, redisDb int) CheckTask {
	return func(timeout time.Duration) bool {
		key, value := splitKeyWithValue(waitRedisKey)
		client := redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       redisDb,
		})
		fmt.Printf("Checking Redis key on server %v (db: %v) %v = \"%v\": ", redisAddr, redisDb, key, value)
		result, err := client.Get(context.Background(), key).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Println("key not found")
			} else {
				fmt.Printf("error=\"%v\"\n", err)
			}
			return false
		}
		checkResult := result == value
		fmt.Printf("current value=\"%v\" (result: %v)\n", result, checkResult)
		return checkResult

	}
}
