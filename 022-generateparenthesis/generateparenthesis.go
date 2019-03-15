package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	} else if n == 1 {
		return []string{"()"}
	} else {
		remain := n << 1
		var reduce func(parenthesis [][]byte, lRemain, rRemain int) [][]byte
		reduce = func(parenthesis [][]byte, lRemain, rRemain int) [][]byte {
			if lRemain == 0 && rRemain == 0 {
				return parenthesis
			}
			ret := make([][]byte, 0)
			if lRemain > 0 {
				lr := make([][]byte, 0, len(parenthesis))
				for _, v := range parenthesis {
					v2 := make([]byte, len(v))
					copy(v2, v)
					v2 = append(v2, '(')
					lr = append(lr, v2)
				}
				ret = append(ret, reduce(lr, lRemain-1, rRemain)...)
			}
			if rRemain > lRemain {
				lr := make([][]byte, 0, len(parenthesis))
				for _, v := range parenthesis {
					v2 := make([]byte, len(v))
					copy(v2, v)
					v2 = append(v2, ')')
					lr = append(lr, v2)
				}
				ret = append(ret, reduce(lr, lRemain, rRemain-1)...)
			}
			return ret
		}
		init := make([][]byte, 1)
		init[0] = []byte{'('}
		parenthesises := reduce(init, remain>>1-1, remain>>1)
		res := make([]string, 0, len(parenthesises))
		for _, v := range parenthesises {
			res = append(res, string(v))
		}
		return res
	}
}

