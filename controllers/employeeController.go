package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/farzamalam/go-employee-details/models"
	"github.com/farzamalam/go-employee-details/utils"
)

var CreateEmployee = func(w http.ResponseWriter, r *http.Request) {
	employee := &models.Employee{}

	err := json.NewDecoder(r.Body).Decode(employee)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Invalid request"))
		return
	}
	resp := employee.Create()
	utils.Respond(w, resp)
}

var GetEmployees = func(w http.ResponseWriter, r *http.Request) {
	resp := utils.Message(true, "Success")
	resp["data"] = models.GetEmployees()
	utils.Respond(w, resp)
}

var GetEmployee = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in id paramter in url"))
	}
	resp := utils.Message(true, "Success")
	resp["data"] = models.GetEmployee(id)
	utils.Respond(w, resp)
}

var UpdateEmployee = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in id parameter in url"))
	}
	employee := &models.Employee{}
	err = json.NewDecoder(r.Body).Decode(employee)
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in the decoding json."))
	}
	resp := utils.Message(true, "Success")
	resp["data"] = employee.UpdateEmployee(id)

	utils.Respond(w, resp)
}

var DeleteEmployee = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		utils.Respond(w, utils.Message(false, "Error in the id parameter of url"))
	}
	ok := models.DeleteEmployee(id)
	if !ok {
		utils.Respond(w, utils.Message(false, "Error while deleting the entry"))
	}
	utils.Respond(w, utils.Message(true, "Employee successfully deleted."))
}
