package main

import "fmt"

func main() {
	fmt.Println(isPalindrome("ABC"))
	fmt.Println(isPalindrome("AABCC"))
	fmt.Println(isPalindrome("AAABAAA"))
	fmt.Println(isPalindrome("A"))
}

func isPalindrome(s string) bool {
	n := len(s) - 1
	for i := 0; i <= n/2; i++ {
		if s[i] != s[n-i] {
			return false
		}
	}
	return true
}
