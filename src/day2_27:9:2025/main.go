package main

import "fmt"

func main() {
	fmt.Println("this is Day two")

	//^ Variables
	//* variables creation types: single,Compound,block,create&assign

	//& Single variable creation
	var age = 49 //! type annotation is opption
	fmt.Println("my age is", age)

	var name string = "damas" //! once annotated you can't change the value type on reassignment
	fmt.Println("my name is", name)
	//~ example:  name = 45 ----> this is an error

	myName := "damas" //! create&assign type
	fmt.Println("my name is", myName)

	//& compound variable cretion(decaletion)
	var a, b, c = 1, 2, "dajo"
	fmt.Println("a:", a, "b:", b, "c:", c) //= a:1 b:2 c: dajo

	z, c := 4, "jose" //! c variable is been reassigned to a new string NOTE: you can reassign it to another type apart from

	//~ exmple: h , c := 3 ,4  //! c:  initial type is string, in this case it is string so you can't reassign it to int(it won't work)

	fmt.Println("z:", z, " c:", c)

	//& block variable creation

	var (
		r        = 1
		t        = 2
		y string = "GoLang"
	)

	fmt.Println(r, t, y)

	//* variables can be assigned to other variables
	o := 4
	o = 5
	n := o
	o = 67               //!once a change the value of o
	fmt.Println("n:", n) //=5
	fmt.Println("o:", o) //=67

	//* variables that are declared, but not assigned will have a default value
	/* default values
	~strings: ""
	~int:0
	~others:nil
	*/

	//* comma,ok
	g := 5
	//! g := 3 --> this is not okay, you can't declar a variable two time

	m, v := 6, 7
	x, v := 4, 4 //! this is okay, since your using comma, it is like your reassinging the variabale "v"

	fmt.Println("g:", g, "m:", m, "x:", x, "v:", v)

	//* constants (const)

	const MaxSpeed = 43.5

	//* naming constants(UpperCamelCase) and variables& function(lowerCamelCase)

	//^Functions

	/* funcion syntax
	~ func name(param1 type, param2 type , param..n type) return types {
	~
	~ fanc body
	~
	~
	~ }

	~ func name() (type1,type2,type3){
	~
	~ fanc body
	~
	~
	~ }
	*/

	//~ examples

}
