package models

import (
	"fmt"
	"strings"

	"github.com/faralam/go-contacts/utils"
)

type Employee struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	MobileNumber string `json:"mobile_number"`
	Department   string `json:"department"`
}

func (employee *Employee) Validate() (map[string]interface{}, bool) {
	if string(employee.ID) == "" {
		return utils.Message(false, "Invalid Employee ID"), false
	}
	if employee.Name == "" {
		return utils.Message(false, "Invalid Employee Name"), false
	}
	if !strings.Contains(employee.Email, "@") {
		return utils.Message(false, "Invalid email address"), false
	}
	return utils.Message(true, "Success"), true
}

func (employee *Employee) Create() map[string]interface{} {
	resp, ok := employee.Validate()
	if !ok {
		return resp
	}
	GetDB().Create(employee)
	resp["employee"] = employee
	return resp
}

func GetEmployees() []*Employee {
	employees := make([]*Employee, 0)
	err := GetDB().Table("employee").Find(&employees).Error
	if err != nil {
		return nil
		fmt.Println("Error while calling the employee table")
	}
	return employees
}
