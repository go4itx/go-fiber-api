package client

import (
	"fmt"
	"home/pkg/code"
	"home/pkg/e"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

var (
	client = fiber.AcquireClient()
)

// HttpGet send request
func HttpGet(url string, query ...string) (data []byte, err error) {
	agent := client.Get(url)
	if len(query) > 0 {
		for _, v := range query {
			agent.QueryString(v) //eg: v = "foo=bar"
		}
	}

	if err = agent.Parse(); err != nil {
		return
	}

	statusCode, data, _ := agent.Bytes()
	if statusCode != http.StatusOK {
		err = e.NewError(code.ParamsIsInvalid, fmt.Sprintf("request %s error:%d", url, statusCode))
		return
	}

	return
}
