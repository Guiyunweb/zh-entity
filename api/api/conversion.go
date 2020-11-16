package api

import (
	"github.com/kataras/iris/v12/context"
	"shiki-template/model"
	"shiki-template/service"
)

func JavaConversion(context *context.Context) {
	dto := &model.EntityDTO{}
	if err := context.ReadJSON(dto); err != nil {
		panic(err.Error())
	} else {
		translate := service.JavaConversion(dto)
		successData(context, translate)
	}
}

func SQLConversion(context *context.Context) {
	dto := &model.EntityDTO{}
	if err := context.ReadJSON(dto); err != nil {
		panic(err.Error())
	} else {
		translate := service.SQLConversion(dto)
		successData(context, translate)
	}
}
