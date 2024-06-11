package utils

type responseSuccess struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

type responseFailed struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func ResponseSuccess(code int, data interface{}) responseSuccess {
	return responseSuccess{
		Code: code,
		Data: data,
	}
}

func ResponseFailed(code int, message interface{}) responseFailed {
	return responseFailed{
		Code:    code,
		Message: message,
	}
}
