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
	rubbishFieldNames          = builderx.RawFieldNames(&Rubbish{})
	rubbishRows                = strings.Join(rubbishFieldNames, ",")
	rubbishRowsExpectAutoSet   = strings.Join(stringx.Remove(rubbishFieldNames, "`create_time`", "`update_time`"), ",")
	rubbishRowsWithPlaceHolder = strings.Join(stringx.Remove(rubbishFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheRubbishIdPrefix = "cache#rubbish#id#"
)

type (
	RubbishModel interface {
		Insert(data Rubbish) (sql.Result, error)
		FindOne(id int64) (*Rubbish, error)
		GetAllByCartId(cartId int64) ([]Rubbish, error)
		FindOneById(cartId int64, rubbishType int64) (*Rubbish, error)
		Update(data Rubbish) error
		Delete(id int64) error
	}

	defaultRubbishModel struct {
		sqlc.CachedConn
		table string
	}

	Rubbish struct {
		Id            int64 `db:"id"`
		CartId        int64 `db:"cartId"`
		RubbishType   int64 `db:"rubbishType"`
		RubbishAmount int64 `db:"rubbishAmount"`
	}
)

func NewRubbishModel(conn sqlx.SqlConn, c cache.CacheConf) RubbishModel {
	return &defaultRubbishModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`rubbish`",
	}
}

func (m *defaultRubbishModel) Insert(data Rubbish) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, rubbishRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.Id, data.CartId, data.RubbishType, data.RubbishAmount)

	return ret, err
}

func (m *defaultRubbishModel) FindOne(id int64) (*Rubbish, error) {
	rubbishIdKey := fmt.Sprintf("%s%v", cacheRubbishIdPrefix, id)
	var resp Rubbish
	err := m.QueryRow(&resp, rubbishIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rubbishRows, m.table)
		return conn.QueryRow(v, query, id)
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

func (m *defaultRubbishModel) FindOneById(cartId int64, rubbishType int64) (*Rubbish, error) {
	var res Rubbish
	query := fmt.Sprintf("select %s from %s where `cartId` = ? AND `rubbishType` = ?", rubbishRows, m.table)

	err := m.QueryRowNoCache(&res, query, cartId, rubbishType)
	switch err {
	case nil:
		return &res, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRubbishModel)GetAllByCartId(cartId int64) ([]Rubbish, error) {
	var res []Rubbish
	query := fmt.Sprintf("select %s from %s where `cartId` = ?", rubbishRows, m.table)

	err := m.QueryRowsNoCache(&res, query, cartId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (m *defaultRubbishModel) Update(data Rubbish) error {
	rubbishIdKey := fmt.Sprintf("%s%v", cacheRubbishIdPrefix, data.Id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rubbishRowsWithPlaceHolder)
		return conn.Exec(query, data.CartId, data.RubbishType, data.RubbishAmount, data.Id)
	}, rubbishIdKey)
	return err
}

func (m *defaultRubbishModel) Delete(id int64) error {

	rubbishIdKey := fmt.Sprintf("%s%v", cacheRubbishIdPrefix, id)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.Exec(query, id)
	}, rubbishIdKey)
	return err
}

func (m *defaultRubbishModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheRubbishIdPrefix, primary)
}

func (m *defaultRubbishModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rubbishRows, m.table)
	return conn.QueryRow(v, query, primary)
}
