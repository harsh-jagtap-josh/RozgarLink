# RozgarLink

## Problem Statement

<p>
RozgarLink is a platform designed to connect daily wage workers and semi-skilled individuals in India with potential employers. 

It provides a simple, accessible solution to address the challenges of unemployment among workers in non-technical sectors like construction, hospitality, and small-scale industries. 

The platform simplifies the hiring process, ensures fair wages, and creates transparency in the job market, thereby empowering workers and improving employer access to skilled labor.
</p>

## Project Documentation

[docs](https://docs.google.com/document/d/1MSxSgyEqYstvGXIdVld4aU8We1fJNNqVOYyKqGm4ehw/edit?usp=sharing)

## APIs

#### Worker
1. <b>List Workers </b> : `GET http://localhost:8080/worker`
2. <b>Get Worker Details API</b> : `GET http://localhost:8080/worker/{worker_id}`
3. <b>Edit Worker Details API</b> : `PUT http://localhost:8080/worker/{worker_id}`
4. <b>Delete Worker  API</b> : `DELETE http://localhost:8080/worker/{worker_id}`
5. <b>Create New Worker Account API</b> : `POST http://localhost:8080/worker`

#### Employer

1. <b>List Employers</b> : `GET http://localhost:8080/employer`
2. <b>Create a New Employer API</b> : `POST http://localhost:8080/employer`
1. <b>Get Employer Details API</b> : `GET http://localhost:8080/employer/{employer_id}`
1. <b>Update Employer Details API</b> : `PUT http://localhost:8080/employer/{employer_id}`
1. <b>Details Employer Account API</b> : `DELETE http://localhost:8080/employer/{employer_id}`

#### Job
1. <b>List Jobs</b> : `GET http://localhost:8080/job`
2. <b>Create a New Job API</b> : `POST http://localhost:8080/jobs`
3. <b>Get Job Details API</b> : `GET http://localhost:8080/jobs/{job_id}`
4. <b>Update Job Details API</b> : `PUT http://localhost:8080/jobs/{job_id}`
5. <b>Details Job Details API</b> : `DELETE http://localhost:8080/jobs/{job_id}`
6. <b>List Jobs by Employer ID</b> : `GET http://localhost:8080/employer/{employer_id}/jobs`


#### Applications

1. <b>List Applications</b> : `GET http://localhost:8080/applications`
2. <b>Create a New Application API</b> : `POST http://localhost:8080/applications`
3. <b>Get Application Details API</b> : `GET http://localhost:8080/applications/{application_id}`
4. <b>Update Application Details API</b> : `PUT http://localhost:8080/applications/{application_id}`
5. <b>Details Application Details API</b> : `DELETE http://localhost:8080/applications/{application_id}`
6. <b>List Applications by Worker ID</b> : `GET http://localhost:8080/worker/{worker_id}/applications`
7. <b>List Applications by Job ID</b> : `GET http://localhost:8080/jobs/{job_id}/applications`


#### Sectors

1. <b>List Sectors</b> : `GET http://localhost:8080/sectors`
2. <b>Create a New Sector API</b> : `POST http://localhost:8080/sectors`
3. <b>Get Sector Details API</b> : `GET http://localhost:8080/sectors/{sectors_id}`
4. <b>Update Sector Details API</b> : `PUT http://localhost:8080/sectors/{sectors_id}`
4. <b>Delete Sector Details API</b> : `DELETE http://localhost:8080/sectors/{sectors_id}`



## Postman Collection


[RozgarLink_Postman](RozgarLink.postman_collection.json)


## Project File Structure

```
├── cmd
│   └── main.go
├── internal
│   ├── app
│   │   ├── application
│   │   │   ├── domain.go
│   │   │   ├── handler.go
│   │   │   ├── helper.go
│   │   │   └── service.go
│   │   ├── auth
│   │   │   ├── domain.go
│   │   │   ├── handler.go
│   │   │   └── service.go
│   │   ├── employer
│   │   │   ├── domain.go
│   │   │   ├── handler.go
│   │   │   ├── helper.go
│   │   │   └── service.go
│   │   ├── job
│   │   │   ├── domain.go
│   │   │   ├── handler.go
│   │   │   ├── helper.go
│   │   │   └── service.go
│   │   ├── sector
│   │   │   ├── domain.go
│   │   │   ├── handler.go
│   │   │   ├── helper.go
│   │   │   └── service.go
│   │   └── worker
│   │   │   ├── domain.go
│   │   │   ├── handler.go
│   │   │   ├── helper.go
│   │   │   └── service.go
│   │   ├── dependencies.go
│   │   └── router.go
│   │
│   ├── pkg
│   │   ├── apperrors
│   │   │   └── errors.go
│   │   ├── logger
│   │   │   └── logger.go
│   │   ├── middleware
│   │   │   ├── jwt.go
│   │   │   └── middleware.go
│   │   └── utils
│   │       ├── bcrypt.go
│   │       └── userValidation.go
│   │
│   └── repo
│       ├── address.go
│       ├── application.go
│       ├── auth.go
│       ├── base.go
│       ├── domain.go
│       ├── employer.go
│       ├── helpers.go
│       ├── job.go
│       ├── sectors.go
│       └── worker.go
│
├── docs
│   ├── RozgarLink Documentation.docx
│   └── sampleDoc.txt
│
├── config.go
├── go.mod
├── go.sum
├── logfile
└── README.md
```
