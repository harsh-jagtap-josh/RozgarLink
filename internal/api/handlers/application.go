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

type FetchApplicationResponseType struct {
	Message      string
	Applications []domain.Application
}

type HandleApplicationByIdResponseType struct {
	Message     string
	Application domain.Application
}

type DeleteApplicationResponseType struct {
	Message string
	Id      int
}

func HandleCreateApplication(w http.ResponseWriter, r *http.Request) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error", http.StatusInternalServerError)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}
	var application domain.Application
	err = json.Unmarshal(body, &application)
	if err != nil {
		http.Error(w, "internal error:"+err.Error(), http.StatusBadRequest)
	}
	_, err = postgres.CreateApplication(db.DB, application)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}

	newResponse := HandleApplicationByIdResponseType{Message: "worker details updated successfully", Application: application}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleApplicationByID(w http.ResponseWriter, r *http.Request) {
	req := r.Method

	vars := mux.Vars(r) // returns route variables
	id := vars["application_id"]
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
		HandleGetApplicationByID(w, r, id_int)
	case http.MethodPut:
		HandleUpdateApplicationByID(w, r, id_int)
	case http.MethodDelete:
		HandleDeleteApplicationByID(w, r, id_int)
	default:
		http.Error(w, "inavlid request method", http.StatusMethodNotAllowed)
		return
	}
}

func HandleGetApplicationByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	application, err := postgres.GetApplicationByID(db.DB, id)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no application found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	newResponse := HandleApplicationByIdResponseType{Message: "application retrieved successfully", Application: *application}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)

}

func HandleUpdateApplicationByID(w http.ResponseWriter, r *http.Request, id int) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}
	var application domain.Application
	err = json.Unmarshal(body, &application)
	if err != nil {
		http.Error(w, "internal error:"+err.Error(), http.StatusBadRequest)
	}
	updatedApplication, err := postgres.UpdateApplication(db.DB, application)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no application found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := HandleApplicationByIdResponseType{Message: "application details updated successfully", Application: updatedApplication}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleDeleteApplicationByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error: "+err.Error(), http.StatusInternalServerError)
	}

	rowsCount, err := postgres.DeleteApplication(db.DB, id)
	if rowsCount == 0 {
		http.Error(w, "application with ID not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := DeleteJobResponseType{
		Message: "application deleted successfully",
		Id:      id,
	}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}
