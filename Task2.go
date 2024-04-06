
// task 2:- Develop a system to manage employees and their departments. 
// Each employee has a name, age, and salary. Each department has a name, a list of employees,
// and a method to calculate the average salary of its employees. Additionally, create methods to 
// add and remove employees from departments and to give a raise to an employee.

package main

import "fmt"

type Employee struct {
    Name   string
    Age    int
    Salary float64
}

type Department struct {
    Name      string
    Employees []Employee 
}

func (d *Department) AddEmployee(emp Employee) {
    d.Employees = append(d.Employees, emp)
}

func (d *Department) RemoveEmployee(emp Employee) {
    for i, e := range d.Employees {
        if e == emp {
            d.Employees = append(d.Employees[:i], d.Employees[i+1:]...)
            break
        }
    }
}

func (d *Department) AverageSalary() float64 {
    if len(d.Employees) == 0 {
        return 0.0
    }
    totalSalary := 0.0
    for _, emp := range d.Employees {
        totalSalary += emp.Salary
    }
    return totalSalary / float64(len(d.Employees))
}

func (emp *Employee) GiveRaise(raiseAmount float64) {
    emp.Salary += raiseAmount
}

func mainn() {
    emp1 := Employee{Name: "Megha", Age: 30, Salary: 50000}
    emp2 := Employee{Name: "Jayesh", Age: 35, Salary: 60000}
    dept := Department{Name: "Data Science"}
    dept.AddEmployee(emp1)
    dept.AddEmployee(emp2)
    fmt.Printf("Average Salary of %s Department: $%.2f\n", dept.Name, dept.AverageSalary())
    emp1.GiveRaise(5000)
    fmt.Printf("%s's updated salary: $%.2f\n", emp1.Name, emp1.Salary)
    dept.RemoveEmployee(emp2)
    fmt.Printf("Updated Average Salary of %s Department: $%.2f\n", dept.Name, dept.AverageSalary())
}
