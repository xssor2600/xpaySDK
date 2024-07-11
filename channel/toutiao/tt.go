package toutiao

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xssor2600/xpaySDK/config"
	"github.com/xssor2600/xpaySDK/dto"
	"github.com/xssor2600/xpaySDK/utils"
	"io"
	"net/http"
	"strings"
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
		OrderEntrySchema: dto.SkuSchema{
			Path:   order.OrderEntrySchema.Path,
			Params: order.OrderEntrySchema.Params,
		},
		LimitPayWayList: nil,
		PayScene:        order.PayScene,
		GoodsDetail:     order.GoodsDetail,
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
		dto.PayScene(payReq.PayScene),
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

func (tt *TTradeApi) CallToutiaoTradeAPI(ctx context.Context, method string, reqURL string, pm utils.ParamMap, token string, resp interface{}) error {
	appId := strings.TrimSpace(tt.ToutiaoConfig.AppId)
	nonceStr, timestamp := utils.GetNonceStr(), utils.GetTimeStamp()
	keyVersion := utils.GetKeyVersion(tt.ToutiaoConfig.KeyVersion)
	body := utils.JonsObject(pm)

	keyBytes, err := utils.ReadPemFile(fmt.Sprintf("config/channel_config/%s", tt.ToutiaoConfig.MerchantPrivateKey))
	if err != nil {
		panic("keyBytes err")
	}
	rsaPrivateKey, keyErr := utils.ParsePKCS1PrivateKey(keyBytes)
	if keyErr != nil {
		panic("rsaPrivateKey err")
	}

	signValue, signErr := utils.GenChannelSign(method, reqURL, timestamp, nonceStr, body, rsaPrivateKey)
	if signErr != nil {
		return errors.New("sign err")
	}

	// 完整http请求地址
	fullRequestUrl := fmt.Sprintf("%s%s", tt.ToutiaoConfig.ApisUrl.BaseUrl, reqURL)

	header := map[string]string{
		"access-token":       token,
		"Content-Type":       "application/json",
		"Byte-Authorization": utils.GetByteAuth(appId, nonceStr, timestamp, keyVersion, signValue),
	}
	responseData, err := utils.DoHttpRequestJson(ctx, method, fullRequestUrl, pm, header)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(responseData, &resp); err != nil {
		return errors.New("json.Unmarshal err")
	}
	return nil
}

// 回掉验证签名
func (tt *TTradeApi) CallBackSignVerify(request http.Request) error {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}
	req := new(dto.ChannelCallBackCommonMsg)
	if err := json.Unmarshal(data, req); err != nil {
		return err
	}
	// 请求头内容获取
	byteTimestamp, byteNonceStr := request.Header.Get("Byte-Timestamp"), request.Header.Get("Byte-Nonce-Str")
	byteSignature := request.Header.Get("Byte-Signature")

	// 获取平台公钥
	keyBytes, err := utils.ReadPemFile(fmt.Sprintf("config/channel_config/%s", tt.ToutiaoConfig.AppPublicKey))
	if err != nil {
		panic("keyBytes err")
	}

	if signErr := utils.CheckChannelSign(byteTimestamp, byteNonceStr, string(data), byteSignature, strings.TrimSpace(string(keyBytes))); signErr != nil {
		return signErr
	}

	return nil
}
