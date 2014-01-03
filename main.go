package main

import (
	"fmt"
	"models"
	"runtime"
	"service"
)

func main() {
	conf := models.MqConfig{}
	conf.LoadConfig()

	num := conf.Workprocess
	runtime.GOMAXPROCS(conf.Workprocess)

	c := make(chan models.Mqbody, num)
	for i := 0; i < num; i++ {
		go service.GoProcess(c, conf.Mysql_url, i)
	}

	fmt.Println("start ", num, " jobs")

	var mq models.Mqbody
	client, _ := service.OpenRedis(conf)
	for {
		value, e := service.Getjob(client)
		if e == nil {
			if len(value) == 2 {
				mq = models.Mqbody{}
				mq.Init(string(value[1]))
				fmt.Println(mq)
				c <- mq
			}
		}
	}

}
