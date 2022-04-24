package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	main "github.com/kgysf/go-mirai-http"
	"strings"
)

type HttpMethod string

const (
	HttpMethodGet  HttpMethod = "GET"
	HttpMethodPost HttpMethod = "POST"
)

// HttpBody Post 请求主体
type HttpBody interface{}

// HttpParams Get 请求参数
type HttpParams map[string]string

// HttpHeader 请求Header
type HttpHeader map[string]string

// HttpRequest 发送 Http 请求
// 第一个参数为请求地址
func HttpRequest[T any](url string, args ...any) (result T, err error) {
	var params HttpParams
	var body HttpBody
	var header HttpHeader
	var method = HttpMethodGet

	// 判断动态参数类型
	for _, el := range args {
		if p, ok := el.(HttpParams); ok {
			params = p
		} else if h, ok := el.(HttpHeader); ok {
			header = h
		} else if m, ok := el.(HttpMethod); ok {
			method = m
		} else if b, ok := el.(HttpBody); ok {
			body = b
		}
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		_url := fmt.Sprintf("http://%s:%d", main.Host, main.Port)
		if !strings.HasPrefix(url, "/") {
			url = _url + url
		} else {
			url = _url + "/" + url
		}
	}

	req := resty.New().R()
	req.SetHeader("Content-Type", "application/json")
	if params != nil {
		req.SetQueryParams(params)
	}
	if header != nil {
		req.SetHeaders(header)
	}
	if method == HttpMethodPost && body != nil {
		req.SetBody(body)
	}

	var rep *resty.Response
	var _err error
	switch method {
	case HttpMethodGet:
		rep, _err = req.Get(url)
	case HttpMethodPost:
		rep, _err = req.Post(url)
	}

	if _err != nil {
		err = _err
	} else if rep.StatusCode() != 200 {
		err = fmt.Errorf("error! Http code: %d, message: %s", rep.StatusCode(), rep.Status())
	} else {
		var ret T // 自动序列化
		_err := json.Unmarshal(rep.Body(), &ret)
		if _err != nil {
			err = _err
			return
		}
		result = ret
	}
	return
}

// HttpGet 发送 Http Get 请求
func HttpGet[T any](url string, args ...any) (result T, err error) {
	for i, el := range args {
		if _, ok := el.(HttpMethod); ok {
			args[i] = HttpMethodGet
		}
	}
	return HttpRequest[T](url, args...)
}

// HttpPost 发送 Http Post 请求
func HttpPost[T any](url string, body HttpBody, args ...any) (result T, err error) {
	flag := false
	for i, el := range args {
		if _, ok := el.(HttpMethod); ok {
			args[i] = HttpMethodPost
			flag = true
		}
	}
	if !flag {
		args = append(args, HttpMethodPost)
	}
	args = append(args, body)
	return HttpRequest[T](url, args...)
}
