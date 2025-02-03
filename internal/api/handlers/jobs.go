package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository/postgres"
)

type FetchJobResponseType struct {
	Message string
	Jobs    []domain.Job
}

type HandleJobByIdResponseType struct {
	Message string
	Job     domain.Job
}

type FetchJobsType struct {
	Message      string
	Applications []domain.Application
}

type DeleteJobResponseType struct {
	Message string
	Id      int
}

func HandleCreateJob(w http.ResponseWriter, r *http.Request) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error", http.StatusInternalServerError)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}
	var job domain.Job
	err = json.Unmarshal(body, &job)
	if err != nil {
		http.Error(w, "internal error:"+err.Error(), http.StatusBadRequest)
	}
	_, err = postgres.CreateJob(db.DB, job)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}

	newResponse := HandleJobByIdResponseType{Message: "job created successfully", Job: job}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleFetchJobs(w http.ResponseWriter, r *http.Request) {

	db, err := repository.ConnectDB()
	if err != nil {
		fmt.Println("Database error Handler")
		return
	}

	queryParams := r.URL.Query()

	// retrieving all query params
	title := queryParams.Get("title")
	// location := queryParams.Get("location")
	sector := queryParams.Get("sector")
	wage := queryParams.Get("wage")
	rating := queryParams.Get("rating")

	jobs, err := postgres.GetFilteredJobs(db.DB, &title, &sector, &wage, &rating)

	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := FetchJobResponseType{Message: "", Jobs: jobs}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleJobApplications(w http.ResponseWriter, r *http.Request) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "error connecting to database", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r) // returns route variables
	id := vars["job_id"]
	id_int, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid ID format provided", http.StatusBadRequest)
		return
	}
	applications, err := postgres.GetApplicationsByJobID(db.DB, id_int)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no applications found with job ID", http.StatusNotFound)
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := FetchApplicationsType{
		Message:      "applications fetched successfully",
		Applications: *applications,
	}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleJobByID(w http.ResponseWriter, r *http.Request) {
	req := r.Method

	vars := mux.Vars(r) // returns route variables
	id := vars["job_id"]
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
		HandleGetJobByID(w, r, id_int)
	case http.MethodPut:
		HandleUpdateJobByID(w, r, id_int)
	case http.MethodDelete:
		HandleDeleteJobByID(w, r, id_int)
	default:
		http.Error(w, "inavlid request method", http.StatusMethodNotAllowed)
		return
	}
}

func HandleGetJobByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}
	job, err := postgres.GetJobByID(db.DB, id)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no job found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error while fetching data from database", http.StatusInternalServerError)
		return
	}
	newResponse := HandleJobByIdResponseType{Message: "job profile retrieved successfully", Job: *job}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)

}

func HandleUpdateJobByID(w http.ResponseWriter, r *http.Request, id int) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error", http.StatusInternalServerError)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}
	var job domain.Job
	err = json.Unmarshal(body, &job)
	if err != nil {
		http.Error(w, "internal error:"+err.Error(), http.StatusBadRequest)
	}
	updatedJob, err := postgres.UpdateJob(db.DB, job)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no job found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := HandleJobByIdResponseType{Message: "job details updated successfully", Job: updatedJob}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleDeleteJobByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "error connecting to db", http.StatusInternalServerError)
	}

	rowsCount, err := postgres.DeleteJob(db.DB, id)
	if rowsCount == 0 {
		http.Error(w, "job with ID not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := DeleteJobResponseType{
		Message: "job deleted successfully",
		Id:      id,
	}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}
