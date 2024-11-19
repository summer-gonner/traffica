package errorx

const DefaultCode = "111111"

type CodeError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type CodeErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func NewCodeError(code, msg string) error {
	return &CodeError{Code: code, Message: msg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(DefaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.Message
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		Code:    e.Code,
		Message: e.Message,
	}
}
