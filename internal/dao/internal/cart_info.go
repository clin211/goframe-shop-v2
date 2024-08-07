// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CartInfoDao is the data access object for table cart_info.
type CartInfoDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns CartInfoColumns // columns contains all the column names of Table for convenient usage.
}

// CartInfoColumns defines and stores column names for table cart_info.
type CartInfoColumns struct {
	Id             string // 购物车表
	UserId         string //
	GoodsId        string // 商品spu
	Count          string // 商品数量
	CreatedAt      string //
	UpdatedAt      string //
	DeletedAt      string //
	GoodsOptionsId string // 商品规格sku
}

// cartInfoColumns holds the columns for table cart_info.
var cartInfoColumns = CartInfoColumns{
	Id:             "id",
	UserId:         "user_id",
	GoodsId:        "goods_id",
	Count:          "count",
	CreatedAt:      "created_at",
	UpdatedAt:      "updated_at",
	DeletedAt:      "deleted_at",
	GoodsOptionsId: "goods_options_id",
}

// NewCartInfoDao creates and returns a new DAO object for table data access.
func NewCartInfoDao() *CartInfoDao {
	return &CartInfoDao{
		group:   "default",
		table:   "cart_info",
		columns: cartInfoColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CartInfoDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CartInfoDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CartInfoDao) Columns() CartInfoColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CartInfoDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CartInfoDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CartInfoDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
