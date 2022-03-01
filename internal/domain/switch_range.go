package domain

// SwitchRange 切换范围的文件路径，切换明细
type SwitchRange struct {
	FilePath      string
	SwitchDetails []string
}

// SwitchDomain 切换域
type SwitchDomain struct {
	Id             int64
	AppDomain      string
}

type RangeRepo interface {
	AddSwitchRangeByModuleId(moduleId int64, filePath string, switchDetails []string) (int64, error)
	GetModulesByDomainId(domainId int64) ([]*Module, error)
	GetModuleByModuleId(moduleId int64) (Module, error)
	GetSwitchDomainById(id int64) (SwitchDomain, error)
}

func (m Module) AddSwitchRangeByModuleId(repo RangeRepo, moduleId int64, filePath string, switchDetails []string) (int64, error) {
	return repo.AddSwitchRangeByModuleId(moduleId, filePath, switchDetails)
}

func (m Module) GetModulesByDomainId(repo RangeRepo, domainId int64) ([]*Module, error) {
	return repo.GetModulesByDomainId(domainId)
}

func (m Module) GetModuleByModuleId(repo RangeRepo, moduleId int64) (Module, error) {
	return repo.GetModuleByModuleId(moduleId)
}

func (m Module) GetSwitchDomainById(repo RangeRepo, id int64) (SwitchDomain, error) {
	return repo.GetSwitchDomainById(id)
}
