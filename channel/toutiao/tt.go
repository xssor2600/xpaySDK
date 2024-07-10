package toutiao

import (
	"context"
	"github.com/xssor2600/xpaySDK/config"
	"github.com/xssor2600/xpaySDK/dto"
)

type TTradeApi struct {
	ToutiaoConfig config.ToutiaoConfig `json:"toutiao_config"`
}

func (tt *TTradeApi) CreateOrder(ctx context.Context, order *dto.ToutiaoPayReq) (interface{}, error) {

	return nil, nil
}
