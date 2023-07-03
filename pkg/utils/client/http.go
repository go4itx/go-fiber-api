package client

import (
	"encoding/json"
	"fmt"
	"home/pkg/code"
	"home/pkg/e"
	"log"
	"net"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/valyala/fasthttp"
)

var (
	timeout   = 1 * time.Minute
	userAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/89.0.4389.82 Safari/537.36"
)

type request struct {
	url         string
	method      string
	queryString string
	retry       bool
	agent       *fiber.Agent
	timeout     time.Duration
	userAgent   string
}

// Request send a http request
func Request(url string, method ...string) request {
	httpMethod := fiber.MethodGet
	if len(method) > 0 && method[0] != "" {
		httpMethod = strings.ToUpper(method[0])
	}

	agent := fiber.AcquireAgent()
	return request{
		url:         url,
		method:      httpMethod,
		queryString: "",
		retry:       false,
		agent:       agent,
		timeout:     timeout,
		userAgent:   userAgent,
	}
}

// Debug mode enables logging request and response detail
func (req request) Debug() request {
	req.agent.Debug()
	return req
}

// Retry controls whether a retry should be attempted after an error.
func (req request) Retry(b bool) request {
	req.retry = b
	return req
}

// QueryString sets the URI query string.
func (req request) QueryString(m map[string]interface{}) request {
	str := ""
	for k, v := range m {
		str += fmt.Sprintf("%s=%v&", k, v)
	}

	if req.queryString != "" {
		req.queryString += "&" + str
	} else {
		req.queryString = str
	}

	req.queryString = strings.TrimRight(req.queryString, "&")
	return req
}

// Form sends a form request by setting the Content-Type header to application/x-www-form-urlencoded.
func (req request) Form(m map[string]interface{}) request {
	args := fiber.AcquireArgs()
	for k, v := range m {
		args.Set(k, fmt.Sprintf("%v", v))
	}

	req.agent.Form(args)
	fiber.ReleaseArgs(args)
	return req
}

// JSON sends a JSON request by setting the Content-Type header to application/json.
func (req request) JSON(v interface{}) request {
	if str, ok := v.(string); ok {
		if json.Valid([]byte(str)) {
			json.Unmarshal([]byte(str), &v)
		}
	}

	req.agent.JSON(v)
	return req
}

// XML sends an XML request by setting the Content-Type header to application/xml.
func (req request) XML(v interface{}) request {
	req.agent.XML(v)
	return req
}

// Set sets the given 'key: value' header.
func (req request) Set(k, v string) request {
	req.agent.Set(k, v)
	return req
}

// Cookie sets one 'key: value' cookie.
func (req request) Cookie(k, v string) request {
	req.agent.Cookie(k, v)
	return req
}

// Cookies sets multiple 'key: value' cookies.
func (req request) Cookies(kv ...string) request {
	req.agent.Cookies(kv...)
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

// SetResponse sets custom response for the Agent instance.
func (req request) SetResponse(customResp *fiber.Response) request {
	req.agent.SetResponse(customResp)
	return req
}

// Result returns
func (req request) Result(v ...interface{}) (bytes []byte, err error) {
	defer req.agent.ConnectionClose()
	req.agent.Timeout(req.timeout)
	req.agent.UserAgent(req.userAgent)

	ar := req.agent.Request()
	ar.SetRequestURI(req.url)
	ar.Header.SetMethod(req.method)
	if req.queryString != "" {
		req.agent.QueryString(req.queryString)
	}

	if err = req.agent.Parse(); err != nil {
		err = e.NewError(code.ParamsIsInvalid, err.Error())
		return
	}

	req.agent.Dial = func(addr string) (net.Conn, error) {
		return fasthttp.DialTimeout(addr, req.timeout)
	}

	req.agent.RetryIf(func(r *fiber.Request) bool {
		return req.retry
	})

	var (
		errs       []error
		statusCode int
	)

	if len(v) > 0 {
		statusCode, bytes, errs = req.agent.Struct(v[0])
	} else {
		statusCode, bytes, errs = req.agent.Bytes()
	}

	if statusCode != fiber.StatusOK {
		if statusCode == 0 {
			statusCode = 400
		}

		err = e.NewError(statusCode, req.url+" "+utils.StatusMessage(statusCode))
		log.Println("error: ", statusCode, err)
		return
	}

	if len(errs) > 0 {
		log.Println("errs", errs)
		err = e.NewError(code.ParamsIsInvalid, errs[0].Error())
		return
	}

	return
}
