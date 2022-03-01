package handlers

import (
	"dcswitch/internal/inbound/http/payloads"
	"dcswitch/internal/outbound/db"
	"dcswitch/internal/service"
	"dcswitch/pkg/mysql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetAllSwitchModules(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/modules/{versionId} GetAllModulesParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetAllModulesResp
	//  500:
	vars := mux.Vars(r)
	versionId, err := strconv.ParseInt(vars["versionId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.DrillModuleDBRepo{}
	svs := service.SwitchDrillService{SwRepo: repo}
	vResp := payloads.GetAllModulesResp{}
	modules, err := svs.GetAllModulesByVersionId(versionId)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Modules = modules

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

func GetAllSwitchDetails(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/details/{versionId} GetAllDetailsParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetAllDetailsResp
	//  500:
	vars := mux.Vars(r)
	versionId, err := strconv.ParseInt(vars["versionId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.DrillModuleDBRepo{}
	svs := service.SwitchDrillService{SwRepo: repo}
	vResp := payloads.GetAllDetailsResp{}
	details, err := svs.GetAllDetailsByVersionId(versionId)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Details = details

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

func GetProgressByModuleId(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/module/progress/{moduleId} GetModuleProgressParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetModuleProgressResp
	//  500:
	vars := mux.Vars(r)
	moduleId, err := strconv.ParseInt(vars["moduleId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.DrillModuleDBRepo{}
	svs := service.SwitchDrillService{SwRepo: repo}
	vResp := payloads.GetModuleProgressResp{}
	progress, err := svs.GetProgressByModuleId(moduleId)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Progress = progress

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

func GetProgressByVersionId(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/version/progress/{versionId} GetVersionProgressParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetVersionProgressResp
	//  500:
	vars := mux.Vars(r)
	versionId, err := strconv.ParseInt(vars["versionId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.DrillModuleDBRepo{}
	svs := service.SwitchDrillService{SwRepo: repo}
	vResp := payloads.GetVersionProgressResp{}
	progress, err := svs.GetProgressByVersionId(versionId)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Progress = progress

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

func GetFailedDetailsByVersionId(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/failed/details/{versionId} GetModuleDetailsParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetModuleDetailsResp
	//  500:
	vars := mux.Vars(r)
	versionId, err := strconv.ParseInt(vars["versionId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.DrillModuleDBRepo{}
	svs := service.SwitchDrillService{SwRepo: repo}
	vResp := payloads.GetModuleDetailsResp{}
	details, err := svs.GetFailedDetailsByVersionId(versionId)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Details = details

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

func GetDetailsByModuleId(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/failed/details/{moduleId} GetModuleDetailsParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetModuleDetailsResp
	//  500:
	vars := mux.Vars(r)
	moduleId, err := strconv.ParseInt(vars["moduleId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.DrillModuleDBRepo{}
	svs := service.SwitchDrillService{SwRepo: repo}
	vResp := payloads.GetModuleDetailsResp{}
	details, err := svs.GetDetailsByModuleId(moduleId)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Details = details

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

func EditDetailCommentById(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/detail/editComment/{id} EditSwitchDetailParam
	//
	// 编辑准备项
	//
	// Responses:
	//  200:
	//  404:
	//  500:
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}
	aaa := r.URL.Query()
	defer r.Body.Close()
	//body, _ := ioutil.ReadAll(r.Body)
	sv := payloads.EditSwitchDetailParam{}
	//err = json.Unmarshal(body, &sv.Body)
	//if err != nil {
	//	setReqBodyError(err, string(body), w)
	//	return
	//}
	sv.Body.Comment = aaa.Get("comment")

	repo := db.DrillModuleDBRepo{}
	svs := service.SwitchDrillService{SwRepo: repo}
	//_, err = svs.EditPrepareById(id, aaa.Get("title"), aaa.Get("content"))
	_, err = svs.EditDetailCommentById(id, sv.Body.Comment)
	switch err.(type) {
	case mysql.NotFoundError:
		setNotFound(err, w)
		return
	}
	if err != nil {
		setServerError(err, w)
		return
	}
	setGetSuccess(w)
}

func EditDetailEndTimeById(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/detail/editEndTIme/{id} EditSwitchDetailParam
	//
	// 编辑准备项
	//
	// Responses:
	//  200:
	//  404:
	//  500:
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}
	aaa := r.URL.Query()
	defer r.Body.Close()
	sv := payloads.EditSwitchDetailParam{}
	sv.Body.EndTime = aaa.Get("endTime")

	repo := db.DrillModuleDBRepo{}
	svs := service.SwitchDrillService{SwRepo: repo}
	_, err = svs.EditDetailEndTimeById(id, sv.Body.EndTime)
	switch err.(type) {
	case mysql.NotFoundError:
		setNotFound(err, w)
		return
	}
	if err != nil {
		setServerError(err, w)
		return
	}
	setGetSuccess(w)
}
