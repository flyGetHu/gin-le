package entity

import "net/http"

type Result struct {
	Code uint        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func (r Result) Ok(data interface{}) Result {
	r.Code = http.StatusOK
	r.Data = data
	return r
}

func (r Result) Error(msg string) Result {
	r.Code = http.StatusInternalServerError
	r.Msg = msg
	return r
}

func (r Result) ErrorData(msg string, data interface{}) Result {
	r.Code = http.StatusInternalServerError
	r.Msg = msg
	r.Data = data
	return r
}

func (r Result) ErrorCode(code uint, msg string, data interface{}) Result {
	r.Code = code
	r.Msg = msg
	r.Data = data
	return r
}
