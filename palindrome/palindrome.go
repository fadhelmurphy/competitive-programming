package main
import "fmt"

func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	s := "katak"
	if isPalindrome(s) {
		fmt.Println(s, "adalah palindrome")
	} else {
		fmt.Println(s, "bukan palindrome")
	}
}
