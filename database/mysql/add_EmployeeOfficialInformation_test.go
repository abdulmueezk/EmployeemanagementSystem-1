package mysql

import (
	"EmployeemanagementSystem/models"
	"testing"
)

func TestDb_addDetails(t *testing.T) {
	type args struct {
		employee *models.Employee
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - add employee details in db",
			args:    args{employee: &models.Employee{ID:"123",UserId: "31601", Salary: "20000", Department: "backend", Designation: "Intern" , TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: false, LifeInsurance: false}},
			wantErr: false,
		},
		{
			name:    "Failed - add employee details in db",
			args:    args{employee: &models.Employee{ Salary: "20000", Department: "backend", Designation: "Intern" , TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: true, LifeInsurance: false}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Db_addDetails(tt.args.employee); (err != nil) != tt.wantErr {
				t.Errorf("Db_addDetails() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
