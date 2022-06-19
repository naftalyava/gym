package main

import (
	"fmt"
)

func isIsomorphic(s string, t string) bool {

	var hash = make(map[byte]byte)
	var hash_b = make(map[byte]byte)

	if len(s) != len(t) {
		return false
	}

	if len(s) == 0 && len(t) == 0 {
		return true
	}

	for i, _ := range s {

		mapping, found := hash[s[i]]

		if !found && s[i] == t[i] {
			hash[s[i]] = t[i]
			continue
		} else if found == true && mapping == t[i] {
			continue
		} else if found {
			return false
		} else {
			hash[s[i]] = t[i]
		}
	}

	for _, e := range hash {
		_, found := hash_b[e]
		if found {
			return false
		}
		hash_b[e] = e
	}

	return true

}

func main() {
	var s = "babc"
	var t = "baba"

	fmt.Println(isIsomorphic(s, t))
}

/*
	aac
	aca

*/
