package main

import (
	"fmt"
)

func isSubsequence(s string, t string) bool {

	if len(s) == 0 && len(t) >= 0 {
		return true
	}

	if len(s) > 0 && len(t) == 0 {
		return false
	}

	if s[0] == t[0] {
		return isSubsequence(s[1:], t[1:])
	} else {
		return isSubsequence(s, t[1:])
	}
}

func main() {
	var s = "abc"
	var t = "taxxbce"

	fmt.Println(isSubsequence(s, t))
}
