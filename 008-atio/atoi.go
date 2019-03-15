package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(myAtio(" "))
	//fmt.Println(myAtio("13213123123213"))
	//fmt.Println(myAtio("-42"))
	//fmt.Println(myAtio("42"))
	//fmt.Println(myAtio("words and 987"))
	//fmt.Println(myAtio("    42"))
	//fmt.Println(myAtio("4193 with words"))
	//fmt.Println(myAtio("-91283472332"))
	fmt.Println(math.MaxInt32)
	fmt.Println(myAtio("2147483648"))
}

func myAtio(str string) int {
	sli := []byte(str)
	var answer int32
	for len(sli) > 0 && sli[0] == ' ' {
		sli = sli[1:]
	}
	if len(sli) == 0 {
		return 0
	}
	var symbol int32
	symbol = 1
	if sli[0] == '-' {
		symbol = -1
		sli = sli[1:]
	} else if sli[0] == '+' {
		symbol = 1
		sli = sli[1:]
	}
	for i := 0; i < len(sli); i++ {
		c := sli[i]
		if c < '0' || c > '9' {
			return int(symbol * answer)
		} else {
			b := int32(c - '0')
			num := answer*10 + b
			fmt.Println(num)
			if (num-b)/10 != answer || num < 0 {
				if symbol == 1 {
					return math.MaxInt32
				} else {
					return math.MinInt32
				}
			} else {
				answer = num
			}
		}
	}
	answer *= symbol
	return int(answer)
}
