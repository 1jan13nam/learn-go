package main

import (
	"encoding/json"
	"fmt"
)

type Names struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	//fmt.Println(ss)

	ss := []byte(s)
	//fmt.Println(ss)

	var names []Names

	err := json.Unmarshal(ss, &names)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Printf("%+v", names)

	for i, v := range names {
		fmt.Println("=======For Person: ==========", i)
		fmt.Println("\t", v.First, v.Last, v.Age)
		for ii, vv := range v.Sayings {
			fmt.Println("Saying number:", ii)
			fmt.Println(vv)
		}

	}

}
