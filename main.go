package main

import (
	"fmt"
	"service"
)

func main() {
	fmt.Println("hello")
	value, e := service.Getjob()
	if e == nil {
		if len(value) == 2 {
			fmt.Println(fmt.Sprintf("%s", value[1]))
		}
	}
}
