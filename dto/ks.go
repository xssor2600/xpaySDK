package dto

type KsPayReq struct {
	OpenId  string `json:"open_id"`
	OrderId string `json:"order_id"`
	Money   int64  `json:"money"`
	Subject string `json:"subject"`
	Detail  string `json:"detail"`
}

type CreatePayResp struct {
	Result    int64  `json:"result"`
	ErrorMsg  string `json:"error_msg"`
	OrderInfo struct {
		OrderNo        string `json:"order_no"`
		OrderInfoToken string `json:"order_info_token"`
	} `json:"order_info,omitempty"`
}

type NotifyReq struct {
	Data      NotifyDataInfo `json:"data" schema:"data"`
	BizType   string         `json:"biz_type,omitempty" schema:"biz_type"`
	MessageId string         `json:"message_id,omitempty" schema:"message_id"`
	AppId     string         `json:"app_id,omitempty" schema:"app_id"`
	Timestamp int64          `json:"timestamp,omitempty" schema:"timestamp"`
}

type NotifyDataInfo struct {
	Channel     string `json:"channel,omitempty"`      // UNKNOWN - 未知｜WECHAT-微信｜ALIPAY-支付宝
	OutOrderNo  string `json:"out_order_no,omitempty"` // 商户系统内部订单号
	Attach      string `json:"attach,omitempty"`       // 预下单时携带的开发者自定义信息
	Status      string `json:"status,omitempty"`       // 订单支付状态 。 取值： PROCESSING-处理中｜SUCCESS-成功｜FAILED-失败 |  PROCESSING-处理中，SUCCESS-成功，FAILED-失败
	KsOrderNo   string `json:"ks_order_no,omitempty"`  // 快手小程序平台订单号
	OrderAmount int64  `json:"order_amount,omitempty"` // 订单金额
	TradeNo     string `json:"trade_no,omitempty"`     // 用户侧支付页交易单号 ( 支付渠道交易订单)
	//ExtraInfo       string `json:"extra_info,omitempty"`       // 订单来源信息，同支付查询接口
	EnablePromotion bool `json:"enable_promotion,omitempty"` // 是否参与分销，true:分销，false:非分销
	PromotionAmount int  `json:"promotion_amount,omitempty"` // 预计分销金额，单位：分

	// 退款
	KsRefundNo   string `json:"ks_refund_no,omitempty"`  // 快手小程序平台退款单号。
	OutRefundNo  string `json:"out_refund_no,omitempty"` // 开发者的退款单号。
	OutSettleNo  string `json:"out_settle_no,omitempty"` // 快手小程序平台订单号
	RefundAmount int64  `json:"refund_amount,omitempty"` // 订单金额
	SettleAmount int64  `json:"settle_amount,omitempty"` // 订单金额
	KsSettleNo   string `json:"ks_settle_no,omitempty"`  // 订单金额

	SkuId  string `json:"sku_id,omitempty"`  // 用户购的一个sku_id，该id是通过支付面板拉起时传入的skuIdList中的sku_id（手支付）
	OpenId string `json:"open_id,omitempty"` // 下单用户id
	Origin string `json:"origin,omitempty"`  // 快手回调的来源字段
}
