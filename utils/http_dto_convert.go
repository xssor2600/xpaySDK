package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/gorilla/schema"
	"io/ioutil"
	"net/http"
)

var (
	URIDecoder = schema.NewDecoder()
)

func ReadRequestParser(ctx context.Context,raw *http.Request,reqDto interface{}) error{
	bodyBytes := make([]byte,0)
	if raw.Body != nil {
		bodyBytes,_ := ioutil.ReadAll(raw.Body)
		defer func() {
			raw.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}()
	}
	method := raw.Method
	switch method {
	case "POST":
		err := URIDecoder.Decode(reqDto,raw.URL.Query())
		if len(bodyBytes) == 0 {
			return err
		}
		derr := JsonUnMashObject(bodyBytes,reqDto)
		if derr != nil {
			return derr
		}
	case "GET":
		derr := URIDecoder.Decode(reqDto,raw.URL.Query())
		if derr != nil {
			return derr
		}
	}
	return nil
}



func JsonUnMashObject(b []byte,obj interface{}) error {
	return json.Unmarshal(b,&obj)
}

func JsonMashObject(resp interface{}) string {
	if r,err := json.Marshal(resp);err == nil {
		return string(r)
	}
	return ""
}



