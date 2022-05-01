package http

import (
	"github.com/toorux/go-mirai-http/internal/pkg"
	"strconv"

	//"fmt"
	. "github.com/toorux/go-mirai-http/model"
)

type FileListResult HttpResult[[]File]
type FileInfoResult HttpResult[File]

// FileList 查看文件列表
//  使用此方法查看文件列表, 建议使用 FileListForGroup
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#查看文件列表
// ⚠️警告: 目前官方仅支持群文件操作，所有好友相关字段均为保留字段
func FileList(sessionKey string, file FileParams, offset int, size int) (result FileListResult) {
	const url = "/file/list"
	params := pkg.HttpParams{
		"sessionKey":       sessionKey,
		"id":               file.Id,
		"path":             file.Path,
		"target":           strconv.Itoa(file.Target),
		"group":            strconv.Itoa(file.Group),
		"qq":               strconv.Itoa(file.Qq),
		"withDownloadInfo": strconv.FormatBool(file.WithDownloadInfo),
		"offset":           strconv.Itoa(offset),
		"size":             strconv.Itoa(size),
	}
	ret, _err := pkg.HttpGet[FileListResult](url, params)
	if _err != nil {
		result = FileListResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// FileListForGroup 查看群文件列表
//  使用此方法查看群文件列表，此函数为 FileList 函数的封装
//  - id: 文件id,空串为根目录,id和path都为空时读取根目录
//  - path: 文件夹路径, 文件夹允许重名, 不保证准确, 准确定位使用 id
//  - group: 要查看的群群号
//  - withDownloadInfo: 是否附带文件下载信息, 此接口不建议附带
//  - offset: 分页偏移
//  - size: 分页大小
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#查看文件列表
func FileListForGroup(sessionKey string, id string, path string, group int, withDownloadInfo bool, offset int, size int) (result FileListResult) {
	fileParam := FileParams{
		Id:               id,
		Path:             path,
		Target:           group,
		WithDownloadInfo: withDownloadInfo,
	}
	return FileList(sessionKey, fileParam, offset, size)
}

// FileInfo 获取文件信息
//  使用此方法获取文件信息, 建议使用 FileInfoForId
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#获取文件信息
// ⚠️警告: 目前官方仅支持群文件操作，所有好友相关字段均为保留字段
func FileInfo(sessionKey string, file FileParams) (result FileInfoResult) {
	const url = "/file/info"
	params := pkg.HttpParams{
		"sessionKey":       sessionKey,
		"id":               file.Id,
		"path":             file.Path,
		"target":           strconv.Itoa(file.Target),
		"group":            strconv.Itoa(file.Group),
		"qq":               strconv.Itoa(file.Qq),
		"withDownloadInfo": strconv.FormatBool(file.WithDownloadInfo),
	}
	ret, _err := pkg.HttpGet[FileInfoResult](url, params)
	if _err != nil {
		result = FileInfoResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// FileInfoForId 获取文件信息
//  使用此方法获取文件信息，此函数为 FileInfo 函数的封装
//  - id: 文件id
//  - group: 要查看的群群号
//  - withDownloadInfo: 是否附带文件下载信息, 此接口不建议附带
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#获取文件信息
func FileInfoForId(sessionKey string, id string, group int, withDownloadInfo bool) (result FileInfoResult) {
	fileParam := FileParams{
		Id:               id,
		Target:           group,
		WithDownloadInfo: withDownloadInfo,
	}
	return FileInfo(sessionKey, fileParam)
}

// FileInfoForPath 获取文件信息
//  使用此方法获取文件信息，此函数为 FileInfo 函数的封装
//  - id: 文件id
//  - group: 要查看的群群号
//  - withDownloadInfo: 是否附带文件下载信息, 此接口不建议附带
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#获取文件信息
// ⚠️警告: 由于path的不准确性，建议使用 FileInfoForId 查询
func FileInfoForPath(sessionKey string, path string, group int, withDownloadInfo bool) (result FileInfoResult) {
	fileParam := FileParams{
		Path:             path,
		Target:           group,
		WithDownloadInfo: withDownloadInfo,
	}
	return FileInfo(sessionKey, fileParam)
}

// FileMkdir 创建文件夹
//  使用此方法创建文件夹, 建议使用 FileMkdirInRoot
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#创建文件夹
//
// ⚠️警告: 目前官方仅支持群文件操作，所有好友相关字段均为保留字段
//
// ⚠️警告: 此接口似乎也存在bug， 不建议使用
func FileMkdir(sessionKey string, file FileParams, directoryName string) (result FileInfoResult) {
	const url = "/file/mkdir"
	body := map[string]any{
		"sessionKey":    sessionKey,
		"id":            file.Id,
		"target":        file.Target,
		"group":         file.Group,
		"qq":            file.Qq,
		"directoryName": directoryName,
	}
	ret, _err := pkg.HttpPost[FileInfoResult](url, body)
	if _err != nil {
		result = FileInfoResult{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// FileMkdirInRoot 在群文件根目录创建文件夹
//  使用此方法在群文件根目录创建文件夹，此函数为 FileMkdir 函数的封装
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#获取文件信息
// ⚠️警告: 此接口似乎也存在bug， 不建议使用
func FileMkdirInRoot(sessionKey string, group int, directoryName string) (result FileInfoResult) {
	fileParam := FileParams{
		Id:     "",
		Path:   "",
		Target: group,
		Group:  group,
	}
	return FileMkdir(sessionKey, fileParam, directoryName)
}

// FileDelete 删除文件
//  使用此方法删除文件, 建议使用 FileDeleteForId
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#删除文件
// ⚠️警告: 目前官方仅支持群文件操作，所有好友相关字段均为保留字段
func FileDelete(sessionKey string, file FileParams) (result HttpResult[any]) {
	const url = "/file/delete"
	body := map[string]any{
		"sessionKey": sessionKey,
		"id":         file.Id,
		"target":     file.Target,
		"group":      file.Group,
		"qq":         file.Qq,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// FileDeleteForId 根据文件id删除文件
//  使用此方法根据文件id删除文件, 此函数为 FileDelete 函数封装
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#删除文件
func FileDeleteForId(sessionKey string, id string, group int) (result HttpResult[any]) {
	fileParam := FileParams{
		Id:     id,
		Target: group,
	}
	return FileDelete(sessionKey, fileParam)
}

// FileMove 移动文件
//  使用此方法移动文件
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#移动文件
// ⚠️警告: 目前官方仅支持群文件操作，所有好友相关字段均为保留字段
func FileMove(sessionKey string, file FileParams, moveTo string, moveToPath string) (result HttpResult[any]) {
	const url = "/file/move"
	body := map[string]any{
		"sessionKey": sessionKey,
		"id":         file.Id,
		"path":       file.Path,
		"target":     file.Target,
		"group":      file.Group,
		"qq":         file.Qq,
		"moveTo":     moveTo,
		"moveToPath": moveToPath,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// FileRename 重命名文件
//  使用此方法重命名文件
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#重命名文件
//
// ⚠️警告: 目前官方仅支持群文件操作，所有好友相关字段均为保留字段
//
// ⚠️警告: 此接口权限貌似存在一些问题
func FileRename(sessionKey string, file FileParams, renameTo string) (result HttpResult[any]) {
	const url = "/file/rename"
	body := map[string]any{
		"sessionKey": sessionKey,
		"id":         file.Id,
		"path":       file.Path,
		"target":     file.Target,
		"group":      file.Group,
		"qq":         file.Qq,
		"renameTo":   renameTo,
	}
	ret, _err := pkg.HttpPost[HttpResult[any]](url, body)
	if _err != nil {
		result = HttpResult[any]{Code: -10000, Msg: _err.Error(), Err: _err}
	} else {
		result = ret
	}
	return
}

// FileRenameForId 根据id重命名文件
//  使用此方法重命名文件, 此函数为 FileRename 函数的封装
//  用法同官方文档，返回值增加msg字段，错误代码-10000为自定义错误， 错误信息见msg字段
//  NOTE: https://docs.mirai.mamoe.net/mirai-api-http/adapter/HttpAdapter.html#重命名文件
// ⚠️警告: 此接口权限貌似存在一些问题
func FileRenameForId(sessionKey string, id string, renameTo string, group int) (result HttpResult[any]) {
	fileParam := FileParams{
		Id:     id,
		Target: group,
	}
	return FileRename(sessionKey, fileParam, renameTo)
}
