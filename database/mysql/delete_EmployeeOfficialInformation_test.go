package mysql

import (
	"EmployeemanagementSystem/models"
	"testing"
)

func TestDb_DeleteDetails(t *testing.T) {
	type args struct {
		employee *models.Employee
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - delete employee details in db",
			args:    args{employee: &models.Employee{UserId: "31601",JobType: "Left"}},
			wantErr: false,
		},
		{
			name:    "Failed - delete employee details in db",
			args:    args{employee: &models.Employee{ Salary: "20000", Department: "backend", Designation: "Intern" , TeamLead: "Kashif Ali", JobType: "Left", HealthInsurance: false, LifeInsurance: false}},
			wantErr: false,
		},
		{
			name:    "Failed - delete employee details in db",
			args:    args{employee: &models.Employee{ TeamLead: "Kashif Ali", JobType: "Left", HealthInsurance: false, LifeInsurance: false}},
			wantErr: false,
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Db_DeleteDetails(tt.args.employee); (err != nil) != tt.wantErr {
				t.Errorf("Db_DeleteDetails() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
