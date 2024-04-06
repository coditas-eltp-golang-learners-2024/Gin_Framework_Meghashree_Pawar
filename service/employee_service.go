package service

import (
    "API/model"
    "API/config"
)

// GetAllEmployees retrieves all employees from the database.
func GetAllEmployees() ([]model.Employee, error) {
    // Query to retrieve all employees from the database
    rows, err := config.DB.Query("SELECT id, first_name, last_name, age, salary FROM employees")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    // Define a slice to store employees
    var employees []model.Employee

    // Iterate over the rows and scan each employee
    for rows.Next() {
        var emp model.Employee
        err := rows.Scan(&emp.ID, &emp.FirstName, &emp.LastName, &emp.Age, &emp.Salary)
        if err != nil {
            return nil, err
        }
        employees = append(employees, emp)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return employees, nil
}


func CreateEmployee(emp model.Employee) error {
   
    stmt, err := config.DB.Prepare("INSERT INTO employees(first_name, last_name, age, salary) VALUES(?, ?, ?, ?)")
    if err != nil {
        return err
    }
    defer stmt.Close()

    // Execute the prepared statement to insert employee
    _, err = stmt.Exec(emp.FirstName, emp.LastName, emp.Age, emp.Salary)
    if err != nil {
        return err
    }

    return nil
}
