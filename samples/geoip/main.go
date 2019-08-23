package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient(nil, "127.0.0.1:7051", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	reply, err := client.GetGeoIP(context.Background(),
		"47.90.46.159",
		"ipvoid")

	if err != nil {
		fmt.Printf("GetGeoIP %v", err)

		return
	}

	if reply != nil {
		fmt.Printf("%v", reply)
	}

	return
}
