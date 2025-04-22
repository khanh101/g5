package g5

import (
	"encoding/json"
)

type Handler = func(ctx *Context)
type Context struct {
	Input  []byte
	Code   int
	Output []byte
}

func (ctx *Context) BindJSON(o any) error {
	return json.Unmarshal(ctx.Input, o)
}

func (ctx *Context) JSON(code int, o any) {
	ctx.Code = code
	if o != nil {
		ctx.Output, _ = json.Marshal(o)
	}
}
