package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient(nil, "127.0.0.1:7051", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	// err := client.AlimamaKeepalive(context.Background(),
	// 	3*60*1000)

	// if err != nil {
	// 	fmt.Printf("AlimamaKeepalive %v", err)

	// 	return
	// }

	reply1, err := client.AlimamaSearch(context.Background(),
		"森贝儿家族",
		3*60*1000)

	if err != nil {
		fmt.Printf("AlimamaSearch %v", err)

		return
	}

	if reply1 != nil {
		fmt.Printf("\n%v", reply1)
	}

	// reply2, err := client.AlimamaGetTop(context.Background(),
	// 	3*60*1000)

	// if err != nil {
	// 	fmt.Printf("AlimamaGetTop %v", err)

	// 	return
	// }

	// if reply2 != nil {
	// 	fmt.Printf("\n%v", reply2)
	// }

	return
}
