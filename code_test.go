package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	test1 := 121
	test2 := -121
	test3 := 10
	test4 := -101
	//  str5 := "-91283472332"
	// 	str6 := "+1"
	// 	str7 := "+-12"
	// 	str8 := "00000-42a1234"
	// 	str9 := "9223372036854775808"
	// 	str10 := "-   234"

	got1 := isPalindrome(test1)
	want1 := true
	if got1 != want1 {
		t.Errorf("test 1 got %t want %t", got1, want1)
	}

	got2 := isPalindrome(test2)
	want2 := false
	if got2 != want2 {
		t.Errorf("test 2 got %t want %t", got2, want2)
	}

	got3 := isPalindrome(test3)
	want3 := false
	if got3 != want3 {
		t.Errorf("test 3 got %t want %t", got3, want3)
	}

	got4 := isPalindrome(test4)
	want4 := false
	if got4 != want4 {
		t.Errorf("test 4 got %t want %t", got4, want4)
	}

	// got5 := myAtoi(str5)
	// want5 := -2147483648
	// if got5 != want5 {
	// 	t.Errorf("5 got %d want %d", got5, want5)
	// }

	// got6 := myAtoi(str6)
	// want6 := 1
	// if got6 != want6 {
	// 	t.Errorf("6 got %d want %d", got6, want6)
	// }

	// got7 := myAtoi(str7)
	// want7 := 0
	// if got7 != want7 {
	// 	t.Errorf("7 got %d want %d", got7, want7)
	// }

	// got8 := myAtoi(str8)
	// want8 := 0
	// if got8 != want8 {
	// 	t.Errorf("8 got %d want %d", got8, want8)
	// }

	// got9 := myAtoi(str9)
	// want9 := 2147483647
	// if got9 != want9 {
	// 	t.Errorf("9 got %d want %d", got9, want9)
	// }

	// got10 := myAtoi(str10)
	// want10 := 0
	// if got10 != want10 {
	// 	t.Errorf("10 got %d want %d", got10, want10)
	// }
}
