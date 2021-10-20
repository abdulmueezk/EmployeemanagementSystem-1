package mysql

import (
	"EmployeemanagementSystem/models"
	"fmt"
	"strconv"
)

func Db_addDetails(employee *models.Employee) error{
	healthInsurance, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	db := Sqlclient()
	result, err := db.Query("INSERT INTO employeeOfficial (salary,designation,department,teamLead,jobType,healthInsurance,lifeInsurance,userId)VALUES ( '"+employee.Salary+"' ,'"+employee.Designation+"' ,'"+employee.Department+"','"+employee.TeamLead+"','"+employee.JobType+"','"+fmt.Sprintf("%v",healthInsurance)+"','"+fmt.Sprintf("%v",lifeInsurance)+"','"+employee.UserId+"')")
		if err != nil {
			return err
		}
		defer result.Close()
	/*for result.Next() {
		err := result.Scan(&employee.Salary, &employee.Designation, &employee.Department, &employee.TeamLead, &employee.JobType, &employee.HealthInsurance, &employee.LifeInsurance, &employee.UserId)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println(result) */
	return nil
}
