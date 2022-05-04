package main

import "fmt"

func main() {
	nums := []int{}
	for i := 1; i <= 10; i++ {
		nums = append(nums, i)
	}
	for _, num := range nums {
		if num%2 == 0 {
			fmt.Printf("%v is even\n", num)
		} else {
			fmt.Printf("%v is odd\n", num)
		}
	}
}
