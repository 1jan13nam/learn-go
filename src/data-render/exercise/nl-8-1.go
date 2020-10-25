package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "Nancy",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	b, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)

	// your code goes here
}
