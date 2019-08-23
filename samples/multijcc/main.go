package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	cfg, err := jccclient.LoadConfig("./config.yaml")
	if err != nil {
		fmt.Printf("LoadConfig %v", err)

		return
	}

	mgr, err := jccclient.NewClientMgr(cfg)
	if err != nil {
		fmt.Printf("NewClientMgr %v", err)

		return
	}

	lst := []string{"47.90.46.159", "47.90.46.158"}

	endchan := make(chan int)
	tasknums := 0

	for _, v := range lst {
		cc := mgr.GetClient("localhost", "www.ipvoid.com")
		if cc == nil {
			fmt.Printf("GetClient %v", "localhost")

			return
		}

		go func(client *jccclient.Client, ipaddr string) {
			reply, err := client.GetGeoIP(context.Background(),
				ipaddr,
				"ipvoid")

			if err != nil {
				fmt.Printf("GetGeoIP %v", err)

				// return
			}

			if reply != nil {
				fmt.Printf("%v", reply)
			}

			endchan <- 1
		}(cc, v)
	}

	for {
		<-endchan
		tasknums++
		if tasknums == len(lst) {
			break
		}
	}

	return
}
