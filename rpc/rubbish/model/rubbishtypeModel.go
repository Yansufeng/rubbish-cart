package model

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tal-tech/go-zero/core/stores/cache"
	"github.com/tal-tech/go-zero/core/stores/sqlc"
	"github.com/tal-tech/go-zero/core/stores/sqlx"
	"github.com/tal-tech/go-zero/core/stringx"
	"github.com/tal-tech/go-zero/tools/goctl/model/sql/builderx"
)

var (
	rubbishtypeFieldNames          = builderx.RawFieldNames(&Rubbishtype{})
	rubbishtypeRows                = strings.Join(rubbishtypeFieldNames, ",")
	rubbishtypeRowsExpectAutoSet   = strings.Join(stringx.Remove(rubbishtypeFieldNames, "`create_time`", "`update_time`"), ",")
	rubbishtypeRowsWithPlaceHolder = strings.Join(stringx.Remove(rubbishtypeFieldNames, "`rubbishType`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheRubbishtypeRubbishTypePrefix = "cache#rubbishtype#rubbishType#"
)

type (
	RubbishtypeModel interface {
		Insert(data Rubbishtype) (sql.Result, error)
		FindOne(rubbishType int64) (*Rubbishtype, error)
		Update(data Rubbishtype) error
		Delete(rubbishType int64) error
		GetAll() (map[int64]string, error)
		GetAllToList() ([]Rubbishtype, error)
	}

	defaultRubbishtypeModel struct {
		sqlc.CachedConn
		table string
	}

	Rubbishtype struct {
		RubbishType int64  `db:"rubbishType"`
		RubbishName string `db:"rubbishName"`
		IconId      int64  `db:"iconId"`
	}
)

func NewRubbishtypeModel(conn sqlx.SqlConn, c cache.CacheConf) RubbishtypeModel {
	return &defaultRubbishtypeModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`rubbishtype`",
	}
}

func (m *defaultRubbishtypeModel) Insert(data Rubbishtype) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, rubbishtypeRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.RubbishType, data.RubbishName, data.IconId)

	return ret, err
}

func (m *defaultRubbishtypeModel) FindOne(rubbishType int64) (*Rubbishtype, error) {
	rubbishtypeRubbishTypeKey := fmt.Sprintf("%s%v", cacheRubbishtypeRubbishTypePrefix, rubbishType)
	var resp Rubbishtype
	err := m.QueryRow(&resp, rubbishtypeRubbishTypeKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `rubbishType` = ? limit 1", rubbishtypeRows, m.table)
		return conn.QueryRow(v, query, rubbishType)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRubbishtypeModel) Update(data Rubbishtype) error {
	rubbishtypeRubbishTypeKey := fmt.Sprintf("%s%v", cacheRubbishtypeRubbishTypePrefix, data.RubbishType)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `rubbishType` = ?", m.table, rubbishtypeRowsWithPlaceHolder)
		return conn.Exec(query, data.RubbishName, data.IconId, data.RubbishType)
	}, rubbishtypeRubbishTypeKey)
	return err
}

func (m *defaultRubbishtypeModel) Delete(rubbishType int64) error {

	rubbishtypeRubbishTypeKey := fmt.Sprintf("%s%v", cacheRubbishtypeRubbishTypePrefix, rubbishType)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `rubbishType` = ?", m.table)
		return conn.Exec(query, rubbishType)
	}, rubbishtypeRubbishTypeKey)
	return err
}

func (m *defaultRubbishtypeModel) GetAll() (map[int64]string, error) {
	var res []Rubbishtype
	query := fmt.Sprintf("select %s from %s", rubbishtypeRows, m.table)

	err := m.QueryRowsNoCache(&res, query)
	if err != nil {
		return nil, err
	}

	mapRes := map[int64]string{}
	for _, item := range res {
		mapRes[item.RubbishType] = item.RubbishName
	}

	return mapRes, nil
}

func (m *defaultRubbishtypeModel) GetAllToList() ([]Rubbishtype, error) {
	var res []Rubbishtype
	query := fmt.Sprintf("select %s from %s", rubbishtypeRows, m.table)

	err := m.QueryRowsNoCache(&res, query)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *defaultRubbishtypeModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRubbishtypeRubbishTypePrefix, primary)
}

func (m *defaultRubbishtypeModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `rubbishType` = ? limit 1", rubbishtypeRows, m.table)
	return conn.QueryRow(v, query, primary)
}
