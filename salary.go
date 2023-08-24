package main

import "fmt"

// Employee represents an employee in a company
type Employee struct {
	ID       int
	Name     string
	Position string
}

// SalaryCalculator calculates the salary of employees
type SalaryCalculator struct {
	Employees []Employee
}

// CalculateSalary calculates the salary for all employees
func (sc *SalaryCalculator) CalculateSalary() []int {
	salaries := make([]int, len(sc.Employees))
	for i, emp := range sc.Employees {
		if emp.Position == "Manager" {
			salaries[i] = sc.calculateManagerSalary(emp)
		} else if emp.Position == "Developer" {
			salaries[i] = sc.calculateDeveloperSalary(emp)
		} else {
			salaries[i] = sc.calculateInternSalary(emp)
		}
	}
	return salaries
}

func (sc *SalaryCalculator) calculateManagerSalary(emp Employee) int {
	// Complex calculation for manager salary
	return emp.ID * 5000
}

func (sc *SalaryCalculator) calculateDeveloperSalary(emp Employee) int {
	// Complex calculation for developer salary
	return emp.ID * 3000
}

func (sc *SalaryCalculator) calculateInternSalary(emp Employee) int {
	// Complex calculation for intern salary
	return emp.ID * 1000
}

func main() {
	employee1 := Employee{ID: 1, Name: "John Doe", Position: "Manager"}
	employee2 := Employee{ID: 2, Name: "Jane Smith", Position: "Developer"}
	employee3 := Employee{ID: 3, Name: "Mike Johnson", Position: "Intern"}

	salaryCalculator := &SalaryCalculator{Employees: []Employee{employee1, employee2, employee3}}

	salaries := salaryCalculator.CalculateSalary()
	fmt.Println(salaries)
}
