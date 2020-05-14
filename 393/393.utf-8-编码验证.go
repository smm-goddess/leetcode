// package main
// import "fmt"
/*
 * @lc app=leetcode.cn id=393 lang=golang
 *
 * [393] UTF-8 编码验证
 */

// @lc code=start
func utf8Length(data int) int {
	data = data & 0xFF
	switch {
	case data < 0x80 : return 1
	case data >= 0xC0 && data <= 0xDF: return 2
	case data >= 0xE0 && data <= 0xEF: return 3
	case data >= 0xF0 && data <= 0xF7: return 4
	}
	return -1
}

func validRemains(data int) bool {
	data = data & 0xFF
	return data > 0x7F && data < 0xC0
}

func validUtf8(data []int) bool {
		if len(data) == 0 {
			return true
		} else {
			length := utf8Length(data[0])
			if length == -1 {
				return false
			} else {
				if length > len(data) {
					return false
				}
				for i := 1;i < length; i++{
					if !validRemains(data[i]) {
						return false
					}
				}
				if !validUtf8(data[length:]) {
					return false
				}
			}
			return true
		}
}
// @lc code=end

// func main() {
// 	fmt.Println(validUtf8([]int{237}))
// 	fmt.Println(validUtf8([]int{197,130,1}))
// 	fmt.Println(validUtf8([]int{235,140,4}))
// }