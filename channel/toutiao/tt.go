package toutiao

import (
	"context"
	"fmt"
	"github.com/xssor2600/xpaySDK/config"
	"github.com/xssor2600/xpaySDK/dto"
	"github.com/xssor2600/xpaySDK/utils"
	"net/http"
)

type TTradeApi struct {
	ToutiaoConfig config.ToutiaoConfig `json:"toutiao_config"`
}

func (tt *TTradeApi) CreateOrder(ctx context.Context, order *dto.ToutiaoPayReq) (interface{}, error) {
	payReq := dto.ToutiaoPayReq{
		Uid:              order.Uid,
		Subject:          order.Subject,
		SkuList:          order.SkuList,
		OutOrderNo:       order.OutOrderNo,
		TotalAmount:      order.TotalAmount,
		PayExpireSeconds: order.PayExpireSeconds,
		PayNotifyUrl:     "",
		MerchantUid:      order.MerchantUid,
		OrderEntrySchema: dto.SkuSchema{},
		LimitPayWayList:  nil,
		PayScene:         "",
		GoodsDetail:      order.GoodsDetail,
	}
	orderParams := dto.NewTradeCreateOrder(
		payReq.OutOrderNo,
		payReq.TotalAmount,
		payReq.PayExpireSeconds,
		payReq.MerchantUid,
		dto.PayNotifyUrl("", &tt.ToutiaoConfig),
		dto.SkuList(payReq.Uid, payReq.Subject, payReq.GoodsDetail, payReq.SkuList[0]),
		dto.LimitPayWayList(&payReq),
		dto.OrderEntrySchema(payReq.GoodsDetail, payReq.SkuList[0]),
		dto.TotalAmount(payReq.TotalAmount, payReq.GoodsDetail),
		dto.PayScene(""),
	)

	keyBytes, err := utils.ReadPemFile(fmt.Sprintf("config/channel_config/%s", tt.ToutiaoConfig.MerchantPrivateKey))
	if err != nil {
		panic("keyBytes err")
	}
	rsaPrivateKey, keyErr := utils.ParsePKCS1PrivateKey(keyBytes)
	if keyErr != nil {
		panic("rsaPrivateKey err")
	}

	nonceStr, timestamp := utils.GetNonceStr(), utils.GetTimeStamp()
	createOrderBody := utils.JonsObject(orderParams)
	orderSign, _ := utils.GenChannelSign(http.MethodPost, "/requestOrder", timestamp, nonceStr, createOrderBody, rsaPrivateKey)

	return dto.TradeCreateOrderResp{
		Data:              createOrderBody,
		ByteAuthorization: utils.GetByteAuth(tt.ToutiaoConfig.AppId, nonceStr, timestamp, utils.GetKeyVersion(tt.ToutiaoConfig.KeyVersion), orderSign),
	}, nil
}
