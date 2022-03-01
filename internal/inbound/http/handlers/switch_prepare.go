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

func GetSwitchPreparationById(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/preparation/{id} GetPrepareParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetPrepareResp
	//  500:
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.PrepareDBRepo{}
	svs := service.SwitchPrepareService{SwRepo: repo}
	vResp := payloads.GetPrepareResp{}
	prepare, err := svs.GetPrepareById(id)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Prepare = prepare

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

func EditSwitchPreparationById(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/preparation/edit/{id} EditPrepareParam
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
	sv := payloads.EditPrepareParam{}
	//err = json.Unmarshal(body, &sv.Body)
	//if err != nil {
	//	setReqBodyError(err, string(body), w)
	//	return
	//}
	sv.Body.Title = aaa.Get("title")
	sv.Body.Content = aaa.Get("content")

	repo := db.PrepareDBRepo{}
	svs := service.SwitchPrepareService{SwRepo: repo}
	//_, err = svs.EditPrepareById(id, aaa.Get("title"), aaa.Get("content"))
	_, err = svs.EditPrepareById(id, sv.Body.Title, sv.Body.Content)
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
