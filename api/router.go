package api

import (
	_"database/sql"
	_"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/employees", GetEmployees)
	r.GET("/employees/:id", GetEmployee)
	r.POST("/employees", CreateEmployee)
	r.PUT("/employees/:id", UpdateEmployee)
	r.DELETE("/employees/:id", DeleteEmployee)
}


