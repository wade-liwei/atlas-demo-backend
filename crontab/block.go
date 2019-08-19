package crontab

import (
	"encoding/json"
	"fmt"
	"github.com/wade-liwei/atlas-demo-backend/dao"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strings"
)

func UpdateBlockList() {
	fmt.Printf("run  UpdateBlockList\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api-r.bitcoinchain.com/v1/blocks/100", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request to server err: %v \n", err.Error())
		return
	}

	if strings.Compare(resp.Status, "200 OK") != 0 {
		fmt.Printf("response status:  %v \n", resp.Status)
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}

	res := dao.BlockList{}

	err = json.Unmarshal(respBody, &res)
	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}
	err = res.Set()

	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}
}
