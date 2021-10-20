package models

import (
	"github.com/fatih/structs"
)

type Employee struct{
	ID     string `json:"id" structs:"id"`
	UserId string `json:"user_id" structs:"user_id"`
	Salary string `json:"salary" structs:"salary"`
	Designation     	string `json:"designation" structs:"designation"`
	Department    		string `json:"department" structs:"department"`
	TeamLead   			string `json:"team_lead" structs:"team_lead"`
	JobType   			string `json:"job_type" structs:"job_type"`
	HealthInsurance 	bool `json:"health_insurance" structs:"health_insurance"`
	LifeInsurance   	bool `json:"life_insurance" structs:"life_insurance"`
}
//Map function returns map values.
func (e *Employee) Map() map[string]interface{} {
	return structs.Map(e)
}
// Names function returns field names.
func (e *Employee) Names() []string {
	fields := structs.Fields(e)
	names := make([]string, len(fields))
	for i, field := range fields {
		name := field.Name()
		tagName := field.Tag(structs.DefaultTagName)
		if tagName != "" {
			name = tagName
		}
		names[i] = name
	}
	return names
}
