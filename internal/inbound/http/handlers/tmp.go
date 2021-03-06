package handlers

import (
	"dcswitch/internal/config"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// swagger:parameters GetMockSlow
type GetMockSlowParam struct {
	// an id of user info
	//
	// Required: true
	// in: path
	Id int `json:"id"`
}

// swagger:response GetMockSlowResponse
type GetMockSlowResponse struct {
	// GetMockSlow Response
	// in: body
	Body struct {
		Id int `json:"id"`
	}
}

func GetMockSlow(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /{id} GetMockSlow
	//
	// 慢查询模拟
	//
	// Responses:
	//  200: GetMockSlowResponse
	vars := mux.Vars(r)
	code := http.StatusOK
	w.WriteHeader(code)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		logrus.Errorf("参数类型错误 %v\n", vars["id"])
	}
	res := GetMockSlowResponse{}
	res.Body.Id = id
	ret, _ := json.Marshal(res.Body)
	_, err = fmt.Fprintln(w, string(ret))
	time.Sleep(3 * time.Second)
	if err != nil {
		logrus.Error(fmt.Sprintf("Error: %v\n", err))
	}
}

// swagger:response GetConfigResp
type GetConfigResp struct {
	// GetConfig Response
	// in: body
	Body struct {
	}
}

// swagger:parameters GetConfigParam
type GetConfigParam struct {
}

func GetConfig(w http.ResponseWriter, r *http.Request) {
	// swagger:route GET /config GetConfigParam
	//
	// 获取配置数据
	//
	// Responses:
	//  200: GetConfigResp
	code := http.StatusOK
	w.WriteHeader(code)
	conf, err := json.Marshal(config.GlobalConf)
	if err != nil {
		logrus.Error(fmt.Sprintf("Error: %v\n", err))
	}
	_, err = fmt.Fprintln(w, string(conf))
	if err != nil {
		logrus.Error(fmt.Sprintf("Error: %v\n", err))
	}
}
