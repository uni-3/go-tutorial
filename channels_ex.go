package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	arr_x := s[:len(s)/2]
	arr_y := s[len(s)/2:]

	go sum(arr_x, c)
	go sum(arr_y, c)
	y, x := <-c, <-c // receive from c

	fmt.Println("sum x is ", arr_x)
	fmt.Println("sum y is ", arr_y)
	fmt.Println(x, y, x+y)
}
