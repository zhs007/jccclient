package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient(nil, "127.0.0.1:7051", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	reply, err := client.GetJDActive(context.Background(),
		"3nTQQZ66AGtiwwtRcikGFnT1DVjX/index.html",
		3*60*1000)

	if err != nil {
		fmt.Printf("GetJDActive %v", err)

		return
	}

	if reply != nil {
		fmt.Printf("%v", reply)
	}

	reply1, err := client.GetJDProduct(context.Background(),
		"100006585530.html",
		3*60*1000)

	if err != nil {
		fmt.Printf("GetJDProduct %v", err)

		return
	}

	if reply1 != nil {
		fmt.Printf("%v", reply1)
	}

	return
}
