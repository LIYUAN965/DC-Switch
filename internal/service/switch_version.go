package service

import (
	"dcswitch/internal/domain"
)

type SwitchVersionService struct {
	SwRepo domain.SwitchVersionRepo
}

func (s SwitchVersionService) GetAll() ([]domain.SwitchVersion, error) {
	sw := domain.SwitchVersion{}
	return sw.GetAll(s.SwRepo)
}

func (s SwitchVersionService) Add(versionDate string) error {
	v := domain.SwitchVersion{VersionDate: versionDate}
	return v.Add(s.SwRepo)
}

func (s SwitchVersionService) EditVersionDateById(id int64, versionDate string) (int64, error) {
	return s.SwRepo.EditVersionDateById(id, versionDate)
}

func (s SwitchVersionService) GetVersionById(id int64) (domain.SwitchVersion, error) {
	return s.SwRepo.GetVersionById(id)
}
