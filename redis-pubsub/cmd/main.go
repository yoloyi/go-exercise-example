package main

import (
	"fmt"
	"log"
	"sync"

	example "github.com/yoloyi/redis-example"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go func(i int) {
			client := example.InitRedis()
			defer wg.Done()
			pb := client.Subscribe("wait-close")
			fmt.Println("阻塞，等待读取 Channel 信息")
			for {
				select {
				case mg := <-pb.Channel():
					if mg.Payload == "close" {
						wg.Done()
					} else {
						log.Println("接受 channel 信息", mg.Payload)
					}
				default:
				}
			}
		}(i)
	}

	wg.Wait()
	log.Println("结束 channel")
}
