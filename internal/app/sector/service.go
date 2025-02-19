package sector

import (
	"context"

	"github.com/harsh-jagtap-josh/RozgarLink/internal/repo"
)

type sectorService struct {
	sectorRepo repo.SectoreStorer
}

type Service interface {
	CreateNewSector(ctx context.Context, sectorData Sector) (Sector, error)
	FetchSectorById(ctx context.Context, sectorId int) (Sector, error)
	UpdateSectorById(ctx context.Context, sectorData Sector) (Sector, error)
	DeleteSectorById(ctx context.Context, sectorId int) (int, error)
	FetchAllSectors(ctx context.Context) ([]Sector, error)
}

func NewService(sectorRepo repo.SectoreStorer) Service {
	return &sectorService{
		sectorRepo: sectorRepo,
	}
}

func (sectorS *sectorService) CreateNewSector(ctx context.Context, sectorData Sector) (Sector, error) {
	sectorRepoObj := MapSectorServiceToRepo(sectorData)

	createdSector, err := sectorS.sectorRepo.CreateNewSector(ctx, sectorRepoObj)
	if err != nil {
		return Sector{}, err
	}

	sector := MapSectorRepoToService(createdSector)

	return sector, nil
}

func (sectorS *sectorService) FetchSectorById(ctx context.Context, sectorId int) (Sector, error) {

	sector, err := sectorS.sectorRepo.FetchSectorById(ctx, sectorId)
	if err != nil {
		return Sector{}, err
	}

	fetchedSector := MapSectorRepoToService(sector)
	return fetchedSector, nil
}

func (sectorS *sectorService) UpdateSectorById(ctx context.Context, sectorData Sector) (Sector, error) {
	sectorRepoObj := MapSectorServiceToRepo(sectorData)

	updatedSector, err := sectorS.sectorRepo.UpdateSectorById(ctx, sectorRepoObj)
	if err != nil {
		return Sector{}, err
	}

	sector := MapSectorRepoToService(updatedSector)

	return sector, nil
}

func (sectorS *sectorService) DeleteSectorById(ctx context.Context, sectorId int) (int, error) {
	id, err := sectorS.sectorRepo.DeleteSectorById(ctx, sectorId)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (sectorS *sectorService) FetchAllSectors(ctx context.Context) ([]Sector, error) {
	var sectors []Sector
	repoSectors, err := sectorS.sectorRepo.FetchAllSectors(ctx)
	if err != nil {
		return []Sector{}, err
	}
	for _, val := range repoSectors {
		sectors = append(sectors, MapSectorRepoToService(val))
	}

	return sectors, nil
}
