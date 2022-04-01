package main 

import (
	"fmt"
	"Append_Controler/controller"
)

func main() {
	a := []int{1, 2, 3}
	a = Append(a, 4)
	fmt.Println(a)
	b := a[:3]
	b = append(b, 5)
	fmt.Println(b)
}

