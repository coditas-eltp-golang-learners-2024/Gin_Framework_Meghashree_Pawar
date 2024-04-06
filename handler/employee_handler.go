// handler.go
package handler

import (
    "github.com/gin-gonic/gin"
    "API/model"
    "API/service"
	"net/http"
)

func GetEmployeesHandler(c *gin.Context) {
    // Retrieve all employees
    employees, err := service.GetAllEmployees()
    if err != nil {
        // Handle error
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    // Respond with employees as JSON
    c.JSON(http.StatusOK, employees)
}

func CreateEmployeeHandler(c *gin.Context) {
    // Parse request body
    var employee model.Employee
    if err := c.BindJSON(&employee); err != nil {
        // Handle error
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    // Create employee
    if err := service.CreateEmployee(employee); err != nil {
        // Handle error
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    // Send success response
    c.Status(http.StatusCreated)
}
