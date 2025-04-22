package main

import (
	"g5/pkg/g5"
	"net/http"
)

func main() {
	engine := g5.Default()

	engine.GET("ping", func(ctx *g5.Context) {
		ctx.JSON(http.StatusOK, map[string]string{"message": "pong"})
	})

	engine.POST("hello", func(ctx *g5.Context) {
		var body struct {
			Name string `json:"name"`
		}
		if err := ctx.BindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
			return
		}
		ctx.JSON(http.StatusOK, map[string]string{"greeting": "Hello, " + body.Name})
	})

	engine.Run()

}
