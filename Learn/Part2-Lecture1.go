package main

import "fmt"
// Object oriented programming
// Convention: capitalize first letter of public fields
type Shark struct {
	Name string
	Age int
}

// Declare a public method
// This is called a receiver method
func (s *Shark) Bite(){
	fmt.Printf("%v says CHOMP!\n", s.Name)
}

// Because function in Go are pass by value//

func (s *Shark) ChangeName(newName string) {
	s.Name = newName
}

func main(){
	fmt.Println("Part 2: Object Oriented Programming");
	shark1 := Shark{"Dinh", 22};
	shark2 := Shark{"Tam", 21};
	fmt.Println(shark1.Name)
	fmt.Print(shark1.Age)
	shark1.Greet(&shark2)
}

// Receiver methods can take in other objects as well
func (s *Shark) Greet(s2 *Shark){
	if (s.Age < s2.Age){
		fmt.Printf("%v says your majesty\n", s.Name)
	} else {
		fmt.Printf("%v says yo what's up %v\n", s.Name, s2.Name)
	}
}

func sharks(){
	shark1 := Shark{"Bruce", 32}
	shark2 := Shark{"Sharkira", 40}
	shark1.Bite();
	shark1.ChangeName("Lee")
	shark1.Greet(&shark2); //pass in pointer
	shark2.Greet(&shark1);
}

//Channels are a way to pass messages across goroutines
func channels()  {
	ch := make(chan int)
	// Launch a gorotine using an anonymous function
	go func() {
		i := 1
		for{
			//This line block until someone consumes from the channel
			ch <- i * i;
			i ++
		}
	}()
	//Extract first 10 squared number from the channel
	for i := 0; i < 10; i ++{
		//This line block until someone sends into the channel
		fmt.Printf("The next squared number is %v\n", <-ch)
	}
}