package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient(nil, "127.0.0.1:7051", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	// reply1, err := client.MountainStealsSale(context.Background(),
	// 	"promo/msbf19",
	// 	3*60*1000)

	// if err != nil {
	// 	fmt.Printf("MountainStealsSale %v", err)

	// 	return
	// }

	// if reply1 != nil {
	// 	fmt.Printf("\n%v", reply1)
	// }

	reply2, err := client.MountainStealsProduct(context.Background(),
		"canada-goose-men-s-waffle-slouchy-beanie_10360288",
		3*60*1000)

	if err != nil {
		fmt.Printf("MountainStealsProduct %v", err)

		return
	}

	if reply2 != nil {
		fmt.Printf("\n%v", reply2)
	}

	return
}
