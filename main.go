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

	num := conf.Go_no
	runtime.GOMAXPROCS(conf.Workprocess)

	c := make(chan models.Mqbody, num) //mq chan
	mqidc := make(chan string, 50)     //mqid chan
	for i := 0; i < num; i++ {
		go service.GoProcess(c, conf.Mysql_url, mqidc, i)
	}

	//redis回调gorountiner

	go service.FinishJob(conf, mqidc)
	fmt.Println("start ", num, " jobs")
	for i := 0; i < 3; i++ {
		go getjob(conf, c)
	}
	select {}
}

func getjob(conf models.MqConfig, c chan models.Mqbody) {
	client, _ := service.OpenRedis(conf)

	for {
		value, e := service.Getjob(client)
		if e == nil {
			if len(value) == 2 {
				mq := models.Mqbody{}
				mq.Init(string(value[1]))
				//fmt.Println(mq)
				c <- mq
			}
		}
	}
}
