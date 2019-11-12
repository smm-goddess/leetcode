package masking_personal_information

import "testing"

func TestMask(t *testing.T) {
	if maskPII("LeetCode@LeetCode.com") != "l*****e@leetcode.com" {
		t.Error("leetcode error")
	}
	if maskPII("AB@qq.com") != "a*****b@qq.com" {
		t.Error("AB@qq.com")
	}
	if maskPII("1(234)567-890") != "***-***-7890" {
		t.Error("1(234) error")
	}
	if maskPII("86-(10)12345678") != "+**-***-***-5678" {
		t.Error("86 error")
	}
	if maskPII("aa@bb.ca") != "a*****a@bb.ca" {
		t.Error("aa@bb.ca")
	}
}
