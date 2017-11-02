package http

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/keegancsmith/shell"
)

type Request struct {
	Type    string `json:"type"`
	Cmd     string `json:"cmd"`
	Timeout int    `json:"timeout"`
}

type Response struct {
	Errno  int         `json:errono`
	Errmsg string      `json:message`
	Body   interface{} `json:body`
}

func MakeResponse(errno int, errmsg string, body interface{}) Response {
	return Response{
		Errno:  errno,
		Errmsg: errmsg,
		Body:   body,
	}
}

func (r *Request) Execute() Response {
	switch r.Type {
	case "shell":
		// set default timeout as 10s
		if r.Timeout == 0 {
			r.Timeout = 10
		}
		cmd := shell.Commandf(r.Cmd)
		cmd.Start()

		doneCh := make(chan error, 1)
		go func() {
			doneCh <- cmd.Wait()
		}()

		select {
		case <-time.After(time.Duration(r.Timeout) * time.Second):
			if err := cmd.Process.Kill(); err != nil {
				return MakeResponse(1, err.Error(), nil)
			}
			return MakeResponse(1, "process timeout", nil)
		case err := <-doneCh:
			if err != nil {
				return MakeResponse(1, err.Error(), nil)
			} else {
				out, err := cmd.CombinedOutput()
				if err != nil {
					return MakeResponse(1, err.Error(), nil)
				} else {
					return MakeResponse(0, "OK", string(out))
				}
			}
		}
	default:
		return MakeResponse(1, "Unknown command", nil)
	}
}

func RunHandler(c *gin.Context) {
	var req Request

	c.Bind(&req)
	resp := req.Execute()
	c.JSON(200, resp)
}
