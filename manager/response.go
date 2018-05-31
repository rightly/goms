package manager

type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func ok(remote string) *Response {
	return &Response{
		Code:200,
		Message:remote + " receive success",
	}
}

func fail(remote string) *Response {
	return &Response{
		Code:400,
		Message:remote + " receive fail",
	}
}