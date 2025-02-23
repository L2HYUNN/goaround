/*
Functions
A function can take zero or more arguments.

In this example, add takes two parameters of type int.

Notice that the type comes after the variable name.

(For more about why types look the way they do, see the article on Go's declaration syntax.)
*/
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

/*
Functions continued
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we shortened

x int, y int
to

x, y int
*/

// shortcut
// func add(x, y int) int {
// 	return x + y;
// }

// func main() {
// 	fmt.Println(add(42, 13))
// }

/*
Multiple results
A function can return any number of results.

The swap function returns two strings.
*/
func swap(x, y string) (string, string) {
	return y, x
}

// func main() {
// 	a, b := swap("hello", "world")
// 	fmt.Println(a, b)
// }

/*
Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A return statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here.

They can harm readability in longer functions.
*/

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(split(17))
}
