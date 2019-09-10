package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient(nil, "127.0.0.1:7051", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	reply, err := client.GetJRJFunds(context.Background(), 3*60*1000)

	if err != nil {
		fmt.Printf("GetJRJFunds %v", err)

		return
	}

	if reply != nil {
		fmt.Printf("%v", reply)
	}

	reply1, err := client.GetJRJFund(context.Background(),
		"110011", 3*60*1000)

	if err != nil {
		fmt.Printf("GetJRJFund %v", err)

		return
	}

	if reply1 != nil {
		fmt.Printf("%v", reply1)
	}

	reply2, err := client.GetJRJFundManager(context.Background(),
		"110011", 3*60*1000)

	if err != nil {
		fmt.Printf("GetJRJFundManager %v", err)

		return
	}

	if reply2 != nil {
		fmt.Printf("%v", reply2)
	}

	reply3, err := client.GetJRJFundValue(context.Background(),
		"110011", "2019", 3*60*1000)

	if err != nil {
		fmt.Printf("GetJRJFundValue %v", err)

		return
	}

	if reply3 != nil {
		fmt.Printf("%v", reply3)
	}

	return
}
