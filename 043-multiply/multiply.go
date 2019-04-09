package multiply
/*
Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, also represented as a string.

Example 1:

Input: num1 = "2", num2 = "3"
Output: "6"
Example 2:

Input: num1 = "123", num2 = "456"
Output: "56088"
Note:

1. The length of both num1 and num2 is < 110.
2. Both num1 and num2 contain only digits 0-9.
3. Both num1 and num2 do not contain any leading zero, except the number 0 itself.
4. You must not use any built-in BigInteger library or convert the inputs to integer directly.
*/
func multiply(num1 string, num2 string) string {
  if len(num1) == 0 || len(num2) == 0 || num1[0] == '0' || num2[0] == '0' {
    return "0"
  }
  base := make([]byte, len(num1))
  for i := 0; i < len(num1); i++{
    base[len(num1) - i - 1] = num1[i] - '0'
  } 
  multi := make([][]byte, len(num2))
  for i := 0 ; i < len(num2); i++ {
    zeros := len(num2) - i - 1
    multiRow := make([]byte, zeros + len(num1))
    num := num2[i] - '0'
    for j := 0; j < len(num1); j++{
      multiRow[j + zeros] = num * base[j]
    }
    multi[i] = multiRow
  }
  maxLength := 0
  for _, mul := range multi {
    if maxLength < len(mul) {
      maxLength = len(mul)
    }
  }
  sum, result, digit := 0, 0, 1
  for i := 0; i< maxLength; i++ {
    for _, v := range multi {
      if len(v) > i {
        sum += int(v[i])
      }
    }
    result = digit * (sum % 10) + result
    digit, sum = digit * 10, sum / 10
  }
  for sum > 0 {
    result =  digit * (sum % 10) + result
    digit, sum = digit * 10, sum / 10
  }
  res := make([]byte, 0)
  for result > 0 {
    res = append([]byte{byte(result % 10) + '0'}, res...)
    result = result / 10
  }
  return string(res)
}
