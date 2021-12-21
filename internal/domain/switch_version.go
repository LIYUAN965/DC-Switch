package domain

import "time"

// SwitchVersion 切换演练版本
type SwitchVersion struct {
	Id   int64
	Name string
	Time time.Time
}

func (s SwitchVersion) GetAll(repo SwitchVersionRepo) ([]SwitchVersion, error) {
	return repo.GetAll()
}

func (s SwitchVersion) Add(repo SwitchVersionRepo) error {
	_, err := repo.Add(s)
	return err
}

func (s SwitchVersion) EditName(repo SwitchVersionRepo, id int64, name string) (int64, error) {
	return repo.EditName(id, name)
}
