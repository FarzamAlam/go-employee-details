package models

import (
	"fmt"
	"strings"

	"github.com/farzamalam/go-employee-details/utils"
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
	err := GetDB().Table("employees").Find(&employees).Error
	if err != nil {
		return nil
		fmt.Println("Error while calling the employee table", err)
	}
	return employees
}

func GetEmployee(id int) *Employee {
	employee := &Employee{}
	err := GetDB().Table("employees").First(employee).Error
	if err != nil {
		return nil
		fmt.Println("Error in GetEmployee():", err)
	}
	return employee
}

func (employee *Employee) UpdateEmployee(id int) *Employee {

	if id != employee.ID {
		return nil
	}
	GetDB().Table("employees").First(&employee)
	GetDB().Table("employees").Save(&employee)
	return employee
}
func DeleteEmployee(id int) bool {
	if id < 1 {
		return false
	}
	GetDB().Table("employees").Delete(id)
	return true
}
