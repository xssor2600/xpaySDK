package utils

import (
	"encoding/json"
	"errors"
	"net/url"
	"sort"
	"strings"
)

type ParamMap map[string]interface{}

// Set 设置参数
func (pm ParamMap) Set(key string, value interface{}) ParamMap {
	pm[key] = value
	return pm
}

func (pm ParamMap) SetBodyMap(key string, value func(b ParamMap)) ParamMap {
	_bm := make(ParamMap)
	value(_bm)
	pm[key] = _bm
	return pm
}

// Get 获取参数
func (pm ParamMap) Get(key string) string {
	if pm == nil {
		return ""
	}
	value, ok := pm[key]
	if !ok {
		return ""
	}
	v, ok := value.(string)
	if !ok {
		return convertToString(value)
	}
	return v
}

// GetInterface 获取原始参数
func (pm ParamMap) GetInterface(key string) interface{} {
	if pm == nil {
		return nil
	}
	return pm[key]
}

// Remove 删除参数
func (pm ParamMap) Remove(key string) {
	delete(pm, key)
}

// Reset 置空BodyMap
func (pm ParamMap) Reset() {
	for k := range pm {
		delete(pm, k)
	}
}

func (pm ParamMap) ToJsonString() (jb string) {
	bs, err := json.Marshal(pm)
	if err != nil {
		return ""
	}
	jb = string(bs)
	return jb
}

// Unmarshal to struct or slice point
func (pm ParamMap) Unmarshal(ptr interface{}) (err error) {
	bs, err := json.Marshal(pm)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, ptr)
}

// EncodeSignParams ("bar=baz&foo=quux") sorted by key.
func (pm ParamMap) EncodeSignParams() string {
	if pm == nil {
		return ""
	}
	var (
		buf     strings.Builder
		keyList []string
	)
	for k := range pm {
		keyList = append(keyList, k)
	}
	sort.Strings(keyList)
	for _, k := range keyList {
		if v := pm.Get(k); v != "" && k != "sign" {
			buf.WriteString(k)
			buf.WriteByte('=')
			buf.WriteString(v)
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}

// EncodeURLParams  ("bar=baz&foo=quux") sorted by key.
func (pm ParamMap) EncodeURLParams() string {
	if pm == nil {
		return ""
	}
	var (
		buf  strings.Builder
		keys []string
	)
	for k := range pm {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		if v := pm.Get(k); v != "" {
			buf.WriteString(url.QueryEscape(k))
			buf.WriteByte('=')
			buf.WriteString(url.QueryEscape(v))
			buf.WriteByte('&')
		}
	}
	if buf.Len() <= 0 {
		return ""
	}
	return buf.String()[:buf.Len()-1]
}

func (pm ParamMap) CheckEmptyError(keys ...string) error {
	var emptyKeys []string
	for _, k := range keys {
		if v := pm.Get(k); v == "" {
			emptyKeys = append(emptyKeys, k)
		}
	}
	if len(emptyKeys) > 0 {
		return errors.New(strings.Join(emptyKeys, ", ") + " : cannot be empty")
	}
	return nil
}

// UnEncodeSignParams 参数字符串转Map
func (pm ParamMap) UnEncodeSignParams(params string) ParamMap {
	paramList := strings.Split(params, "&")
	for _, item := range paramList {
		index := strings.Index(item, "=")
		if index == -1 {
			continue
		}
		k := item[0:index]
		v := item[index+1:]
		pm.Set(k, v)
	}
	return pm
}

func convertToString(v interface{}) (str string) {
	if v == nil {
		return ""
	}
	var (
		bs  []byte
		err error
	)
	if bs, err = json.Marshal(v); err != nil {
		return ""
	}
	str = string(bs)
	return
}

func JonsObject(obj interface{}) string {
	if obj == nil {
		return ""
	}
	if res, err := json.Marshal(obj); err == nil {
		return string(res)
	}
	return ""
}
