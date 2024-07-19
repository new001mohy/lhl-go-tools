package response

import (
	"errors"
	"github.com/zeromicro/go-zero/rest/httpx"
	lhlweberr "lhl-go-tools/web/error"
	"net/http"
)

type Body struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Response(w http.ResponseWriter, response interface{}, err error) {
	var body Body
	if err != nil {
		if errors.Is(err, lhlweberr.ErrSystemError{}) {
			body.Code = 0
			body.Message = "服务器忙，请稍后重试"
		} else {
			body.Code = 0
			body.Message = err.Error()
		}
	} else {
		body.Code = 1
		body.Data = response
	}
	httpx.OkJson(w, body)
}
