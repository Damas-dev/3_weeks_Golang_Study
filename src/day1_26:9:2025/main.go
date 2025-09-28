package main

import "fmt"

func main() {
	fmt.Println("hello")

	//^ Data type: string, int, float, bool

	//* integer:signed integers & unsigned integers
	//& signed integers: int(same as int 32) int8(+-2^7),int16(+-2^15),int32,int64
	//& unsigned integers: Uint(same as int 32), int8(+-2^8-1),int16(+-2^16-1),int32,int64,byte(same as int8)

	var name string = "john"
	var age int = 30
	var height float64 = 5.9
	var isStudent bool = true

	// Print variables

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)
	fmt.Println("Height:", height)
	fmt.Println("Is Student:", isStudent)

	//^ Type Aliases: it is simply assigning a new name to an existing type

	//examples of type aleases
	type UserName string
	type UserId byte
	type Speed float64
	type Velocity Speed

	//? how to initialise type aleases
	// UserId(1)
	// UserName("Damas")
	// Velocity(445.75)

	//^ Runes(char in C-lang): 'a' ,'b' '\n',`∫`,`ß`
	//! we use backtick(``) to represent non-alphanumeric characters except '\n' as shown above

}
