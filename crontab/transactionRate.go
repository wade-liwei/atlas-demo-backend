package crontab

//https://www.blockchain.com/en/charts/transactions-per-second?timespan=1year

import (
	//"encoding/json"
	"fmt"
	"github.com/wade-liwei/atlas-demo-backend/dao"
	"io/ioutil"
	"net/http"
	//"net/url"
	"strings"
)

func UpdateTransactionRate() {
	fmt.Printf("run  UpdateTransactionRate\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.blockchain.info/charts/transactions-per-second?timespan=1year&format=json", nil)
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

	//fmt.Printf("transactionRate: %s \n", respBody)

	res := dao.TransactionRate{}
	err = res.Set(respBody)

	if err != nil {
		fmt.Printf("insert into redis err:  %v \n", err.Error())
		return
	}
}
