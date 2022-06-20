package main
// the above line means compiler will create executable file for us

// fmt is for formatting strings and printing messages to a console
import "fmt"

// all functions start with keyword "func" and have parens after name
// only one main function in entire APPLICATION, not one per file
func main() {

	// STRINGS (must be double quotes)
	// these first two lines do the same thing, the first one just explicitly types it
	var name1 string = "Hello"
	var name2 = "Hello"
	// we also don't need to give a string a value, default is empty string
	var name3 string

	fmt.Println(name1, name2, name3)

	// change variables like this 
	name1 = "cheese"

	// shorthand for string declaration, ':=' is used when initially declaring a variable
	// cannot use ':=' outside of a function
	name4 := "world"
	fmt.Println(name4)

	// INTEGERS
	var age1 int = 20
	var age2 int = 30
	age3 := 40

	fmt.Println(age1, age2, age3)
	
	// BITS & MEMORY
	// golang.org/pkg/builtin
	var num1 int8 = 25
	var num2 int16 = -129
	var num3 uint8 = 25

	fmt.Println(num1, num2, num3)

	// FLOATS
	var score1 float32 = -15.3535
	var score2 float64 = 2343223432432432423423423432423423423423423.2342
	// default type is float64
	score3 := 1.5
	
	fmt.Println(score1, score2, score3)
	
}