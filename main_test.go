package main

import (
	"context"
	"fmt"
	"github.com/xssor2600/xpaySDK/channel"
	"github.com/xssor2600/xpaySDK/channel/kuaishou"
	"github.com/xssor2600/xpaySDK/config"
	"github.com/xssor2600/xpaySDK/dto"
	"github.com/xssor2600/xpaySDK/utils"
	"log"
	"testing"
)

func Test_channelConfig(t *testing.T) {
	if v, ok := config.ChannelConfigMap.Load(config.CHANNEL_KUAISHOU); ok {
		if kc, ok := v.(*config.KuaishouConfig); ok {
			fmt.Println(kc.AppId)
			fmt.Println(kc.ApisUrl.EpayCreateOrder)
			fmt.Println(kc.ApiNotifyUrl.PayNotify)
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
	ks, err := handler.GetChannelHandler(context.Background(), &config.ToutiaoConfig{})
	if err != nil {
		log.Fatalln(err)
	}
	if ksApi, ok := ks.(*kuaishou.KsApi); ok {
		ksOrderReq := dto.KsPayReq{
			OpenId:  "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX",
			OrderId: "XXXXXXXXXXXXXXXXXXXXX",
			Money:   10,
			Subject: "test",
			Detail:  "test",
		}
		orderResp, err := ksApi.EpayCreateOrder(context.Background(), &ksOrderReq)
		fmt.Println(utils.JsonMashObject(orderResp), err)
	}

}
