package controllers

import (
	"github.com/faralam/go-contacts/utils"
	"net/http"

	"github.com/faralam/go-contacts/models"
	"github.com/faralam/go-employee-details/models"
)

var CreateEmployee = func(w http.ResponseWriter, w *http.Request) {
	employee := &models.Employee{}

	err := json.NewDecoder(r.Body).Decode(employee)
	if err != nil{
		utils.Respond(w,utils.Message(false,"Invalid request"))
		return
	}
	resp := employee.CreateEmployee()
	u.Respond(w,resp)
}
