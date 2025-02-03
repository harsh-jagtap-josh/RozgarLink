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

type FetchSectorResponseType struct {
	Message string
	Sectors []domain.Sector
}

type HandleSectorByIdResponseType struct {
	Message string
	Sector  domain.Sector
}

func HandleCreateSector(w http.ResponseWriter, r *http.Request) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error", http.StatusInternalServerError)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}
	var sector domain.Sector
	err = json.Unmarshal(body, &sector)
	if err != nil {
		http.Error(w, "internal error:"+err.Error(), http.StatusBadRequest)
	}
	_, err = postgres.CreateSector(db.DB, sector)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}

	newResponse := HandleSectorByIdResponseType{Message: "sector details updated successfully", Sector: sector}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleFetchSectors(w http.ResponseWriter, r *http.Request) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	sectors, err := postgres.GetAllSectors(db.DB)

	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := FetchSectorResponseType{Message: "", Sectors: sectors}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleSectorByID(w http.ResponseWriter, r *http.Request) {
	req := r.Method

	vars := mux.Vars(r) // returns route variables
	id := vars["sector_id"]
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
		HandleGetSectorByID(w, r, id_int)
	case http.MethodPut:
		HandleUpdateSectorByID(w, r, id_int)
	case http.MethodDelete:
		HandleDeleteSectorByID(w, r, id_int)
	default:
		http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		return
	}
}

func HandleGetSectorByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error: "+err.Error(), http.StatusInternalServerError)
		return
	}
	sector, err := postgres.GetSectorByID(db.DB, id)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no sector found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "Error while fetching data from database", http.StatusInternalServerError)
		return
	}
	newResponse := HandleSectorByIdResponseType{Message: "sector profile retrieved successfully", Sector: *sector}
	handleMarhalAndResponse(w, http.StatusOK, newResponse)

}

func HandleUpdateSectorByID(w http.ResponseWriter, r *http.Request, id int) {

	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error", http.StatusInternalServerError)
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
	}
	var sector domain.Sector
	err = json.Unmarshal(body, &sector)
	if err != nil {
		http.Error(w, "internal error:"+err.Error(), http.StatusBadRequest)
	}
	updatedSector, err := postgres.UpdateSector(db.DB, sector)
	if errors.Is(err, sql.ErrNoRows) {
		http.Error(w, "no sector found with ID", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := HandleSectorByIdResponseType{Message: "job details updated successfully", Sector: updatedSector}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}

func HandleDeleteSectorByID(w http.ResponseWriter, r *http.Request, id int) {
	db, err := repository.ConnectDB()
	if err != nil {
		http.Error(w, "internal db error: "+err.Error(), http.StatusInternalServerError)
	}

	rowsCount, err := postgres.DeleteSector(db.DB, id)
	if rowsCount == 0 {
		http.Error(w, "sector with ID not found", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, "internal error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	newResponse := DeleteJobResponseType{
		Message: "sector deleted successfully",
		Id:      id,
	}

	handleMarhalAndResponse(w, http.StatusOK, newResponse)
}
