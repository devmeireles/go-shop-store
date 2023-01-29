package utils

type ResponseMsg struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func ResError(err error) ResponseMsg {
	res := ResponseMsg{
		Success: false,
		Message: err.Error(),
	}

	return res
}

func ResSuccess(data interface{}) ResponseMsg {
	res := ResponseMsg{
		Success: true,
		Data:    data,
	}

	return res
}
