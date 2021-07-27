package controller


type ResponseBaseModel struct{
	Success bool `json:"success"`
	Code int	`json:"code"`
	Message string	`json:"message"`
	Version string 	`json:"versiont"`
}

type ResponseDataModel struct{
	Size	int64	`json:"size"`
	Mime	string	`json:"mime"`
	FileID	string 	`json:"fileId"`
}

type ResponseModel struct {
	ResponseBaseModel
	Data ResponseDataModel	`json:"data"`
}

