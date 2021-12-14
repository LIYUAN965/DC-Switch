package http

import "dcswitch/internal/domain"

// swagger:parameters GetAllVersionsParam
type GetAllVersionsParam struct{}

// swagger:response GetAllVersionsResp
type GetAllVersionsResp struct {
	// in: body
	Body struct {
		Versions []domain.SwitchVersion `json:"versions"`
	}
}

// swagger:parameters EditVersionName
type EditVersionName struct {
	// Required: true
	// in: path
	Id int64 `json:"id"`
	// Required: true
	// in: body
	Body struct {
		Name string `json:"name"`
	}
}
