package main

import (
	"fmt"
	"time"
)

func main0() {
	var data int
	go func() { data++ }()
	time.Sleep(1 * time.Second)
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
