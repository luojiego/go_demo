package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//从苹果服务器获取数据是否合法

const (
	receipt = `你要验证的凭证`
	url     = `https://buy.itunes.apple.com/verifyReceipt`
)

type ReceiptData struct {
	Receipt string `json:"receipt-data"`
}

type ReceiptResponseData struct {
	OriginalPurchaseDatePst string `json:"original_purchase_date_pst"` //购买时间、太平洋标准时间
	PurchaseDateMs          string `json:"purchase_date_ms"`           //购买时间毫秒
	UniqueIdentifier        string `json:"unique_identifier"`          //唯一标识符
	OriginalTransactionId   string `json:"original_transaction_id"`    //原始交易id
	Bvrs                    string `json:"bvrs"`                       //iphone程序的版本号
	TransactionId           string `json:"transaction_id"`             //交易的标识
	Quantity                string `json:"quantity"`                   //购买商品的数量
	UniqueVendorIdentifier  string `json:"unique_vendor_identifier"`   //开发商交易ID
	ItemId                  string `json:"item_id"`                    //App store用来标识程序的字符串
	OriginalPurchaseDate    string `json:"original_purchase_date"`     //原始购买时间
	IsInIntroOfferPeriod    string `json:"is_in_intro_offer_period"`   //
	ProductId               string `json:"product_id"`                 //商品的标识
	PurchaseDate            string `json:"purchase_date"`              //购买时间
	IsTrialPeriod           string `json:"is_trial_period"`            //
	PurchaseDatePst         string `json:"purchase_date_pst"`          //太平洋标准时间
	Bid                     string `json:"bid"`                        //iphone程序的bundle标识
	OriginalPurchaseDateMs  string `json:"original_purchase_date_ms"`  //毫秒
}

type ResponseData struct {
	Receipt ReceiptResponseData `json:"receipt"`
	Status  int                 `json:"status"`
}

func getResult(request []byte, url string) (receiptResponse ResponseData, err error) {
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(request))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return receiptResponse, err
	}
	defer resp.Body.Close()
	log.Printf("%+v", resp)

	//not equal status ok: 200
	if resp.StatusCode != http.StatusOK {
		return receiptResponse, fmt.Errorf(http.StatusText(resp.StatusCode))
	}

	log.Printf("\n resp[%v]", resp)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return receiptResponse, err
	}

	if err := json.Unmarshal(body, &receiptResponse); err != nil {
		return receiptResponse, err
	}

	log.Println(string(body))
	return receiptResponse, nil
}

func main() {
	req := ReceiptData{Receipt: receipt}
	data, _ := json.Marshal(req)

	getResult(data, url)
}
