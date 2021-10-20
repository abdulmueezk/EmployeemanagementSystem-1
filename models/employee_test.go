package models

import (
	"fmt"
	"reflect"
	"testing"
)

func TestEmployee_Map(t *testing.T) {
	type fields struct {
		ID              string
		UserId          string
		Salary          string
		Designation     string
		Department      string
		TeamLead        string
		JobType         string
		HealthInsurance bool
		LifeInsurance   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name: " success - student struct to map",
			fields: fields{
				ID:              "12345",
				UserId:          "abc123",
				Salary:          "20000",
				Designation:     "Intern",
				Department:      "BackEnd",
				TeamLead:        "Kashif ALi",
				JobType:         "Intern",
				HealthInsurance: false,
				LifeInsurance:   false,
			},
			want: map[string]interface{}{
				"id":              "12345",
				"user_id":          "abc123",
				"salary":          "20000",
				"designation":     "Intern",
				"department":      "BackEnd",
				"team_lead":        "Kashif ALi",
				"job_type":         "Intern",
				"health_insurance": false,
				"life_insurance":   false,
			},
		},
		{
			name: " success - student struct to map",
			fields: fields{
				ID:              "12345",
				UserId:          "abc123",
				Salary:          "20000",
				Designation:     "Intern",
				Department:      "BackEnd",
				TeamLead:        "Kashif ALi",
				JobType:         "Intern",
				HealthInsurance: false,
				LifeInsurance:   false,
			},
			want: map[string]interface{}{
				"id":              "12345",
				"user_id":          "abc123",
				"salary":          "20000",
				"designation":     "Intern",
				"department":      "BackEnd",
				"team_lead":        "Kashif ALi",
				"job_type":         "Intern",
				"health_insurance": false,
				"life_insurance":   false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Employee{
				ID:              tt.fields.ID,
				UserId:          tt.fields.UserId,
				Salary:          tt.fields.Salary,
				Designation:     tt.fields.Designation,
				Department:      tt.fields.Department,
				TeamLead:        tt.fields.TeamLead,
				JobType:         tt.fields.JobType,
				HealthInsurance: tt.fields.HealthInsurance,
				LifeInsurance:   tt.fields.LifeInsurance,
			}
			fmt.Println("GOT:",e.Map())
			fmt.Println("WANT:",tt.want)
			if got := e.Map(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmployee_Names(t *testing.T) {
	type fields struct {
		ID              string
		UserId          string
		Salary          string
		Designation     string
		Department      string
		TeamLead        string
		JobType         string
		HealthInsurance bool
		LifeInsurance   bool
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "success",
			fields: fields{
				UserId:          "123abc",
				Salary:          "25000",
				Designation:     "Intern",
				Department:      "BackEnd",
				TeamLead:        "Kashif Ali",
				JobType:         "Intern",
				HealthInsurance: false,
				LifeInsurance:   false,
			},
			want: []string{"id","user_id", "salary", "designation", "department", "team_lead", "job_type", "health_insurance", "life_insurance"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Employee{
				ID:              tt.fields.ID,
				UserId:          tt.fields.UserId,
				Salary:          tt.fields.Salary,
				Designation:     tt.fields.Designation,
				Department:      tt.fields.Department,
				TeamLead:        tt.fields.TeamLead,
				JobType:         tt.fields.JobType,
				HealthInsurance: tt.fields.HealthInsurance,
				LifeInsurance:   tt.fields.LifeInsurance,
			}
			if got := e.Names(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Names() = %v, want %v", got, tt.want)
			}
		})
	}
}
