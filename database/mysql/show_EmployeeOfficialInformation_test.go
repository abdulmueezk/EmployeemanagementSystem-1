package mysql

import (
	"EmployeemanagementSystem/models"
	"testing"
)

func TestDb_ShowDetails(t *testing.T) {
	type args struct {
		employee *models.Employee
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - show employee details in db",
			args:    args{employee: &models.Employee{UserId: "31601", Salary: "20000", Department: "backend", Designation: "Intern" , TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: false, LifeInsurance: false}},
			wantErr: false,
		},
		{
			name:    "Failed - show employee details in db",
			args:    args{employee: &models.Employee{ TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: false, LifeInsurance: false}},
			wantErr: false,
		},
		{
			name:    "success - show employee details in db",
			args:    args{employee: &models.Employee{UserId: "31601", Salary: "", Department: "backend", Designation: "Intern" , TeamLead: "Kashif Ali", JobType: "Intern", HealthInsurance: false, LifeInsurance: false}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Db_ShowDetails(tt.args.employee); (err != nil) != tt.wantErr {
				t.Errorf("Db_ShowDetails() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
