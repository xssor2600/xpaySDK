package utils

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"sort"
)

func KsSignFromMap(s map[string]interface{}, key string) string {
	sorted_keys := make([]string, 0)
	for k := range s {
		sorted_keys = append(sorted_keys, k)
	}
	sort.Strings(sorted_keys)
	var signStrings string
	for _, k := range sorted_keys {
		value := fmt.Sprintf("%v", s[k])
		if k != "" && value != "" && k != "sign" && k != "-" {
			signStrings = signStrings + k + "=" + value + "&"
		}
	}
	if key != "" && len(signStrings) > 0 {
		signStrings = signStrings[0:len(signStrings)-1] + key
	}
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(signStrings))
	cipherStr := md5Ctx.Sum(nil)
	upperSign := hex.EncodeToString(cipherStr)
	return upperSign
}