package restoreipaddresses

import (
	"fmt"
	"strconv"
	"strings"
)

/*
 * @lc app=leetcode id=93 lang=golang
 *
 * [93] Restore IP Addresses
 *
 * https://leetcode.com/problems/restore-ip-addresses/description/
 *
 * algorithms
 * Medium (31.02%)
 * Total Accepted:    135.4K
 * Total Submissions: 434.9K
 * Testcase Example:  '"25525511135"'
 *
 * Given a string containing only digits, restore it by returning all possible
 * valid IP address combinations.
 *
 * Example:
 *
 *
 * Input: "25525511135"
 * Output: ["255.255.11.135", "255.255.111.35"]
 *
 *
 */
func restoreIpAddresses(s string) []string {
	ips := restore(s, 0)
	var result []string
	for _, ip := range ips {
		if len(ip) == 4 {
			result = append(result, strings.Join(ip, "."))
		}
	}
	fmt.Println(result)
	return result
}

func restore(s string, depth int) [][]string {
	if len(s) == 0 {
		return [][]string{}
	}
	if depth == 3 {
		if validIp(s) {
			return [][]string{{s}}
		} else {
			return [][]string{}
		}
	} else {
		s1 := s[0:1]
		var result [][]string
		if validIp(s1) {
			ips := restore(s[1:], depth+1)
			for _, ip := range ips {
				result = append(result, append([]string{s1}, ip...))
			}
			if len(s) >= 2 {
				s2 := s[0:2]
				if validIp(s2) {
					ips = restore(s[2:], depth+1)
					for _, ip := range ips {
						result = append(result, append([]string{s2}, ip...))
					}
					if len(s) >= 3 {
						s3 := s[0:3]
						if validIp(s3) {
							ips = restore(s[3:], depth+1)
							for _, ip := range ips {
								result = append(result, append([]string{s3}, ip...))
							}
						}
					}
				}
			}
		}
		return result
	}
}

func validIp(s string) bool {
	if ip, err := strconv.Atoi(s); err == nil && ((s[0] != '0' && ip > 0 && ip <= 255) || (s == "0")) {
		return true
	}
	return false
}
