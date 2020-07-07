package main

import (
	"context"
	"fmt"
	ap "google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
)

const configFile = "./账号的配置json.json"
const packageName = "包名"

func main() {
	ctx := context.Background()
	service, err := ap.NewService(ctx, option.WithCredentialsFile(configFile))
	if err != nil {
		panic(err)
	}

	fmt.Println(service.Reviews.List(packageName))
	do, err := service.Inappproducts.List(packageName).Do()
	if err != nil {
		panic(err)
	}

	for _, v := range do.Inappproduct {
		bytes, _ := v.MarshalJSON()
		fmt.Println(string(bytes))
	}

	// 退款的数据
	fmt.Println(service.Purchases.Voidedpurchases.List(packageName))
	// 验证 请正确填充 productId 和 token
	get := service.Purchases.Products.Get(packageName, "", "")
	purchase, err := get.Do()
	if err != nil {
		panic(err)
	}
	// acknowledge 订单数据 请正确填充 productId 和 token
	service.Purchases.Products.Acknowledge(packageName, "", "", &ap.ProductPurchasesAcknowledgeRequest{
		DeveloperPayload: purchase.DeveloperPayload,
	})
}
