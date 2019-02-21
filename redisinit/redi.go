package redisinit

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"

	"log"
	"time"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

// X 全局DB
var Re *redis.Pool

func init() {
	var err error
	cfg, err := ini.Load("/home/lzw/DarkShell/src/gitlab.com/z547743799/iriscontent/config/redis.ini")
	if err != nil {
		log.Fatal(err)
	}
	url := cfg.Section("redis").Key("url").Value()
	redisMaxIdle,_:=strconv.Atoi(cfg.Section("redis").Key("redisMaxIdle").Value())

	   Re=&redis.Pool{
		MaxIdle:     redisMaxIdle,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
		c, err := redis.DialURL(url)
		if err != nil {
		return nil, fmt.Errorf("redis connection error: %s", err)
	}
		//验证redis密码
		//	if _, authErr := c.Do("AUTH", RedisPassword); authErr != nil {
		//		return nil, fmt.Errorf("redis auth password error: %s", authErr)
		//	}
		return c, err
	},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
		_, err := c.Do("PING")
		if err != nil {
		return fmt.Errorf("ping redis error: %s", err)
	}
		return nil
	},
	}

}
