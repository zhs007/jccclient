package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient(nil, "127.0.0.1:7051", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	reply, err := client.GetSteepAndCheapProducts(context.Background(),
		"rc/arcteryx-on-sale",
		0, 3*60*1000)

	if err != nil {
		fmt.Printf("GetSteepAndCheapProducts %v", err)

		return
	}

	if reply != nil {
		fmt.Printf("%v", reply)
	}

	reply1, err := client.GetSteepAndCheapProduct(context.Background(),
		"arc-teryx-rho-lt-zip-neck-top-womens?skid=ARC3698-HARCOR-XL",
		3*60*1000)

	if err != nil {
		fmt.Printf("GetSteepAndCheapProducts %v", err)

		return
	}

	if reply1 != nil {
		fmt.Printf("%v", reply1)
	}

	return
}
