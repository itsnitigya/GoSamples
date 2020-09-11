package main

import (
    "errors"
	"fmt"
	"time"
)

func add(a, b int) int { // Basic function definition
    return a + b 
}

func multiple_return() (int, int) { // Multiple returns in a function
    return 3, 7
}

func sum(nums ...int) { // ... Vardiac function that can have any trailing number of arguments
    fmt.Print(nums, " ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func fact(n int) int { // Recursive function
    if n == 0 {
        return 1
    }
    return n * fact(n-1)
}

func zeroptr(iptr *int) { // Pointer of the variable
    *iptr = 0
}

type person struct { // Struct
    name string
    age  int
}


type rect struct {
    width, height int
}

func (r *rect) area() int { // This area method has a receiver type of *rect
    return r.width * r.height
}

type geometry interface { // Interfaces are named collections of method signatures.
    area() float64
    perim() float64
}

type circle struct {
    radius float64
}

func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
    return 2 * math.Pi * c.radius
}

func measure(g geometry) {
    fmt.Println(g)
    fmt.Println(g.area())
    fmt.Println(g.perim())
}


func f1(arg int) (int, error) { // Error handling
    if arg == 42 {
        return -1, errors.New("can't work with 42")
    }
    return arg + 3, nil
}

func main() {
	fmt.Println("hello world") // Printing
	
	var b, c int = 1, 2 // Initialising the variables
	fmt.Println(b, c)
	
	b := 1 // Short hand declaration

	const n = 500000000 // Declaring constants

	for i := 0  j <= 10; i++ { // For loop
		if i % 2 == 0 {
			continue
		}
        fmt.Println(i)
	}

	if num := 9; num < 0 {   // If else mulitple statements
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
	}
	
	switch time.Now().Weekday() { // Switch statements
		case time.Saturday, time.Sunday:
			fmt.Println("It's the weekend")
		default:
			fmt.Println("It's a weekday")
    }

	var a [5]int
	fmt.Println("empty array:", a)

	b := [5]int{1, 2, 3, 4, 5} // Declare and initialize an array in one line
	t := []string{"g", "h", "i"}

	s := make([]string, 3)
	fmt.Println("empty string:", s)

	s = append(s, "d") // Appending a string 

	c := make([]string, len(s))
	copy(c, s)	// Copying a string

	slice := s[2:5]
	fmt.Println("slice of a string:", slice)

	twoD := make([][]int, 3) // 2D array
    for i := 0; i < 3; i++ { 
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
	fmt.Println("2d: ", twoD) 
	
	m := make(map[string]int) // Unordered map
	m["k1"] = 7

	v1 := m["k1"]
	fmt.Println("v1: ", v1)
	delete(m, "k2")

	for k, v := range m { // Range operator
        fmt.Printf("%s -> %s\n", k, v)
	}
	
	res := add(1, 2) // Use of functions
	fmt.Println("1+2 =", res)
	
	a, b := multiple_return() // Multiple returns
	fmt.Println(a, b)

	sum(1, 2) // Same function, multiple arguments
    sum(1, 2, 3)

	i := 1
	zeroptr(&i) // Passing by references

	s := person{name: "Sean", age: 50}  // Defining a struct

	r := rect{width: 10, height: 5}
	fmt.Println("area: ", r.area()) // Calling the method on the struct instance 
	
	c := circle{radius: 5}
	measure(c) // Interface
	
}
