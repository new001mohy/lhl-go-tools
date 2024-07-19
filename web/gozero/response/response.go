package response

import (
	"errors"
	lhlweberr "github.com/new001mohy/lhl-go-tools/web/error"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

type Body struct {
	Data    any    `json:"data"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func Response(w http.ResponseWriter, response any, err error) {
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
