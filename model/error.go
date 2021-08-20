package model

// 状态码
const (
	StatusJson           = 1
	StatusForm           = 2
	StatusImgDecode      = 3
	StatusImgIsType      = 4
	StatusFileSeek       = 5
	StatusFileMd5        = 6
	StatusMkdir          = 7
	StatusOpenFile       = 8
	StatusImgEncode      = 9
	StatusImgNotFound    = 10
	StatusUrlNotFound    = 11
	StatusImgCompression = 12
)

var statusText = map[int]string{
	StatusJson:           "json打包失败",
	StatusForm:           "表单字段 userfile 缺少",
	StatusImgDecode:      "图片解码不符合",
	StatusImgIsType:      "图片类型不符合",
	StatusFileSeek:       "设置文件读写位置失败",
	StatusFileMd5:        "计算文件MD5失败",
	StatusMkdir:          "创建目录失败",
	StatusOpenFile:       "文件创建失败",
	StatusImgEncode:      "图片生成失败",
	StatusImgNotFound:    "没有找到图片",
	StatusUrlNotFound:    "Url Not Found",
	StatusImgCompression: "图片压缩失败",
	200:                  "OK",
	404:                  "404 Not Found!",
	500:                  "服务器错误",
}

// StatusText return status text.
func StatusText(code int) string {
	return statusText[code]
}
