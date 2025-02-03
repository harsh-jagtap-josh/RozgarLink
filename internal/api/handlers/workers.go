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

type FetchWorkerResponseType struct {
	Message string
	Workers []domain.Worker
}

type HandleByIdResponseType struct {
	Message string
	Worker  domain.Worker
}

type FetchApplicationsType struct {
	Message      string
	Applications []domain.Application
}

type DeleteWorkerResponseType struct {
	Message string
	Id      int
}

func HandleFetchWorkers(w http.ResponseWriter, r *http.Request) {

	db, err := repository.ConnectDB()
	if err != nil {
		fmt.Println("Database error Handler")
		return
	}

	queryParams := r.URL.Query()

	// retrieving all query params
	isAvailable := queryParams.Get("isAvailable")
	name := queryParams.Get("name")
	sector := queryParams.Get("sector")
	jobsWorked := queryParams.Get("jobsWorked")
	rating := queryParams.Get("rating")

	workers, err := postgres.GetFilteredWorkers(db.DB, &isAvailable, &name, &sector, &rating, &jobsWorked)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := FetchWorkerResponseType{Message: "", Workers: workers}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleWorkerApplications(w http.ResponseWriter, r *http.Request) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "error connecting to database", http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r) // returns route variables
	id := vars["worker_id"]
	id_int, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid ID format provided", http.StatusBadRequest)
		return
	}
	applications, err := postgres.GetApplicationsByWorkerID(db.DB, id_int)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no applications found with worker ID", http.StatusNotFound)
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

func HandleWorkerByID(w http.ResponseWriter, r *http.Request) {
	req := r.Method

	vars := mux.Vars(r) // returns route variables
	id := vars["worker_id"]
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
		HandleGetWorkerByID(w, r, id_int)
	case http.MethodPut:
		HandleUpdateWorkerByID(w, r, id_int)
	case http.MethodDelete:
		HandleDeleteWorkerByID(w, r, id_int)
	default:
		http.Error(w, "inavlid request method", http.StatusMethodNotAllowed)
		return
	}
}

func HandleGetWorkerByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "DB Error", http.StatusInternalServerError)
		return
	}
	worker, err := postgres.GetWorkerByID(db.DB, id)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no worker found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error while fetching data from database", http.StatusInternalServerError)
		return
	}
	newResponse := HandleByIdResponseType{Message: "worker profile retrieved successfully", Worker: *worker}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)

}

func HandleUpdateWorkerByID(w http.ResponseWriter, r *http.Request, id int) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error", http.StatusInternalServerError)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}
	var worker domain.Worker
	err = json.Unmarshal(body, &worker)
	if err != nil {
		http.Error(w, "internal error:"+err.Error(), http.StatusBadRequest)
	}
	updatedWorker, err := postgres.UpdateWorker(db.DB, worker)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no worker found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := HandleByIdResponseType{Message: "worker details updated successfully", Worker: updatedWorker}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleDeleteWorkerByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "error connecting to db", http.StatusInternalServerError)
	}

	rowsCount, err := postgres.DeleteWorker(db.DB, id)
	if rowsCount == 0 {
		http.Error(w, "worker with ID not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := DeleteWorkerResponseType{
		Message: "worker deleted successfully",
		Id:      id,
	}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func handleMarhalAndResponse(w http.ResponseWriter, statusCode int, data any) {
	httpResponse, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(statusCode)
	w.Write(httpResponse)
}
