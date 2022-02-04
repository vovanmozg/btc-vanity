package btcvanity

import (
	"strings"
	"regexp"
)
// 1Ej6c1YnJYMoNDynBaWPUzMrb2GxUCTc99
// isMatch checks if the btc wallet matches the prefix of pattern
func isMatch(pattern string, addr string) bool {
	// No digits:
	// re, _ := regexp.Compile(`[a-zA-Z]{33}`)

	// Digits:
	//re, _ := regexp.Compile(`[0-9]{34}`)

	// Part of has - alternating vowels and consonants
	// re, _ := regexp.Compile(`(?i)([aouie][bcdfghjklmnpqrstvwxz]){7}`)

	// Cyrilic-like letters
	// re, _ := regexp.Compile(`[ABCEHKMOPTX]{8}`)

	// Consecutive capital letters
	// re, _ := regexp.Compile(`[A-Z]{20}`)

	// Word from any-case letters 
	//re, _ := regexp.Compile(`(?i)wozniak`)

	re, _ := regexp.Compile(pattern)
	
	matched := re.MatchString(addr)
	return matched

	addr = strings.ToLower(addr[1 : len(addr)-1])
	pattern = strings.ToLower(pattern)
	return strings.HasPrefix(addr, pattern)
}
