package main

import (
    "fmt"
)

func main(){
    var a int64 = 0x7fffffffffffffff
    fmt.Printf("number %d",a)
}

func myAtio(s string) int {
    sli := []byte(s)
    var answer int64 = 0
    for sli[0] == ' '{
        sli = sli[1:]
    }
    symbol := 1
    maxPos := 0x7FFFFFFFFFFFFFFF >> 2
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
            return symbol * answer
        } else {
            b := int(c - '0')
            max := maxPos
            for j := 0; j < 5; j++ {
               max = max - answer
            }
            max = max << 2
            if symbol == 1 {
                if b > max {
                    return maxPos
                }
            } else {
                if b > max + 1 {
                    return maxPos
                }
            }
        }
    }
    return answer
}
