package service

import (
	"dcswitch/internal/domain"
)

type SwitchDrillService struct {
	SwRepo domain.DrillRepo
}

func (s SwitchDrillService) GetAllModulesByVersionId(versionId int64) ([]domain.Module, error) {
	return s.SwRepo.GetAllModulesByVersionId(versionId)
}

func (s SwitchDrillService) GetAllDetailsByVersionId(versionId int64) ([]domain.Detail, error) {
	return s.SwRepo.GetAllDetailsByVersionId(versionId)
}

func (s SwitchDrillService) GetProgressByModuleId(moduleId int64) (domain.ModuleProgress, error) {
	return s.SwRepo.GetProgressByModuleId(moduleId)
}

func (s SwitchDrillService) GetProgressByVersionId(versionId int64) (domain.ModuleProgress, error) {
	return s.SwRepo.GetProgressByVersionId(versionId)
}

func (s SwitchDrillService) GetFailedDetailsByVersionId(versionId int64) ([]domain.Detail, error) {
	return s.SwRepo.GetFailedDetailsByVersionId(versionId)
}

func (s SwitchDrillService) GetDetailsByModuleId(moduleId int64) ([]domain.Detail, error) {
	return s.SwRepo.GetDetailsByModuleId(moduleId)
}

func (s SwitchDrillService) EditDetailCommentById(id int64, comment string) (int64, error) {
	return s.SwRepo.EditDetailCommentById(id, comment)
}

func (s SwitchDrillService) EditDetailEndTimeById(id int64, endTime string) (int64, error) {
	return s.SwRepo.EditDetailEndTimeById(id, endTime)
}
