package service

import (
	"EmployeemanagementSystem/database/mysql"
	"EmployeemanagementSystem/models"
)

func Service_DeleteDetails(employee *models.Employee) error {
	return  mysql.Db_DeleteDetails(employee)
}
