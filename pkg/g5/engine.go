package g5

import (
	"encoding/json"
	"syscall/js"
)

type Engine struct {
	HandlerMap map[string]Handler
}

func Default() *Engine {
	return &Engine{
		HandlerMap: make(map[string]Handler),
	}
}

func (e *Engine) POST(path string, handler Handler) {
	e.HandlerMap["POST/"+path] = handler
}

func (e *Engine) GET(path string, handler Handler) {
	e.HandlerMap["GET/"+path] = handler
}

type Input struct {
	Path string `json:"path"`
	Data string `json:"data"`
}

type Output struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

func (e *Engine) Run() {
	js.Global().Set("g5", js.FuncOf(func(this js.Value, p []js.Value) any {
		if len(p) == 0 {
			return js.ValueOf("no input")
		}
		input := p[0].String()

		output := func(input string) string {
			var i Input
			err := json.Unmarshal([]byte(input), &i)
			if err != nil {
				panic(err)
			}
			ctx := &Context{
				Input:  []byte(i.Data),
				Code:   0,
				Output: nil,
			}
			e.HandlerMap[i.Path](ctx)
			var o Output = Output{
				Code: ctx.Code,
				Data: string(ctx.Output),
			}
			b, err := json.Marshal(o)
			if err != nil {
				panic(err)
			}
			return string(b)
		}(input)
		return output
	}))
	select {}
}
