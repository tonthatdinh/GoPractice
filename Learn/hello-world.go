package main

import (
	"errors"
	"fmt"
)

func main(){
	fmt.Println("Hello World");
	basic();
	fmt.Println(add(1, 2));
	fmt.Println(divide(3, 4));
	loops()
	slices()
	maps()
	squarenumber() // Ex1: Print the first 10 squared numbers
	fibonaccinumber() //Ex2: Print the first 10 fibonacci numbers
}

func fibonaccinumber() {
	fmt.Println("The first 10 Fibonacci numbers")
	prev2 := 1
	prev1 := 1
	for i:= 1; i < 10; i ++{
		current := prev2 + prev1;
		fmt.Println(current);
		prev1 = prev2
		prev2 = current
	}
}


func squarenumber() {
	fmt.Println("The first 10 square numbers")
	for i:= 1; i < 10; i ++{
		fmt.Printf("%v ", i * i);
	}
	fmt.Println("---------")
}

func maps() {
	myMap := make(map[string]int)
	myMap["yellow"] = 1
	myMap["magic"] = 2
	myMap["amsterdam"] = 3
	fmt.Println(myMap)
	myMap["magic"] = 100
	delete(myMap, "amsterdam")
	fmt.Println(myMap)
	for index, value := range myMap{
		fmt.Printf("%v: %v\n", index, value)
	}
	fmt.Print("Map size = %v\n", len(myMap))
}

func slices() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice);
	fmt.Println(slice[2: 5]) // 3, 4, 5
	fmt.Println(slice[5:]) // 6, 7, 8
	fmt.Println(slice[:3]) // 1, 2, 3
	slice2 := make([]string, 3);
	slice2[0] = "tic"
	slice2[1] = "tac"
	slice2[2] = "toe"
	fmt.Print(slice2);
	slice2 = append(slice2, "tom");
	slice2 = append(slice2, "radar");
	fmt.Println(slice2)
	for index, value := range slice2{
		fmt.Printf("%v: %v\n", index, value);
	}
	fmt.Printf("Slice length = %v\n", len(slice2));
}


func loops() {
	//For loop
	for i:= 0; i < 10; i ++{
		fmt.Print(".");
	}
	//While loop
	sum := 1
	for sum < 1000{
		sum *= 2;
	}
	fmt.Printf("The sum is %v\n", sum);
}


// Function that returns two things; error is nil if sucessful
func divide(x int, y int) (float64, error) {
	if y == 0{
		return 0.0, errors.New("Divide by zero")
	}
	//Cast x and y to float64 before dividing
	return float64(x)/float64(y), nil
}


// Function declaration; takes in 2 ints and outputs an int
func add(x int, y int) int {
	return x + y;
}


func basic() {
	//Declare x as a variable, intialized to 0
	var x int;
	//Decalare y as a variable, intialized to 2
	var y int = 2;
	//Decalre z as a variable, initialized to 4
	//This syntax can only be used in a function
	z := 4;
	// Assign values to variables
	x = 1;
	y = 2;
	z = x + 2 * y + 3;
	// Print the variables, just use %v for most types
	fmt.Printf("x = %v, y = %v, z = %v\n", x, y, z)
}
