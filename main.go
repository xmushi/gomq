package main

import (
	"fmt"
	"models"
	"runtime"
	"service"
)

func main() {
	numcpu := runtime.NumCPU() - 1
	runtime.GOMAXPROCS(numcpu)

	c := make(chan models.Mqbody)
	for i := 0; i < numcpu; i++ {
		go service.GoProcess(c, i)
	}

	fmt.Println("start ", numcpu, " jobs")

	var mq models.Mqbody
	client, _ := service.OpenRedis()
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
