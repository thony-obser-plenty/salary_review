<?php

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
interface EmployeeInterface {
    public function getId(): int;
    public function getName(): string;
}

// Added salary interface because salary class violated several SOLID-Principles
// -> Single responsibility principle. Salary class did not only calculate but also check what type of employee it was
// -> Open/closed principle. Salary class was not open for extension. If we wanted to add a new employee we would have to change the salary class
interface SalaryInterface {
    public function calculateSalary(): float;
}

// implemented employee interface
// Removed position property because it was redundant
// Added types
class Employee implements EmployeeInterface {
    private int $id;
    private string $name;

    public function __construct($id, $name) {
        $this->id = $id;
        $this->name = $name;
    }

    public function getId(): int{
        return $this->id;
    }

    public function getName(): string {
        return $this->name;
    }
}

// Added manager subclass
// Subclasses implement salary interface
class Manager extends Employee implements SalaryInterface {
    public function __construct($id, $name) {
        parent::__construct($id, $name);
    }

    public function calculateSalary(): float
    {
        return $this->getId() * 5000;
    }
}

// Added developer subclass
// Subclasses implement salary interface
class Developer extends Employee implements SalaryInterface {
    public function __construct($id, $name) {
        parent::__construct($id, $name);
    }

    public function calculateSalary(): float
    {
        return $this->getId() * 3000;
    }
}

// Added intern subclass
// Subclasses implement salary interface
class Intern extends Employee implements SalaryInterface {
    public function __construct($id, $name) {
        parent::__construct($id, $name);
    }

    public function calculateSalary(): float
    {
        return $this->getId() * 1000;
    }
}

// Removed SalaryCalculator class because it was redundant. In a more complex scenario it would be useful to outsource the calculation to a separate class.
$salaries = [];
$salaries[] = (new Manager(1, 'John Doe'))->calculateSalary();
$salaries[] = (new Developer(2, 'Jane Smith'))->calculateSalary();
$salaries[] = (new Intern(3, 'Mike Johnson'))->calculateSalary();

print_r($salaries);
