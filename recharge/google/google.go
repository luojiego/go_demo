package main

import (
	"context"
	"fmt"
	ap "google.golang.org/api/androidpublisher/v3"
	"google.golang.org/api/option"
)

const configFile = "./gam-finance-f4b7c4fbf9bc.json"
const packageName = "com.lt.guard_and_merge"
const productId = "com.gam.big_package_1"
const token = `ddcnlfjeeoilmkdhghelofnj.AO-J1OzHOEAXFbp85m3xce9urh9zqSonnan_MQyuUnLxlgvRVbNkU59vk1aGVN23F9k8I_1Gts9Mj1icilqXA0ZzlSsip_MYksLVbIUUHwfNntKUWG8wd67JlUZfyoKmKIA93elceeXO`
const token2 = `mfkkdeehhlmabcahipfffgke.AO-J1Oyr8p_lMEaolUiuj2nWwkKyMNt0oRJsLGe-wIrO1YpAHwIgC96JEy6sgyompDFO6p5Bx2m1yeQy8x7wqPr7Op_DzPFleKhnmEPd9ksMYUD2al19PTR2C11d5ZN_RBUDHMU7oftV`

func main() {
	ctx := context.Background()
	service, err := ap.NewService(ctx, option.WithCredentialsFile(configFile))
	if err != nil {
		panic(err)
	}

	/*fmt.Println(service.Reviews.List(packageName))
	do, err := service.Inappproducts.List(packageName).Do()
	if err != nil {
		panic(err)
	}

	for _, v := range do.Inappproduct {
		bytes, _ := v.MarshalJSON()
		fmt.Println(string(bytes))
	}*/

	//call := ap.NewPurchasesProductsService(service).Get(packageName, productId, token)
	// 退款的数据
	// fmt.Println(service.Purchases.Voidedpurchases.List(packageName))
	// 验证 请正确填充 productId 和 token
	purchase, err := service.Purchases.Products.Get(packageName, productId, token2).Do()
	if err != nil {
		panic(err)
	}
	bytes, _ := purchase.MarshalJSON()
	fmt.Println("order: ", string(bytes))
	// acknowledge 订单数据 请正确填充 productId 和 token
	/*err = service.Purchases.Products.Acknowledge(packageName, productId, token, &ap.ProductPurchasesAcknowledgeRequest{
		DeveloperPayload: purchase.DeveloperPayload,
	}).Do()*/

	err = service.Purchases.Products.Acknowledge(packageName, productId, token2, nil).Do()
	if err != nil {
		panic(err)
	}
}
