package respons

type respon struct {
	Meta meta
	Data interface{}
}

type meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponsApi(message string, code int, status string, data interface{}) respon {
	meta := meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	respon := respon{
		Meta: meta,
		Data: data,
	}

	return respon
}
