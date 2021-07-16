package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

func InitREDIS(){
	r:= redis.NewClient(&redis.Options{
		Addr:     os.Getenv("ADDR"),
		Password: "", //
		DB:       0,  // use default DB
	})
	_, err := r.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	rdb =r
}
