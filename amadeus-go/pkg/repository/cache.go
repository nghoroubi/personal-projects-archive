package repository

import (
    "fmt"
    "github.com/go-redis/redis"
)

var goRedis *redis.Client

// InitRedis , get the redis address and
// initializes the client for using in project
func InitRedis(addr string) {
    goRedis = redis.NewClient(&redis.Options{
        Network: "tcp",
        Addr:    addr,
        OnConnect: func(conn *redis.Conn) error {
            fmt.Println("Connected successfully to redis server")
            return nil
        },
        Password: "",
    })
    goRedis.FlushAll()
}

// Cache , returns the instance for use
func Cache() *redis.Client {
    return goRedis
}
