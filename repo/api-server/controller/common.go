package controller

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"go-scribblog/repo/api-server/view"
	"go-scribblog/repo/api-server/view/errs"
	"go-scribblog/repo/log"
	"go.uber.org/zap"
	"reflect"
)

// ParseRequest parse the body of request to the object req
func ParseRequest(ctx iris.Context, req interface{}) error {
	if err := ctx.ReadJSON(req); err != nil {
		body, _ := ctx.GetBody()
		log.Error("parse request error", zap.String("body", string(body)), zap.Error(err))
		return processErr(req, err)
	}
	return nil
}

// ProcessErr go validator参数校验器自定义规则及提示
func processErr(u interface{}, err error) error {
	if err == nil { //如果为nil 说明校验通过
		return nil
	}
	invalid, ok := err.(*validator.InvalidValidationError) //如果是输入参数无效，则直接返回输入参数错误
	if ok {
		return fmt.Errorf("输入参数错误：" + invalid.Error())
	}
	validationErrs, ok := err.(validator.ValidationErrors) //断言是ValidationErrors
	if !ok {
		return err
	}
	for _, validationErr := range validationErrs {
		fieldName := validationErr.Field() //获取是哪个字段不符合格式
		typeOf := reflect.TypeOf(u)
		// 如果是指针，获取其属性
		if typeOf.Kind() == reflect.Ptr {
			typeOf = typeOf.Elem()
		}
		if field, ok1 := typeOf.FieldByName(fieldName); ok1 { //通过反射获取filed
			if errorInfo, ok2 := field.Tag.Lookup("reg_err_info"); ok2 { // 获取field对应的reg_error_info tag值
				return fmt.Errorf(fieldName + ":" + errorInfo) // 返回错误
			}
		} else {
			return err
		}
	}
	return err
}

func ReturnRequestParseError(ctx iris.Context, msg string) {
	ctx.JSON(view.RespWithError(errs.WrapError(errs.RequestParseError, msg), nil))
}
