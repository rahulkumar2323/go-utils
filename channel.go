package main

import (
	"fmt"
	"sync"
)

func main()  {
	mychan := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go func(wg *sync.WaitGroup, ch chan int){
		mychan <- 5
		mychan <- 6
		mychan <- 3
		close(mychan)
		wg.Done()
	}(wg, mychan)

	go func(wg *sync.WaitGroup, ch chan int){
		for {
			val, isopen:= <-mychan
			if !isopen{
				break
			}
			fmt.Println(val)
		}
		wg.Done()
	}(wg, mychan)

	wg.Wait()
}
