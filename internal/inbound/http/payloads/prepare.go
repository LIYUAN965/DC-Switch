package payloads

import "dcswitch/internal/domain"

// swagger:parameters GetAllPreparesParam
type GetAllPreparesParam struct{
	// Required: true
	// in: path
	VersionId int64 `json:"versionId"`
}

// swagger:response GetAllPreparesResp
type GetAllPreparesResp struct {
	// in: body
	Body struct {
		Prepares []domain.Prepare `json:"prepares"`
	}
}
