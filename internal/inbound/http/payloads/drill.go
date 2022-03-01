package payloads

import "dcswitch/internal/domain"

// swagger:parameters GetAllModulesParam
type GetAllModulesParam struct{
	// Required: true
	// in: path
	VersionId int64 `json:"versionId"`
}

// swagger:response GetAllModulesResp
type GetAllModulesResp struct {
	// in: body
	Body struct {
		Modules []domain.Module `json:"modules"`
	}
}

// swagger:parameters GetAllDetailsParam
type GetAllDetailsParam struct{
	// Required: true
	// in: path
	VersionId int64 `json:"versionId"`
}

// swagger:response GetAllDetailsResp
type GetAllDetailsResp struct {
	// in: body
	Body struct {
		Details []domain.Detail `json:"details"`
	}
}

// swagger:parameters GetModuleProgressParam
type GetModuleProgressParam struct{
	// Required: true
	// in: path
	ModuleId int64 `json:"moduleId"`
}

// swagger:response GetModuleProgressResp
type GetModuleProgressResp struct {
	// in: body
	Body struct {
		Progress domain.ModuleProgress `json:"progress"`
	}
}

// swagger:parameters GetVersionProgressParam
type GetVersionProgressParam struct{
	// Required: true
	// in: path
	VersionId int64 `json:"versionId"`
}

// swagger:response GetVersionProgressResp
type GetVersionProgressResp struct {
	// in: body
	Body struct {
		Progress domain.ModuleProgress `json:"progress"`
	}
}

// swagger:parameters GetModulesProgressParam
type GetModulesProgressParam struct{
	// Required: true
	// in: path
	VersionId int64 `json:"versionId"`
}

// swagger:response GetModulesProgressResp
type GetModulesProgressResp struct {
	// in: body
	Body struct {
		Progress domain.ModuleProgress `json:"progress"`
	}
}

// swagger:parameters GetModuleDetailsParam
type GetModuleDetailsParam struct{
	// Required: true
	// in: path
	ModuleId int64 `json:"moduleId"`
}

// swagger:response GetModuleDetailsResp
type GetModuleDetailsResp struct {
	// in: body
	Body struct {
		Details []domain.Detail `json:"details"`
	}
}
