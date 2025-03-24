package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := &sync.WaitGroup{}
	s := 0
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(val int, ch chan<- int) {
			ch <- val
			s++
		}(i, ch)

		go func(ch <-chan int, wg *sync.WaitGroup) {
			defer wg.Done()
			fmt.Println(<-ch)
		}(ch, wg)

	}
	wg.Wait()
}

// ch := make(chan int)
// 	wg := &sync.WaitGroup{}
// 	wg.Add(10)
// 	for i := 0; i < 10; i++ {
// 		go func(ch chan<- int) {
// 			ch <- i
// 		}(ch)

// 		go func(ch <-chan int, wg *sync.WaitGroup) {
// 			defer wg.Done()
// 			fmt.Println(<-ch)
// 		}(ch, wg)
// 	}
// 	wg.Wait()

// ch := make(chan int)

// 	for i := 0; i < 10; i++ {
// 		wg := &sync.WaitGroup{}
// 		wg.Add(1)
// 		go func(ch chan<- int) {
// 			ch <- i
// 		}(ch)

// 		go func(ch <-chan int, wg *sync.WaitGroup) {
// 			defer wg.Done()
// 			fmt.Println(<-ch)
// 		}(ch, wg)
// 		wg.Wait()
// 	}
