//--Summary:
//  The existing program is used to restrict access to a resource
//  based on day of the week and user role. Currently, the program
//  allows any user to access the resource. Implement the functionality
//  needed to grant resource access using any combination of `if`, `else if`,
//  and `else`.
//
//--Requirements:
//? Use the accessGranted() and accessDenied() functions to display informational messages
//? Access at any time: Admin, Manager
//? Access weekends: Contractor
//? Access weekdays: Member
//? Access Mondays, Wednesdays, and Fridays: Guest

package main

import "fmt"

// Days of the week
const (
	Monday    = 0
	Tuesday   = 1
	Wednesday = 2
	Thursday  = 3
	Friday    = 4
	Saturday  = 5
	Sunday    = 6
)

// User roles
const (
	Admin      = 10
	Manager    = 20
	Contractor = 30
	Member     = 40
	Guest      = 50
)

func accessGranted() {
	fmt.Println("Access Granted")
}

func accessDenied() {
	fmt.Println("Access Denied")
}

func isWeakDay(day int) bool {
	if day <= Friday {
		return true
	}
	return false

}

func main() {
	// The day and role. Change these to check your work.
	today, role := Friday, Member

	//! if role <= 20 {
	//! 	accessGranted()
	//! } else if role == 30 && today >= 5 {
	//! 	accessGranted()
	//! } else if role == 40 && today <= 4 {
	//! 	accessGranted()
	//! } else if role == 50 && (today == 0 || today == 2 || today == 4) {
	//! 	accessGranted()
	//! } else {
	//! 	accessDenied()
	//! }

	//* collection no magic numbers

	if role == Admin || role == Manager {
		accessGranted()
	} else if role == Contractor && !isWeakDay(today) {
		accessGranted()
	} else if role == Member && isWeakDay(today) {
		accessGranted()
	} else if role == Guest && (today == Monday || today == Wednesday || today == Friday) {
		accessGranted()
	} else {
		accessDenied()
	}

	//! i will retter implove this by adding passwords and user imput

	//?Question: Employee Salary & Bonus Evaluator

	//? Write a Go program that evaluates an employee’s net salary based on their performance and years of experience.

	//? Requirements:

	//? Create a function bonus(baseSalary float64, years int) float64 that returns a bonus amount based on experience:

	//? If years ≥ 10 → 20% of base salary

	//? If years ≥ 5 → 10% of base salary

	//? Otherwise → 5% of base salary

	//? In main(), let the user input:

	//? Employee name

	//? Base salary

	//? Years of experience

	//? Performance rating (Excellent, Good, Average, Poor)

	//? Calculate net salary as:

	//? net = baseSalary + bonus(baseSalary, years)

	//? Apply performance adjustment with if / else if:

	//? Excellent → +10% extra on net

	//? Good → +5% extra on net

	//? Average → no change

	//? Poor → -5% penalty

	//? Print:

	//? Employee name

	//? Base Salary

	//? Bonus

	//? Net Salary after performance adjustment

	//? Final Status:

	//? "Outstanding" if net ≥ 1,000,000

	//? "Well Settled" if net ≥ 500,000

	//? "Needs Growth" otherwise

	//? 🔹 Example Run
	//? Enter Employee Name: Alice
	//? Enter Base Salary: 800000
	//? Enter Years of Experience: 12
	//? Enter Performance (Excellent/Good/Average/Poor): Excellent

	// ? Employee: Alice
	// ? Base Salary: 800000
	// ? Bonus: 160000
	// ? Net Salary after Performance: 1056000
	// ? Final Status: Outstanding
}
