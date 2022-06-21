// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// NftEntityDao is the data access object for table nft_entity.
type NftEntityDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns NftEntityColumns // columns contains all the column names of Table for convenient usage.
}

// NftEntityColumns defines and stores column names for table nft_entity.
type NftEntityColumns struct {
	Id         string // 主键
	EntityId   string // 实体id##
	EntityType string // 实体类型，1:装备##
	Token      string // 区块链token##
	CreatedAt  string // 创建时间##
	UpdatedAt  string // 创建时间##
	DeletedAt  string // 创建时间##
}

//  nftEntityColumns holds the columns for table nft_entity.
var nftEntityColumns = NftEntityColumns{
	Id:         "id",
	EntityId:   "entity_id",
	EntityType: "entity_type",
	Token:      "token",
	CreatedAt:  "created_at",
	UpdatedAt:  "updated_at",
	DeletedAt:  "deleted_at",
}

// NewNftEntityDao creates and returns a new DAO object for table data access.
func NewNftEntityDao() *NftEntityDao {
	return &NftEntityDao{
		group:   "default",
		table:   "nft_entity",
		columns: nftEntityColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *NftEntityDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *NftEntityDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *NftEntityDao) Columns() NftEntityColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *NftEntityDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *NftEntityDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *NftEntityDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}