package main

import "fmt"

func fourSum(nums []int, target int) [][]int {

}

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

func main() {
	test1 := make([]int, 0)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 100; i++ {
	// 	test1 = append(test1, rand.Intn(101)-50)
	// }
	fmt.Println(fourSum(test1, 1))
}
