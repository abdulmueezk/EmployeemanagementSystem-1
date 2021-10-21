package handler

import (
	"EmployeemanagementSystem/models"
	"EmployeemanagementSystem/service"
	"encoding/json"
	"net/http"
)

func Handler_AddDetails (w http.ResponseWriter, r *http.Request) {
	var employee *models.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		http.Error(w, "Error decoidng response object", http.StatusBadRequest)
		panic(err)
	}
	err := service.Service_AddDetails(employee)
	if err != nil{
		json.NewEncoder(w).Encode(400)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(201)
}
