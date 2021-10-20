package mysql

import (
	"EmployeemanagementSystem/models"
	"fmt"
	"strconv"
)

func Db_DeleteDetails(employee *models.Employee) error{
	var abcd error
	db := Sqlclient()
	query := "UPDATE employeeOfficial SET jobType='"+employee.JobType+"' where userid='"+employee.UserId +"'"
	querychk := "select * from employeeOfficial where userid ='" +employee.UserId + "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		panic(err)
	}
	defer resultchk.Close()
	var emp models.Employee
	healthInsurance, err := strconv.Atoi(strconv.FormatBool(emp.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(emp.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}

	for resultchk.Next() {
		err := resultchk.Scan(&emp.Salary, &emp.Designation, &emp.Department, &emp.TeamLead, &emp.JobType, &healthInsurance, &lifeInsurance, &emp.UserId)
		if err != nil {
			panic(err.Error())
		}
	}
	if employee.UserId == emp.UserId {
		result, err := db.Exec(query)
		if err != nil {
			panic(err)
		}
		defer db.Close()
		fmt.Println(result)
	} else {
		var x int
		err := fmt.Errorf("math: square root of negative number %v", x)
		abcd = err
	}
	return abcd
}
