package main

import "fmt"

func main(){
	//fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations("235"))
	fmt.Println(letterCombinations(""))
}



func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	getNumbers := func(c byte) []byte {
		switch c {
		case '2':
			return []byte{'a','b','c'}
		case '3':
			return []byte{'d', 'e', 'f'}
		case '4':
			return []byte{'g', 'h', 'i'}
		case '5':
			return []byte{'j', 'k', 'l'}
		case '6':
			return []byte{'m', 'n', 'o'}
		case '7':
			return []byte{'p', 'q', 'r', 's'}
		case '8':
			return []byte{'t', 'u', 'v'}
		case '9':
			return []byte{'w', 'x', 'y', 'z'}
		default:
			return []byte{' '}
		}
	}

	reduce := func(a [][]byte, b []byte) [][]byte{
		r := make([][]byte, 0, len(a) * len(b))
		length := len(a[0])
		for i1 := range a {
			for i2 := range(b) {
				sli := make([]byte, length)
				copy(sli,a[i1])
				r = append(r, append(sli, b[i2]))
			}
		}
		return r
	}

	byteArray := make([][]byte, 0, len(digits))
	for i := range(digits) {
		byteArray = append(byteArray, getNumbers(digits[i]))
	}

	init := make([][]byte, 0, len(byteArray[0]))
	for _, c := range(byteArray[0]) {
		init = append(init, []byte{c})
	}

	byteArray = byteArray[1:]

	for len(byteArray) > 0 {
		init = reduce(init, byteArray[0])
		byteArray = byteArray[1:]
	}

	result := make([]string, 0, len(init))
	for i := range(init) {
		result = append(result, string(init[i]))
	}
	return result

}

