package main

import (
	"fmt"
	"time"
)

func main() {
	worlds()
	time.Sleep(2 * time.Second)
}

func worlds() {
	for x := 0; x < 10; x++ {
		go world(x)
	}
}

func world(i int) {
	fmt.Println("hallo world ", i)
}
