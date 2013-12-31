package test

import (
	"models"
	"service"
	"testing"
)

func Test_testdb(t *testing.T) {
	mq := models.Mqbody{}
	mq.Init("tv1;123;aaaaa")

	db := service.Opendb()

	_, e := service.Process(&mq, db)

	if e != nil {
		t.Error(e)
	}

}
