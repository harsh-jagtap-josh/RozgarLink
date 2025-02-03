package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository/auth"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repository/domain"
	"golang.org/x/crypto/bcrypt"
)

type workerRegisterResponse struct {
	message string
	token   string
	worker  domain.Worker
}

type employerRegisterResponse struct {
	message  string
	token    string
	employer domain.Employer
}

type LoginResponse struct {
	message string
	token   string
	user    any
}

// RegisterWorkerHandler handles worker registration.
func RegisterWorkerHandler(w http.ResponseWriter, r *http.Request) {

	db, err := repository.ConnectDB()
	if err != nil {
		fmt.Println("Database error Handler")
		return
	}
	var worker domain.Worker
	if err := json.NewDecoder(r.Body).Decode(&worker); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	hashedPassword, err := HashPassword(worker.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	worker.Password = hashedPassword

	workerID, err := auth.CreateWorker(db.DB, worker)
	if err != nil {
		http.Error(w, "Failed to register worker", http.StatusInternalServerError)
		return
	}
	token, err := GenerateJWT(workerID, "worker")
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	newResponse := workerRegisterResponse{"worker created successfully", token, worker} // not returning the entitiy stored in database rather the entity
	handleMarhalAndResponse(w, http.StatusCreated, newResponse)
}

// RegisterEmployerHandler handles employer registration.
func RegisterEmployerHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.ConnectDB()
	if err != nil {
		fmt.Println("Database error Handler")
		return
	}
	var employer domain.Employer
	if err := json.NewDecoder(r.Body).Decode(&employer); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	hashedPassword, err := HashPassword(employer.Password)
	if err != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}
	employer.Password = hashedPassword

	employerID, err := auth.CreateEmployer(db.DB, employer)
	if err != nil {
		http.Error(w, "Failed to register employer", http.StatusInternalServerError)
		return
	}

	token, err := GenerateJWT(employerID, "employer")
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	newResponse := employerRegisterResponse{"worker created successfully", token, employer} // not returning the entitiy stored in database rather the entity
	handleMarhalAndResponse(w, http.StatusCreated, newResponse)
}

// LoginHandler handles login for workers and employers.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	db, err := repository.ConnectDB()
	if err != nil {
		fmt.Println("Database error Handler")
		return
	}
	var loginData struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	var user auth.User
	if loginData.Role == "worker" {
		user, err = auth.GetWorkerByEmail(db.DB, loginData.Email)
	} else if loginData.Role == "employer" {
		user, err = auth.GetEmployerByEmail(db.DB, loginData.Email)
	} else {
		http.Error(w, "Invalid role", http.StatusBadRequest)
		return
	}

	if err == sql.ErrNoRows || CheckPasswordHash(loginData.Password, user.Password) {
		http.Error(w, "Invalid email or password", http.StatusUnauthorized)
		return
	}

	token, err := GenerateJWT(user.ID, loginData.Role)
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}
	newResponse := LoginResponse{"worker created successfully", token, user}
	handleMarhalAndResponse(w, http.StatusCreated, newResponse)

}

var secretKey = []byte("your_secret_key")

// GenerateJWT generates a JWT token with user role.
func GenerateJWT(userID int, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(secretKey)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
