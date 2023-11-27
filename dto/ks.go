package dto



type KsPayReq struct {
	OpenId string `json:"open_id"`
	OrderId string `json:"order_id"`
	Money int64 `json:"money"`
	Subject string `json:"subject"`
	Detail string `json:"detail"`
}


type CreatePayResp struct {
	Result    int64  `json:"result"`
	ErrorMsg  string `json:"error_msg"`
	OrderInfo struct {
		OrderNo        string `json:"order_no"`
		OrderInfoToken string `json:"order_info_token"`
	} `json:"order_info,omitempty"`
}