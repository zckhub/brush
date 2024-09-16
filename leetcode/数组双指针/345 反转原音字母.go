package main

import "strings"

func reverseVowels(s string) string {
	sList := []byte(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for left < right && !inVowel(sList[left]) {
			left++
		}
		for right > left && !inVowel(sList[right]) {
			right--
		}
		sList[left], sList[right] = sList[right], sList[left]
		left++
		right--

	}
	return string(sList)
}

func inVowel(s byte) bool {
	if strings.Contains("AEIOUaeiou", string(s)) {
		return true
	}
	return false
}
