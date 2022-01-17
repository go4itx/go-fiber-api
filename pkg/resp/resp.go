package resp

import (
	"home/pkg/code"
	"time"

	"github.com/gofiber/fiber/v2"
)

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
		ret    result
		ok     bool
		err    error
		fe     *fiber.Error
		length = len(data)
		max    = length - 1
	)

	if length == 0 {
		ret = success()
		goto END
	}

	// fiber Error
	if fe, ok = data[max].(*fiber.Error); ok && fe != nil {
		return fe
	}

	// golang error
	if err, ok = data[max].(error); ok && err != nil {
		return err
	}

	if length == 1 {
		if data[max] == nil || err == fe {
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
			"items": data[0],
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
