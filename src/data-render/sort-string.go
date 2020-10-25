package main

import (
	"fmt"
	"sort"
)

func main() {
	as := []string{"dil", "deewana", "bin", "dil", "sajana", "ke", "maane", "ke", "na"}
	sort.Strings(as)
	fmt.Println(as)
	rd := remdup(as)
	fmt.Println(rd)

}

func remdup(aus []string) []string {
	result := []string{}
	mp := map[string]bool{}

	for i, v := range aus {
		if mp[aus[i]] == true {
			fmt.Println("dup found at", i, "Value", v)
		} else {
			mp[aus[i]] = true
			result = append(result, aus[i])
		}
	}

	return result

}
