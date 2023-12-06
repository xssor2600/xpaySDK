package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var (
	configPath        string
	defaultConfigPath = "config/channel_config"
)

var (
	once             sync.Once
	ChannelConfigMap sync.Map

	// channel config object.
	alipayConfig AlipayConfig
	wechatConfig WechatConfig
	ksConfig     KuaishouConfig
	ggConfig     GoogleConfig
	ttConfig     ToutiaoConfig
)

func init() {
	workPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	configPath = filepath.Join(workPath, defaultConfigPath)
	log.Println(configPath)
}

// load channel json config
func init() {
	once.Do(func() {
		//LoadJsonConfig(&alipayConfig,"alipay.json")
		//LoadJsonConfig(&wechatConfig,"wechat.json")
		//LoadJsonConfig(&ttConfig,"toutiao.json")
		//LoadJsonConfig(&ggConfig,"google.json")
		LoadJsonConfig(&ksConfig, "kuaishou.json")
		//ChannelConfigMap.Store(CHANNEL_ALIPAY,&alipayConfig)
		//ChannelConfigMap.Store(CHANNEL_WECHAT,&alipayConfig)
		//ChannelConfigMap.Store(CHANNEL_TOUTIAO,&ttConfig)
		//ChannelConfigMap.Store(CHANNEL_GOOGLE,&ggConfig)
		ChannelConfigMap.Store(CHANNEL_KUAISHOU, &ksConfig)
	})
}

type GlobalConfig interface {
	GetChannelConfig(channel string) (interface{}, error)
}

func LoadJsonConfig(config interface{}, fileName string) {
	var err error
	var decoder *json.Decoder

	file := OpenFile(fileName)
	defer file.Close()
	decoder = json.NewDecoder(file)
	if err = decoder.Decode(config); err != nil {
		log.Fatal(err)
	}
}

func OpenFile(filename string) *os.File {
	fullPath := filepath.Join(configPath, filename)

	var file *os.File
	var err error

	if file, err = os.Open(fullPath); err != nil {
		msg := fmt.Sprintf("Can not load config at %s. Error: %v", fullPath, err)
		log.Fatal(msg)
	}

	return file
}
