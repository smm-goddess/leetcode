package findrepeateddnasequences

/*
 * @lc app=leetcode id=187 lang=golang
 *
 * [187] Repeated DNA Sequences
 *
 * https://leetcode.com/problems/repeated-dna-sequences/description/
 *
 * algorithms
 * Medium (35.65%)
 * Total Accepted:    123.4K
 * Total Submissions: 344.4K
 * Testcase Example:  '"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"'
 *
 * All DNA is composed of a series of nucleotides abbreviated as A, C, G, and
 * T, for example: "ACGAATTCCG". When studying DNA, it is sometimes useful to
 * identify repeated sequences within the DNA.
 * 
 * Write a function to find all the 10-letter-long sequences (substrings) that
 * occur more than once in a DNA molecule.
 * 
 * Example:
 * 
 * 
 * Input: s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
 * 
 * Output: ["AAAAACCCCC", "CCCCCAAAAA"]
 * 
 * 
 */
func findRepeatedDnaSequences(s string) []string {
	length := len(s)
	if length <= 10 {
		return []string{}
	}
	bs := []byte(s)
	sumMap := make(map[int][][]byte)
	resultStrings := make([]string, 0)
	for i := 0; i < length-9; i++ {
		stringBytes := bs[i : i+10]
		sum := byteSum(stringBytes)
		if l, ok := sumMap[sum]; ok {
			alreadyIn := false
			for i := range l {
				if bytesEqual(l[i], stringBytes) {
					alreadyIn = true
					alreadyInResult := false
					str := string(stringBytes)
					for _, result := range resultStrings {
						if str == result {
							alreadyInResult = true
							break
						}
					}
					if !alreadyInResult {
						resultStrings = append(resultStrings, str)
					}
					break
				}
			}
			if !alreadyIn {
				sumMap[sum] = append(l, stringBytes)
			}
		} else {
			sumMap[sum] = [][]byte{stringBytes}
		}
	}

	return resultStrings
}

func byteSum(a []byte) (r int) {
	for i := range a {
		r += int(a[i])
	}
	return
}

func bytesEqual(a, b []byte) bool {
	for i, c := range a {
		if b[i] != c {
			return false
		}
	}
	return true
}
