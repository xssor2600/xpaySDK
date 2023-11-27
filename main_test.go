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
	ctx := context.Background()
	ksConfig := &config.KuaishouConfig{}
	handler := channel.InstanceChannelHandler{ChannelName: config.CHANNEL_KUAISHOU}
	ks, err := handler.GetChannelHandler(context.Background(), ksConfig)
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
		orderResp, err := ksApi.EpayCreateOrder(ctx, &ksOrderReq)
		fmt.Println(utils.JsonMashObject(orderResp), err)
	}

}
