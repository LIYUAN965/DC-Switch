package domain

import (
	"github.com/sirupsen/logrus"
	"time"
)

// 业务域

// SwitchVersion 切换演练版本
type SwitchVersion struct {
	Id   int64
	Name string
	Time time.Time
}

func (s SwitchVersion) GetAll(repo SwitchVersionRepo) ([]SwitchVersion, error) {
	return repo.GetAll()
}

func (s SwitchVersion) Add(repo SwitchVersionRepo) error {
	_, err := repo.Add(s)
	return err
}

func (s SwitchVersion) EditName(repo SwitchVersionRepo, id int64, name string) (int64, error) {
	return repo.EditName(id, name)
}

type BizDomain struct {
	domainID int
	name     string
}

// DutyUser 值班人员信息
type DutyUser struct {
	Version SwitchVersion
	User
}

// User 用户信息， 域数据权限校验
type User struct {
	SsoUm     string
	UserName  string
	BizDomain string
}

// SwitchModule 切换模块
type SwitchModule struct {
	Version      SwitchVersion
	ProgressRate float64
	ModuleName   string
	BizDomain    string
	Details      []ModuleDetail
	File         string // 切换模块明细附件
}

// SwitchPreparation 切换准备项
type SwitchPreparation struct {
	Version   SwitchVersion
	BizDomain string
	Name      string
	Status    string // 准备完成、未完成
	Sequence  int    // 顺序，int从大到小，按顺序执行
}

// BreakDown cat故障
type BreakDown struct {
	Version SwitchVersion
	Content string
	File    string // cat截图
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
