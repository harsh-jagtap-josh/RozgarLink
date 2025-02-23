package sector

import "github.com/harsh-jagtap-josh/RozgarLink/internal/repo"

func MapSectorRepoToService(sector repo.Sector) Sector {
	return Sector{
		ID:          sector.ID,
		Name:        sector.Name,
		Description: sector.Description,
	}
}

func MapSectorServiceToRepo(sector Sector) repo.Sector {
	return repo.Sector{
		ID:          sector.ID,
		Name:        sector.Name,
		Description: sector.Description,
	}
}
