package manager

type Response struct {
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

func setResponse(reqId, message string) *Response {
	return &Response{
		RequestId: reqId,
		Message:   message,
	}
}