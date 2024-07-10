package dto

import (
	"fmt"
	"github.com/xssor2600/xpaySDK/config"
	"strings"
)

type ToutiaoPayReq struct {
	Uid              int64                     `json:"uid"`
	Subject          string                    `json:"subject"`
	SkuList          []SkuInfo                 `json:"skuList"`
	OutOrderNo       string                    `json:"outOrderNo"`
	TotalAmount      int64                     `json:"totalAmount"`      // 订单总金额
	PayExpireSeconds int64                     `json:"payExpireSeconds"` // 支付超时时间，单位秒，例如 300 表示 300 秒后过期；不传或传 0 会使用默认值 300，不能超过48小时。
	PayNotifyUrl     string                    `json:"payNotifyUrl"`
	MerchantUid      string                    `json:"merchantUid"`
	OrderEntrySchema SkuSchema                 `json:"orderEntrySchema"`
	LimitPayWayList  []int64                   `json:"limitPayWayList"`    // 屏蔽的支付方式，当开发者没有进件某个支付渠道，可在下单时屏蔽对应的支付方式。如：[1, 2]表示屏蔽微信和支付宝 10-抖音支付
	PayScene         string                    `json:"payScene,omitempty"` // 指定支付场景，这个是H5支付常用
	GoodsDetail      []ToutiaoTradeGoodsDetail `json:"goodsDetail"`
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

type ToutiaoTradeGoodsDetail struct {
	GoodsId         string    `json:"goods_id" validate:"required"`           // 商品ID
	GoodsName       string    `json:"goods_name" validate:"required"`         // 商品名称
	GoodsPrice      int64     `json:"price" validate:"required"`              //价格
	Quantity        int64     `json:"quantity" validate:"required"`           // 数量
	ImageList       []string  `json:"image_list" validate:"required,max=512"` // 商品图片链接，长度 <= 512 字节
	GoodsType       int       `json:"sku_type"`                               // 商品类型 301：虚拟工具类商品
	GoodsDetailPage GoodsPage `json:"sku_detail_page"`                        //商品详情页链接
}

type GoodsPage struct {
	Path   string `json:"path"`                               // 小程序xxx详情页跳转路径，没有前导的“/”，路径后不可携带query参数
	Params string `json:"params" validate:"required,max=512"` // xx情页路径参数，自定义的json结构，内部为k-v结构
}

// https://developer.open-douyin.com/docs/resource/zh-CN/mini-app/develop/api/industry/general_trade/create_order/requestOrder
// TradeCreateOrderData tt.requestOrder.Data  下单对象
type TradeCreateOrderData struct {
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

type Option func(*TradeCreateOrderData)

func NewTradeCreateOrder(orderId string, totalAmount int64, payExpiredSeconds int64, MerchantUid string, options ...Option) *TradeCreateOrderData {
	createOrder := &TradeCreateOrderData{
		OutOrderNo:       orderId,
		PayExpireSeconds: payExpiredSeconds,
		//PayNotifyUrl:     PayNotifyUrl,
		MerchantUid: MerchantUid,
		TotalAmount: totalAmount,
	}

	for _, op := range options {
		op(createOrder)
	}
	return createOrder
}

func SkuList(uid int64, subject string, skuInfos []ToutiaoTradeGoodsDetail, skuInfo SkuInfo) Option {
	return func(order *TradeCreateOrderData) {
		if len(skuInfos) > 0 {
			for _, sku := range skuInfos {
				order.SkuList = []SkuInfo{{
					SkuId: func(skuId string) string {
						if skuId != "" {
							return skuId
						} else {
							return fmt.Sprintf("%d", uid)
						}
					}(sku.GoodsId),
					Price:      sku.GoodsPrice,
					Quantity:   sku.Quantity,
					Title:      sku.GoodsName,
					ImageList:  sku.ImageList, // 商品图片链接，长度 <= 512 字节
					Type:       int64(sku.GoodsType),
					TagGroupId: skuInfo.TagGroupId,
					EntrySchema: SkuSchema{
						Path:   strings.TrimLeft(sku.GoodsDetailPage.Path, "/"),
						Params: sku.GoodsDetailPage.Params,
					},
				}}
			}

		} else {
			order.SkuList = []SkuInfo{{
				SkuId:    fmt.Sprintf("%d", uid),
				Price:    order.TotalAmount,
				Quantity: 1,
				Title: func(title string) string {
					if subject != "" {
						return subject
					} else {
						return "短剧充值"
					}
				}(subject),
				ImageList:  []string{skuInfo.ImageList[0]}, // 商品图片链接，长度 <= 512 字节
				Type:       401,
				TagGroupId: "tag_group_7272625659888041996",
				EntrySchema: SkuSchema{
					Path:   strings.TrimLeft(skuInfo.EntrySchema.Path, "/"),
					Params: skuInfo.EntrySchema.Params,
				},
			}}
		}
	}
}

func TotalAmount(payMoney int64, skuInfos []ToutiaoTradeGoodsDetail) Option {
	return func(order *TradeCreateOrderData) {
		if len(skuInfos) > 0 {
			order.TotalAmount = skuInfos[0].GoodsPrice
		} else {
			order.TotalAmount = payMoney
		}
	}
}

func PayScene(tradeType string) Option {
	return func(order *TradeCreateOrderData) {
		// H5支付传，指定交易场景即可
		if tradeType == "H5" {
			order.PayScene = "IM"
		}
	}
}

// 校验params是否是有效的json
func OrderEntrySchema(skuInfos []ToutiaoTradeGoodsDetail, skuInfo SkuInfo) Option {
	return func(order *TradeCreateOrderData) {
		if len(skuInfos) > 0 {
			order.OrderEntrySchema = SkuSchema{
				Path:   strings.TrimLeft(skuInfos[0].GoodsDetailPage.Path, "/"),
				Params: skuInfos[0].GoodsDetailPage.Params,
			}
		} else {
			order.OrderEntrySchema = SkuSchema{
				Path:   strings.TrimLeft(skuInfo.EntrySchema.Path, "/"),
				Params: skuInfo.EntrySchema.Params,
			}
		}
	}
}

func LimitPayWayList(payReq *ToutiaoPayReq) Option {
	return func(order *TradeCreateOrderData) {
		if len(payReq.LimitPayWayList) > 0 {
			order.LimitPayWayList = payReq.LimitPayWayList
		} else {
			order.LimitPayWayList = make([]int64, 0)
		}
	}
}

func PayNotifyUrl(payUrl string, ttConfig *config.ToutiaoConfig) Option {
	return func(order *TradeCreateOrderData) {
		if len(payUrl) > 0 {
			order.PayNotifyUrl = payUrl
		} else {
			order.PayNotifyUrl = ttConfig.ApiNotifyUrl.PayNotify
		}
	}
}

type TradeCreateOrderResp struct {
	Data              string `json:"data"`
	ByteAuthorization string `json:"byteAuthorization"`
}
