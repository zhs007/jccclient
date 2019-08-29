package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
	jarviscrawlercore "github.com/zhs007/jccclient/proto"
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

	mgr.AddTask(&jccclient.Tags{
		Tag: "localhost",
	},
		&jccclient.Task{
			Callback: func(ctx context.Context, task *jccclient.Task, err error, reply *jarviscrawlercore.ReplyCrawler) {
				if err != nil {
					fmt.Printf("callback %v", err)

					return
				}

				ap, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Geoip)
				if !isok {
					fmt.Printf("callback %v", jccclient.ErrNoReplyAnalyzePage)

					return
				}

				fmt.Printf("callback %v", ap.Geoip)
			},
			RetryNums: 3,
			GeoIP: &jccclient.TaskGeoIP{
				IP: "47.90.46.159",
			},
		})

	mgr.AddTask(&jccclient.Tags{
		Tag: "localhost",
	},
		&jccclient.Task{
			Callback: func(ctx context.Context, task *jccclient.Task, err error, reply *jarviscrawlercore.ReplyCrawler) {
				if err != nil {
					fmt.Printf("callback %v", err)

					return
				}

				ap, isok := (reply.CrawlerResult).(*jarviscrawlercore.ReplyCrawler_Geoip)
				if !isok {
					fmt.Printf("callback %v", jccclient.ErrNoReplyAnalyzePage)

					return
				}

				fmt.Printf("callback %v", ap.Geoip)
			},
			RetryNums: 3,
			GeoIP: &jccclient.TaskGeoIP{
				IP: "47.90.46.158",
			},
		})

	mgr.StartAllTasks(context.Background())

	// lst := []string{"47.90.46.159", "47.90.46.158"}

	// endchan := make(chan int)
	// tasknums := 0

	// for _, v := range lst {
	// 	cc := mgr.GetClient("localhost", "www.ipvoid.com")
	// 	if cc == nil {
	// 		fmt.Printf("GetClient %v", "localhost")

	// 		return
	// 	}

	// 	go func(client *jccclient.Client, ipaddr string) {
	// 		reply, err := client.GetGeoIP(context.Background(),
	// 			ipaddr,
	// 			"ipvoid")

	// 		if err != nil {
	// 			fmt.Printf("GetGeoIP %v", err)

	// 			// return
	// 		}

	// 		if reply != nil {
	// 			fmt.Printf("%v", reply)
	// 		}

	// 		endchan <- 1
	// 	}(cc, v)
	// }

	// for {
	// 	<-endchan
	// 	tasknums++
	// 	if tasknums == len(lst) {
	// 		break
	// 	}
	// }

	return
}
