package main

import (
	"testing"
)

func TestMyatoi(t *testing.T) {
	str1 := "42"
	str2 := "   -42"
	str3 := "4193 with words"
	str4 := "words and 987"
	str5 := "-91283472332"
	str6 := "+1"
	str7 := "+-12"
	str8 := "00000-42a1234"
	str9 := "9223372036854775808"
	str10 := "-   234"

	got1 := myAtoi(str1)
	want1 := 42
	if got1 != want1 {
		t.Errorf("1 got %d want %d", got1, want1)
	}

	got2 := myAtoi(str2)
	want2 := -42
	if got2 != want2 {
		t.Errorf("2 got %d want %d", got2, want2)
	}

	got3 := myAtoi(str3)
	want3 := 4193
	if got3 != want3 {
		t.Errorf("3 got %d want %d", got3, want3)
	}

	got4 := myAtoi(str4)
	want4 := 0
	if got4 != want4 {
		t.Errorf("4 got %d want %d", got4, want4)
	}

	got5 := myAtoi(str5)
	want5 := -2147483648
	if got5 != want5 {
		t.Errorf("5 got %d want %d", got5, want5)
	}

	got6 := myAtoi(str6)
	want6 := 1
	if got6 != want6 {
		t.Errorf("6 got %d want %d", got6, want6)
	}

	got7 := myAtoi(str7)
	want7 := 0
	if got7 != want7 {
		t.Errorf("7 got %d want %d", got7, want7)
	}

	got8 := myAtoi(str8)
	want8 := 0
	if got8 != want8 {
		t.Errorf("8 got %d want %d", got8, want8)
	}

	got9 := myAtoi(str9)
	want9 := 2147483647
	if got9 != want9 {
		t.Errorf("9 got %d want %d", got9, want9)
	}

	got10 := myAtoi(str10)
	want10 := 0
	if got10 != want10 {
		t.Errorf("10 got %d want %d", got10, want10)
	}
}
