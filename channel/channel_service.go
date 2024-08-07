package channel

import (
	"context"
	"github.com/xssor2600/xpaySDK/channel/kuaishou"
	"github.com/xssor2600/xpaySDK/channel/toutiao"
	"github.com/xssor2600/xpaySDK/config"
)

type InstanceChannelFactory struct {
	ChannelName string `json:"channel_name"`
}

func (ich *InstanceChannelFactory) GetChannelHandler(ctx context.Context, globalConfig config.GlobalConfig) (interface{}, error) {
	switch ich.ChannelName {
	case config.CHANNEL_ALIPAY:

	case config.CHANNEL_WECHAT:

	case config.CHANNEL_GOOGLE:

	case config.CHANNEL_PAYPAL:

	case config.CHANNEL_TOUTIAO:
		ttc, err := globalConfig.GetChannelConfig(ich.ChannelName)
		if err != nil {
			return nil, err
		}
		if trade, ok := ttc.(*config.ToutiaoConfig); ok {
			return &toutiao.TTradeApi{*trade}, nil
		}
		break

	case config.CHANNEL_KUAISHOU:
		ksc, err := globalConfig.GetChannelConfig(ich.ChannelName)
		if err != nil {
			return nil, err
		}
		if kc, ok := ksc.(*config.KuaishouConfig); ok {
			return &kuaishou.KsApi{KsConfig: *kc}, nil
		}
		break
	}
	return nil, nil
}
