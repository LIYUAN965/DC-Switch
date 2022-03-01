package payloads

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

// swagger:response GetVersionResp
type GetVersionResp struct {
	// in: body
	Body struct {
		Version domain.SwitchVersion `json:"version"`
	}
}

// swagger:parameters GetVersionParam
type GetVersionParam struct {
	// Required: true
	// in: path
	Id int64 `json:"id"`
}

// swagger:parameters EditVersionDate
type EditVersionDate struct {
	// Required: true
	// in: path
	Id int64 `json:"id"`
	// Required: true
	// in: body
	Body struct {
		Date string `json:"date"`
	}
}

// swagger:parameters ModuleDetailTaskParam
type ModuleDetailTaskParam struct {
	// Required: true
	// in: body
	// example: {"type": "start", "name": "DBASwitch1"}
	Body struct {
		Name string `json:"name"`
		// Required: trues
		// pattern: start|success|fail
		Type string `json:"type"` // start/success
	}
}

// swagger:response ModuleDetailTaskResp
type ModuleDetailTaskResp struct {
	// in: body
	Body struct {
		Result string `json:"result"`
	}
}

func (p ModuleDetailTaskParam) CheckParam() bool {
	checkPass := false
	typeOptions := []string{"start", "success", "fail"}
	for _, v := range typeOptions {
		if p.Body.Type == v {
			checkPass = true
			return checkPass
		}
	}
	return checkPass
}
