package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	//contact   contactInfo
	contactInfo
}

type contactInfo struct {
	email   string
	pinCode int
}

func main() {

	// First way to declare and assign value
	myVar1 := person{"Gautam", "Singh", contactInfo{"gs7876@rediffmail.com", 560078}}

	fmt.Printf("Type of myVar1 %T\n", myVar1)
	fmt.Printf("Value of myVar1 %v\n", myVar1)
	fmt.Printf("%+v %+v", myVar1, myVar1.contactInfo)
	fmt.Printf("firtName=%v lastName=%v email=%v pinCode=%v\n\n\n\n", myVar1.firstName, myVar1.lastName, myVar1.contactInfo.email, myVar1.contactInfo.pinCode)

	// Second way to declare and assign value
	var myVar2 person
	fmt.Printf("This is empty struct %+v\n", myVar2)

	myVar2.firstName = "Akshat"
	myVar2.lastName = "Singh"
	myVar2.contactInfo.email = "as7876@rediffmail.com"
	myVar2.contactInfo.pinCode = 560078
	fmt.Printf("Type of myVar2 %T\n", myVar2)
	fmt.Printf("Value of myVar2 %v\n", myVar2)
	fmt.Printf("%+v %+v", myVar2, myVar2.contactInfo)
	fmt.Printf("firtName=%v lastName=%v email=%v pinCode=%v\n\n\n\n", myVar2.firstName, myVar2.lastName, myVar2.contactInfo.email, myVar2.contactInfo.pinCode)

}
