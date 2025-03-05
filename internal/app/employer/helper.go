package employer

import "github.com/harsh-jagtap-josh/RozgarLink/internal/repo"

func MapRepoToServiceDomain(employer repo.Employer) Employer {
	return Employer{
		ID:        employer.ID,
		Name:      employer.Name,
		ContactNo: employer.ContactNo,
		Email:     employer.Email,
		Type:      EmployerType(employer.Type),
		Sectors:   employer.Sectors,
		Location: Address{
			ID:      employer.Location,
			Details: employer.Details,
			Street:  employer.Street,
			City:    employer.City,
			State:   employer.State,
			Pincode: employer.Pincode,
		},
		IsVerified:   employer.IsVerified,
		Rating:       employer.Rating,
		WorkersHired: employer.WorkersHired,
		CreatedAt:    employer.CreatedAt,
		UpdatedAt:    employer.UpdatedAt,
		Language:     employer.Language,
	}
}

func MapServiceToRepoDomain(employer Employer) repo.Employer {
	return repo.Employer{
		ID:           employer.ID,
		Name:         employer.Name,
		ContactNo:    employer.ContactNo,
		Email:        employer.Email,
		Type:         repo.EmployerType(employer.Type),
		Sectors:      employer.Sectors,
		Password:     employer.Password,
		Location:     employer.Location.ID,
		IsVerified:   employer.IsVerified,
		Rating:       employer.Rating,
		WorkersHired: employer.WorkersHired,
		CreatedAt:    employer.CreatedAt,
		UpdatedAt:    employer.UpdatedAt,
		Language:     employer.Language,
		Details:      employer.Location.Details,
		Street:       employer.Location.Street,
		City:         employer.Location.City,
		State:        employer.Location.State,
		Pincode:      employer.Location.Pincode,
	}
}
