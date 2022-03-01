package payloads

import "dcswitch/internal/domain"

// swagger:parameters GetModulesParam
type GetModulesParam struct{
	// Required: true
	// in: path
	DomainId int64 `json:"domainId"`
}

// swagger:response GetModulesResp
type GetModulesResp struct {
	// in: body
	Body struct {
		Modules []*domain.Module `json:"modules"`
	}
}

// swagger:parameters GetModuleParam
type GetModuleParam struct{
	// Required: true
	// in: path
	ModuleId int64 `json:"moduleId"`
}

// swagger:response GetModuleResp
type GetModuleResp struct {
	// in: body
	Body struct {
		Module domain.Module `json:"module"`
	}
}

// swagger:parameters GetSwitchDomainParam
type GetSwitchDomainParam struct{
	// Required: true
	// in: path
	Id int64 `json:"id"`
}

// swagger:response GetSwitchDomainResp
type GetSwitchDomainResp struct {
	// in: body
	Body struct {
		SwitchDomain domain.SwitchDomain `json:"switchDomain"`
	}
}
