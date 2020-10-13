package main

import "fmt"

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}

	var digInterval = []int{3, 3, 3, 3, 3, 4, 3, 4}
	var digStart = []int{0, 3, 6, 9, 12, 15, 19, 22}

	var result = make([]string, 0)
	for _, d := range digits {
		letters := make([]string, 0)
		dig := int(d) - 48
		for i := 0; i < digInterval[dig-2]; i++ {
			letters = append(letters, string(97+digStart[dig-2]+i))
		}
		if len(result) == 0 {
			result = letters
			continue
		}
		var temp = make([]string, 0)
		for _, char1 := range result {
			for _, char2 := range letters {
				temp = append(temp, char1+char2)
			}
		}
		result = temp
	}
	return result
}

// func abs(a int) int {
// 	if a < 0 {
// 		return -a
// 	}
// 	return a
// }

func main() {
	// test1 := make([]int, 0)
	// rand.Seed(time.Now().UnixNano())
	// for i := 0; i < 100; i++ {
	// 	test1 = append(test1, rand.Intn(101)-50)
	// }
	test1 := "7"
	fmt.Println(letterCombinations(test1))
}
