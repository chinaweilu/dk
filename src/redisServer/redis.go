package redisServer

import (
	"conf"
	"strconv"

	"gopkg.in/redis.v3"
)

type redisHandle struct {
	status bool
	handle *redis.Client
}

type redisConf struct {
	connect int
}

var redisMap map[string][]redisHandle

func NewServer() {
	redisConfig := &redisConf{}
	redisConfig.connect, _ = strconv.Atoi(conf.GetIni("redisConnect", "connect"))
	redisMap = make(map[string][]redisHandle)
	for _, address := range conf.GetRedisServer() {
		redisTemp := make([]redisHandle, redisConfig.connect)
		for i := 0; i < redisConfig.connect; i++ {
			client := redisHandle{status: false}
			client.handle = redis.NewClient(&redis.Options{
				Addr:     address,
				Password: "",
				DB:       0,
			})
			client.handle.Ping().Result()
			redisTemp[i] = client
		}
		redisMap[address] = redisTemp
	}
}
