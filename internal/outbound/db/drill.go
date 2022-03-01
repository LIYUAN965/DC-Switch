package db

import (
	"dcswitch/internal/domain"
	"dcswitch/pkg/mysql"
	log "github.com/sirupsen/logrus"
)

type DrillModuleDBRepo struct{}

func (repo DrillModuleDBRepo) GetAllModulesByVersionId(versionId int64) ([]domain.Module, error) {
	var modules []domain.Module
	sql := "SELECT id, module FROM switch_module WHERE del_flag = 0 AND switch_id IN (SELECT id FROM version_domain_switch WHERE version_domain_id IN (SELECT id FROM version_domain WHERE version_id = ?))"
	versionRepo := SwitchVersionDBRepo{}
	version_exists, _ := versionRepo.CheckExist(versionId)
	if !version_exists {
		return modules, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, versionId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"versionId":  versionId,
		}).Error("Query Error")
		return modules, err
	}
	defer rows.Close()
	for rows.Next() {
		var module domain.Module
		err := rows.Scan(&module.Id, &module.Module)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return modules, err
		}
		modules = append(modules, module)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return modules, err
	}
	return modules, nil
}

func (repo DrillModuleDBRepo) GetAllDetailsByVersionId(versionId int64) ([]domain.Detail, error) {
	var details []domain.Detail
	sql := "SELECT d.id, m.id, m.module, d.detail, d.status from switch_module_detail AS d, switch_module AS m WHERE d.del_flag=0 AND d.module_id=m.id AND d.module_id IN (SELECT id FROM switch_module WHERE del_flag=0 AND switch_id IN (SELECT id FROM version_domain_switch WHERE version_domain_id IN (SELECT id FROM version_domain WHERE version_id=?)))"
	versionRepo := SwitchVersionDBRepo{}
	version_exists, _ := versionRepo.CheckExist(versionId)
	if !version_exists {
		return details, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, versionId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"versionId":  versionId,
		}).Error("Query Error")
		return details, err
	}
	defer rows.Close()
	for rows.Next() {
		var detail domain.Detail
		err := rows.Scan(&detail.Id, &detail.Module.Id, &detail.Module.Module, &detail.Detail, &detail.Status)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return details, err
		}
		details = append(details, detail)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return details, err
	}
	return details, nil
}

func (repo DrillModuleDBRepo) ModuleCheckExist(id int64) (bool, error) {
	sql := "SELECT id from switch_module WHERE id = ? AND del_flag = 0"
	exists, err := mysql.DB.RowExists(sql, id)
	if !exists {
		return false, mysql.NotFoundError{}
	}
	return exists, err
}

func (repo DrillModuleDBRepo) ModuleDetailsCheckExist(moduleId int64) (bool, error) {
	sql := "SELECT id FROM switch_module_detail WHERE module_id = ? AND del_flag = 0"
	exists, err := mysql.DB.RowExists(sql, moduleId)
	if !exists {
		return false, mysql.NotFoundError{}
	}
	return exists, err
}

func (repo DrillModuleDBRepo) DetailCheckExist(id int64) (bool, error) {
	sql := "SELECT id FROM switch_module_detail WHERE id = ? AND del_flag = 0"
	exists, err := mysql.DB.RowExists(sql, id)
	if !exists {
		return false, mysql.NotFoundError{}
	}
	return exists, err
}

func (repo DrillModuleDBRepo) GetProgressByModuleId(moduleId int64) (domain.ModuleProgress, error) {
	var progress domain.ModuleProgress
	sql := "SELECT status, count(*) FROM switch_module_detail WHERE del_flag=0 AND module_id=? GROUP BY status"
	moduleRepo := DrillModuleDBRepo{}
	moduleExists, _ := moduleRepo.ModuleCheckExist(moduleId)
	if !moduleExists {
		return progress, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, moduleId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"moduleId":  moduleId,
		}).Error("Query Error")
		return progress, err
	}
	defer rows.Close()
	for rows.Next() {
		var detail domain.ModuleStatus
		err := rows.Scan(&detail.Status, &detail.Count)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return progress, err
		}
		progress.ModuleStatus = append(progress.ModuleStatus, detail)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return progress, err
	}


	sql = "SELECT MIN(begin_time), MAX(end_time), TIMEDIFF(MAX(end_time), MIN(begin_time)) FROM switch_module_detail WHERE del_flag=0 AND module_id=?"
	rows, err = db.Query(sql, moduleId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"moduleId":  moduleId,
		}).Error("Query Error")
		return progress, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&progress.BeginTime, &progress.EndTime, &progress.Duration)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return progress, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return progress, err
	}

	sql = "SELECT module FROM switch_module WHERE id=?"
	rows, err = db.Query(sql, moduleId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"moduleId":  moduleId,
		}).Error("Query Error")
		return progress, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&progress.Module)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return progress, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return progress, err
	}
	return progress, nil
}

