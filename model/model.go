package model

import "encoding/json"

const version = "v0.1.0"

type ResponseBaseModel struct {
	Success bool   `json:"success"`
	Code    int    `json:"code"`
	Message string `json:"message"`
	Version string `json:"version"`
}

type FileInfoModel struct {
	Size   uint   `json:"size"`
	Mime   string `json:"mime"`
	FileID string `json:"fileId"`
}

type ResponseModel struct {
	ResponseBaseModel
	Data []*FileInfoModel `json:"data"`
}

func (res *ResponseModel) ResponseJson(code int, status bool, infos []*FileInfoModel) []byte {
	res.Version = version
	res.Success = status
	res.Code = code
	res.Message = StatusText(code)
	res.Data = infos

	data, err := json.Marshal(res)
	if err != nil {
		return nil
	}

	return data
}
