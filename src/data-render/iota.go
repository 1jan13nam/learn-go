package main

import "fmt"

const (
	a = iota
	b
	f
	d = iota
	e
	k
)

const (
	// kb = 1024
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 4
	fmt.Println(a, b, f)
	fmt.Println(d, e, k)
	fmt.Printf("%T\n%T\n%T\n", a, b, f)
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b\n", gb, gb)

	kk := 42

	fmt.Printf("%d\t%b\t\t%#x\n", kk, kk, kk)

	if kk == 42 {
		fmt.Println("Good")
	}

	if kk <= 42 {
		fmt.Println("Good")
	}

	if kk >= 42 {
		fmt.Println("Good")
	}

	if kk != 43 {
		fmt.Println("Good")
	}

	if kk < 45 {
		fmt.Println("Good")
	}
	if kk > 41 {
		fmt.Println("Good")
	}

	const (
		bb     = 34
		cc int = 43
	)

	fmt.Println(bb, cc)

	trf := 42

	fmt.Printf("%d\t\t%b\t\t%#x\n", trf, trf, trf)

	kjh := trf << 1
	fmt.Printf("%d\t\t%b\t\t%#x\n", kjh, kjh, kjh)

	var jklj string
	jklj = `Love go
	or dont `
	fmt.Println(jklj)

	const (
		yr1 = 2010
		yr2 = yr1 - iota
		yr3 = yr1 - iota
		yr4 = yr1 - iota
	)

	fmt.Println(yr1, yr2, yr3, yr4)

}
