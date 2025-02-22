package job

import (
	"net/url"
	"strconv"
	"time"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/app/worker"
	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

func MapJobRepoStructToService(job repo.Job) Job {
	return Job{
		ID:              job.ID,
		EmployerID:      job.EmployerID,
		Title:           job.Title,
		RequiredGender:  job.RequiredGender,
		Description:     job.Description,
		DurationInHours: job.DurationInHours,
		SkillsRequired:  job.SkillsRequired,
		Sectors:         job.Sectors,
		Wage:            job.Wage,
		Vacancy:         job.Vacancy,
		Location: worker.Address{
			ID:      job.Location,
			Details: job.Details,
			Street:  job.Street,
			City:    job.City,
			State:   job.State,
			Pincode: job.Pincode,
		},
		Date:      job.Date,
		StartHour: job.StartHour,
		EndHour:   job.EndHour,
		CreatedAt: job.CreatedAt,
		UpdatedAt: job.UpdatedAt,
	}
}

func MapJobServiceStructToRepo(job Job) repo.Job {
	return repo.Job{
		ID:              job.ID,
		EmployerID:      job.EmployerID,
		Title:           job.Title,
		RequiredGender:  job.RequiredGender,
		Description:     job.Description,
		DurationInHours: job.DurationInHours,
		SkillsRequired:  job.SkillsRequired,
		Sectors:         job.Sectors,
		Wage:            job.Wage,
		Vacancy:         job.Vacancy,
		Location:        job.Location.ID,
		Date:            job.Date,
		StartHour:       job.StartHour,
		EndHour:         job.EndHour,
		CreatedAt:       job.CreatedAt,
		UpdatedAt:       job.UpdatedAt,
		Details:         job.Location.Details,
		Street:          job.Location.Street,
		City:            job.Location.City,
		State:           job.Location.State,
		Pincode:         job.Location.Pincode,
	}
}

func MapJobFilterServiceToRepo(filters JobFilters) repo.JobFilters {
	return repo.JobFilters{
		Title:     filters.Title,
		Sector:    filters.Sector,
		WageMin:   filters.WageMin,
		WageMax:   filters.WageMax,
		StartDate: filters.StartDate,
		EndDate:   filters.EndDate,
		City:      filters.City,
		Gender:    filters.Gender,
	}
}

func retrieveQueryParams(queryParams url.Values) JobFilters {
	jobFilters := JobFilters{}

	title := queryParams.Get("title")
	sector := queryParams.Get("sector")
	wageMin := queryParams.Get("wage_min")
	wageMax := queryParams.Get("wage_max")
	startDate := queryParams.Get("start_date")
	endDate := queryParams.Get("end_date")
	city := queryParams.Get("city")
	gender := queryParams.Get("required_gender")

	if title != "" {
		jobFilters.Title = title
	}
	if sector != "" {
		jobFilters.Sector = sector
	}
	if wageMin != "" {
		if min, err := strconv.Atoi(wageMin); err == nil {
			jobFilters.WageMin = min
		}
	}
	if wageMax != "" {
		if max, err := strconv.Atoi(wageMax); err == nil {
			jobFilters.WageMax = max
		}
	}
	if startDate != "" {
		if parsed, err := time.Parse("2006-01-02", startDate); err == nil {
			jobFilters.StartDate = parsed
		}
	}
	if endDate != "" {
		if parsed, err := time.Parse("2006-01-02", endDate); err == nil {
			jobFilters.EndDate = parsed
		}
	}
	if city != "" {
		jobFilters.City = city
	}
	if gender != "" {
		jobFilters.Gender = gender
	}

	return jobFilters
}
