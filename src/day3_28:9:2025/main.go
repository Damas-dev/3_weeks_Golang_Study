package main

import "fmt"

//* if else qn
//? Write a Go program that evaluates a student’s overall performance based on marks in 3 subjects: Math, Science, and English.

//? Requirements:

//? Create a function average(m1, m2, m3 int) float64 that calculates and returns the average marks.

//? In main(), read the three subject marks from the user.

//? Use if / else if / else with the average() function inside the condition to classify the student:

//? Distinction if average ≥ 75

//? First Class if average ≥ 60

//? Pass if average ≥ 40

//? Fail otherwise

//? Additionally, add one more condition: if the student fails any single subject (mark < 40), then print "Fail (because of one subject)" regardless of average.

func average(m1, m2, m3 int) float32 {
	return float32(m1+m2+m3) / 3
}

func main() {
	fmt.Println("this is day3")

	//* if else solution
	var m1, m2, m3 int
	fmt.Print("please enter you marks in Math(%): ")
	fmt.Scan(&m1)
	fmt.Print("Science(%): ")
	fmt.Scan(&m2)
	fmt.Print("English(%): ")
	fmt.Scan(&m3)

	// calculating average
	avg := average(m1, m2, m3)
	fmt.Println("Average:", avg)

	// if subject < 35 = Status: Fail
	if m1 < 35 {
		fmt.Println("Status: Fail(Math < 35)")
		return
	} else if m2 < 35 {
		fmt.Println("Status: Fail(Science < 35)")
		return
	} else if m3 < 35 {
		fmt.Println("Status: Fail(English < 35)")
		return
	}
	// cheaking average to compute status
	if avg >= 75 && avg <= 100 {
		fmt.Println("Stutus: Distinction")
	} else if avg >= 60 {
		fmt.Println("Status: First Class")
	} else if avg >= 40 {
		fmt.Println("Status: Pass")
	} else {
		fmt.Println("Status: fail(average < 40)")
	}

	//^ switch

	//* basic case
	var subject string = "Math"
	switch subject {
	case "Math":
		fmt.Println("numbers")
	case "English":
		fmt.Println("language")

	case "Geo":
		fmt.Println("maps")
	default:
		fmt.Println("not Math,English or Geo")
	}

	//* coditional case

	switch age := 18; {
	case age > 18:
		fmt.Println("old enogh to vote")
		fallthrough //* fallthrough
	case age < 18:
		fmt.Println("you can't vote")
	case age == 18:
		fmt.Println("you can vote")
	default:
		fmt.Println("pls enter you age in years")
	}

	//* case list

	var number int
	fmt.Println("Positive interger is a whole number from 0 to infinity")
	fmt.Print("to know the type, please enter any a positive interger  from 0 to 15: ")
	fmt.Scan(&number)

	switch {
	case number%2 == 0:
		fmt.Println("your number is Even")

	case number%2 != 0:
		fmt.Println("your number is Odd")

	}
	switch number {
	case 2, 3, 5, 7, 11, 13:
		fmt.Println("your number is Prime")
	default:
		fmt.Println("!but it is not a positive interger from 0 - 15")

	}

}
