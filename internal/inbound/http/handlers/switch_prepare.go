package handlers

import (
	"dcswitch/internal/inbound/http/payloads"
	"dcswitch/internal/outbound/db"
	"dcswitch/internal/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllSwitchPreparations(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/preparations/{versionId} GetAllPreparesParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetAllPreparesResp
	//  500:
	vars := mux.Vars(r)
	versionId, err := strconv.ParseInt(vars["versionId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.PrepareDBRepo{}
	svs := service.SwitchPrepareService{SwRepo: repo}
	vResp := payloads.GetAllPreparesResp{}
	prepares, err := svs.GetAllPreparesByVersionId(versionId)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Prepares = prepares

	resp, err := json.Marshal(vResp.Body)
	if err != nil {
		setServerError(err, w)
		return
	}
	_, err = fmt.Fprintf(w, string(resp))
	if err != nil {
		setServerError(err, w)
		return
	}
	setGetSuccess(w)
}
