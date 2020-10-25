package main

import "oop/employee"

func main() {
	e := employee.Employee{
		FirstName:   "Sam",
		LastName:    "Sharma",
		TotalLeaves: 30,
		LeavesTaken: 20,
	}
	e.LeavesRemaining()
}
