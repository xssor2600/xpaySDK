package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func DoHttpRequestJson(ctx context.Context,method string,requestUrl string,bodyMap map[string]interface{},header map[string]string)([]byte,error){
	client := http.Client{}
	by,_ := json.Marshal(bodyMap)
	req,err := http.NewRequest(method,requestUrl,bytes.NewBuffer(by))
	if err != nil {
		return nil,err
	}
	req.Header.Set("content-type","application/json")
	// set header
	if len(header) > 0 {
		for k := range header {
			req.Header.Set(k,header[k])
		}
	}
	resp,err := client.Do(req)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	rb,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil,err
	}
	return rb,nil






}


