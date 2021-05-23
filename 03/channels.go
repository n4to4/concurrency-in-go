package main

import "fmt"

func main0() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	fmt.Println(<-stringStream)
}

func main1() {
	intStream := make(chan int)
	close(intStream)

	integer, ok := <-intStream
	fmt.Printf("(%v): %v\n", ok, integer)
}

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream {
		fmt.Printf("%v ", integer)
	}
	fmt.Println("")
}
