package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	ch := make(chan int)

	for _, num := range numbers {
		go func(n int) {
			sum += n
			ch <- sum
		}(num)
	}

	var finalSum int
	for range numbers {
		finalSum = <-ch
		fmt.Println("Current Sum:", finalSum)
	}

	fmt.Println("Final Sum:", finalSum)
}
