package config

import (
	"fmt"
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

func (appConfig *AppConfig) SetupRedis() {
	pool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", appConfig.RedisConfig.Host+":"+appConfig.RedisConfig.Port)
			if err != nil {
				log.Fatalf("Unable to dial redis: %s", err.Error())
			}
			return c, err
		},
	}
	appConfig.RedisConfig.Pool = pool

	// get a connection from the pool (redis.Conn)
	conn := pool.Get()
	// use defer to close the connection when the function completes
	defer conn.Close()

	// call Redis PING command to test connectivity
	s, err := redis.String(conn.Do("PING"))
	if err != nil {
		log.Fatalf("Unable to ping redis: %s", err.Error())
		panic(err.Error())
	} else {
		fmt.Printf("PING Response = %s\n", s)
	}
}
