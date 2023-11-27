package channel

import (
	"context"
	"github.com/xssor2600/xpaySDK/channel/kuaishou"
	"github.com/xssor2600/xpaySDK/config"
)

// InstanceChannelCenter integration channel service center
type InstanceChannelCenter interface {
	GetChannelInstance(ctx context.Context, config config.GlobalConfig) (interface{}, error)
}

type InstanceChannelHandler struct {
	ChannelName string `json:"channel_name"`
}

func (ich *InstanceChannelHandler) GetChannelHandler(ctx context.Context, globalConfig config.GlobalConfig) (interface{}, error) {
	switch ich.ChannelName {
	case config.CHANNEL_ALIPAY:

	case config.CHANNEL_WECHAT:

	case config.CHANNEL_GOOGLE:

	case config.CHANNEL_PAYPAL:

	case config.CHANNEL_TOUTIAO:

	case config.CHANNEL_KUAISHOU:
		ksc, err := globalConfig.GetChannelConfig(ich.ChannelName)
		if err != nil {
			return nil, err
		}
		if kc, ok := ksc.(*config.KuaishouConfig); ok {
			return &kuaishou.KsApi{*kc}, nil
		}
	}

	return nil, nil
}
