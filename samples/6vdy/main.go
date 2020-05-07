package main

import (
	"context"
	"fmt"

	"github.com/zhs007/jccclient"
)

func main() {
	client := jccclient.NewClient(nil, "127.0.0.1:7052", "wzDkh9h2fhfUVuS9jZ8uVbhV3vC5AWX3")

	// reply1, err := client.P6vdyMovies(context.Background(), "https://www.6vdy.org",
	// 	3*60*1000)

	// if err != nil {
	// 	fmt.Printf("P6vdyMovies %v", err)

	// 	return
	// }

	// if reply1 != nil {
	// 	fmt.Printf("\n%v", reply1)
	// }

	reply2, err := client.P6vdyMovie(context.Background(), "https://www.6vdy.org/dianshiju/oumeiju/13261.html",
		3*60*1000)

	if err != nil {
		fmt.Printf("P6vdyMovie %v", err)

		return
	}

	if reply2 != nil {
		fmt.Printf("\n%v", reply2)
	}

	return
}
