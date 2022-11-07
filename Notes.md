# Notes

## Go File Anatomy

- every file in go MUST have a package name

```go
package example
```

or

```go
package main
```

The above is the main package name. This is required for any go program to run.

- import statements:
    - import statements

```go
import "fmt"
```

alternatively you can use the following for multiple imports.

```go
import (
	"fmt"
	"errors"
)
```

## Access Go Docs in Terminal

- you can access go docs in the terminal by typing "go doc" and the thing you would like to lookup

example:

```bash
$ go doc fmt
```

## Formatting Code

```bash
$ go fmt .
```

Or you can specify a file

```bash
$ go fmt pat/to/file
```

## Variables

- The "var" statement declares a list of variables.
- If all variables share the same type you can include the type at the end

ex:

```go
var num1, num2, num3 int
```

- A var declaration can include initializers, one per variable
- If an initializer is present, the type can be omitted.

```go
// With Type
var a, b int = 1, 2

// Without Type
var
```

- Go has a shorthand method of declaring variables using ":=" and omitting the "var" keyword.
- IMPORTANT: This short hand style can only be used within a function

```go
func main() {
	name := "JR Frazier"
	fmt.Println(name)
}
```

- Variables can be initialized with out defining a value.  When that is done Go will assign that variable is zero most value.
- When doing this a type must be given to the variable

```go
var name string
var num int
```

## What's the difference between "Const and Var"?

## Why would I use one over the other?

## Basic Types in Go

```go
// Boolean "bool"
var foo bool = True
var bar bool = False

// String "string"
var baz string = "This is a string"

// Numbers "int int int8 int16 int32 int64"
var num int = 1
```

## Control Structures

# !NEED TO FINISH!

- Control Structures consist of *if*, *for loops*, and *switch* statements

### IF Statement:

-

## Functions

---

- Functions can take zero or more arguments
- Parameters must be defined with their expected data type

```go
// here you can see that both parameters are required to be numbers
func add(x int, y int) int {
	return x + y
}

// if both of the parameters are the same type you can omit the type for the
// first paramater
func add(x, y int) int {
	return x + y
}
```

- Just like TypeScript the return type must also be defined between the parenthesis and the curly brackets of the function body
- A function can also return multiple results

```go
func printMulti(name string, age int) (string, int) {
	return name, age
}
```

- You can also instantiate your return variable in the return return statement

```go
// As you can see below there "personName" and "personAge" have been defined in
// the return statement. Therefore you do not have to use the keyword "var" or the ":="
// to assigne values to "personName" and "personAge"
func printMulti(name string, age int) (personName string, personAge int) {
	personName = name
	personAge = age
	// since the return variables have been defined in the return statement, you
	// do not have to define them in your return
	return
} 
```

### Exporting Functions

- In Go when you want you export a function you canalize the first letter of the func

```go
func Add(nums ...int) int {
	total := 0
	for _, v := range nums {
		printNum(v)
		total += v
	}
	return total
} 
```

### Variadic Functions

- A function that can take an infinite number of arguments of a specific type.

```go
// This will return an arry of numbers that are provided in the argument
func printMulti(age ...int) (ages []int) {
	ages = age
	return
}

printMulti(20, 21) // [20, 21]
```

### Variadic Parameter

- On the other hand the Parameter can be "Variadic" meaning it can be of any number and *type.*

```go
func printMulti(age ...interface{}) (ages []int) {
	ages = age
	return
}
 
```

### Anonymous Function

- Functions that have no name

Example:

```go
package main

import (
	"fmt"
)

func main() {
	// Below is an anon func
	func(x int) {
		fmt.Println(x * 2)
	}(2)
}

// This will evaluate to 4 whan ran
```

### Function Expressions

- when you assign a function to a variable

Example

```go
package main

import "fmt"

func main() {
  // Below is a func expression
	twoTimes := func(x int) {
		fmt.Println(x * 2)
	}
	twoTimes(2)
}
// This will eval to 4
```

## Slices and Arrays

---

- All values within the array must be of the same type

```go
// Here is an example of an instantiation of an arrya with 5 intigers
var arr1 [5]int // [0,0,0,0,0]

// You can instantiate an array with values like this
var arr2 = [5]int{1, 2, 3 ,4, 5}
var arr3 = [...]int{1, 2, 3, 4, 5}

// or withouth defining the number of elementes in the array like this
var arr4 = []int{1, 2, 3, 4, 5}
```

Great Explanation of Slices [https://blog.golang.org/slices-intro](https://blog.golang.org/slices-intro)

# Bit Shifting in Go

Bit shifting is when you shift binary digits to the left or right. We can use bit shifting, along with iota, to build some creative constants.

- Article on the subject - [https://t.co/2146j97dQZ?amp=1](https://t.co/2146j97dQZ?amp=1)

# Structs

---

## Anonymous Struct

```go
package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last string
		age int
	}{
		first: "JR",
		last: "Frazier",
		age: 40,
	}

	fmt.Println(p1)
}
```

[Notes From Course](https://www.notion.so/Notes-From-Course-22f4b43bfa2a49c880b1511715930166)

[Articles](https://www.notion.so/Articles-d19dd1d08d174d3082f70d447e0132be)

[Go Templates](https://www.notion.so/Go-Templates-7f7aae7becde4039b896a7309188f881)