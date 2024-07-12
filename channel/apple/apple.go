package apple

// StoreKit v2 support:  https://developer.apple.com/videos/play/wwdc2023/10141

type AppleConfig struct {
	Package                string `json:"package"` // 包名
	AppId                  string `json:"app_id"`  // 应用id
	KitApi                 string `json:"kit_api"` // api
	SandboxKitApi          string `json:"sandbox_kit_api"`
	ReceiptApi             string `json:"receipt_api"` // 票据解析地址
	SandboxReceiptApi      string `json:"sandbox_receipt_api"`
	OpenShareCode          string `json:"open_share_code"`          // 共享密钥
	Products               string `json:"products"`                 // apple后台配置商品信息json格式
	PublishId              string `json:"publish_id"`               // 发行方
	PrivateKey             string `json:"private_key"`              // 密钥
	SignKey                string `json:"sign_key"`                 // 签名私钥
	ServerNotificationsUrl string `json:"server_notifications_url"` // App Store 服务器主动通知开发者服务器的 API。比如退款通知、订阅商品续费成功通知等 (V1 版本：响应内容是 JSON 格式的数据/V2 版本：响应内容是由 App Store 签名的JSON Web签名（JWS）格式。)
}
