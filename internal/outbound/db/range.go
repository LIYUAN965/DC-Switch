package db

import (
	"dcswitch/internal/domain"
	"dcswitch/pkg/mysql"
	log "github.com/sirupsen/logrus"
)

type RangeDBRepo struct{}

func (repo RangeDBRepo) AddSwitchRangeByModuleId(moduleId int64, filePath string, switchDetails []string) (int64, error) {
	moduleRepo := DrillModuleDBRepo{}
	module_exists, _ := moduleRepo.ModuleCheckExist(moduleId)
	if !module_exists {
		return -1, mysql.NotFoundError{}
	}

	db := mysql.DB.GetConn()
	sql := "UPDATE switch_module SET file = ? WHERE id = ?"
	_, err := db.Exec(sql, filePath, moduleId)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"moduleId":  moduleId,
			"filePath":  filePath,
		}).Error(err.Error())
		return -1, err
	}

	details_exists, _ := moduleRepo.ModuleDetailsCheckExist(moduleId)
	if details_exists {
		sql := "UPDATE switch_module_detail SET del_flag = 1 WHERE module_id = ?"
		_, err := db.Exec(sql, moduleId)
		if err != nil {
			log.WithFields(log.Fields{
				"sql": sql,
				"moduleId":  moduleId,
			}).Error(err.Error())
			return -1, err
		}
	}
	for _, switchDetail := range switchDetails {
		sql := "INSERT INTO switch_module_detail (module_id, detail) VALUES (?, ?)"
		_, err := db.Exec(sql, moduleId, switchDetail)
		if err != nil {
			log.WithFields(log.Fields{
				"sql":          sql,
				"moduleId": moduleId,
				"switchDetail": switchDetail,
			}).Error(err.Error())
			return -1, err
		}
	}
	return 1, nil
}

func (repo RangeDBRepo) DomainCheckExist(domainId int64) (bool, error) {
	sql := "SELECT id FROM version_domain WHERE id = ? AND del_flag = 0"
	exists, err := mysql.DB.RowExists(sql, domainId)
	if !exists {
		return false, mysql.NotFoundError{}
	}
	return exists, err
}

func (repo RangeDBRepo) GetModulesByDomainId(domainId int64) ([]*domain.Module, error) {
	var modules []*domain.Module
	rangeRepo := RangeDBRepo{}
	domain_exists, _ := rangeRepo.DomainCheckExist(domainId)
	if !domain_exists {
		return modules, mysql.NotFoundError{}
	}

	db := mysql.DB.GetConn()
	sql := "SELECT id, module, file FROM switch_module WHERE switch_id in (SELECT id FROM version_domain_switch WHERE version_domain_id=? AND del_flag=0)"
	rows, err := db.Query(sql, domainId)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"domainId":  domainId,
		}).Error(err.Error())
		return modules, err
	}

	defer rows.Close()
	for rows.Next() {
		var module domain.Module
		err := rows.Scan(&module.Id, &module.Module, &module.File)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return modules, err
		}
		modules = append(modules, &module)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return modules, err
	}
	return modules, nil
}

func (repo RangeDBRepo) GetModuleByModuleId(moduleId int64) (domain.Module, error) {
	var module domain.Module
	moduleRepo := DrillModuleDBRepo{}
	module_exists, _ := moduleRepo.ModuleCheckExist(moduleId)
	if !module_exists {
		return module, mysql.NotFoundError{}
	}

	db := mysql.DB.GetConn()
	sql := "SELECT id, module, file FROM switch_module WHERE id=? AND del_flag=0"
	rows, err := db.Query(sql, moduleId)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"moduleId":  moduleId,
		}).Error(err.Error())
		return module, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&module.Id, &module.Module, &module.File)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return module, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return module, err
	}
	return module, nil
}

func (repo RangeDBRepo) GetSwitchDomainById(id int64) (domain.SwitchDomain, error) {
	var switchDomain domain.SwitchDomain
	rangeRepo := RangeDBRepo{}
	domain_exists, _ := rangeRepo.DomainCheckExist(id)
	if !domain_exists {
		return switchDomain, mysql.NotFoundError{}
	}

	db := mysql.DB.GetConn()
	sql := "SELECT app_domain FROM version_domain WHERE id = ? AND del_flag = 0"
	rows, err := db.Query(sql, id)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"domainId":  id,
		}).Error(err.Error())
		return switchDomain, err
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&switchDomain.AppDomain)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return switchDomain, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return switchDomain, err
	}
	return switchDomain, nil
}
