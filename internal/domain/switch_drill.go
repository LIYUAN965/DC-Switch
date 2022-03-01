package domain

type Module struct {
	Id         int64
	Module     string
	File       string
	Details    []DetailExcel
}

type DetailExcel struct {
	Detail     []string
}

type Detail struct {
	Id         int64
	Module     Module
	Detail     string
	Status     int64
	BeginTime  string
	EndTime    string
	Duration   string
	Comment    string
}

type ModuleProgress struct {
	Module           string
	BeginTime        *string
	EndTime          *string
	Duration         *string
	ModuleStatus     []ModuleStatus

}

type ModuleStatus struct {
	Status     int64
	Count      int64
}

type ModulesProgressResp struct {
	ModulesProgress     []ModulesProgress
}

type ModulesProgress struct {
	Id          int64
	Name        string
	BeginTime   string
	EndTime     string
	Duration    string
	Count       int64
	Status      string
	Percent     int64
	content     string
}

type DrillRepo interface {
	GetAllModulesByVersionId(versionId int64) ([]Module, error)
	GetAllDetailsByVersionId(versionId int64) ([]Detail, error)
	GetProgressByModuleId(moduleId int64) (ModuleProgress, error)
	GetProgressByVersionId(versionId int64) (ModuleProgress, error)
	GetFailedDetailsByVersionId(versionId int64) ([]Detail, error)
	GetDetailsByModuleId(moduleId int64) ([]Detail, error)
	EditDetailCommentById(id int64, comment string) (int64, error)
	EditDetailEndTimeById(id int64, endTime string) (int64, error)
}

func (m Module) GetAllModulesByVersionId(repo DrillRepo, versionId int64) ([]Module, error) {
	return repo.GetAllModulesByVersionId(versionId)
}

func (m Module) GetAllDetailsByVersionId(repo DrillRepo, versionId int64) ([]Detail, error) {
	return repo.GetAllDetailsByVersionId(versionId)
}

func (m Module) GetProgressByModuleId(repo DrillRepo, moduleId int64) (ModuleProgress, error) {
	return repo.GetProgressByModuleId(moduleId)
}

func (m Module) GetProgressByVersionId(repo DrillRepo, versionId int64) (ModuleProgress, error) {
	return repo.GetProgressByVersionId(versionId)
}

func (m Module) GetFailedDetailsByVersionId(repo DrillRepo, versionId int64) ([]Detail, error) {
	return repo.GetFailedDetailsByVersionId(versionId)
}

func (m Module) GetDetailsByModuleId(repo DrillRepo, moduleId int64) ([]Detail, error) {
	return repo.GetDetailsByModuleId(moduleId)
}

func (m Module) EditDetailCommentById(repo DrillRepo, id int64, comment string) (int64, error) {
	return repo.EditDetailCommentById(id, comment)
}

func (m Module) EditDetailEndTimeById(repo DrillRepo, id int64, endTime string) (int64, error) {
	return repo.EditDetailEndTimeById(id, endTime)
}
