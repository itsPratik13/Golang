package main

import (
	"fmt"
)

func Sum[T ~[]int | ~[5]int](numbers T) int {
	sum := 0
	// for i:=0;i<5;i++{
	// 	sum+=numbers[i]
	// }
	 for i := 0; i < len(numbers); i++ {
        sum += numbers[i]
    }

	return sum
}

func main() {
	numbers := []int{1, 2, 3}
	fmt.Println(Sum(numbers))
}
