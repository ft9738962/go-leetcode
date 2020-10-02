package main

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	test1S := "aa"
	test1P := "a"
	test2S := "aa"
	test2P := "a*"
	test3S := "ab"
	test3P := ".*"
	test4S := "aab"
	test4P := "c*a*b"
	test5S := "mississippi"
	test5P := "mis*is*p*."
	test6S := "aab"
	test6P := "c*a**b"
	test7S := "aaa"
	test7P := "a*a"
	test8S := "aaa"
	test8P := "ab*a*c*a"
	// 	str10 := "-   234"

	got1 := isMatch(test1S, test1P)
	want1 := false
	if got1 != want1 {
		t.Errorf("test 1 got %t want %t", got1, want1)
	}

	got2 := isMatch(test2S, test2P)
	want2 := true
	if got2 != want2 {
		t.Errorf("test 2 got %t want %t", got2, want2)
	}

	got3 := isMatch(test3S, test3P)
	want3 := true
	if got3 != want3 {
		t.Errorf("test 3 got %t want %t", got3, want3)
	}

	got4 := isMatch(test4S, test4P)
	want4 := true
	if got4 != want4 {
		t.Errorf("test 4 got %t want %t", got4, want4)
	}

	got5 := isMatch(test5S, test5P)
	want5 := false
	if got5 != want5 {
		t.Errorf("test 5 got %t want %t", got5, want5)
	}

	got6 := isMatch(test6S, test6P)
	want6 := false
	if got6 != want6 {
		t.Errorf("test 6 got %t want %t", got6, want6)
	}

	got7 := isMatch(test7S, test7P)
	want7 := true
	if got7 != want7 {
		t.Errorf("test 7 got %t want %t", got7, want7)
	}

	got8 := isMatch(test8S, test8P)
	want8 := true
	if got8 != want8 {
		t.Errorf("test 8 got %t want %t", got8, want8)
	}

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
