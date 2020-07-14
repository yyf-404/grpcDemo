package yy

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

//var ctx = context.Background()
//func RedisConnect(){
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "101.37.79.100:6380",
//		Password: "66538269y", // no password set
//		DB:       0,  // use default DB
//	})
//	pong,err :=rdb.Ping(ctx).Result()
//	fmt.Println(pong,err)
//}
var ctx = context.Background()

func ExampleNewClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.37.79.100:6380",
		Password: "66538269y", // no password set
		DB:       0,  // use default DB
	})
	pong, err := rdb.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
func ExampleClient() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "101.37.79.100:6380",
		Password: "66538269y", // no password set
		DB:       0,  // use default DB
	})
	err := rdb.Set( "key", "aaa", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get( "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get( "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}