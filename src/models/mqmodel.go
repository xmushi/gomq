package models

import (
	// "fmt"
	// "github.com/go-sql-driver/mysql"
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

func process() {

}
