package crontab

import (
	"github.com/robfig/cron"
	"io/ioutil"
	"net/http"
)

const (
	ApiKey = "b5cd3d03-b51f-434a-a12a-0bdc68383c5e"
)

// c := cron.New()
// c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
// c.AddFunc("@hourly",      func() { fmt.Println("Every hour") })
// c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
// c.Start()
// c.AddFunc("@daily", func() { fmt.Println("Every day") })

var Crontab *cron.Cron

func init() {
	Crontab = cron.New()
	Crontab.AddFunc("@every 5m", UpdateBtcBanner)

	Crontab.Start()
}

func HttpGet(urlStr string) ([]byte, error) {
	resp, err := http.Get(urlStr)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
