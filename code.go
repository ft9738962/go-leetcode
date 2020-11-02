package main

import "fmt"

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

func findSubstring(s string, words []string) []int {
	cws := &[]string{}
	combineWords("", words, cws)
	subStringDict := map[string]int{}
	for _, word := range *cws {
		if _, ok := subStringDict[word]; ok == true {
			continue
		}
		subStringDict[word] = -1
		l := len(word)
		for i := 0; i < len(s)-l+1; i++ {
			if string(s[i:i+l]) == word {
				subStringDict[word] = i
				break
			}
		}
	}
	result := []int{}
	for _, ind := range subStringDict {
		if ind > -1 {
			result = append(result, ind)
		}
	}
	return result
}

func combineWords(prefix string,
	words []string, cws *[]string) {
	if len(words) == 1 {
		*cws = append(*cws, prefix+words[0])
	} else {
		for i := 0; i < len(words); i++ {
			if i == 0 {
				combineWords(prefix+string(words[i]),
					words[i+1:],
					cws)
			} else if i == len(words)-1 {
				combineWords(prefix+string(words[i]),
					words[:i],
					cws)
			} else {
				combineWords(prefix+string(words[i]),
					append(words[:i], words[i+1:]...),
					cws)
			}
		}
	}
}

func printVals(head *ListNode) {
	fmt.Println("start")
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
	fmt.Println("end")
}

func main() {
	// test1 := []int{1, 1}
	// test2 := []int{7, -3}
	// test3 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	// a6 := ListNode{
	// 	Val:  6,
	// 	Next: nil,
	// }
	// test5 := []int{-2147483648, 2}
	// var c1 *ListNode

	s := "barfoofoobarthefoobarman"
	words := []string{"bar", "foo", "the"}
	// d1 := ListNode{
	// 	Val:  2,
	// 	Next: nil,
	// }

	// d2 := ListNode{
	// 	Val:  6,
	// 	Next: nil,
	// }

	// d3 := ListNode{
	// 	Val:  1130,
	// 	Next: nil,
	// }

	// a5.Next = &a6

	// d1.Next = &d2
	// d2.Next = &d3

	// lists := []*ListNode{&a1, &b1, c1, &d1}
	// lists := []*ListNode{&a1, &b1, &d1}
	// lists := []*ListNode{}

	// var emp *ListNode

	// for _, n := range lists {
	// 	for n.Next != nil {
	// 		fmt.Println(n.Val)
	// 		n = n.Next
	// 	}
	// 	fmt.Println(n.Val)

	// head := mergeTwoLists(&a1, &b1)
	// answer1 := test1[:removeDuplicates(test1)]
	// answer2 := test2[:removeDuplicates(test2)]
	// answer3 := test3[:removeDuplicates(test3)]
	// fmt.Printf("answer1 %v\n", answer1)
	// fmt.Printf("answer2 %v\n", answer2)
	// fmt.Printf("answer3 %v\n", answer3)
	fmt.Println(findSubstring(s, words))
	// fmt.Println(r2)

}
