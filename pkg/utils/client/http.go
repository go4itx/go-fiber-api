package client

import (
	"fmt"
	"home/pkg/code"
	"home/pkg/e"
	"log"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

var (
	timeout   = 30 * time.Second
	userAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/110.0.0.0 Safari/537.36"
)

type request struct {
	err       error
	agent     *fiber.Agent
	timeout   time.Duration
	userAgent string
}

// Request send a http request
func Request(url string, method ...string) (req request) {
	methodName := fiber.MethodGet
	if len(method) > 0 && method[0] != "" {
		methodName = strings.ToUpper(method[0])
	}

	client := fiber.AcquireClient()
	switch methodName {
	case fiber.MethodHead:
		req.agent = client.Head(url)
	case fiber.MethodGet:
		req.agent = client.Get(url)
	case fiber.MethodPost:
		req.agent = client.Post(url)
	case fiber.MethodPut:
		req.agent = client.Put(url)
	case fiber.MethodPatch:
		req.agent = client.Patch(url)
	case fiber.MethodDelete:
		req.agent = client.Delete(url)
	default:
		req.err = fmt.Errorf("the method is not supported: %s", methodName)
	}

	return
}

func (req request) Debug() request {
	req.agent.Debug()
	return req
}

// QueryString sets the URI query string.
func (req request) QueryString(m map[string]interface{}) request {
	str := ""
	for k, v := range m {
		str += fmt.Sprintf("%s=%v&", k, v)
	}

	req.agent.QueryString(strings.TrimRight(str, "&"))
	return req
}

// JSON sends a JSON request by setting the Content-Type header to application/json.
func (req request) JSON(v interface{}) request {
	req.agent.JSON(v)
	return req
}

// XML sends an XML request by setting the Content-Type header to application/xml.
func (req request) XML(v interface{}) request {
	req.agent.XML(v)
	return req
}

// Form sends a form request by setting the Content-Type header to application/x-www-form-urlencoded.
func (req request) Form(param fiber.Map) request {
	args := fiber.AcquireArgs()
	for k, v := range param {
		args.Set(k, fmt.Sprintf("%v", v))
	}

	req.agent.Form(args)
	fiber.ReleaseArgs(args)
	return req
}

// Timeout sets request timeout duration.
func (req request) Timeout(duration time.Duration) request {
	req.timeout = duration
	return req
}

// UserAgent sets User-Agent header value.
func (req request) UserAgent(userAgent string) request {
	req.userAgent = userAgent
	return req
}

// Result returns
func (req request) Result(v ...interface{}) (bytes []byte, err error) {
	defer req.agent.ConnectionClose()
	if req.err != nil {
		err = e.NewError(code.ParamsIsInvalid, req.err.Error())
		return
	}

	if req.timeout != 0 {
		timeout = req.timeout
	}

	if req.userAgent != "" {
		userAgent = req.userAgent
	}

	req.agent.Timeout(timeout)
	req.agent.UserAgent(userAgent)
	if err = req.agent.Parse(); err != nil {
		err = e.NewError(code.ParamsIsInvalid, err.Error())
		return
	}

	var (
		errs       []error
		statusCode int
	)

	if len(v) > 0 {
		statusCode, bytes, errs = req.agent.Struct(v[0])
	} else {
		statusCode, bytes, errs = req.agent.Bytes()
	}

	if len(errs) > 0 {
		log.Println("errs", errs)
		err = e.NewError(code.ParamsIsInvalid, errs[0].Error())
		return
	}

	if statusCode != fiber.StatusOK {
		err = fiber.NewError(statusCode)
		return
	}

	return
}
