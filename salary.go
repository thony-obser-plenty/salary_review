package main

import "fmt"

// I check for:
// Readability
// Understandability
// Performance
// Error handling
// Bugs & Logical flaws
// DRY-Principle -> Don't write redundant code
// SOP-Principle -> Each thing does its own thing
// KISS-Principle -> Try to keep it simple
// SOLID-Principle -> https://en.wikipedia.org/wiki/SOLID
// POLA-Principle -> Code should work as one would expect
// YAGNI-Principle -> Don't code stuff you might not need

// Added employee interface for
// -> Type safety.
// -> Polymorphism and modularity.
type EmployeeInterface interface {
	getId() int
	getName() string
}

// Added salary interface because salary class violated several SOLID-Principles
// -> Single responsibility principle. Salary function did not only calculate but also check what type of employee it was
// -> Open/closed principle. Salary function was not open for extension. If we wanted to add a new employee we would have to change the salary function
type SalaryInterface interface {
	calculateSalary() float64
}

// implemented employee interface
// Removed position property because it was redundant
// Added types
type Employee struct {
	Id   int
	Name string
}

// Added manager subtype
type Manager struct {
	Employee
}

func (manager Manager) getId() int {
	return manager.Id
}

func (manager Manager) getName() string {
	return manager.Name
}

// Subclasses implement salary interface
func (manager Manager) calculateSalary() float64 {
	return float64(manager.Id * 5000)
}

// Added developer subtype
type Developer struct {
	Employee
}

func (developer Developer) getId() int {
	return developer.Id
}

func (developer Developer) getName() string {
	return developer.Name
}

// Subclasses implement salary interface
func (developer Developer) calculateSalary() float64 {
	return float64(developer.Id * 3000)
}

// Added intern subtype
type Intern struct {
	Employee
}

func (intern Intern) getId() int {
	return intern.Id
}

func (intern Intern) getName() string {
	return intern.Name
}

// Subclasses implement salary interface
func (intern Intern) calculateSalary() float64 {
	return float64(intern.Id * 1000)
}

// Removed SalaryCalculator because it was redundant. In a more complex scenario it would be useful to outsource the calculation to a separate class.
func main() {
	var manager EmployeeInterface
	manager = Manager{Employee{Id: 1, Name: "John Doe"}}

	var developer EmployeeInterface
	developer = Developer{Employee{Id: 2, Name: "Jane Smith"}}

	var intern EmployeeInterface
	intern = Intern{Employee{Id: 3, Name: "Mike Johnson"}}

	var salaries []float64

	salaries = append(salaries, manager.(SalaryInterface).calculateSalary())
	salaries = append(salaries, developer.(SalaryInterface).calculateSalary())
	salaries = append(salaries, intern.(SalaryInterface).calculateSalary())

	fmt.Println(salaries)
}
