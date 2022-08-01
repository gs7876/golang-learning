//Struct with receiver function
//update the value of struct

package main

import "fmt"

type person1 struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo1
}

type contactInfo1 struct {
	email   string
	pinCode int
}

func main() {

	// First way to declare and assign value
	myVar := person1{"Gautam", "Singh", contactInfo1{"gs7876@rediffmail.com", 560078}}

	myVar.updateName("Akshat")
	myVar.print()

	// Way 1
	//myVarPointer := &myVar
	//myVarPointer.updateNameP("Akshat")

	myVar.updateNameP("Akshat")
	myVar.print()

}

// p person1 is example of receiver function p is giving access to updateName
func (p person1) updateName(newFirstname string) {
	fmt.Println("update from updateName")
	p.firstName = newFirstname
}

//pointerToPerson *person1 is example of receiver function using pointer
//Update using pointer
func (pointerToPerson *person1) updateNameP(newFirstname string) {
	fmt.Println("update from updateNameP")
	(*pointerToPerson).firstName = newFirstname

}

// p person1 is example of receiver function p is giving access to updateName
func (p person1) print() {
	fmt.Printf("%+v\n", p)
}
