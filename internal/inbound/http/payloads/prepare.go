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

// swagger:parameters GetPrepareParam
type GetPrepareParam struct{
	// Required: true
	// in: path
	Id int64 `json:"id"`
}

// swagger:response GetPrepareResp
type GetPrepareResp struct {
	// in: body
	Body struct {
		Prepare domain.Prepare `json:"prepare"`
	}
}

// swagger:parameters EditPrepareParam
type EditPrepareParam struct {
	// Required: true
	// in: path
	Id int64 `json:"id"`
	// Required: true
	// in: body
	Body struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
}
