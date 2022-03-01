package handlers

import (
	"dcswitch/internal/inbound/http/payloads"
	"dcswitch/internal/outbound/db"
	"dcswitch/internal/service"
	"dcswitch/pkg/mysql"
	"dcswitch/pkg/upload"
	"encoding/json"

	//"dcswitch/pkg/upload"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/tealeg/xlsx"
	//"log"
	"net/http"
	"strconv"
)


func SwitchRangeUpload(w http.ResponseWriter, r *http.Request) {
	// 接受文件
	file, header, err := r.FormFile("file")
	if err != nil {
		// ignore the error handler
	}

	//filePath := "20220224140952/DB.xlsx"
	//fullFilePath := "/Users/liyuan/LY/projects/uploads/dcswitch/20220224140952/DB.xlsx"
	// 上传文件
	filePath, fullFilePath := upload.Upload(file, header.Filename)
	//log.Printf("uploaded file path is: %s", filePath)

	// 解析excel
	output, err := xlsx.FileToSlice(fullFilePath)
	if err != nil {
		panic(err.Error())
	}

	// 读取excel中的第一列切换明细
	var switchDetails []string
	for  _, row := range output[0] {
        fmt.Println(row[0])
		switchDetails = append(switchDetails, row[0])
	}

	vars := mux.Vars(r)
	moduleId, err := strconv.ParseInt(vars["moduleId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}
	repo := db.RangeDBRepo{}
	svs := service.SwitchRangeService{SwRepo: repo}
	_, err = svs.AddSwitchRangeByModuleId(moduleId, filePath, switchDetails)
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

func GetModulesByDomainId(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/modules/domain/{domainId} GetModulesParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetModulesResp
	//  500:
	vars := mux.Vars(r)
	domainId, err := strconv.ParseInt(vars["domainId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.RangeDBRepo{}
	svs := service.SwitchRangeService{SwRepo: repo}
	vResp := payloads.GetModulesResp{}
	modules, err := svs.GetModulesByDomainId(domainId)
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

func GetModuleByModuleId(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/module/domain/{moduleId} GetModuleParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetModuleResp
	//  500:
	vars := mux.Vars(r)
	moduleId, err := strconv.ParseInt(vars["moduleId"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.RangeDBRepo{}
	svs := service.SwitchRangeService{SwRepo: repo}
	vResp := payloads.GetModuleResp{}
	module, err := svs.GetModuleByModuleId(moduleId)
	if err != nil {
		setServerError(err, w)
		return
	}

	vResp.Body.Module = module
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

func GetSwitchDomainById(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /switch/domain/{id} GetSwitchDomainParam
	//
	// 获取所有切换版本信息
	//
	// Responses:
	//  200: GetSwitchDomainResp
	//  500:
	vars := mux.Vars(r)
	domainId, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		setParamsError(err, vars, w)
		return
	}

	repo := db.RangeDBRepo{}
	svs := service.SwitchRangeService{SwRepo: repo}
	vResp := payloads.GetSwitchDomainResp{}
	switchDomain, err := svs.GetSwitchDomainById(domainId)
	if err != nil {
		setServerError(err, w)
		return
	}
	vResp.Body.SwitchDomain = switchDomain
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
