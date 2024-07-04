package api

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

func LoadRedis() {

	// vars
	Address := os.Getenv("REDIS_ADDRESS")
	Pass := os.Getenv("REDIS_PASS")

	rdb := redis.NewClient(&redis.Options{
		Addr: Address,
		Password: Pass,
		DB: 0, // 0 = default db
	})

	err := rdb.Set(ctx, "key", "value", 0).Err()

	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

    val2, err := rdb.Get(ctx, "key2").Result()
    if err == redis.Nil {
        fmt.Println("key2 does not exist")
    } else if err != nil {
        panic(err)
    } else {
        fmt.Println("key2", val2)
    }
}


