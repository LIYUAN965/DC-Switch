package db

import (
	"dcswitch/internal/domain"
	"dcswitch/pkg/mysql"
	log "github.com/sirupsen/logrus"
)

type ModuleDetailDBRepo struct{}

func (repo ModuleDetailDBRepo) GetModuleDetail(name string) (domain.ModuleDetail, error) {
	// TODO: 未实现
	var detail domain.ModuleDetail
	sql := "SELECT id FROM switch_module_detail WHERE detail = ?"
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, name)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"detail":  name,
		}).Error("Query Error")
		return detail, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&detail.Id)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return detail, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return detail, err
	}
	return detail, nil
}

func (repo ModuleDetailDBRepo) ModuleDetailStart(detailID int64, detail domain.ModuleDetail) (int64, error) {
	// TODO: 未实现
	return 0, nil
}

func (repo ModuleDetailDBRepo) ModuleDetailSuccess(detailID int64, detail domain.ModuleDetail) (int64, error) {
	// TODO: 未实现
	return 0, nil
}

func (repo ModuleDetailDBRepo) ModuleDetailFail(detailID int64, detail domain.ModuleDetail) (int64, error) {
	// TODO: 未实现
	return 0, nil
}


func (repo ModuleDetailDBRepo) ModifyModuleDetailStatus(detailID int64, status int64) (int64, error) {
	sql := "UPDATE switch_module_detail SET status = ? WHERE id = ?"
	db := mysql.DB.GetConn()
	res, err := db.Exec(sql, status, detailID)
	if err != nil {
		log.WithFields(log.Fields{
			"sql":  sql,
			"id":   detailID,
			"status": status,
		}).Error(err.Error())
		return -1, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	return rowsAffected, nil
}



