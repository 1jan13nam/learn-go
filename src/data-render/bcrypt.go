package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `Password@1234`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	loginPword1 := `Password@1234`

	er := bcrypt.CompareHashAndPassword(bs, []byte(loginPword1))

	if er != nil {
		fmt.Println("You can't Login:", err)
		return
	}
	fmt.Println("Logged In")

}
