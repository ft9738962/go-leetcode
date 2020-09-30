package main

import "fmt"

func reverse(x int) int {
	if x == 0 {
		return 0
	}
	var r int

	for x != 0 {
		r = x%10 + r*10
		x = x / 10
	}
	if r > 1<<31-1 || r < -(1<<31) {
		return 0
	}
	return r

}

func main() {
	fmt.Println(reverse(-321))
}
