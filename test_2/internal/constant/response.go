package constant

const (
	CodeSuccess     = 200
	CodeBadRequest  = 400
	CodeServerError = 500
)

var messageText = map[int]string{
	CodeSuccess:     `{"Response":"True"}`,
	CodeBadRequest:  `{"Response":"False","Error":"Bad Request"}`,
	CodeServerError: `{"Response":"False","Error":"Internal Servcer Error"}`,
}

func ErrorMessage(code int) string {
	return messageText[code]
}

func GetError(code int) (string, int) {
	return ErrorMessage(code), code
}
