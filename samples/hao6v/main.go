package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient(nil, "127.0.0.1:7052", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	// reply1, err := client.Hao6vNewest(context.Background(),
	// 	3*60*1000)

	// if err != nil {
	// 	fmt.Printf("Hao6vNewest %v", err)

	// 	return
	// }

	// if reply1 != nil {
	// 	fmt.Printf("\n%v", reply1)
	// }

	reply2, err := client.Hao6vRes(context.Background(), "http://www.hao6v.com/dy/2020-02-18/JianDieBenSe.html",
		3*60*1000)

	if err != nil {
		fmt.Printf("Hao6vRes %v", err)

		return
	}

	if reply2 != nil {
		fmt.Printf("\n%v", reply2)
	}

	return
}
