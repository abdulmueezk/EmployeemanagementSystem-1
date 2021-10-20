package mysql

import (
	"EmployeemanagementSystem/models"
	"testing"
)

func TestDb_UpdateDetails(t *testing.T) {
	type args struct {
		employee *models.Employee
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - update employee details in db",
			args:    args{employee: &models.Employee{UserId: "31601", Salary: "20000", Department: "backend", Designation: "Intern" , TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: false, LifeInsurance: false}},
			wantErr: false,
		},
		{
			name:    "Failed - update employee details in db",
			args:    args{employee: &models.Employee{Salary: "30000", Department: "Backend", Designation: "Intern" , TeamLead: "Abdul Moeid",JobType: "Intern", HealthInsurance: false, LifeInsurance: false}},
			wantErr: false,
		},
		{
			name:    "Failed - update employee details in db",
			args:    args{employee: &models.Employee{Salary: "40000", TeamLead: "Kashif Ali", JobType: "Regular", HealthInsurance: true, LifeInsurance: true}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Db_UpdateDetails(tt.args.employee); (err != nil) != tt.wantErr {
				t.Errorf("Db_UpdateDetails() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
