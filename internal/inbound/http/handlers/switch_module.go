package handlers

import (
	"dcswitch/internal/inbound/http/payloads"
	"dcswitch/internal/outbound/db"
	"dcswitch/internal/service"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func CreateModuleDetailTask(w http.ResponseWriter, r *http.Request) {
	// swagger:route POST /task/switch/module/detail ModuleDetailTaskParam
	//
	// 模块明细 start/success/fail
	//
	// Responses:
	//  200:ModuleDetailTaskResp
	//  403:
	//  500:
	defer r.Body.Close()
	body, _ := ioutil.ReadAll(r.Body)
	md := payloads.ModuleDetailTaskParam{}
	err := json.Unmarshal(body, &md.Body)
	if err != nil {
		setReqBodyError(err, string(body), w)
		return
	}
	checkPass := md.CheckParam()
	if !checkPass {
		setReqBodyError(err, string(body), w)
		return
	}
	// TODO: REPO未写具体逻辑
	repo := db.ModuleDetailDBRepo{}
	svc := service.SwitchModuleSvc{SmRepo: repo}
	vResp := payloads.ModuleDetailTaskResp{}
	vResp.Body.Result = "failed"
	///start|success|fail/
	switch md.Body.Type {
	case "start":
		err = svc.Start(md.Body.Name)
	case "success":
		err = svc.Success(md.Body.Name)
	case "fail":
		err = svc.Fail(md.Body.Name)
	default:
		setServerError(fmt.Errorf("type must in (start|success|fail)"), w)
		return
	}
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.Result = "success"

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
	setPostSuccess(w)
}
