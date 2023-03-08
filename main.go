package main

import (
	"fmt"
	"time"
	"yampc/client"
)

func main() {
	c := client.New("127.0.0.1:2137")
	// time.Sleep(time.Second)
	// c.PushAlbum(14)
	// time.Sleep(time.Second)
	for _, s := range c.Queue().Songs {
		fmt.Println(s)
	}
	c.Seek(false, time.Second*20)
	c.SetPause(true)
	// c.PLay()
	// fmt.Println("hello, world")
}
