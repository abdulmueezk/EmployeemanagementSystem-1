package service

import (
	"EmployeemanagementSystem/database/mysql"
	"EmployeemanagementSystem/models"
)

func Service_ShowDetails(employee *models.Employee) error {
	return mysql.Db_ShowDetails(employee)
}
