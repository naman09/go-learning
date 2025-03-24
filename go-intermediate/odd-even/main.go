package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	sourceChan := make(chan int)
	evenChan := make(chan int)
	oddChan := make(chan int)
	stopSplitter := make(chan struct{})
	evenSumChan := make(chan int)
	oddSumChan := make(chan int)
	stopEvenSum := make(chan struct{})
	stopOddSum := make(chan struct{})

	s1Fin := source("data1.dat", sourceChan)
	s2Fin := source("data2.dat", sourceChan)

	wq := sync.WaitGroup{}

	wq.Add(1)
	go func() {
		for i := 0; i < 2; i++ {
			select {
			case <-s1Fin:
				fmt.Println("Source 1 fin")
			case <-s2Fin:
				fmt.Println("Source 2 fin")
			}
		}
		stopSplitter <- struct{}{}
		wq.Done()
	}()

	wq.Add(1)
	go func() {
		splitter(sourceChan, evenChan, oddChan, stopSplitter, stopEvenSum, stopOddSum)
		wq.Done()
	}()

	wq.Add(1)
	go func() {
		go sum(evenChan, evenSumChan, stopEvenSum)
		wq.Done()
	}()

	wq.Add(1)
	go func() {
		sum(oddChan, oddSumChan, stopOddSum)
		wq.Done()
	}()

	wq.Add(1)
	go func() {
		merger(oddSumChan, evenSumChan)
		wq.Done()
	}()

	wq.Wait()

}

func source(fileName string, sourceChan chan int) chan struct{} {
	//read and put it in source channel
	sourceFin := make(chan struct{})
	go func() {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalln(err)
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			txt := scanner.Text()
			if no, err := strconv.Atoi(txt); err == nil {
				sourceChan <- no
			}
		}
		sourceFin <- struct{}{}
	}()
	return sourceFin
}

func splitter(sourceChan chan int, evenChan chan int, oddChan chan int, stop chan struct{}, stopEvenSum chan struct{}, stopOddSum chan struct{}) {
	//stop when source is finished
LOOP:
	for {
		select {
		case no := <-sourceChan:
			if no%2 == 0 {
				evenChan <- no
			} else {
				oddChan <- no
			}
		case <-stop:
			stopEvenSum <- struct{}{}
			stopOddSum <- struct{}{}
			break LOOP
		}
	}

}

func sum(numberChan chan int, sumChan chan int, stop chan struct{}) {
	total := 0
LOOP:
	for {
		select {
		case no := <-numberChan:
			total += no
		case <-stop:
			sumChan <- total
			break LOOP

		}
	}
}

func merger(oddSumChan chan int, evenSumChan chan int) {
	for i := 0; i < 2; i++ {
		select {
		case oddSum := <-oddSumChan:
			fmt.Println("Odd number sum: ", oddSum)
		case evenSum := <-evenSumChan:
			fmt.Println("Even number sum: ", evenSum)
		}
	}
}
