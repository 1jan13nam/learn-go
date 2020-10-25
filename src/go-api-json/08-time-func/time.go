package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println("Time now:", time.Now())
	fmt.Println("Time formatted", t.Format(time.RFC3339))
}
