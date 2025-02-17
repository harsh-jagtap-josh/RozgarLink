package employer

import "github.com/harsh-jagtap-josh/RozgarLink/internal/repo"

func MapRepoToServiceDomain(employer repo.EmployerResponse) Employer {
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
