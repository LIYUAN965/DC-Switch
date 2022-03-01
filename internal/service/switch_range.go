package service

import (
	"dcswitch/internal/domain"
	"github.com/tealeg/xlsx"
)

type SwitchRangeService struct {
	SwRepo domain.RangeRepo
}

func (s SwitchRangeService) AddSwitchRangeByModuleId(moduleId int64, filePath string, switchDetails []string) (int64, error) {
	return s.SwRepo.AddSwitchRangeByModuleId(moduleId, filePath, switchDetails)
}

func (s SwitchRangeService) GetModulesByDomainId(domainId int64) ([]*domain.Module, error) {
	modules, err := s.SwRepo.GetModulesByDomainId(domainId)
	var fileFolder = "/Users/liyuan/LY/projects/uploads/dcswitch/"
	for _, module := range modules {
		var tableData []domain.DetailExcel
		if module.File != "" {
			fullFilePath := fileFolder + module.File
			// 解析excel
			output, err := xlsx.FileToSlice(fullFilePath)
			if err != nil {
				panic(err.Error())
			}

			// 读取每一行数据
			for  _, row := range output[0] {
				var rowData domain.DetailExcel
				// 读取每一列数据
				for _, cell := range row {
					rowData.Detail = append(rowData.Detail, cell)
				}
				tableData = append(tableData, rowData)
				//println(fmt.Sprintf("tableData:%d, %s", len(tableData), tableData))
			}
		}
		module.Details = tableData
		//println(fmt.Sprintf("module.Details:%d, %s", len(module.Details), module.Details))
	}
	//println(fmt.Sprintf("modules:%+v, %+v", modules[0], modules[1]))

	return modules, err
}

func (s SwitchRangeService) GetModuleByModuleId(moduleId int64) (domain.Module, error) {
	module, err := s.SwRepo.GetModuleByModuleId(moduleId)
	var fileFolder = "/Users/liyuan/LY/projects/uploads/dcswitch/"
	var tableData []domain.DetailExcel
	if module.File != "" {
		fullFilePath := fileFolder + module.File
		// 解析excel
		output, err := xlsx.FileToSlice(fullFilePath)
		if err != nil {
			panic(err.Error())
		}

		// 读取每一行数据
		for  _, row := range output[0] {
			var rowData domain.DetailExcel
			// 读取每一列数据
			for _, cell := range row {
				rowData.Detail = append(rowData.Detail, cell)
			}
			tableData = append(tableData, rowData)
			//println(fmt.Sprintf("tableData:%d, %s", len(tableData), tableData))
		}
	}
	module.Details = tableData

	return module, err
}

func (s SwitchRangeService) GetSwitchDomainById(id int64) (domain.SwitchDomain, error) {
	return s.SwRepo.GetSwitchDomainById(id)
}
