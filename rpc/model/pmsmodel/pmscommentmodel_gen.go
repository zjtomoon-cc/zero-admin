// Code generated by goctl. DO NOT EDIT!

package pmsmodel

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
	pmsCommentFieldNames          = builder.RawFieldNames(&PmsComment{})
	pmsCommentRows                = strings.Join(pmsCommentFieldNames, ",")
	pmsCommentRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsCommentFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	pmsCommentRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsCommentFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	pmsCommentModel interface {
		Insert(ctx context.Context, data *PmsComment) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsComment, error)
		Update(ctx context.Context, data *PmsComment) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsCommentModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsComment struct {
		Id               int64          `db:"id"`
		ProductId        sql.NullInt64  `db:"product_id"`
		MemberNickName   sql.NullString `db:"member_nick_name"`
		ProductName      sql.NullString `db:"product_name"`
		Star             sql.NullInt64  `db:"star"`      // 评价星数：0->5
		MemberIp         sql.NullString `db:"member_ip"` // 评价的ip
		CreateTime       sql.NullTime   `db:"create_time"`
		ShowStatus       sql.NullInt64  `db:"show_status"`
		ProductAttribute sql.NullString `db:"product_attribute"` // 购买时的商品属性
		CollectCouont    sql.NullInt64  `db:"collect_couont"`
		ReadCount        sql.NullInt64  `db:"read_count"`
		Content          sql.NullString `db:"content"`
		Pics             sql.NullString `db:"pics"`        // 上传图片地址，以逗号隔开
		MemberIcon       sql.NullString `db:"member_icon"` // 评论用户头像
		ReplayCount      sql.NullInt64  `db:"replay_count"`
	}
)

func newPmsCommentModel(conn sqlx.SqlConn) *defaultPmsCommentModel {
	return &defaultPmsCommentModel{
		conn:  conn,
		table: "`pms_comment`",
	}
}

func (m *defaultPmsCommentModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsCommentModel) FindOne(ctx context.Context, id int64) (*PmsComment, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsCommentRows, m.table)
	var resp PmsComment
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

func (m *defaultPmsCommentModel) Insert(ctx context.Context, data *PmsComment) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsCommentRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.MemberNickName, data.ProductName, data.Star, data.MemberIp, data.ShowStatus, data.ProductAttribute, data.CollectCouont, data.ReadCount, data.Content, data.Pics, data.MemberIcon, data.ReplayCount)
	return ret, err
}

func (m *defaultPmsCommentModel) Update(ctx context.Context, data *PmsComment) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsCommentRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.ProductId, data.MemberNickName, data.ProductName, data.Star, data.MemberIp, data.ShowStatus, data.ProductAttribute, data.CollectCouont, data.ReadCount, data.Content, data.Pics, data.MemberIcon, data.ReplayCount, data.Id)
	return err
}

func (m *defaultPmsCommentModel) tableName() string {
	return m.table
}