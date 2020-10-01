package main

func myAtoi(str string) int {
	var num, flag int
	// flag用来记录是否有数字出现，如果出现了则后续再出现空格表示数字结束
	sign := 1
	//sign记录正负号

	for _, n := range str {
		if n == ' ' && flag == 0 {
			continue
		} else if n >= '0' && n <= '9' {
			num = num*10 + int(n-48)
			flag = 1
			if num > 1<<31-1 && sign == 1 {
				return 1<<31 - 1
			} else if num > 1<<31 && sign == -1 {
				return -(1 << 31)
			}
		} else if n == '+' && flag == 0 {
			flag = 1
		} else if n == '-' && flag == 0 {
			sign = -1
			flag = 1
		} else {
			break
		}
	}
	return num * sign
}

// func main() {
// 	fmt.Print(myAtoi("9223372036854775808"))
// }
