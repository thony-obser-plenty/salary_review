<?php

// Employee class represents an employee in a company
class Employee {
    private $id;
    private $name;
    private $position;

    public function __construct($id, $name, $position) {
        $this->id = $id;
        $this->name = $name;
        $this->position = $position;
    }

    public function getId() {
        return $this->id;
    }

    public function getName() {
        return $this->name;
    }

    public function getPosition() {
        return $this->position;
    }
}

// SalaryCalculator class calculates the salary of employees
class SalaryCalculator {
    private $employees = [];

    public function addEmployee(Employee $employee) {
        $this->employees[] = $employee;
    }

    public function calculateSalary() {
        $salaries = [];
        foreach ($this->employees as $employee) {
            if ($employee->getPosition() === 'Manager') {
                $salaries[] = $this->calculateManagerSalary($employee);
            } elseif ($employee->getPosition() === 'Developer') {
                $salaries[] = $this->calculateDeveloperSalary($employee);
            } else {
                $salaries[] = $this->calculateInternSalary($employee);
            }
        }
        return $salaries;
    }

    private function calculateManagerSalary($employee) {
        // Complex calculation for manager salary
        return $employee->getId() * 5000;
    }

    private function calculateDeveloperSalary($employee) {
        // Complex calculation for developer salary
        return $employee->getId() * 3000;
    }

    private function calculateInternSalary($employee) {
        // Complex calculation for intern salary
        return $employee->getId() * 1000;
    }
}

$employee1 = new Employee(1, 'John Doe', 'Manager');
$employee2 = new Employee(2, 'Jane Smith', 'Developer');
$employee3 = new Employee(3, 'Mike Johnson', 'Intern');

$calculator = new SalaryCalculator();
$calculator->addEmployee($employee1);
$calculator->addEmployee($employee2);
$calculator->addEmployee($employee3);

$salaries = $calculator->calculateSalary();
print_r($salaries);
