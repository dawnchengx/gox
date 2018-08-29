package gox

import (
    "github.com/gomodule/redigo/redis"
)

var (
    redisType string = ""
    redisIpAndPort string = ""
    redisAuth = ""
)

func InitConn(rType string, ipAndPort string, auth string) {
    redisType = rType
    redisIpAndPort = ipAndPort
    redisAuth = auth
}

func RedisConn() (r redis.Conn, err error) {
    r, err = redis.Dial(redisType, redisIpAndPort)
    if err != nil {
        return
    }
    if "" != redisAuth {
        r.Do("auth", redisAuth)
    }
    return
}

