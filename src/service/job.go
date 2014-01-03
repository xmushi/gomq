package service

import (
	"fmt"
	"github.com/xmushi/Go-Redis"
	"log"
	"models"
)

const (
	_KEY         = "tv"
	_MQ_BODY_KEY = "mq"
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

	value, err = client.Brpop(_KEY,
		_REDIS_TIMEOUT)
	if err != nil {
		log.Println("error get key")
		return nil, err
	}
	return value, nil
}

/*
 *删除redis里面备份的mq字段
 */
func FinishJob(conf models.MqConfig, c chan string) {
	// client:=O
	client, _ := OpenRedis(conf)
	for {
		v, ok := <-c
		if ok {
			fmt.Println(v)
			client.Hdel(_MQ_BODY_KEY, v)
		}
	}
}
