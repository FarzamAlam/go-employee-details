package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/farzamalam/go-employee-details/controllers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	e := godotenv.Load()
	if e != nil {
		fmt.Println("Error while loading the environment variables", e)
	}
	port := os.Getenv("PORT")
	router.HandleFunc("/api/v1/employees", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/v1/employees", controllers.GetEmployees).Methods("GET")
	router.HandleFunc("/api/v1/employees/{id}", controllers.GetEmployee).Methods("GET")
	router.HandleFunc("/api/v1/employees/{id}", controllers.UpdateEmployee).Methods("PUT")
	router.HandleFunc("/api/v1/employees/{id}", controllers.DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
