package service

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"models"
)

const (
	_MYSQLURL   = "root:jiagoubu123456@tcp(192.168.33.11:3306)/shimingtest"
	_INSERT_SQL = "insert into message_info(mq_id,tvname,content,create_time) values(?, ?, ?, NOW())"
)

func Opendb() *sql.DB {
	db, err := sql.Open("mysql", _MYSQLURL)
	if err != nil {
		//panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		fmt.Println(err.Error())
	}
	return db
}

func Process(job *models.Mqbody, db *sql.DB) (ret bool, err error) {
	stmtIns, err := db.Prepare(_INSERT_SQL)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	defer stmtIns.Close()

	_, err = stmtIns.Exec(job.Mqid, job.Tv, job.Msg) // Insert tuples (i, i^2)
	if err != nil {
		return false, err
	}
	return true, nil
}

func GoProcess(c chan models.Mqbody, jobid int) {
	db := Opendb()
	fmt.Println("jobid:", jobid)
	for {
		select {
		case v, ok := <-c:
			if !ok {
				continue
			}
			_, e := Process(&v, db)
			if e == nil {
				fmt.Println(jobid, " process", v.Mqid)
			}
		}
	}
}
