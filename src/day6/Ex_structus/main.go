//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing its coordinates
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

func main() {
	fmt.Println("this is Day six")

	//^ Rectangle Area and Perimeter

	//* structure
	type Rectangle struct {
		x1, y1 float64
		x2, y2 float64
	}

	//* functions
	area := func(r Rectangle) float64 {
		length := r.x2 - r.x1
		width := r.y2 - r.y1
		return length * width
	}

	perimeter := func(r Rectangle) float64 {
		length := r.x2 - r.x1
		width := r.y2 - r.y1
		return 2 * (length + width)
	}

	doubleSize := func(r *Rectangle) {
		r.x2 = r.x1 + 2*(r.x2-r.x1)
		r.y2 = r.y1 + 2*(r.y2-r.y1)
	}

	//* rectangle instance
	rect := Rectangle{x1: 0, y1: 0, x2: 10, y2: 5}

	fmt.Println("\nOriginal Rectangle:")
	fmt.Println("Area:", area(rect))
	fmt.Println("Perimeter:", perimeter(rect))

	//* double size
	doubleSize(&rect)

	fmt.Println("\nAfter Doubling Size:")
	fmt.Println("Area:", area(rect))
	fmt.Println("Perimeter:", perimeter(rect))
}
