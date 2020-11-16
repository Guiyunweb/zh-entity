package http

import (
	"github.com/kataras/iris/v12"
	"shiki-template/api"
)

func LoadRouter(app *iris.Application) {
	test(app)
}

func test(app *iris.Application) {
	app.Post("/conversion/java", api.JavaConversion)
}
