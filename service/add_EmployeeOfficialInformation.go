package service

import (
	"EmployeemanagementSystem/database/mysql"
	"EmployeemanagementSystem/models"
)

func Service_AddDetails(employee *models.Employee) error{
	return  mysql.Db_addDetails(employee)
}
