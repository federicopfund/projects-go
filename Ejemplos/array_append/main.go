package main 

import (
	"fmt"
	_"array_append/controller"
	 "sync"
	
)

func main() {
	a := []int{1, 2, 3}
	a = append(a, 6)
	fmt.Println(a)
	b := a[:3]
	b = append(b, 5)
	fmt.Println(b)
	ch := make(chan int, 4)
	ch <- 1
	ch <- 2
	close(ch)
	c := <-ch
	d := <-ch
	fmt.Println(c, d)
	
	
	var count int
	var wg sync.WaitGroup

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()

	}
	wg.Wait()
	fmt.Println(count)
	for i := range fibs(5) {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}




