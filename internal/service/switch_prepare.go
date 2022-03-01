package service

import (
	"dcswitch/internal/domain"
)

type SwitchPrepareService struct {
	SwRepo domain.PrepareRepo
}

func (s SwitchPrepareService) GetAllPreparesByVersionId(versionId int64) ([]domain.Prepare, error) {
	return s.SwRepo.GetAllPreparesByVersionId(versionId)
}

func (s SwitchPrepareService) GetPrepareById(id int64) (domain.Prepare, error) {
	return s.SwRepo.GetPrepareById(id)
}

func (s SwitchPrepareService) EditPrepareById(id int64, title string, content string) (int64, error) {
	return s.SwRepo.EditPrepareById(id, title, content)
}



//func (s SwitchPrepareService) Add(domainId int64, title string) error {
//	v := domain.SwitchPrepare{VersionDate: versionDate}
//	return v.Add(s.SwRepo)
//}
//
//func (s SwitchPrepareService) DeleteById(id int64) (int64, error) {
//	return s.SwRepo.DeleteById(id)
//}
//
//func (s SwitchPrepareService) EditTitleById(id int64, title string) (int64, error) {
//	return s.SwRepo.EditTitleById(id, title)
//}
//
//func (s SwitchPrepareService) GetAllByDomainId(domainId int64) (int64, error) {
//	return s.SwRepo.GetAllByDomainId(id)
//}
