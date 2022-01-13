package domain

//version_domain 切换域
type Domain struct {
	Id              int64
	Version         SwitchVersion
	Type            string
	AppDomain       string
	DomainArchitect string
	DelFlag         int64
}

//version_domain_prepare 切换准备项
type Prepare struct {
	Id         int64
	Domain     Domain
	Title      string
	Content    string
	Status     int64
	CreateUser string
	CreateTime string
	UpdateUser string
	UpdateTime string
	DelFlag    int64
	Sequence   int64
}

type PrepareRepo interface {
	GetAllPreparesByVersionId(id int64) ([]Prepare, error)
}

func (s Prepare) GetAllPreparesByVersionId(repo PrepareRepo, versionId int64) ([]Prepare, error) {
	return repo.GetAllPreparesByVersionId(versionId)
}