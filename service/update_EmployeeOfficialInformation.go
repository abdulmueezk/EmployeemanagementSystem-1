package service

import (
	"EmployeemanagementSystem/database/mysql"
	"EmployeemanagementSystem/models"
)

func Service_UpdateDetails(employee *models.Employee) error {
	return mysql.Db_UpdateDetails(employee)
}
