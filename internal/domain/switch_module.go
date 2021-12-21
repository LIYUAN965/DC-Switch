package domain

import (
	"github.com/sirupsen/logrus"
	"time"
)

// SwitchModule 切换模块
type SwitchModule struct {
	Version      SwitchVersion
	ProgressRate float64
	ModuleName   string
	BizDomain    string
	Details      []ModuleDetail
	File         string // 切换模块明细附件
}

// Progress 模块明细进度
type Progress struct {
	Status    string // success, fail
	StartTime time.Time
	EndTime   time.Time
}

// ModuleDetailFailInfo 模块明细失败原因
type ModuleDetailFailInfo struct {
	Content  string
	ExecUser DutyUser
	Affect   string
}

type ModuleDetail struct {
	Progress
	Name     string
	FailInfo ModuleDetailFailInfo
}

func (m *ModuleDetail) Start() {
	m.StartTime = time.Now()
	logrus.Infof("%v start time: %v\n", m.Name, m.StartTime)
}

func (m *ModuleDetail) End() {
	m.EndTime = time.Now()
	logrus.Infof("%v end time: %v\n", m.Name, m.StartTime)
}

func (m *ModuleDetail) Fail(user DutyUser, affect string, content string) {
	m.FailInfo.ExecUser = user
	m.FailInfo.Affect = affect
	m.FailInfo.Content = content
}
