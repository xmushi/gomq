package service

import (
	"github.com/xmushi/Go-Redis"
	"log"
	"models"
	"fmt"
)

const (
	_KEY = "tv"
)

var _REDIS_TIMEOUT = 0

func OpenRedis(conf models.MqConfig) (ret redis.Client, er error) {
	_REDIS_TIMEOUT = conf.Redis_timeout
	spec := redis.DefaultSpec().Db(conf.Redis_timeout).Host(conf.Redis_host).Port(conf.Redis_port).Db(0)
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Println("failed to create the client", e)
		return nil, e
	}
	return client, e
}

func Getjob(client redis.Client) (value [][]byte, err error) {
fmt.Println("hello")
	value, err = client.Brpop(_KEY,
		_REDIS_TIMEOUT)
	if err != nil {
		log.Println("error get key")
		return nil, err
	}
	return value, nil
}

func FinishJob(client redis.Client, mqid int64) {
	// client
}
