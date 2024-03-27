package main

import "fmt"

type myFunc func(*DB)

type Department struct {
	departName      string
	employeeList    DB
	calculateAvgsal myFunc
}

func calculateAvgsal(db *DB) {
	total := 0.0
	count := 0
	for index, _ := range db.empDetails {
		total += float64(db.empDetails[index].salary)
		count++
	}
	fmt.Println("Average salary: ", total/float64(count))
}

type Employee struct {
	name   string
	age    int
	salary float32
}

type DB struct {
	empDetails []Employee
}

func (e Employee) save(db *DB) {
	db.empDetails = append(db.empDetails, e)

}

func (e *DB) removeEmployee(name string) {
	for index, emp := range e.empDetails {
		if emp.name == name {
			e.empDetails = append(e.empDetails[0:index], e.empDetails[index+1:]...)
		}
	}
}

func (e *DB) raiseSalary(name string, salary float32) {
	for index, _ := range e.empDetails {
		if e.empDetails[index].name == name {
			e.empDetails[index].salary += salary
		}
	}
}

func main() {
	//Creating a database(i.e slice) to store the data of the employee
	db := DB{}

	//Adding values
	emp1 := Employee{name: "Prahlad", age: 21, salary: 40000}
	emp1.save(&db)

	//Adding values
	emp2 := Employee{name: "Rohan", age: 27, salary: 90000}
	emp2.save(&db)
	//fmt.Println(db)

	//Removing the employee
	db.removeEmployee("Prahlad")
	fmt.Println(db)

	//raising the salary by the amount
	db.raiseSalary("Rohan", 5000)
	fmt.Println(db)

	//To calculate the average salary
	department1 := Department{"Golang", db, calculateAvgsal}
	department1.calculateAvgsal(&department1.employeeList)

}
