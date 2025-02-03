package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository/postgres"
)

type FetchEmployerResponseType struct {
	Message   string
	Employers []domain.Employer
}

type HandleEmpByIdResponseType struct {
	Message  string
	Employer domain.Employer
}

type DeleteEmployerResponseType struct {
	Message string
	Id      int
}

type FetchEmployerJobsType struct {
	Message string
	Jobs    []domain.Job
}

// func HandleFetchEmployers(w http.ResponseWriter, r *http.Request) {

// 	db, err := repository.ConnectDB()
// 	if err != nil {
// 		fmt.Println("Database error Handler")
// 		return
// 	}

// 	queryParams := r.URL.Query()

// 	// retrieving all query params
// 	isAvailable := queryParams.Get("isAvailable")
// 	name := queryParams.Get("name")
// 	sector := queryParams.Get("sector")
// 	jobsWorked := queryParams.Get("jobsWorked")
// 	rating := queryParams.Get("rating")

// 	workers, err := postgres.GetFilteredWorkers(db.DB, &isAvailable, &name, &sector, &rating, &jobsWorked)

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	newResponse := FetchWorkerResponseType{Message: "", Workers: workers}
// 	handleMarhalAndResponse(w, http.StatusOK, newResponse)
// }

func HandleEmployerByID(w http.ResponseWriter, r *http.Request) {
	req := r.Method

	vars := mux.Vars(r) // returns route variables
	id := vars["employer_id"]
	id_int, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "invalid ID format provided", http.StatusBadRequest)
		return
	}

	if len(id) == 0 {
		http.Error(w, "ID cannot be empty", http.StatusBadRequest)
		return
	}

	switch req {
	case http.MethodGet:
		HandleGetEmployerByID(w, r, id_int)
	case http.MethodPut:
		HandleUpdateEmployerByID(w, r, id_int)
	case http.MethodDelete:
		HandleDeleteEmployerByID(w, r, id_int)
	default:
		http.Error(w, "inavlid request method", http.StatusMethodNotAllowed)
		return
	}
}

func HandleEmployerJobs(w http.ResponseWriter, r *http.Request) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "error connecting to database", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r) // returns route variables
	id := vars["employer_id"]
	id_int, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid ID format provided", http.StatusBadRequest)
		return
	}
	jobs, err := postgres.GetJobsByEmployerId(db.DB, id_int)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no jobs found with employer ID", http.StatusNotFound)
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := FetchEmployerJobsType{
		Message: "jobs fetched successfully",
		Jobs:    *jobs,
	}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleGetEmployerByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	employer, err := postgres.GetEmployerByID(db.DB, id)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no employer found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newResponse := HandleEmpByIdResponseType{Message: "employer profile retrieved successfully", Employer: *employer}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)

}

func HandleUpdateEmployerByID(w http.ResponseWriter, r *http.Request, id int) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error: "+err.Error(), http.StatusInternalServerError)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}
	var employer domain.Employer
	err = json.Unmarshal(body, &employer)
	if err != nil {
		http.Error(w, "internal error:"+err.Error(), http.StatusBadRequest)
	}
	updatedEmployer, err := postgres.UpdateEmployer(db.DB, employer)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no employer found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := HandleEmpByIdResponseType{Message: "employer details updated successfully", Employer: updatedEmployer}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleDeleteEmployerByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error: "+err.Error(), http.StatusInternalServerError)
	}

	rowsCount, err := postgres.DeleteEmployer(db.DB, id)
	if rowsCount == 0 {
		http.Error(w, "employer with ID not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := DeleteWorkerResponseType{
		Message: "employer deleted successfully",
		Id:      id,
	}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}
