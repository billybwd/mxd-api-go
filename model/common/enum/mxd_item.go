// ServerArea         
// Career             
// ReqLv              
// StarForce          
// AttStr             
// AttDex             
// AttInt             
// AttLuk             
// PotentialColor     
// PotentialQuality   
// AddPotentialColor  
// AddPotentialQuality
// Hammer             
// Contact            


package enum
// "value": "0",
// "label": "不限"
// }, {
// "value": "1",
// "label": "艾麗亞"
// }, {
// "value": "2",
// "label": "普利特"
// }, {
// "value": "3",
// "label": "琉德"
// }, {
// "value": "4",
// "label": "優依娜"
// }, {
// "value": "5",
// "label": "愛麗西亞"
// }, {
// "value": "6",
// "label": "杀人鲸"
type ServerAreaEnum int

const (
	AREA_ALL ServerAreaEnum = 0
	AREA1   ServerAreaEnum = 1
	AREA2   ServerAreaEnum = 2
	AREA3   ServerAreaEnum = 3
	AREA4   ServerAreaEnum = 4
	AREA5   ServerAreaEnum = 5
	AREA6   ServerAreaEnum = 6
)

func ServerAreaEnumByStatus(status int) (int, string) {
	switch status {
	case 1:
		return 1, "艾麗亞"
	case 2:
		return 2, "普利特"
	case 3:
		return 3, "琉德"
	case 4:
		return 4, "優依娜"
	case 5:
		return 5, "愛麗西亞"
	case 6:
		return 6, "杀人鲸"										
	default:
		return 0, "不限"
	}
}

func (g ServerAreaEnum) Code() int {
	switch g {
	case AREA1:
		return 1
	case AREA2:
		return 2
	case AREA3:
		return 3
	case AREA4:
		return 4
	case AREA5:
		return 5
	case AREA6:
		return 6										
	default:
		return 0
	}
}




type MallOrderStatusEnum int

const (
	DEFAULT                  MallOrderStatusEnum = -9
	ORDER_PRE_PAY            MallOrderStatusEnum = 0
	ORDER_PAID               MallOrderStatusEnum = 1
	ORDER_PACKAGED           MallOrderStatusEnum = 2
	ORDER_EXPRESS            MallOrderStatusEnum = 3
	ORDER_SUCCESS            MallOrderStatusEnum = 4
	ORDER_CLOSED_BY_MALLUSER MallOrderStatusEnum = -1
	ORDER_CLOSED_BY_EXPIRED  MallOrderStatusEnum = -2
	ORDER_CLOSED_BY_JUDGE    MallOrderStatusEnum = -3
)

func GetNewBeeMallOrderStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "待支付"
	case 1:
		return 1, "已支付"
	case 2:
		return 2, "配货完成"
	case 3:
		return 3, "出库成功"
	case 4:
		return 4, "交易成功"
	case -1:
		return -1, "手动关闭"
	case -2:
		return -2, "超时关闭"
	case -3:
		return -3, "商家关闭"
	default:
		return -9, "error"
	}
}

func (g MallOrderStatusEnum) Code() int {
	switch g {
	case ORDER_PRE_PAY:
		return 0
	case ORDER_PAID:
		return 1
	case ORDER_PACKAGED:
		return 2
	case ORDER_EXPRESS:
		return 3
	case ORDER_SUCCESS:
		return 4
	case ORDER_CLOSED_BY_MALLUSER:
		return -1
	case ORDER_CLOSED_BY_EXPIRED:
		return -2
	case ORDER_CLOSED_BY_JUDGE:
		return 3
	default:
		return -9
	}
}



type MallOrderStatusEnum int

const (
	DEFAULT                  MallOrderStatusEnum = -9
	ORDER_PRE_PAY            MallOrderStatusEnum = 0
	ORDER_PAID               MallOrderStatusEnum = 1
	ORDER_PACKAGED           MallOrderStatusEnum = 2
	ORDER_EXPRESS            MallOrderStatusEnum = 3
	ORDER_SUCCESS            MallOrderStatusEnum = 4
	ORDER_CLOSED_BY_MALLUSER MallOrderStatusEnum = -1
	ORDER_CLOSED_BY_EXPIRED  MallOrderStatusEnum = -2
	ORDER_CLOSED_BY_JUDGE    MallOrderStatusEnum = -3
)

func GetNewBeeMallOrderStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "待支付"
	case 1:
		return 1, "已支付"
	case 2:
		return 2, "配货完成"
	case 3:
		return 3, "出库成功"
	case 4:
		return 4, "交易成功"
	case -1:
		return -1, "手动关闭"
	case -2:
		return -2, "超时关闭"
	case -3:
		return -3, "商家关闭"
	default:
		return -9, "error"
	}
}

func (g MallOrderStatusEnum) Code() int {
	switch g {
	case ORDER_PRE_PAY:
		return 0
	case ORDER_PAID:
		return 1
	case ORDER_PACKAGED:
		return 2
	case ORDER_EXPRESS:
		return 3
	case ORDER_SUCCESS:
		return 4
	case ORDER_CLOSED_BY_MALLUSER:
		return -1
	case ORDER_CLOSED_BY_EXPIRED:
		return -2
	case ORDER_CLOSED_BY_JUDGE:
		return 3
	default:
		return -9
	}
}



