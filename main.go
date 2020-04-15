package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// Employee model
type Employee struct {
	ID         int    `json:"id"`
	FirstName  string `json:"firstname"`
	LastName   string `json:"lastname"`
	Department string `json:"department"`
}

var db *sql.DB
var err error

func main() {

	// Open Database Connection
	db, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/employee_api")

	if err != nil {
		panic(err.Error())
	} else {
		fmt.Println("Connected to Database...")
	}

	defer db.Close()

	// Init Router
	r := mux.NewRouter()

	// Route Handlers/Endpoints
	r.HandleFunc("/api/employees", getEmployees).Methods("GET")
	r.HandleFunc("/api/employees/{id}", getEmployee).Methods("GET")
	r.HandleFunc("/api/employees", createEmployee).Methods("POST")
	r.HandleFunc("/api/employees/{id}", updateEmployee).Methods("PUT")
	r.HandleFunc("/api/employees/{id}", deleteEmployee).Methods("DELETE")

	// Run Server
	fmt.Println("Server running on port :3000...")
	log.Fatal(http.ListenAndServe(":3000", handlers.CORS()(r)))
}

// CRUD Operations

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var employees []Employee

	result, err := db.Query("SELECT id, fname, lname, department FROM employees")
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	for result.Next() {
		var employee Employee
		err := result.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Department)
		if err != nil {
			panic(err.Error())
		}

		employees = append(employees, employee)
	}

	json.NewEncoder(w).Encode(employees)
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	result, err := db.Query("SELECT id, fname, lname, department FROM employees WHERE id = ?", params["id"])

	if err != nil {
		panic(err.Error())
	}
	defer result.Close()

	var employee Employee

	for result.Next() {
		err := result.Scan(&employee.ID, &employee.FirstName, &employee.LastName, &employee.Department)
		if err != nil {
			panic(err.Error())
		}
	}
	json.NewEncoder(w).Encode(employee)
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	stmt, err := db.Prepare("INSERT INTO employees(fname, lname, department) VALUES(?, ?, ?)")
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}

	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	firstname := keyVal["firstname"]
	lastname := keyVal["lastname"]
	department := keyVal["department"]
	_, err = stmt.Exec(firstname, lastname, department)
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Employee succesfully created")
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	stmt, err := db.Prepare("UPDATE employees SET fname = ?, lname = ?, department = ? WHERE id = ?")

	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]string)
	json.Unmarshal(body, &keyVal)
	firstname := keyVal["firstname"]
	lastname := keyVal["lastname"]
	department := keyVal["department"]
	_, err = stmt.Exec(firstname, lastname, department, params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Employee with ID = %s succesfully updated", params["id"])

}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM employees WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err.Error())
	}
	fmt.Fprintf(w, "Employee with ID = %s succesfully deleted", params["id"])
}
