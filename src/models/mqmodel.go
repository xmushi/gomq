package models

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Mqbody struct {
	Tv   string
	Mqid int64
	Msg  string
}

func (this *Mqbody) Init(str string) {
	i := 0
	old := 0
	// var e error
	for idx, v := range str {
		if v == ';' {
			if i == 0 {
				this.Tv = string(str[:idx])
				old = idx + 1
				i++
			} else if i == 1 {
				this.Mqid, _ = strconv.ParseInt(string(str[old:idx]), 10, 64)
				this.Msg = string(str[idx+1:])
			}
		}
	}
}

type MqConfig struct {
	Workprocess   int
	Redis_host    string
	Redis_port    int
	Redis_timeout int
	Mysql_url     string
}

func (this *MqConfig) LoadConfig() {
	f, err := os.OpenFile("conf.json", os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		panic("error open conf.json")
	}
	defer f.Close()
	buf := make([]byte, 1024)
	n, _ := f.Read(buf)
	err = json.Unmarshal(buf[:n], this)
	if err != nil {
		panic(err)
		fmt.Println(err)
	}
}