type MallOrderStatusEnum int

const (
	DEFAULT                  MallOrderStatusEnum = -9
	ORDER_PRE_PAY            MallOrderStatusEnum = 0
	ORDER_PAID               MallOrderStatusEnum = 1
	ORDER_PACKAGED           MallOrderStatusEnum = 2
	ORDER_EXPRESS            MallOrderStatusEnum = 3
	ORDER_SUCCESS            MallOrderStatusEnum = 4
	ORDER_CLOSED_BY_MALLUSER MallOrderStatusEnum = -1
	ORDER_CLOSED_BY_EXPIRED  MallOrderStatusEnum = -2
	ORDER_CLOSED_BY_JUDGE    MallOrderStatusEnum = -3
)

func GetNewBeeMallOrderStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "待支付"
	case 1:
		return 1, "已支付"
	case 2:
		return 2, "配货完成"
	case 3:
		return 3, "出库成功"
	case 4:
		return 4, "交易成功"
	case -1:
		return -1, "手动关闭"
	case -2:
		return -2, "超时关闭"
	case -3:
		return -3, "商家关闭"
	default:
		return -9, "error"
	}
}

func (g MallOrderStatusEnum) Code() int {
	switch g {
	case ORDER_PRE_PAY:
		return 0
	case ORDER_PAID:
		return 1
	case ORDER_PACKAGED:
		return 2
	case ORDER_EXPRESS:
		return 3
	case ORDER_SUCCESS:
		return 4
	case ORDER_CLOSED_BY_MALLUSER:
		return -1
	case ORDER_CLOSED_BY_EXPIRED:
		return -2
	case ORDER_CLOSED_BY_JUDGE:
		return 3
	default:
		return -9
	}
}




type MallOrderStatusEnum int

const (
	DEFAULT                  MallOrderStatusEnum = -9
	ORDER_PRE_PAY            MallOrderStatusEnum = 0
	ORDER_PAID               MallOrderStatusEnum = 1
	ORDER_PACKAGED           MallOrderStatusEnum = 2
	ORDER_EXPRESS            MallOrderStatusEnum = 3
	ORDER_SUCCESS            MallOrderStatusEnum = 4
	ORDER_CLOSED_BY_MALLUSER MallOrderStatusEnum = -1
	ORDER_CLOSED_BY_EXPIRED  MallOrderStatusEnum = -2
	ORDER_CLOSED_BY_JUDGE    MallOrderStatusEnum = -3
)

func GetNewBeeMallOrderStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "待支付"
	case 1:
		return 1, "已支付"
	case 2:
		return 2, "配货完成"
	case 3:
		return 3, "出库成功"
	case 4:
		return 4, "交易成功"
	case -1:
		return -1, "手动关闭"
	case -2:
		return -2, "超时关闭"
	case -3:
		return -3, "商家关闭"
	default:
		return -9, "error"
	}
}

func (g MallOrderStatusEnum) Code() int {
	switch g {
	case ORDER_PRE_PAY:
		return 0
	case ORDER_PAID:
		return 1
	case ORDER_PACKAGED:
		return 2
	case ORDER_EXPRESS:
		return 3
	case ORDER_SUCCESS:
		return 4
	case ORDER_CLOSED_BY_MALLUSER:
		return -1
	case ORDER_CLOSED_BY_EXPIRED:
		return -2
	case ORDER_CLOSED_BY_JUDGE:
		return 3
	default:
		return -9
	}
}




type MallOrderStatusEnum int

const (
	DEFAULT                  MallOrderStatusEnum = -9
	ORDER_PRE_PAY            MallOrderStatusEnum = 0
	ORDER_PAID               MallOrderStatusEnum = 1
	ORDER_PACKAGED           MallOrderStatusEnum = 2
	ORDER_EXPRESS            MallOrderStatusEnum = 3
	ORDER_SUCCESS            MallOrderStatusEnum = 4
	ORDER_CLOSED_BY_MALLUSER MallOrderStatusEnum = -1
	ORDER_CLOSED_BY_EXPIRED  MallOrderStatusEnum = -2
	ORDER_CLOSED_BY_JUDGE    MallOrderStatusEnum = -3
)

func GetNewBeeMallOrderStatusEnumByStatus(status int) (int, string) {
	switch status {
	case 0:
		return 0, "待支付"
	case 1:
		return 1, "已支付"
	case 2:
		return 2, "配货完成"
	case 3:
		return 3, "出库成功"
	case 4:
		return 4, "交易成功"
	case -1:
		return -1, "手动关闭"
	case -2:
		return -2, "超时关闭"
	case -3:
		return -3, "商家关闭"
	default:
		return -9, "error"
	}
}

func (g MallOrderStatusEnum) Code() int {
	switch g {
	case ORDER_PRE_PAY:
		return 0
	case ORDER_PAID:
		return 1
	case ORDER_PACKAGED:
		return 2
	case ORDER_EXPRESS:
		return 3
	case ORDER_SUCCESS:
		return 4
	case ORDER_CLOSED_BY_MALLUSER:
		return -1
	case ORDER_CLOSED_BY_EXPIRED:
		return -2
	case ORDER_CLOSED_BY_JUDGE:
		return 3
	default:
		return -9
	}
}