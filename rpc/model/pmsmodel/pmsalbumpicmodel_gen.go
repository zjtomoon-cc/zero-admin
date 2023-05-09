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
	pmsAlbumPicFieldNames          = builder.RawFieldNames(&PmsAlbumPic{})
	pmsAlbumPicRows                = strings.Join(pmsAlbumPicFieldNames, ",")
	pmsAlbumPicRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsAlbumPicFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	pmsAlbumPicRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsAlbumPicFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	pmsAlbumPicModel interface {
		Insert(ctx context.Context, data *PmsAlbumPic) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsAlbumPic, error)
		Update(ctx context.Context, data *PmsAlbumPic) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsAlbumPicModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsAlbumPic struct {
		Id      int64          `db:"id"`
		AlbumId sql.NullInt64  `db:"album_id"`
		Pic     sql.NullString `db:"pic"`
	}
)

func newPmsAlbumPicModel(conn sqlx.SqlConn) *defaultPmsAlbumPicModel {
	return &defaultPmsAlbumPicModel{
		conn:  conn,
		table: "`pms_album_pic`",
	}
}

func (m *defaultPmsAlbumPicModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsAlbumPicModel) FindOne(ctx context.Context, id int64) (*PmsAlbumPic, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsAlbumPicRows, m.table)
	var resp PmsAlbumPic
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

func (m *defaultPmsAlbumPicModel) Insert(ctx context.Context, data *PmsAlbumPic) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, pmsAlbumPicRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.AlbumId, data.Pic)
	return ret, err
}

func (m *defaultPmsAlbumPicModel) Update(ctx context.Context, data *PmsAlbumPic) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsAlbumPicRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.AlbumId, data.Pic, data.Id)
	return err
}

func (m *defaultPmsAlbumPicModel) tableName() string {
	return m.table
}