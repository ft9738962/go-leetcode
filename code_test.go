package main

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	// test1 := "()"
	test2 := "()[]{}"
	// test3 := "(]"
	// test4 := "([)]"
	test5 := "{[]}"
	test6 := "[([]])"
	// test6P := "c*a**b"
	// test7S := "aaa"
	// test7P := "a*a"
	// test8S := "aaa"
	// test8P := "ab*a*c*a"
	// 	str10 := "-   234"

	// got1 := isValid(test1)
	// want1 := true
	// if got1 != want1 {
	// 	t.Errorf("test 1 got %#v want %#v", got1, want1)
	// }

	got2 := isValid(test2)
	want2 := true
	if got2 != want2 {
		t.Errorf("test 2 got %#v want %#v", got2, want2)
	}

	// got3 := isValid(test3)
	// want3 := false
	// if got3 != want3 {
	// 	t.Errorf("test 3 got %#v want %#v", got3, want3)
	// }

	// got4 := isValid(test4)
	// want4 := false
	// if got4 != want4 {
	// 	t.Errorf("test 4 got %#v want %#v", got4, want4)
	// }

	got5 := isValid(test5)
	want5 := true
	if got5 != want5 {
		t.Errorf("test 5 got %#v want %#v", got5, want5)
	}

	got6 := isValid(test6)
	want6 := false
	if got6 != want6 {
		t.Errorf("test 6 got %t want %t", got6, want6)
	}

	// got7 := isMatch(test7S, test7P)
	// want7 := true
	// if got7 != want7 {
	// 	t.Errorf("test 7 got %t want %t", got7, want7)
	// }

	// got8 := isMatch(test8S, test8P)
	// want8 := true
	// if got8 != want8 {
	// 	t.Errorf("test 8 got %t want %t", got8, want8)
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
