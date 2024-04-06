// router.go
package router

import (
    "github.com/gin-gonic/gin"
    "API/handler"
)

func NewRouter() *gin.Engine {
    // Create a new Gin router
    r := gin.Default()

    // Define route for GET request to retrieve employees
    r.GET("/getEmployees", handler.GetEmployeesHandler)

    // Define route for POST request to create an employee
    r.POST("/postEmployees", handler.CreateEmployeeHandler)

    return r
}
