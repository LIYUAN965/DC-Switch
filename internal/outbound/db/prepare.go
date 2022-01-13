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
