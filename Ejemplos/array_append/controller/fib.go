package controller

func fibs(n int) chan int {
	ch := make(chan int)

	go func() {
		defer close(ch) // <1>
		a, b := 1, 1
		for i := 0; i < n; i++ {
			ch <- a
			a, b = b, a+b
		}
	}()

	return ch
}