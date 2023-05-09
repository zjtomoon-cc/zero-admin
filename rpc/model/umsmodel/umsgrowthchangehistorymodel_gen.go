// Code generated by goctl. DO NOT EDIT!

package umsmodel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	umsGrowthChangeHistoryFieldNames          = builder.RawFieldNames(&UmsGrowthChangeHistory{})
	umsGrowthChangeHistoryRows                = strings.Join(umsGrowthChangeHistoryFieldNames, ",")
	umsGrowthChangeHistoryRowsExpectAutoSet   = strings.Join(stringx.Remove(umsGrowthChangeHistoryFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	umsGrowthChangeHistoryRowsWithPlaceHolder = strings.Join(stringx.Remove(umsGrowthChangeHistoryFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	umsGrowthChangeHistoryModel interface {
		Insert(ctx context.Context, data *UmsGrowthChangeHistory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*UmsGrowthChangeHistory, error)
		Update(ctx context.Context, data *UmsGrowthChangeHistory) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUmsGrowthChangeHistoryModel struct {
		conn  sqlx.SqlConn
		table string
	}

	UmsGrowthChangeHistory struct {
		Id          int64          `db:"id"`
		MemberId    sql.NullInt64  `db:"member_id"`
		CreateTime  sql.NullTime   `db:"create_time"`
		ChangeType  sql.NullInt64  `db:"change_type"`  // 改变类型：0->增加；1->减少
		ChangeCount sql.NullInt64  `db:"change_count"` // 积分改变数量
		OperateMan  sql.NullString `db:"operate_man"`  // 操作人员
		OperateNote sql.NullString `db:"operate_note"` // 操作备注
		SourceType  sql.NullInt64  `db:"source_type"`  // 积分来源：0->购物；1->管理员修改
	}
)

func newUmsGrowthChangeHistoryModel(conn sqlx.SqlConn) *defaultUmsGrowthChangeHistoryModel {
	return &defaultUmsGrowthChangeHistoryModel{
		conn:  conn,
		table: "`ums_growth_change_history`",
	}
}

func (m *defaultUmsGrowthChangeHistoryModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultUmsGrowthChangeHistoryModel) FindOne(ctx context.Context, id int64) (*UmsGrowthChangeHistory, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", umsGrowthChangeHistoryRows, m.table)
	var resp UmsGrowthChangeHistory
	err := m.conn.QueryRowCtx(ctx, &resp, query, id)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUmsGrowthChangeHistoryModel) Insert(ctx context.Context, data *UmsGrowthChangeHistory) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, umsGrowthChangeHistoryRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.MemberId, data.ChangeType, data.ChangeCount, data.OperateMan, data.OperateNote, data.SourceType)
	return ret, err
}

func (m *defaultUmsGrowthChangeHistoryModel) Update(ctx context.Context, data *UmsGrowthChangeHistory) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, umsGrowthChangeHistoryRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.MemberId, data.ChangeType, data.ChangeCount, data.OperateMan, data.OperateNote, data.SourceType, data.Id)
	return err
}

func (m *defaultUmsGrowthChangeHistoryModel) tableName() string {
	return m.table
}