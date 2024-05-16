package request

import (
	"main.go/model/common"
	"main.go/model/common/request"
	"main.go/model/manage"
)

type MallGoodsInfoSearch struct {
	manage.MallGoodsInfo
	request.PageInfo
}

// type GoodsInfoAddParam struct {
// 	GoodsName          string `json:"goodsName"`
// 	GoodsIntro         string `json:"goodsIntro"`
// 	GoodsCategoryId    int    `json:"goodsCategoryId"`
// 	GoodsCoverImg      string `json:"goodsCoverImg"`
// 	GoodsCarousel      string `json:"goodsCarousel"`
// 	GoodsDetailContent string `json:"goodsDetailContent"`
// 	OriginalPrice      string `json:"originalPrice"`
// 	SellingPrice       string `json:"sellingPrice"`
// 	StockNum           string `json:"stockNum"`
// 	Tag                string `json:"tag"`
// 	GoodsSellStatus    string `json:"goodsSellStatus"`
// }

type GoodsInfoAddParam struct {
	GoodsName           string `json:"goodsName"`
	GoodsIntro          string `json:"goodsIntro"`
	GoodsCategoryId     int    `json:"goodsCategoryId"`
	GoodsCoverImg       string `json:"goodsCoverImg"`
	GoodsCarousel       string `json:"goodsCarousel"`
	GoodsDetailContent  string `json:"goodsDetailContent"`
	OriginalPrice       string `json:"originalPrice"`
	SellingPrice        string `json:"sellingPrice"`
	StockNum            string `json:"stockNum"`
	Tag                 string `json:"tag"`
	GoodsSellStatus     string `json:"goodsSellStatus"`
	ServerArea          string `json:"serverArea"`
	Career              []string `json:"career"`
	ReqLv               int `json:"reqLv"`
	StarForce           int `json:"starForce"`
	AttStr              int `json:"attStr"`
	AttDex              int `json:"attDex"`
	AttInt              int `json:"attInt"`
	AttLuk              int `json:"attLuk"`
	PotentialColor      string    `json:"potentialColor"`
	PotentialQuality    string    `json:"potentialQuality"`
	AddPotentialColor   string    `json:"addPotentialColor"`
	AddPotentialQuality string    `json:"addPotentialQuality"`
	Hammer              int `json:"hammer"`
	Contact             string `json:"contact"`
}

// GoodsInfoUpdateParam 更新商品信息的入参
type GoodsInfoUpdateParam struct {
	GoodsId             string          `json:"goodsId"`
	GoodsName           string          `json:"goodsName"`
	GoodsIntro          string          `json:"goodsIntro"`
	GoodsCategoryId     int             `json:"goodsCategoryId"`
	GoodsCoverImg       string          `json:"goodsCoverImg"`
	GoodsCarousel       string          `json:"goodsCarousel"`
	GoodsDetailContent  string          `json:"goodsDetailContent"`
	OriginalPrice       string          `json:"originalPrice"`
	SellingPrice        int             `json:"sellingPrice"`
	StockNum            string          `json:"stockNum"`
	Tag                 string          `json:"tag"`
	GoodsSellStatus     int             `json:"goodsSellStatus"`
	ServerArea          string `json:"serverArea"`
	Career              []string `json:"career"`
	ReqLv               int `json:"reqLv"`
	StarForce           int `json:"starForce"`
	AttStr              int `json:"attStr"`
	AttDex              int `json:"attDex"`
	AttInt              int `json:"attInt"`
	AttLuk              int `json:"attLuk"`
	PotentialColor      string    `json:"potentialColor"`
	PotentialQuality    string    `json:"potentialQuality"`
	AddPotentialColor   string    `json:"addPotentialColor"`
	AddPotentialQuality string    `json:"addPotentialQuality"`
	Hammer              int `json:"hammer"`
	Contact             string `json:"contact"`
	UpdateUser          int             `json:"updateUser" form:"updateUser" gorm:"column:update_user;comment:修改者主键id;type:int"`
	UpdateTime          common.JSONTime `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:商品修改时间;type:datetime"`
}

type StockNumDTO struct {
	GoodsId    int `json:"goodsId"`
	GoodsCount int `json:"goodsCount"`
}
