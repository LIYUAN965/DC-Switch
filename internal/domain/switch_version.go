package domain

// SwitchVersion 切换演练版本
type SwitchVersion struct {
	Id           int64
	VersionDate  string
	CreateUser   string
	CreateTime   string
}

type SwitchVersionRepo interface {
	Add(version SwitchVersion) (int64, error)
	GetVersionById(id int64) (SwitchVersion, error)
	GetAll() ([]SwitchVersion, error)
	EditVersionDateById(id int64, versionDate string) (int64, error)
}

func (s SwitchVersion) GetAll(repo SwitchVersionRepo) ([]SwitchVersion, error) {
	return repo.GetAll()
}

func (s SwitchVersion) GetVersionById(repo SwitchVersionRepo, id int64) (SwitchVersion, error) {
	return repo.GetVersionById(id)
}

func (s SwitchVersion) Add(repo SwitchVersionRepo) error {
	_, err := repo.Add(s)
	return err
}

func (s SwitchVersion) EditVersionDateById(repo SwitchVersionRepo, id int64, versionDate string) (int64, error) {
	return repo.EditVersionDateById(id, versionDate)
}
