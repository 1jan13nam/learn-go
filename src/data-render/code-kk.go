package main

import (
	"fmt"
	"sort"
)

func main() {
	as := []string{"dil", "deewana", "bin", "dil", "sajana", "deewana", "maane", "bin", "na"}
	sort.Strings(as)
	fmt.Println(as)
	rd := remdup(as)
	fmt.Println(rd)

}

func remdup(a []string) []string {

	mp := map[string]bool{}
	result := []string{}
	for v := range a {
		if mp[a[v]] == true {
			fmt.Println("...")
		} else {
			mp[a[v]] = true
			result = append(result, a[v])
		}
	}

	return result

}
