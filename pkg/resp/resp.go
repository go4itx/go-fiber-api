package resp

import (
	"github.com/gofiber/fiber/v2"
	"home/pkg/code"
	"time"
)

// result Uniform results
type result struct {
	Code       int         `json:"status"`
	Msg        string      `json:"msg"`
	ServerTime int64       `json:"serverTime"`
	Data       interface{} `json:"data"`
}

//  ErrorHandler unified processing error
func ErrorHandler(ctx *fiber.Ctx, err error) error {
	ret := result{
		Code:       0,
		Msg:        "",
		ServerTime: time.Now().Unix(),
		Data:       "",
	}

	if e, ok := err.(*fiber.Error); ok {
		ret.Code = e.Code
		ret.Msg = e.Message
	} else {
		ret.Code = fiber.StatusInternalServerError
		ret.Msg = err.Error()
	}

	return ctx.Status(fiber.StatusOK).JSON(ret)
}

type Response struct {
	ctx *fiber.Ctx
}

// New return Handle
func New(ctx *fiber.Ctx) Response {
	return Response{ctx: ctx}
}

// JSON handle Result
func (response Response) JSON(data ...interface{}) error {
	var (
		err error
		ok  bool
		l   = len(data)
		max = l - 1
	)

	// fiber Error
	if err, ok = data[max].(*fiber.Error); ok && err != nil {
		return err
	}

	// golang error
	if err, ok = data[max].(error); ok && err != nil {
		return err
	}

	var ret result
	if l == 1 {
		if data[max] == nil {
			ret = success()
			goto END
		}

		ret = success(data[0])
		goto END
	}

	ret = success(data[0:max]...)
	goto END

END:
	return response.ctx.Status(fiber.StatusOK).JSON(ret)
}

// success return ret
func success(data ...interface{}) (ret result) {
	length := len(data)
	ret = result{
		Code:       code.Ok,
		Msg:        code.Value(code.Ok),
		ServerTime: time.Now().Unix(),
		Data:       "",
	}

	switch length {
	case 1:
		if length == 1 {
			ret.Data = data[0]
		}
	case 2, 3:
		tmp := map[string]interface{}{
			"list":  data[0],
			"count": data[1],
			"ext":   "",
		}

		if length == 3 {
			tmp["ext"] = data[2]
		}

		ret.Data = tmp
	}

	return
}
