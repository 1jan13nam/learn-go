package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}
type ByAge []person

type ByName []person

func main() {
	p1 := person{"Jagga", 26}
	p2 := person{"Sonu", 32}
	p3 := person{"Chujja", 64}
	p4 := person{"Titu", 56}

	people := []person{p1, p2, p3, p4}

	fmt.Println("Before:", people)
	sort.Sort(ByAge(people))
	fmt.Println("After:", people)
	sort.Sort(ByName(people))
	fmt.Println("After Name:", people)

	//s := []int{5, 2, 6, 3, 1, 4} // unsorted
	//sort.Ints(s)
	//sort.Sort(sort.Reverse(sort.IntSlice(s)))
	//fmt.Println(s)

}

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].age < a[j].age }

func (bn ByName) Len() int           { return len(bn) }
func (bn ByName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn ByName) Less(i, j int) bool { return bn[i].first < bn[j].first }
