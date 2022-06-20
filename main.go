package main

// the above line means compiler will create executable file for us

// fmt is for formatting strings and printing messages to a console
import (
	"fmt"
	"sort"
	"strings"
)

// FUNCTIONS
func sayGreeting(n string) {
	fmt.Printf("Good morning %v \n", n)
}

// functions can have function param
func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

// functions can have return value located after params
func addTwo(n1 int, n2 int) int {
	return n1 + n2
}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")

	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	} else {
		return initials[0], "_"
	}
}

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

	// printf function used for formatting strings
	fmt.Println("my age is ", age1, " and my name is ", name1)
	// %v is for variables
	fmt.Printf("my age is %v and my name is %v \n", age1, name1)
	// %q puts quotes around strings
	// note this will not work for the "age1" variable as it is an int
	fmt.Printf("my age is %q and my name is %q \n", age1, name1)
	// %T gives type of variable
	fmt.Printf("age is of type %T \n", age1)
	// to output float, use %f
	// use leading decimal point to indicate what to round by
	fmt.Printf("you scored %f points", score1)
	fmt.Printf("you scored %0.1f points", score1)

	// Sprintf (save formatted strings)
	var str = fmt.Sprintf("my age is %v and my name is %v \n", age1, name1)
	fmt.Println("the saved string is: ", str)

	// ARRAYS
	// can only put in up to as many variables as the declared length
	var ages [4]int = [4]int{20, 25, 30}
	// use len function to find array length
	fmt.Println(ages, len(ages))

	// slices (use arrays under the hood)
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 40)

	// slice ranges
	range1 := scores[1:3]
	range2 := scores[2:]
	range1[0] = 3
	range2[0] = 4

	// STANDARD LIB
	// https://golang.org/pkg/
	fmt.Println(strings.Contains(name4, "or"))
	fmt.Println(strings.ReplaceAll(name4, "or", "and"))
	// other strings functions:
	// ToUpper (changes chars to uppercase)
	// Index (finds index of given substring)
	// Split (splits on given substring)

	sort.Ints(scores)
	// other sort functions:
	// SearchInts (searches for given int)
	// Strings (sorts strings)
	// SearchStrings (searches for given strings)
	// etc.

	// FOR LOOPS
	// can be used like while loops in other languages
	x := 0
	for x < 5 {
		x++
	}
	// or the normal way
	for i := 0; i < 5; i++ {
		fmt.Println("wow")
	}

	// using i as index of arrays
	names := []string{"luke", "nick", "dylan"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// for each loop
	for index, value := range names {
		fmt.Printf("the position at index %v is %v", index, value)
	}

	// for each loop without index
	for _, value := range names {
		fmt.Printf("name: %v", value)
	}

	// note: altering values in for each loop doesn't alter values in original slice

	// BOOLEANS
	// same as any other language, no code required

	// FUNCTION CALL
	// multiple return values
	fn1, sn1 := getInitials("luke muratore")
	fmt.Println(fn1, sn1)

	// can call functions in other files
	//sayHello("luke")
	// can access global variables from other files
	/*for _, v := range points {
		fmt.Println(v)
	}
	*/
}
