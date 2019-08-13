package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient("127.0.0.1:7051", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	reply, err := client.AnalyzePage(context.Background(),
		"http://47.90.46.159:8090/game.html?gameCode=nightclub&language=zh_CN&isCheat=true&slotKey=",
		10, 1280, 800, 1.0, false, false)

	if err != nil {
		fmt.Printf("AnalyzePage %v", err)

		return
	}

	if reply != nil {
		fmt.Printf("%v", reply)
	}

	return
}
