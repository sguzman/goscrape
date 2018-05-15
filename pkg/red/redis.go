package red

import (
    "github.com/go-redis/redis"
    "fmt"
)

type KeyVal struct {
    Key string
    Val []byte
}

var (
    Client  *redis.Client
    InRedis chan KeyVal
    Cache   map[string]string
)

func Set() {
    for {
        pair := <- InRedis
        out, err := Client.HSet("ebooks", pair.Key, pair.Val).Result()
        if !out || err != nil {
            fmt.Println(err)
        } else {
            fmt.Printf("%s -> %d\n", pair.Key, len(pair.Val))
        }
    }
}

func redisCache() map[string]string {
    cache, err := Client.HGetAll("ebooks").Result()
    {
        if err != nil {
            panic(err)
        }

        fmt.Printf("Retrieved %d items\n", len(cache))
    }

    return cache
}

func Init() {
    Client = redis.NewClient(&redis.Options{
        Addr: "localhost:6379",
    })

    InRedis = make(chan KeyVal, 1000)
    Cache = redisCache()
}