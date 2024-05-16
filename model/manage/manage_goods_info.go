// // 自动生成模板MallGoodsInfo
package manage

import (
	"main.go/model/common"
)

// MallGoodsInfo 结构体
// 如果含有time.Time 请自行import time包
type MallGoodsInfo struct {
	GoodsId             int             `json:"goodsId" form:"goodsId" gorm:"primarykey;AUTO_INCREMENT"`
	GoodsName           string          `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名;type:varchar(200);"`
	GoodsIntro          string          `json:"goodsIntro" form:"goodsIntro" gorm:"column:goods_intro;comment:商品简介;type:varchar(200);"`
	GoodsCategoryId     int             `json:"goodsCategoryId" form:"goodsCategoryId" gorm:"column:goods_category_id;comment:关联分类id;type:bigint"`
	GoodsCoverImg       string          `json:"goodsCoverImg" form:"goodsCoverImg" gorm:"column:goods_cover_img;comment:商品主图;type:varchar(200);"`
	GoodsCarousel       string          `json:"goodsCarousel" form:"goodsCarousel" gorm:"column:goods_carousel;comment:商品轮播图;type:varchar(500);"`
	GoodsDetailContent  string          `json:"goodsDetailContent" form:"goodsDetailContent" gorm:"column:goods_detail_content;comment:商品详情;type:text;"`
	OriginalPrice       int             `json:"originalPrice" form:"originalPrice" gorm:"column:original_price;comment:商品价格;type:int"`
	SellingPrice        int             `json:"sellingPrice" form:"sellingPrice" gorm:"column:selling_price;comment:商品实际售价;type:int"`
	StockNum            int             `json:"stockNum" form:"stockNum" gorm:"column:stock_num;comment:商品库存数量;type:int"`
	Tag                 string          `json:"tag" form:"tag" gorm:"column:tag;comment:商品标签;type:varchar(20);"`
	GoodsSellStatus     int             `json:"goodsSellStatus" form:"goodsSellStatus" gorm:"column:goods_sell_status;comment:商品上架状态 1-下架 0-上架;type:tinyint"`
	CreateUser          int             `json:"createUser" form:"createUser" gorm:"column:create_user;comment:添加者主键id;type:int"`
	CreateTime          common.JSONTime `json:"createTime" form:"createTime" gorm:"column:create_time;comment:商品添加时间;type:datetime"`
	UpdateUser          int             `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者主键id;type:int"`
	UpdateTime          common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:商品修改时间;type:datetime"`
	// TypeId              int             `json:"typeId" form:"typeId" gorm:"column:type_id;comment:类型id;size:10;"`
	// Deleted             bool            `json:"deleted" form:"deleted" gorm:"column:deleted;comment:是否删除;"`                                                 //是否删除
	ServerArea          int          `json:"serverArea" form:"serverArea" gorm:"column:server_area;comment:区服;size:20;"`                                 //区服
	Career              string          `json:"career" form:"career" gorm:"column:career;comment:可装备职业;size:20;"`                                           //可装备职业
	ReqLv               int             `json:"reqLv" form:"reqLv" gorm:"column:req_lv;comment:装备等级;size:10;"`                                              //装备等级                                               //价格
	StarForce           int             `json:"starForce" form:"starForce" gorm:"column:star_force;comment:星力;size:10;"`                                    //星力
	AttStr              int             `json:"attStr" form:"attStr" gorm:"column:att_str;comment:力量;size:10;"`                                             //力量
	AttDex              int             `json:"attDex" form:"attDex" gorm:"column:att_dex;comment:敏捷;size:10;"`                                             //敏捷
	AttInt              int             `json:"attInt" form:"attInt" gorm:"column:att_int;comment:智力;size:10;"`                                             //智力
	AttLuk              int             `json:"attLuk" form:"attLuk" gorm:"column:att_luk;comment:幸运;size:10;"`                                             //幸运
	PotentialColor      int             `json:"potentialColor" form:"potentialColor" gorm:"column:potential_color;comment:潜能颜色;size:10;"`                   //潜能颜色
	PotentialQuality    int             `json:"potentialQuality" form:"potentialQuality" gorm:"column:potential_quality;comment:潜能属性;size:10;"`             //潜能属性
	AddPotentialColor   int             `json:"addPotentialColor" form:"addPotentialColor" gorm:"column:add_potential_color;comment:附加潜能颜色;size:10;"`       //附加潜能颜色
	AddPotentialQuality int             `json:"addPotentialQuality" form:"addPotentialQuality" gorm:"column:add_potential_quality;comment:附加潜能属性;size:10;"` //附加潜能属性
	Hammer              int             `json:"hammer" form:"hammer" gorm:"column:hammer;comment:锤子;size:10;"`                                              //锤子
	Contact             string          `json:"contact" form:"contact" gorm:"column:contact;comment:;size:50;"`                                             //contact字段
}

// TableName MallGoodsInfo 表名
func (MallGoodsInfo) TableName() string {
	return "tb_newbee_mall_goods_info"
}
