package http

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	main "github.com/kgysf/go-mirai-http"
	. "github.com/kgysf/go-mirai-http/model"
	"strconv"
)

type UploadType string

const (
	UploadTypeFriend UploadType = "friend"
	UploadTypeGroup  UploadType = "group"
	UploadTypeTemp   UploadType = "temp"
)

// UploadImage 图片文件上传
//  使用此方法上传图片文件至服务器并返回ImageId
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#图片文件上传
func UploadImage(sessionKey string, typ UploadType, img string) (result ImageInfo, e error) {
	url := fmt.Sprintf("http://%s:%d/%s", main.Host, main.Port, "uploadImage")
	body := map[string]string{
		"sessionKey": sessionKey,
		"type":       string(typ),
	}

	req := resty.New().R()
	req.SetFile("img", img)
	req.SetFormData(body)
	rep, err := req.Post(url)

	if err != nil {
		e = err
	} else {
		var ret ImageInfo
		_err := json.Unmarshal(rep.Body(), &ret)
		if _err != nil {
			e = _err
		} else {
			result = ret
		}
	}

	return
}

// UploadVoice 语音文件上传
//  使用此方法上传语音文件至服务器并返回VoiceId
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#语音文件上传
func UploadVoice(sessionKey string, voice string) (result VoiceInfo, e error) {
	url := fmt.Sprintf("http://%s:%d/%s", main.Host, main.Port, "uploadVoice")
	body := map[string]string{
		"sessionKey": sessionKey,
		"type":       string(UploadTypeGroup),
	}

	req := resty.New().R()
	req.SetFile("voice", voice)
	req.SetFormData(body)
	rep, err := req.Post(url)

	if err != nil {
		e = err
	} else {
		var ret VoiceInfo
		_err := json.Unmarshal(rep.Body(), &ret)
		if _err != nil {
			e = _err
		} else {
			result = ret
		}
	}

	return
}

// UploadFile 群文件上传
//  使用此方法上传文件至群文件
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#群文件上传
func UploadFile(sessionKey string, file string, group int, path string) (result FileInfoResult, e error) {
	url := fmt.Sprintf("http://%s:%d/%s", main.Host, main.Port, "file/upload")
	body := map[string]string{
		"sessionKey": sessionKey,
		"path":       path,
		"type":       string(UploadTypeGroup),
		"target":     strconv.Itoa(group),
	}

	req := resty.New().R()
	req.SetFile("file", file)
	req.SetFormData(body)
	rep, err := req.Post(url)

	if err != nil {
		result = FileInfoResult{Code: -10000, Msg: err.Error(), Err: err}
	} else {
		var ret FileInfoResult
		_err := json.Unmarshal(rep.Body(), &ret)
		if _err != nil {
			result = FileInfoResult{Code: -10000, Msg: _err.Error(), Err: _err}
		} else {
			result = ret
		}
	}

	return
}
