package util

import (
	"context"
	"encoding/json"

	"github.com/gogf/gf/v2/frame/g"
)

// HttpClient http请求
func HttpClient(ctx context.Context, apiUrl string, method string, reqParams interface{}, headers map[string]string, isDebug bool) (string, error) {
	var rest string

	token := ""
	authorization := ctx.Value("token")
	if authorization != nil {
		token = authorization.(string)
	}

	gClient := g.Client()
	if len(token) > 0 {
		gClient.SetHeader("Authorization", "Bearer "+token)
	}

	for k, v := range headers {
		gClient.SetHeader(k, v)
	}

	if isDebug {
		g.Log().Debug(ctx, "netClient apiUrl:", apiUrl)
		g.Log().Debug(ctx, "netClient method:", method)
		g.Log().Debug(ctx, "netClient params:", reqParams)
		//g.Log().Debug(ctx, "netClient Authorization:", token)
		//g.Log().Debug(ctx, "netClient header:", headers)
		//g.Log().Debug(ctx, "netClient ctx:", ctx)
	}

	if method == "GET" || method == "get" {
		r, err := gClient.Get(ctx, apiUrl)
		if err != nil {
			return rest, err
		}
		defer r.Close()
		return r.ReadAllString(), nil
	}

	if method == "POST" || method == "post" {
		bytes, _ := json.Marshal(reqParams)
		r, err := gClient.Post(ctx, apiUrl, string(bytes))
		if err != nil {
			return rest, err
		}
		defer r.Close()

		return r.ReadAllString(), nil
	}

	if method == "POSTFORM" || method == "postform" {
		r, err := gClient.SetContentType("application/x-www-form-urlencoded").
			Post(ctx, apiUrl, reqParams)

		if err != nil {
			return rest, err
		}
		defer r.Close()

		return r.ReadAllString(), nil
	}

	if method == "DELETE" || method == "delete" || method == "del" || method == "DEL" {
		r, err := gClient.Delete(ctx, apiUrl)
		if err != nil {
			return rest, err
		}
		defer r.Close()
		return r.ReadAllString(), nil
	}

	if method == "PUT" || method == "put" {
		bytes, _ := json.Marshal(reqParams)

		r, err := gClient.Put(ctx, apiUrl, bytes)
		if err != nil {
			return rest, err
		}
		defer r.Close()
		return r.ReadAllString(), nil
	}

	return rest, nil
}
