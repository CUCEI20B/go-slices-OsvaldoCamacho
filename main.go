package main

import "fmt"

func main() {
	var n int
	var auxS int
	var s []int
	var sum int = 0
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&auxS)
		s = append(s, auxS)

	}

	for _, v := range s {
		sum = sum + v

	}
	fmt.Println(sum)
}
