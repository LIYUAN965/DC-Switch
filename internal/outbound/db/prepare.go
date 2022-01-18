package db

import (
	"dcswitch/internal/domain"
	"dcswitch/pkg/mysql"
	log "github.com/sirupsen/logrus"
)

type PrepareDBRepo struct{}

func (repo PrepareDBRepo) GetAllPreparesByVersionId(versionId int64) ([]domain.Prepare, error) {
	var prepares []domain.Prepare
	sql := "SELECT p.id, d.app_domain, p.title, p.status, p.sequence FROM version_domain AS d, version_domain_prepare AS p WHERE d.id = p.version_domain_id AND p.del_flag = 0 AND d.del_flag = 0 AND d.type = '运维' AND d.version_id = ?"
	versionRepo := SwitchVersionDBRepo{}
	version_exists, _ := versionRepo.CheckExist(versionId)
	if !version_exists {
		return prepares, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, versionId)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"versionId":  versionId,
		}).Error("Query Error")
		return prepares, err
	}
	defer rows.Close()
	for rows.Next() {
		var prepare domain.Prepare
		err := rows.Scan(&prepare.Id, &prepare.Domain.AppDomain, &prepare.Title, &prepare.Status, &prepare.Sequence)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return prepares, err
		}
		prepares = append(prepares, prepare)
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return prepares, err
	}
	return prepares, nil
}

func (repo PrepareDBRepo) CheckExist(id int64) (bool, error) {
	sql := "SELECT id FROM version_domain_prepare WHERE id = ?"
	exists, err := mysql.DB.RowExists(sql, id)
	if !exists {
		return false, mysql.NotFoundError{}
	}
	return exists, err
}

func (repo PrepareDBRepo) GetPrepareById(id int64) (domain.Prepare, error) {
	var prepare domain.Prepare
	sql := "SELECT p.id, p.title, p.content FROM version_domain_prepare AS p WHERE p.id = ?"
	prepareRepo := PrepareDBRepo{}
	prepare_exists, _ := prepareRepo.CheckExist(id)
	if !prepare_exists {
		return prepare, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	rows, err := db.Query(sql, id)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err.Error(),
			"sql": sql,
			"id":  id,
		}).Error("Query Error")
		return prepare, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&prepare.Id, &prepare.Title, &prepare.Content)
		if err != nil {
			log.Errorf("could not scan row: %v\n", err)
			return prepare, err
		}
	}
	err = rows.Err()
	if err != nil {
		log.Error(err)
		return prepare, err
	}
	return prepare, nil
}

func (repo PrepareDBRepo) EditPrepareById(id int64, title string, content string) (int64, error) {
	sql := "UPDATE version_domain_prepare SET title = ?, content = ? WHERE id = ?"
	prepareRepo := PrepareDBRepo{}
	prepare_exists, _ := prepareRepo.CheckExist(id)
	if !prepare_exists {
		return -1, mysql.NotFoundError{}
	}
	db := mysql.DB.GetConn()
	res, err := db.Exec(sql, title, content, id)
	if err != nil {
		log.WithFields(log.Fields{
			"sql": sql,
			"id":  id,
			"title":  title,
			"content":  content,
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
