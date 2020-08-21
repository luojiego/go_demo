package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
)

const (
	url = "https://e.dianping.com/flower/delivery/newOrder/search"
	agent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36"
	cookie = ``
)

type requestBody struct {
	Status string `json:"status"`
	SortItem string `json:"sortItem"`
	Source string `json:"source"`
	BeginTime string `json:"beginTime"`
	EndTime string `json:"endTime"`
	Page int `json:"page"`
	PageSize int `json:"pageSize"`
	NewOrderId string `json:"newOrderId"`
	TDealGroupID string `json:"mTDealGroupID"`
	Mobile string `json:"mobile"`
}

type orderStatistic struct {
	Date string `json:"date"`
	//DateDesc string `json:"dateDesc"`
	TotalOrderCount int `json:"totalOrderCount"`
	HandledOrderCount int `json:"handledOrderCount"`
	UnhandledOrderCount int `json:"unhandledOrderCount"`
}

type orderItem struct {
	//OrderId string `json:"newOrderId"`
	//OrderDate string `json:"orderDate"`
	//OrderTime string `json:"orderTime"`
	ArrivalTime string `json:"arrivalTime"`
	BuyerMobile string `json:"buyerMobile"`
	Address string `json:"address"`
	RecMobile string `json:"receiverMobile"`
	RecName string `json:"receiverName"`
	Card string `json:"cardMessage"`
	//DealGroupTitle string `json:"dealGroupTitle"`
	//Count int `json:"quantity"`
}

type responseMsg struct {
	Statistic []orderStatistic `json:"statisticDatas"`
	OrderResult []orderItem `json:"orderResults"`
	TotalPage int `json"totalPage"`
}

type responseData struct {
	Code int `json:"code"`
	Msg responseMsg `json:"msg"`
	Success bool `json:"success"`
}

/*
{"status":"0","sortItem":"1","source":"0",
"beginTime":"2019-04-24 00:00:00","endTime":"2019-05-01 23:59:59",
"page":1,"pageSize":10,"newOrderId":"","mTDealGroupID":"","mobile":""}

*/

func main() {
	page := 0
	if len(os.Args) != 1 {
		page, _ = strconv.Atoi(os.Args[1])
	}
	requestData := requestBody{
		Status:"0",
		SortItem:"1",
		Source:"0",
		BeginTime:"2020-05-03 00:00:00",
		EndTime:"2020-05-10 23:59:59",
		Page:page,
		PageSize:10,
		NewOrderId:"",
		TDealGroupID:"",
		Mobile:"",
	}

	result, err := json.Marshal(&requestData)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(result))
	reqNew := bytes.NewBuffer([]byte(result))
	client := &http.Client{}
	request, _ := http.NewRequest("POST", url, reqNew)
	request.Header.Set("Content-type", "application/json")
	request.Header.Set("User-Agent", agent)
	request.Header.Set("Cookie", cookie)
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	result, err = ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Println(string(result))

	res := responseData{}
	json.Unmarshal(result, &res)

	for _, v := range res.Msg.OrderResult {
		fmt.Printf("%s , %s , %s , %s\n", v.ArrivalTime, v.RecMobile, v.Address, v.BuyerMobile)
	}


	//result , _ = json.MarshalIndent(res, "", "  ")

	//fmt.Println(string(result))
}