func (repo DrillModuleDBRepo) GetProgressByVersionId(versionId int64) (domain.ModuleProgress, error) {
	var progress domain.ModuleProgress
	sql := "SELECT status, count(*) FROM switch_module_detail WHERE del_flag=0 AND module_id IN (SELECT id FROM switch_module WHERE switch_id IN (SELECT id FROM version_domain_switch WHERE version_domain_id IN (SELECT id FROM version_domain WHERE version_id=?)) ) GROUP BY status"
	versionRepo := SwitchVersionDBRepo{}
	versionExists, _ := versionRepo.CheckExist(versionId)
	if !versionExists {
		return progress, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, versionId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"versionId":  versionId,
		}).Error("Query Error")
		return progress, err
	}
	defer rows.Close()
	for rows.Next() {
		var detail domain.ModuleStatus
		err := rows.Scan(&detail.Status, &detail.Count)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return progress, err
		}
		progress.ModuleStatus = append(progress.ModuleStatus, detail)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return progress, err
	}
	return progress, nil
}

func (repo DrillModuleDBRepo) GetFailedDetailsByVersionId(versionId int64) ([]domain.Detail, error) {
	var details []domain.Detail
	sql := "SELECT m.id, m.module, d.detail, d.begin_time, d.end_time, d.comment FROM switch_module AS m, switch_module_detail AS d WHERE m.switch_id IN (SELECT id FROM version_domain_switch WHERE version_domain_id IN (SELECT id FROM version_domain WHERE version_id=?)) AND m.id=d.module_id AND d.del_flag=0 AND d.status=3 ORDER BY d.end_time"
	versionRepo := SwitchVersionDBRepo{}
	version_exists, _ := versionRepo.CheckExist(versionId)
	if !version_exists {
		return details, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, versionId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"versionId":  versionId,
		}).Error("Query Error")
		return details, err
	}
	defer rows.Close()
	for rows.Next() {
		var detail domain.Detail
		err := rows.Scan(&detail.Module.Id, &detail.Module.Module, &detail.Detail, &detail.BeginTime, &detail.EndTime, &detail.Comment)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return details, err
		}
		details = append(details, detail)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return details, err
	}
	return details, nil
}

func (repo DrillModuleDBRepo) GetDetailsByModuleId(moduleId int64) ([]domain.Detail, error) {
	var details []domain.Detail
	sql := "SELECT id, detail, status, IFNULL(TIMEDIFF(end_time, begin_time),'00:00:00'), IFNULL(begin_time,'0000-00-00 00:00:00'), IFNULL(end_time,'0000-00-00 00:00:00'), IFNULL(comment,'') FROM switch_module_detail WHERE del_flag=0 AND module_id=? ORDER BY status DESC, end_time DESC"
	moduleRepo := DrillModuleDBRepo{}
	moduleExists, _ := moduleRepo.ModuleCheckExist(moduleId)
	if !moduleExists {
		return details, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, moduleId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"moduleId":  moduleId,
		}).Error("Query Error")
		return details, err
	}
	defer rows.Close()
	for rows.Next() {
		var detail domain.Detail
		err := rows.Scan(&detail.Id, &detail.Detail, &detail.Status, &detail.Duration, &detail.BeginTime, &detail.EndTime, &detail.Comment)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return details, err
		}
		details = append(details, detail)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return details, err
	}
	return details, nil
}

func (repo DrillModuleDBRepo) EditDetailCommentById(id int64, comment string) (int64, error) {
	sql := "UPDATE switch_module_detail SET comment=? WHERE id=?"
	moduleRepo := DrillModuleDBRepo{}
	detail_exists, _ := moduleRepo.DetailCheckExist(id)
	if !detail_exists {
		return -1, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	res, err := db.Exec(sql, comment, id)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"id":  id,
			"comment":  comment,
		}).Error(err.Error())
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return -1, err
	}
	return rowsAffected, nil
}

func (repo DrillModuleDBRepo) EditDetailEndTimeById(id int64, endTime string) (int64, error) {
	sql := "UPDATE switch_module_detail SET end_time=? WHERE id=?"
	moduleRepo := DrillModuleDBRepo{}
	detail_exists, _ := moduleRepo.DetailCheckExist(id)
	if !detail_exists {
		return -1, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	res, err := db.Exec(sql, endTime, id)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"id":  id,
			"endTime":  endTime,
		}).Error(err.Error())
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Error(err)
		return -1, err
	}
	return rowsAffected, nil
}
