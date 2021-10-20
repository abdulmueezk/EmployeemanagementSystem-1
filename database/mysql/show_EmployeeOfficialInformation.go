package mysql

import (
	"EmployeemanagementSystem/models"
	"fmt"
)

func Db_ShowDetails(employee *models.Employee) error {
	var abcd error
	db := Sqlclient()
	query := "SELECT * FROM employeeOfficial WHERE userid='" + employee.UserId + "'"
	querychk := "select userId from employeeOfficial where userid ='" + employee.UserId + "'"
	resultchk, err := db.Query(querychk)
	if err != nil {
		return err
	}
	defer resultchk.Close()
	var emp models.Employee
	for resultchk.Next() {
		err := resultchk.Scan(&emp.UserId)
		if err != nil {
			return err
		}
	}
	if employee.UserId == emp.UserId {
		result, err := db.Query(query)
		if err != nil {
			return err

		}
		defer result.Close()
		for result.Next() {
			err := result.Scan(&employee.Salary, &employee.Designation, &employee.Department, &employee.TeamLead, &employee.JobType, &employee.HealthInsurance, &employee.LifeInsurance, &employee.UserId)
			if err != nil {
				return err
			}
		}
		defer db.Close()
	} else {
		var x int
		err := fmt.Errorf("math: square root of negative number %v", x)
		abcd = err
	}
return abcd
}
