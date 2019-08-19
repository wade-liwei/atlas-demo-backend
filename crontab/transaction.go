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

func UpdateTransactionList() {
	fmt.Printf("run  UpdateTransactionList\n")
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.smartbit.com.au/v1/blockchain/transactions", nil)
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

	res := dao.TransactionPage{}

	err = json.Unmarshal(respBody, &res)
	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}

	fmt.Printf("transaction: %v \n", res)
	err = res.Set()

	if err != nil {
		fmt.Printf("json unmarshal err:  %v \n", err.Error())
		return
	}
}
