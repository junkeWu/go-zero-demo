// Code generated by goctl. DO NOT EDIT.

package genModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGoZeroUserIdPrefix     = "cache:goZero:user:id:"
	cacheGoZeroUserMobilePrefix = "cache:goZero:user:mobile:"
	cacheGoZeroUserUserIdPrefix = "cache:goZero:user:userId:"
)

type (
	userModel interface {
		Insert(ctx context.Context, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*User, error)
		FindOneByMobile(ctx context.Context, mobile string) (*User, error)
		FindOneByUserId(ctx context.Context, userId string) (*User, error)
		Update(ctx context.Context, data *User) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id       int64  `db:"id"`       // id
		Username string `db:"username"` // 用户名
		Mobile   string `db:"mobile"`   // 电话号码
		UserId   string `db:"user_id"`  // 用户ID
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	goZeroUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, id)
	goZeroUserMobileKey := fmt.Sprintf("%s%v", cacheGoZeroUserMobilePrefix, data.Mobile)
	goZeroUserUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, goZeroUserIdKey, goZeroUserMobileKey, goZeroUserUserIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id int64) (*User, error) {
	goZeroUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, goZeroUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
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

func (m *defaultUserModel) FindOneByMobile(ctx context.Context, mobile string) (*User, error) {
	goZeroUserMobileKey := fmt.Sprintf("%s%v", cacheGoZeroUserMobilePrefix, mobile)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, goZeroUserMobileKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `mobile` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, mobile); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindOneByUserId(ctx context.Context, userId string) (*User, error) {
	goZeroUserUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserUserIdPrefix, userId)
	var resp User
	err := m.QueryRowIndexCtx(ctx, &resp, goZeroUserUserIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `user_id` = ? limit 1", userRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, data *User) (sql.Result, error) {
	goZeroUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, data.Id)
	goZeroUserMobileKey := fmt.Sprintf("%s%v", cacheGoZeroUserMobilePrefix, data.Mobile)
	goZeroUserUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserUserIdPrefix, data.UserId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, userRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.Mobile, data.UserId)
	}, goZeroUserIdKey, goZeroUserMobileKey, goZeroUserUserIdKey)
	return ret, err
}

func (m *defaultUserModel) Update(ctx context.Context, newData *User) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	goZeroUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, data.Id)
	goZeroUserMobileKey := fmt.Sprintf("%s%v", cacheGoZeroUserMobilePrefix, data.Mobile)
	goZeroUserUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserUserIdPrefix, data.UserId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Username, newData.Mobile, newData.UserId, newData.Id)
	}, goZeroUserIdKey, goZeroUserMobileKey, goZeroUserUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}