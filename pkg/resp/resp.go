package resp

import (
	"home/pkg/code"
	"home/pkg/e"
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
func (response Response) JSON(params ...interface{}) (err error) {
	var (
		data interface{}
		ok   bool
		ee   *fiber.Error

		length = len(params)
		max    = length - 1
	)

	if length > 0 {
		// customize Error
		if ee, ok = params[max].(*fiber.Error); !ok {
			// golang error
			if err, ok = params[max].(error); ok && err != nil {
				ee = e.NewError(fiber.StatusInternalServerError, err.Error())
			}
		}
	}

	if ee != nil {
		data = ""
	} else {
		ee = e.NewError(code.OK)
		switch length {
		case 0:
			data = ""
		case 1:
			data = response.format(params[0])
		default:
			data = response.format(params[0:max]...)
		}
	}

	return response.ctx.Status(fiber.StatusOK).JSON(Ret{
		Code:       ee.Code,
		Msg:        ee.Message,
		ServerTime: time.Now().Unix(),
		Data:       data,
	})
}

// format return ret data
func (response Response) format(params ...interface{}) (data interface{}) {
	length := len(params)
	switch length {
	case 1:
		if params[0] == nil {
			data = ""
		} else {
			data = params[0]
		}
	case 2, 3:
		tmp := PaginationRet{
			Items: params[0],
			Count: params[1],
		}

		if length == 3 {
			tmp.Ext = params[2]
		}

		data = tmp
	default:
		data = ""
	}

	return
}
