package service

import (
	"github.com/alphazero/Go-Redis"
	"log"
)

const (
	_KEY           = "tv"
	_HOST          = "192.168.33.11"
	_DB_INDEX      = 0
	_REDIS_TIMEOUT = 1000
	_REDIS_PORT    = 16379
)

func OpenRedis() (ret redis.Client, er error) {
	spec := redis.DefaultSpec().Db(_DB_INDEX).Host(_HOST).Port(_REDIS_PORT)
	client, e := redis.NewSynchClientWithSpec(spec)
	if e != nil {
		log.Println("failed to create the client", e)
		return nil, e
	}
	return client, e
}

func Getjob(client redis.Client) (value [][]byte, err error) {

	value, err = client.Brpop(_KEY, _REDIS_TIMEOUT)
	if err != nil {
		log.Println("error get key")
		return nil, err
	}
	return value, nil
}

// func FinishJob(client redis.Client, mqid int64) {
// 	client
// }
