package dto

type ToutiaoPayReq struct {
	SkuList          []SkuInfo `json:"skuList"`
	OutOrderNo       string    `json:"outOrderNo"`
	TotalAmount      int64     `json:"totalAmount"`      // 订单总金额
	PayExpireSeconds int64     `json:"payExpireSeconds"` // 支付超时时间，单位秒，例如 300 表示 300 秒后过期；不传或传 0 会使用默认值 300，不能超过48小时。
	PayNotifyUrl     string    `json:"payNotifyUrl"`
	MerchantUid      string    `json:"merchantUid"`
	OrderEntrySchema SkuSchema `json:"orderEntrySchema"`
	LimitPayWayList  []int64   `json:"limitPayWayList"`    // 屏蔽的支付方式，当开发者没有进件某个支付渠道，可在下单时屏蔽对应的支付方式。如：[1, 2]表示屏蔽微信和支付宝 10-抖音支付
	PayScene         string    `json:"payScene,omitempty"` // 指定支付场景，这个是H5支付常用
}

type SkuInfo struct {
	SkuId       string    `json:"skuId"`
	Price       int64     `json:"price"`
	Quantity    int64     `json:"quantity"`
	Title       string    `json:"title"`
	ImageList   []string  `json:"imageList"`
	Type        int64     `json:"type"` // 301：虚拟工具类商品
	TagGroupId  string    `json:"tagGroupId"`
	EntrySchema SkuSchema `json:"entrySchema"`
}

type SkuSchema struct {
	Path   string `json:"path"`
	Params string `json:"params"`
}
