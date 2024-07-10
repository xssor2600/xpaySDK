package kuaishou

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
	"time"
)

type KsApi struct {
	KsConfig config.KuaishouConfig `json:"ks_config"`
}

func (ks *KsApi) EpayCreateOrder(ctx context.Context, ksReq *dto.KsPayReq) (interface{}, error) {
	validTime := time.Now().Add(time.Minute*30).Unix() - time.Now().Unix()
	pm := make(utils.ParamMap)
	pm.Set("out_order_no", ksReq.OrderId).
		Set("open_id", ksReq.OpenId).
		Set("total_amount", ksReq.Money).
		Set("subject", ksReq.Subject).
		Set("detail", ksReq.Detail).
		Set("type", ks.KsConfig.CategoryType).
		Set("expire_time", validTime).
		Set("notify_url", ks.KsConfig.ApiNotifyUrl.PayNotify).
		Set("app_id", ks.KsConfig.AppId)
	pm.Set("sign", utils.KsSignFromMap(pm, ks.KsConfig.AppSecret))
	requestUrl := fmt.Sprintf("%s?app_id=%s&access_token=%s", ks.KsConfig.ApisUrl.EpayCreateOrder, ks.KsConfig.AppId, ks.KsConfig.AccessToken)
	respbyte, err := utils.DoHttpRequestJson(ctx, http.MethodPost, requestUrl, pm, nil)
	if err != nil {
		return nil, err
	}
	payResp := dto.CreatePayResp{}
	utils.JsonUnMashObject(respbyte, &payResp)
	return payResp, nil
}

// 回掉验证签名
func (ks *KsApi) CallBackSignVerify(request http.Request) error {
	data, err := io.ReadAll(request.Body)
	if err != nil {
		return err
	}
	req := new(dto.NotifyReq)
	if err := json.Unmarshal(data, req); err != nil {
		return err
	}
	kwaisign := request.Header.Get("kwaisign")
	if dataSign := utils.Md5BackSign(string(data), ks.KsConfig.AppSecret); dataSign != kwaisign {
		return errors.New("sign verify err")
	}
	return nil
}
