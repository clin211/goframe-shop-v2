// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// CartInfo is the golang structure of table cart_info for DAO operations like Where/Data.
type CartInfo struct {
	g.Meta         `orm:"table:cart_info, do:true"`
	Id             interface{} // 购物车表
	UserId         interface{} //
	GoodsId        interface{} // 商品spu
	Count          interface{} // 商品数量
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
	GoodsOptionsId interface{} // 商品规格sku
}
