package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//var jsonBlob = []byte(`[{"First":"Jagga","Last":"Jassos","Age":24},
	//{"First":"Gabbar","Last":"Singh","Age":43}]`)

	s := `[{"First":"Jagga","Last":"Jassos","Age":24},{"First":"Gabbar","Last":"Singh","Age":43}]`
	bs := []byte(s)

	type Person struct {
		First string `json:"First"`
		Last  string `json:"Last"`
		Age   int    `json:"Age"`
	}

	pp := []Person{}
	err := json.Unmarshal(bs, &pp)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", pp)
	fmt.Printf("%T\n", bs)
	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", pp)

	for i, v := range pp {
		fmt.Println(i, v)
		fmt.Println(v.First, v.Last, v.Age)
	}

}
