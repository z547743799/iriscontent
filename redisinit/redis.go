package redisinit

import (
	"github.com/kataras/iris/sessions/sessiondb/redis"
	"github.com/kataras/iris/sessions/sessiondb/redis/service"
	"log"
	"time"

	"github.com/go-ini/ini"
	_ "github.com/go-sql-driver/mysql"
)

// X 全局DB
var IRRe *redis.Database

func init() {
	var err error
	cfg, err := ini.Load("/home/lzw/DarkShell/src/gitlab.com/z547743799/iriscontent/config/redis.ini")
	if err != nil {
		log.Fatal(err)
	}
	Url := cfg.Section("redis").Key("url").Value()

	IRRe=redis.New(service.Config{
		Network:     "tcp",
		Addr:        Url,
		Password:    "",
		Database:    "",
		MaxIdle:     0,
		MaxActive:   0,
		IdleTimeout: time.Duration(5) * time.Minute,
		Prefix:      ""}) // optionally configure the bridge between your redis server
}