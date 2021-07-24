package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go sleepyGopher(i)
	}
	time.Sleep(5 * time.Second)
}

func sleepyGopher(i int) {
	fmt.Println(i)
	time.Sleep(2 * time.Second)
	fmt.Println("... snore ...",i)
}
