package main

import (
	"flag"
	"math/rand"
	"time"

	"github.com/golang/glog"
	"wx.com/redis"
)

func main() {
	flag.Parse()
	r := &redisServer{}
	r.connect()
	str := "abcdefghijklmnopqrstuvwxyz0123456789"
	for {

		for f := 0; f < 200; f++ {
			key := ""
			for i := 0; i < 10; i++ {
				key += string(str[rand.Intn(36)])
			}
			go r.run(key)
		}
		time.Sleep(time.Second * 1)
	}
}

type redisServer struct {
	handle *redis.Ring
}

func (r *redisServer) connect() {
	r.handle = redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"s7001": "192.168.0.117:7001",
			"s7002": "192.168.0.117:7002",
			"s7003": "192.168.0.117:7003",
			"s7004": "192.168.0.117:7004",
		},
		PoolSize:    1000,
		PoolTimeout: 30,
	})
}

func (r *redisServer) run(key string) {
	err := r.handle.Set(key, time.Now().Unix(), time.Second*60).Err()
	if err != nil {
		glog.Error("redis set error : ", err, " key:", key)
	}
}
