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
	pmsBrandFieldNames          = builder.RawFieldNames(&PmsBrand{})
	pmsBrandRows                = strings.Join(pmsBrandFieldNames, ",")
	pmsBrandRowsExpectAutoSet   = strings.Join(stringx.Remove(pmsBrandFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	pmsBrandRowsWithPlaceHolder = strings.Join(stringx.Remove(pmsBrandFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"
)

type (
	pmsBrandModel interface {
		Insert(ctx context.Context, data *PmsBrand) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*PmsBrand, error)
		Update(ctx context.Context, data *PmsBrand) error
		Delete(ctx context.Context, id int64) error
	}

	defaultPmsBrandModel struct {
		conn  sqlx.SqlConn
		table string
	}

	PmsBrand struct {
		Id                  int64          `db:"id"`
		Name                sql.NullString `db:"name"`
		FirstLetter         sql.NullString `db:"first_letter"` // 首字母
		Sort                sql.NullInt64  `db:"sort"`
		FactoryStatus       sql.NullInt64  `db:"factory_status"` // 是否为品牌制造商：0->不是；1->是
		ShowStatus          sql.NullInt64  `db:"show_status"`
		ProductCount        sql.NullInt64  `db:"product_count"`         // 产品数量
		ProductCommentCount sql.NullInt64  `db:"product_comment_count"` // 产品评论数量
		Logo                sql.NullString `db:"logo"`                  // 品牌logo
		BigPic              sql.NullString `db:"big_pic"`               // 专区大图
		BrandStory          sql.NullString `db:"brand_story"`           // 品牌故事
	}
)

func newPmsBrandModel(conn sqlx.SqlConn) *defaultPmsBrandModel {
	return &defaultPmsBrandModel{
		conn:  conn,
		table: "`pms_brand`",
	}
}

func (m *defaultPmsBrandModel) Delete(ctx context.Context, id int64) error {
	query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
	_, err := m.conn.ExecCtx(ctx, query, id)
	return err
}

func (m *defaultPmsBrandModel) FindOne(ctx context.Context, id int64) (*PmsBrand, error) {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", pmsBrandRows, m.table)
	var resp PmsBrand
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

func (m *defaultPmsBrandModel) Insert(ctx context.Context, data *PmsBrand) (sql.Result, error) {
	query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, pmsBrandRowsExpectAutoSet)
	ret, err := m.conn.ExecCtx(ctx, query, data.Name, data.FirstLetter, data.Sort, data.FactoryStatus, data.ShowStatus, data.ProductCount, data.ProductCommentCount, data.Logo, data.BigPic, data.BrandStory)
	return ret, err
}

func (m *defaultPmsBrandModel) Update(ctx context.Context, data *PmsBrand) error {
	query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, pmsBrandRowsWithPlaceHolder)
	_, err := m.conn.ExecCtx(ctx, query, data.Name, data.FirstLetter, data.Sort, data.FactoryStatus, data.ShowStatus, data.ProductCount, data.ProductCommentCount, data.Logo, data.BigPic, data.BrandStory, data.Id)
	return err
}

func (m *defaultPmsBrandModel) tableName() string {
	return m.table
}