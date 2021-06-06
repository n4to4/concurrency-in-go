package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fanoutin1()
}

func fanoutin1() {
	take := func(
		done <-chan interface{},
		valueStream <-chan interface{},
		num int,
	) <-chan interface{} {
		takeStream := make(chan interface{})
		go func() {
			defer close(takeStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case takeStream <- <-valueStream:
				}
			}
		}()
		return takeStream
	}

	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	toInt := func(
		done <-chan interface{},
		valueStream <-chan interface{},
	) <-chan int {
		intStream := make(chan int)
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				case intStream <- v.(int):
				}
			}
		}()
		return intStream
	}

	isPrime := func(n int) bool {
		if n < 0 {
			panic("negative")
		}
		if n < 2 {
			return false
		}
		if n == 2 {
			return true
		}
		//for x := 2; float64(x) <= math.Sqrt(float64(n)); x++ {
		for x := 2; x <= n; x++ {
			if n%x == 0 {
				return false
			}
		}
		return true
	}

	primeFinder := func(
		done <-chan interface{},
		valueStream <-chan int,
	) <-chan interface{} {
		intStream := make(chan interface{})
		go func() {
			defer close(intStream)
			for v := range valueStream {
				select {
				case <-done:
					return
				default:
					//fmt.Println("==> v: ", v)
					if isPrime(v) {
						intStream <- v
					} else {
						continue
					}
				}
			}
		}()
		return intStream
	}

	//rand := func() interface{} { return rand.Intn(50000000) }
	rand := func() interface{} { return rand.Intn(50000) }

	done := make(chan interface{})
	defer close(done)

	start := time.Now()

	randIntStream := toInt(done, repeatFn(done, rand))
	fmt.Println("Primes:")
	for prime := range take(done, primeFinder(done, randIntStream), 10) {
		fmt.Printf("\t%d\n", prime)
	}
	fmt.Printf("Search took: %v", time.Since(start))
}
