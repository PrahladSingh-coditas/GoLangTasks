package main

import "fmt"

type myFunc func(p Person) //Passing p of Person to access age  and name in func

// Person Class
type Person struct {
	name        string
	age         int
	introduceMe myFunc
}

func checkVotingStatus(p Person) {
	if p.age >= 18 {
		fmt.Println("You are eligible to vote.")
	} else {
		fmt.Println("You are not elegible to vote")
	}
}

func main2() {

	introduceMyself := func(p Person) {
		fmt.Printf("Hi there, my name is %v and my age is %v\n", p.name, p.age)
	}

	//Creating an object/instance of struct
	p1 := Person{
		name:        "prahlad",
		age:         15,
		introduceMe: introduceMyself,
	}
	//Introduce person
	p1.introduceMe(p1)

	//check his age criteria for voting
	checkVotingStatus(p1)

	p1.age = 21

	//Updated age voting status
	checkVotingStatus(p1)

}
