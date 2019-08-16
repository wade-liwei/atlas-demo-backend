package crontab

import (
	"fmt"
	"github.com/wade-liwei/atlas-demo-backend/dao"
)

func UpdateBtcBanner() {
	HttpGet("1234")

	if err := (dao.BtcBanner{}.Set()); err != nil {
		fmt.Println(err.Error())
	}
}
