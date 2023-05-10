package main

import "fmt"

type Employee struct{
	name string
	age int
	job string 
	salary float32
}

var emp1 Employee

func getEmployeeName() string{
	return emp1.name
}

func setEmployeeName(name string){
	emp1.name = name
}

func getEmployeeAge() int{
	return emp1.age;
}

func setEmployeeAge(age int){
	emp1.age = age
}

func getEmployeeJob() string{
	return emp1.job
}

func setEmployeeJob(job string){
	emp1.job = job
}

func getEmployeeSalary() float32{
	return emp1.salary
}

func setEmployeeSalary(salary float32){
	emp1.salary = salary
}

func main(){
	var empName string
	var empAge int
	var empJob string
	var empSalary float32
	fmt.Println("enter employee name: ")
	fmt.Scanf("%s",&empName)
	fmt.Println("enter employee age: ")
	fmt.Scanf("%d",&empAge)
	fmt.Println("enter employee job: ")
	fmt.Scanf("%s",&empJob)
	fmt.Println("enter employee salary: ")
	fmt.Scanf("%g",&empSalary)

	setEmployeeName(empName)
	setEmployeeAge(empAge)
	setEmployeeJob(empJob)
	setEmployeeSalary(empSalary)

	fmt.Println("Employee Name: ", getEmployeeName())
	fmt.Println("Employee Age: %d", getEmployeeAge())
	fmt.Println("Employee Job: ", getEmployeeJob())
	fmt.Println("Employee Salary: %g", getEmployeeSalary())
}
