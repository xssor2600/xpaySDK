package main

import (
	"context"
	"fmt"
	"github.com/xssor2600/xpaySDK/channel"
	"github.com/xssor2600/xpaySDK/channel/kuaishou"
	"github.com/xssor2600/xpaySDK/channel/toutiao"
	"github.com/xssor2600/xpaySDK/config"
	"github.com/xssor2600/xpaySDK/dto"
	"github.com/xssor2600/xpaySDK/utils"
	"log"
	"testing"
)

func Test_channelConfig(t *testing.T) {
	//if v, ok := config.ChannelConfigMap.Load(config.CHANNEL_KUAISHOU); ok {
	//	if kc, ok := v.(*config.KuaishouConfig); ok {
	//		fmt.Println(kc.AppId)
	//		fmt.Println(kc.ApisUrl.EpayCreateOrder)
	//		fmt.Println(kc.ApiNotifyUrl.PayNotify)
	//	}
	//}

	if v, ok := config.ChannelConfigMap.Load(config.CHANNEL_TOUTIAO); ok {
		if trade, ok := v.(*config.ToutiaoConfig); ok {
			fmt.Println(trade.AppId)
			fmt.Println(trade.ApisUrl.RefundUrl)
			fmt.Println(trade.ApiNotifyUrl.PayNotify)
		}
	}
}

func Test_channelKsApi(t *testing.T) {
	handler := channel.InstanceChannelFactory{ChannelName: config.CHANNEL_KUAISHOU}
	ks, err := handler.GetChannelHandler(context.Background(), &config.KuaishouConfig{})
	if err != nil {
		log.Fatalln(err)
	}
	if ksApi, ok := ks.(*kuaishou.KsApi); ok {
		ksOrderReq := dto.KsPayReq{
			OpenId:  "f18db489fa3fa6b914bf572468ee6dff",
			OrderId: "KS2022083155556666666",
			Money:   10,
			Subject: "test",
			Detail:  "test",
		}
		orderResp, err := ksApi.EpayCreateOrder(context.Background(), &ksOrderReq)
		fmt.Println(utils.JsonMashObject(orderResp), err)
	}

}

func Test_channelToutiaoApi(t *testing.T) {
	handler := channel.InstanceChannelFactory{ChannelName: config.CHANNEL_TOUTIAO}
	trade, err := handler.GetChannelHandler(context.Background(), &config.ToutiaoConfig{})
	if err != nil {
		log.Fatalln(err)
	}
	if tradeApi, ok := trade.(*toutiao.TTradeApi); ok {
		payReq := dto.ToutiaoPayReq{
			Uid:     0,
			Subject: "",
			SkuList: []dto.SkuInfo{{
				SkuId:       "123",
				Price:       10,
				Quantity:    1,
				Title:       "test",
				ImageList:   nil,
				Type:        10,
				TagGroupId:  "",
				EntrySchema: dto.SkuSchema{},
			}},
			OutOrderNo:       "",
			TotalAmount:      0,
			PayExpireSeconds: 0,
			PayNotifyUrl:     "",
			MerchantUid:      "",
			OrderEntrySchema: dto.SkuSchema{},
			LimitPayWayList:  nil,
			PayScene:         "",
			GoodsDetail: []dto.ToutiaoTradeGoodsDetail{{
				GoodsId:         "1",
				GoodsName:       "1",
				GoodsPrice:      10,
				Quantity:        1,
				ImageList:       nil,
				GoodsType:       0,
				GoodsDetailPage: dto.GoodsPage{},
			}},
		}
		resp, _ := tradeApi.CreateOrder(context.Background(), &payReq)
		fmt.Println(resp)
	}

}
