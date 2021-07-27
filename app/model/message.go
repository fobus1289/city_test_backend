package model

type ResponseError struct {
	Message string `json:"message"`
}

type Response struct {
	Name  string      `json:"-"`
	Data  interface{} `json:"data"`
	Count int64       `json:"count"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

func NewResponse(name string, page, limit int, data interface{}) *Response {
	return &Response{
		Name:  name,
		Page:  page,
		Limit: limit,
		Data:  data,
	}
}
