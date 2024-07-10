package config

import (
	"errors"
)

type AlipayConfig struct {
}

type GoogleConfig struct {
}

type KuaishouConfig struct {
	AppId        string `json:"app_id"`
	AppSecret    string `json:"app_secret"`
	CategoryType int64  `json:"category_type"`
	AccessToken  string `json:"access_token"`
	ApisUrl      struct {
		EpayCreateOrder string `json:"epay_create_order"`
		EpayQueryOrder  string `json:"epay_query_order"`
		Refund          string `json:"refund_url"`
		RefundQuery     string `json:"refund_query"`
		Settle          string `json:"settle_url"`
		SettleQuery     string `json:"settle_query_url"`
	} `json:"apis_url"`
	ApiNotifyUrl struct {
		PayNotify    string `json:"pay_notify"`
		SettleNotify string `json:"settle_notify"`
		RefundNotify string `json:"refund_notify"`
	} `json:"api_notify_url"`
}

type ToutiaoConfig struct {
	AppId              string `json:"app_id"`
	AppPublicKey       string `json:"app_public_key"`
	MerchantPrivateKey string `json:"merchant_private_key"`
	AccessToken        string `json:"access_token"`
	PicUrl             string `json:"pic_url"`
	DetailUrl          string `json:"detail_url"`
	KeyVersion         string `json:"key_version"`
	ProductType        int    `json:"product_type"`
	TagGroupId         string `json:"tag_group_id"`
	ApisUrl            struct {
		BaseUrl        string `json:"base_url"`
		CreateOrder    string `json:"create_order"`
		QueryOrder     string `json:"query_order"`
		RefundUrl      string `json:"refund_url"`
		RefundQuery    string `json:"refund_query"`
		SettleUrl      string `json:"settle_url"`
		SettleQueryUrl string `json:"settle_query_url"`
	} `json:"apis_url"`
	ApiNotifyUrl struct {
		PayNotify    string `json:"pay_notify"`
		SettleNotify string `json:"settle_notify"`
		RefundNotify string `json:"refund_notify"`
	} `json:"api_notify_url"`
}

type WechatConfig struct {
}

type AppleConfig struct {
}

func (gbc *KuaishouConfig) GetChannelConfig(channel string) (interface{}, error) {
	if channel == "" {
		return nil, errors.New("GetChannelConfig channel is empty ")
	}
	if cconfig, ok := ChannelConfigMap.Load(channel); ok {
		return cconfig, nil
	}
	return nil, nil
}

func (tt *ToutiaoConfig) GetChannelConfig(channel string) (interface{}, error) {
	if channel == "" {
		return nil, errors.New("GetChannelConfig channel is empty ")
	}
	if cconfig, ok := ChannelConfigMap.Load(channel); ok {
		return cconfig, nil
	}
	return nil, nil
}
