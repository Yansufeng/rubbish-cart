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
	cartFieldNames          = builderx.RawFieldNames(&Cart{})
	cartRows                = strings.Join(cartFieldNames, ",")
	cartRowsExpectAutoSet   = strings.Join(stringx.Remove(cartFieldNames, "`create_time`", "`update_time`"), ",")
	cartRowsWithPlaceHolder = strings.Join(stringx.Remove(cartFieldNames, "`cartId`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheCartCartIdPrefix = "cache#cart#cartId#"
	cacheCartCartNamePrefix = "cache#cart#cartName#"
)

type (
	CartModel interface {
		Insert(data Cart) (sql.Result, error)
		FindOne(cartId int64) (*Cart, error)
		FindOneByName(cartName string) (*Cart, error)
		Update(data Cart) error
		Delete(cartId int64) error
		GetAll() ([]Cart, error)
	}

	defaultCartModel struct {
		sqlc.CachedConn
		table string
	}

	Cart struct {
		CartId   int64  `db:"cartId"`
		CartName string `db:"cartName"`
		State    string `db:"state"`
	}
)

func NewCartModel(conn sqlx.SqlConn, c cache.CacheConf) CartModel {
	return &defaultCartModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`cart`",
	}
}

func (m *defaultCartModel) Insert(data Cart) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, cartRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.CartId, data.CartName, data.State)

	return ret, err
}

func (m *defaultCartModel) FindOne(cartId int64) (*Cart, error) {
	cartCartIdKey := fmt.Sprintf("%s%v", cacheCartCartIdPrefix, cartId)
	var resp Cart
	err := m.QueryRow(&resp, cartCartIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `cartId` = ? limit 1", cartRows, m.table)
		return conn.QueryRow(v, query, cartId)
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

func (m *defaultCartModel) FindOneByName(cartName string) (*Cart, error) {
	cartCartNameKey := fmt.Sprintf("%s%v", cacheCartCartNamePrefix, cartName)
	var resp Cart
	err := m.QueryRow(&resp, cartCartNameKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `cartName` = ? limit 1", cartRows, m.table)
		return conn.QueryRow(v, query, cartName)
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

func (m *defaultCartModel) Update(data Cart) error {
	cartCartIdKey := fmt.Sprintf("%s%v", cacheCartCartIdPrefix, data.CartId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `cartId` = ?", m.table, cartRowsWithPlaceHolder)
		return conn.Exec(query, data.CartName, data.State, data.CartId)
	}, cartCartIdKey)
	return err
}

func (m *defaultCartModel) Delete(cartId int64) error {
	cartCartIdKey := fmt.Sprintf("%s%v", cacheCartCartIdPrefix, cartId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `cartId` = ?", m.table)
		return conn.Exec(query, cartId)
	}, cartCartIdKey)
	return err
}

func (m *defaultCartModel) GetAll() ([]Cart, error) {
	var resp []Cart
	query := fmt.Sprintf("select %s from %s", cartRows, m.table)

	err := m.QueryRowsNoCache(&resp, query)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *defaultCartModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheCartCartIdPrefix, primary)
}

func (m *defaultCartModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `cartId` = ? limit 1", cartRows, m.table)
	return conn.QueryRow(v, query, primary)
}
