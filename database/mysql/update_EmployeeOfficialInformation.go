package mysql

import (
	"EmployeemanagementSystem/models"
	"fmt"
	"strconv"
)

func Db_UpdateDetails(employee *models.Employee) error{
	healthInsurance, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	lifeInsurance, _ := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	var abcd error
	db := Sqlclient()
	query := "UPDATE employeeOfficial SET salary='"+employee.Salary+"' ,designation='"+employee.Designation+"' ,department='"+employee.Department+"', teamLead='"+employee.TeamLead+"',jobType='"+employee.JobType+"',healthInsurance='"+fmt.Sprintf("%v",healthInsurance)+"',lifeInsurance='"+fmt.Sprintf("%v",lifeInsurance)+"' where userId='"+employee.UserId +"'"
	querychk := "select * from employeeOfficial where userid ='" + employee.UserId + "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		panic(err)
	}
	defer resultchk.Close()
	var emp models.Employee
	healthInsurance1, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	lifeInsurance1, err := strconv.Atoi(strconv.FormatBool(employee.HealthInsurance))
	if err != nil {
		fmt.Errorf("failed to convert string to int, error : %v", err)
	}
	for resultchk.Next() {
		err := resultchk.Scan(&emp.Salary, &emp.Designation, &emp.Department, &emp.TeamLead, &emp.JobType, &healthInsurance1, &lifeInsurance1, &emp.UserId)
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
