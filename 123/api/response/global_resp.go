package response

type GlobalHttpResponse struct {
	Success bool     `json:"success"`
	Code    RespCode `json:"code"`
	Error   any      `json:"error"`
	Data    any      `json:"data"`
}

func GenResp(success bool, respCode RespCode, err any, data any) *GlobalHttpResponse {
	return &GlobalHttpResponse{
		Success: success,
		Code:    respCode,
		Error:   err,
		Data:    data,
	}
}

func SuccessResp() *GlobalHttpResponse {
	return &GlobalHttpResponse{
		Success: true,
		Code:    Success,
		Error:   nil,
		Data:    nil,
	}
}

func SuccessRespWithData(data any) *GlobalHttpResponse {
	return &GlobalHttpResponse{
		Success: true,
		Code:    Success,
		Error:   nil,
		Data:    data,
	}
}

func ErrorResp(err any) *GlobalHttpResponse {
	return &GlobalHttpResponse{
		Success: false,
		Code:    InternalError,
		Error:   err,
		Data:    nil,
	}
}
