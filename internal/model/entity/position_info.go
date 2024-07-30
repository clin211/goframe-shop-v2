// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// PositionInfo is the golang structure for table position_info.
type PositionInfo struct {
	Id        int         `json:"id"        orm:"id"         description:""`
	PicUrl    string      `json:"picUrl"    orm:"pic_url"    description:"图片链接"`
	Link      string      `json:"link"      orm:"link"       description:"跳转链接"`
	Sort      int         `json:"sort"      orm:"sort"       description:"排序"`
	GoodsId   string      `json:"goodsId"   orm:"goods_id"   description:"商品 id"`
	GoodsName string      `json:"goodsName" orm:"goods_name" description:"商品名称"`
	CreatedAt *gtime.Time `json:"createdAt" orm:"created_at" description:""`
	UpdatedAt *gtime.Time `json:"updatedAt" orm:"updated_at" description:""`
	DeletedAt *gtime.Time `json:"deletedAt" orm:"deleted_at" description:""`
}
