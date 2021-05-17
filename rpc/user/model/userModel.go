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
	userFieldNames          = builderx.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`openId`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheUserOpenIdPrefix = "cache#user#openId#"
)

type (
	UserModel interface {
		Insert(data User) (sql.Result, error)
		FindOne(openId string) (*User, error)
		Update(data User) error
		Delete(openId string) error
		GetAll() ([]User, error)
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		UserName string `db:"userName"`
		Purview  int64  `db:"purview"`
		OpenId   string `db:"openId"`
	}
)

func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(data User) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userRowsExpectAutoSet)
	ret, err := m.ExecNoCache(query, data.UserName, data.Purview, data.OpenId)

	return ret, err
}

func (m *defaultUserModel) FindOne(openId string) (*User, error) {
	userOpenIdKey := fmt.Sprintf("%s%v", cacheUserOpenIdPrefix, openId)
	var resp User
	err := m.QueryRow(&resp, userOpenIdKey, func(conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `openId` = ? limit 1", userRows, m.table)
		return conn.QueryRow(v, query, openId)
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

func (m *defaultUserModel) Update(data User) error {
	userOpenIdKey := fmt.Sprintf("%s%v", cacheUserOpenIdPrefix, data.OpenId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `openId` = ?", m.table, userRowsWithPlaceHolder)
		return conn.Exec(query, data.UserName, data.Purview, data.OpenId)
	}, userOpenIdKey)
	return err
}

func (m *defaultUserModel) Delete(openId string) error {

	userOpenIdKey := fmt.Sprintf("%s%v", cacheUserOpenIdPrefix, openId)
	_, err := m.Exec(func(conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `openId` = ?", m.table)
		return conn.Exec(query, openId)
	}, userOpenIdKey)
	return err
}

func (m *defaultUserModel) GetAll() ([]User, error) {
	var resp []User
	query := fmt.Sprintf("select %s from %s", userRows, m.table)

	err := m.QueryRowsNoCache(&resp, query)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheUserOpenIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `openId` = ? limit 1", userRows, m.table)
	return conn.QueryRow(v, query, primary)
}
