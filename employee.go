package employee

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
