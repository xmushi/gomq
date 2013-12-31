package service

import (
	"github.com/alphazero/Go-Redis"
	"log"
)

const (
	_KEY           = "mq_pop_key"
	_HOST          = "192.168.33.11"
	_DB_INDEX      = 2
	_REDIS_TIMEOUT = 1000
)

func Getjob() (value [][]byte, err error) {

	spec := redis.DefaultSpec().Db(_DB_INDEX).Host(_HOST)
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Println("failed to create the client", e)
		return nil, e
	}
	value, err = client.Brpop(_KEY, _REDIS_TIMEOUT)
	if e != nil {
		log.Println("error get key")
		return nil, e
	}
	return value, nil
}
