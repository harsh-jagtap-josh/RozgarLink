package routes

import (
	"github.com/gorilla/mux"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/api/handlers"
)

// Worker API Group

func NewRouter(api *mux.Router) {

	// Authentication Routes
	api.HandleFunc("/register/worker", handlers.RegisterWorkerHandler).Methods("POST")
	api.HandleFunc("/register/employer", handlers.RegisterEmployerHandler).Methods("POST")
	api.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	// Worker Routes
	WorkersRouter := api.PathPrefix("/workers").Subrouter()
	WorkersRouter.HandleFunc("", handlers.HandleFetchWorkers).Methods("GET")
	WorkersRouter.HandleFunc("/{worker_id}", handlers.HandleWorkerByID).Methods("GET", "PUT", "DELETE", "PATCH")
	WorkersRouter.HandleFunc("/{worker_id}/applications", handlers.HandleWorkerApplications).Methods("GET")

	// Employer Routes
	EmployersRouter := api.PathPrefix("/employers").Subrouter()
	EmployersRouter.HandleFunc("/{employer_id}", handlers.HandleEmployerByID).Methods("GET", "PUT", "DELETE", "PATCH")
	EmployersRouter.HandleFunc("/{employer_id}/jobs", handlers.HandleEmployerJobs).Methods("GET")

	// Job Routes
	JobsRouter := api.PathPrefix("/jobs").Subrouter()
	JobsRouter.HandleFunc("", handlers.HandleCreateJob).Methods("POST")
	JobsRouter.HandleFunc("", handlers.HandleFetchJobs).Methods("GET")
	JobsRouter.HandleFunc("/{job_id}", handlers.HandleJobByID).Methods("GET", "PUT", "DELETE", "PATCH")
	JobsRouter.HandleFunc("/{job_id}/applications", handlers.HandleJobApplications).Methods("GET")

	// Sector Routes
	SectorsRouter := api.PathPrefix("/sectors").Subrouter()
	SectorsRouter.HandleFunc("", handlers.HandleCreateSector).Methods("POST")
	SectorsRouter.HandleFunc("", handlers.HandleFetchSectors).Methods("GET")
	SectorsRouter.HandleFunc("/{sector_id}", handlers.HandleSectorByID).Methods("GET", "PUT", "DELETE", "PATCH")

	// Application Routes
	ApplicationsRouter := api.PathPrefix("/applications").Subrouter()
	ApplicationsRouter.HandleFunc("/", handlers.HandleCreateApplication).Methods("GET", "POST")
	ApplicationsRouter.HandleFunc("/{application_id}", handlers.HandleApplicationByID).Methods("GET", "PUT", "DELETE", "PATCH")

}
