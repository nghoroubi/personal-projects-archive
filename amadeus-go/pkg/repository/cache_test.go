package repository

import (
    "fmt"
    "testing"
)

func TestInitRedis(t *testing.T) {
    InitRedis("localhost:6379")
    if goRedis.Ping().Err() != nil {
        t.FailNow()
    }
    str:= goRedis.HMGet("Z", "adult_fare").Val()[0].(string)
    fmt.Println(str)
}

