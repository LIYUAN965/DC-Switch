package domain

type SwitchVersionRepo interface {
	Add(version SwitchVersion) (int64, error)
	Get(id int64) (SwitchVersion, error)
	GetAll() ([]SwitchVersion, error)
	EditName(id int64, name string) (int64, error)
}
