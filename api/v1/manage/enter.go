package manage

import "main.go/service"

type ManageGroup struct {
	ManageAdminUserApi
	ManageGoodsCategoryApi
	ManageGoodsInfoApi
	ManageCarouselApi
	ManageIndexConfigApi
	ManageOrderApi
}

var mallAdminUserService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserService
var mallAdminUserTokenService = service.ServiceGroupApp.ManageServiceGroup.ManageAdminUserTokenService
var mallUserService = service.ServiceGroupApp.ManageServiceGroup.ManageUserService
var mallGoodsCategoryService = service.ServiceGroupApp.ManageServiceGroup.ManageGoodsCategoryService
var fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
var mallGoodsInfoService = service.ServiceGroupApp.ManageServiceGroup.ManageGoodsInfoService
var mallCarouselService = service.ServiceGroupApp.ManageServiceGroup.ManageCarouselService
var mallIndexConfigService = service.ServiceGroupApp.ManageServiceGroup.ManageIndexConfigService
var mallOrderService = service.ServiceGroupApp.ManageServiceGroup.ManageOrderService
