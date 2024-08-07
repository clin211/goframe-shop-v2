// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// EventeECardDao is the data access object for table evente_e_card.
type EventeECardDao struct {
	table   string             // table is the underlying table name of the DAO.
	group   string             // group is the database configuration group name of current DAO.
	columns EventeECardColumns // columns contains all the column names of Table for convenient usage.
}

// EventeECardColumns defines and stores column names for table evente_e_card.
type EventeECardColumns struct {
	Id            string //
	OrgId         string // 主办id
	CardName      string // E通卡名称
	SaleStartDate string // 售卖开始时间
	SaleStopDate  string // 售卖结束时间
	Usage         string // 使用方式 1、预约 2、免预约
}

// eventeECardColumns holds the columns for table evente_e_card.
var eventeECardColumns = EventeECardColumns{
	Id:            "id",
	OrgId:         "org_id",
	CardName:      "card_name",
	SaleStartDate: "sale_start_date",
	SaleStopDate:  "sale_stop_date",
	Usage:         "usage",
}

// NewEventeECardDao creates and returns a new DAO object for table data access.
func NewEventeECardDao() *EventeECardDao {
	return &EventeECardDao{
		group:   "default",
		table:   "evente_e_card",
		columns: eventeECardColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *EventeECardDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *EventeECardDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *EventeECardDao) Columns() EventeECardColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *EventeECardDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *EventeECardDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *EventeECardDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
