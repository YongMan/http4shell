package http

import (
	"bytes"
	"net/http"
	"strings"
	"time"

	"github.com/YongMan/http4shell/config"
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

type HandlerRun struct {
	Config *config.Config
}

func MakeResponse(errno int, errmsg string, body interface{}) Response {
	return Response{
		Errno:  errno,
		Errmsg: errmsg,
		Body:   body,
	}
}

func NewHandlerRun(c *config.Config) *HandlerRun {
	return &HandlerRun{Config: c}
}

func (r *Request) Execute() Response {
	switch r.Type {
	case "shell":
		// set default timeout as 10s
		if r.Timeout == 0 {
			r.Timeout = 10
		}
		cmd := shell.Commandf(r.Cmd)

		var stdout, stderr bytes.Buffer
		cmd.Stdout = &stdout
		cmd.Stderr = &stderr
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
				return MakeResponse(1, err.Error(), string(stderr.Bytes()))
			} else {
				return MakeResponse(0, "OK", string(stdout.Bytes()))
			}
		}
	default:
		return MakeResponse(1, "Unknown command", nil)
	}
}

func (h *HandlerRun) RunHandler(c *gin.Context) {
	var req Request

	c.Bind(&req)

	// check if cmd in cmds whitelist
	valid := false

	cmd := strings.TrimSpace(req.Cmd)

	for _, c := range h.Config.Cmds {
		if c.Cmd.Arg == cmd {
			valid = true
			break
		}
	}
	if !valid {
		c.JSON(http.StatusUnauthorized, MakeResponse(1, "Unauthorized command, please contact admin", nil))
		return
	}

	resp := req.Execute()
	c.JSON(200, resp)
}
