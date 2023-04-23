package middleware

import (
	"blog/utils"
	"github.com/gin-gonic/gin"
	"log"
	"reflect"
	"strconv"
	"strings"
)

func PostWrapper(f interface{}) gin.HandlerFunc {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("handler should be function")
	}
	num := t.NumIn()
	if num == 0 || num > 2 {
		panic("handler parameter number should be 1 or 2")
	}
	if num == 2 {
		if t.In(0) != reflect.TypeOf(&gin.Context{}) {
			panic("handler first parameter type should be *gin.Context")
		}
	}
	return func(ctx *gin.Context) {
		var req interface{}
		if num == 1 {
			req = newReqInstance(t.In(0))
		} else {
			req = newReqInstance(t.In(1))
		}
		//解析json, 如果只有一个参数就不用解析了
		if num != 1 {
			err := ctx.ShouldBindJSON(req)
			if err != nil {
				utils.EndWithError(ctx, err)
				return
			}
		}
		var inValue []reflect.Value
		if num == 1 {
			inValue = []reflect.Value{reflect.ValueOf(req).Elem()}
		} else {
			inValue = []reflect.Value{reflect.ValueOf(ctx), reflect.ValueOf(req).Elem()}
		}
		result := reflect.ValueOf(f).Call(inValue)
		data := result[0].Interface()
		err, _ := result[1].Interface().(error)
		if err == nil {
			if data == nil {
				utils.ResponseNormal(ctx)
				return
			}
			utils.ResponseWithData(ctx, data)
			return
		} else {
			utils.EndWithError(ctx, err)
		}
	}
}

func GetWrapper(f interface{}) gin.HandlerFunc {
	t := reflect.TypeOf(f)
	if t.Kind() != reflect.Func {
		panic("handler should be function")
	}
	num := t.NumIn()
	if num == 0 || num > 2 {
		panic("handler parameter number should be 1 or 2")
	}
	if num == 2 {
		if t.In(0) != reflect.TypeOf(&gin.Context{}) {
			panic("handler first parameter type should be *gin.Context")
		}
	}
	return func(ctx *gin.Context) {
		var req reflect.Value
		var reqType reflect.Type
		if num == 1 {
			reqType = t.In(0)
			req = reflect.New(t.In(0)).Elem()
		} else {
			reqType = t.In(1)
			req = reflect.New(t.In(1)).Elem()
		}
		//解析json, 如果只有一个参数就不用解析了
		for i := 0; num != 1 && i < reqType.NumField(); i++ {
			fieldType := reqType.Field(i)
			jsonTag := strings.Split(fieldType.Tag.Get("json"), ",")[0]
			structValue := ctx.Query(jsonTag)
			switch field := reqType.Field(i); field.Type.Kind() {
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				typedV, _ := strconv.ParseInt(structValue, 10, 64)
				req.FieldByName(fieldType.Name).SetInt(typedV)
			case reflect.String:
				req.FieldByName(fieldType.Name).SetString(structValue)
			case reflect.Uint, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				typedV, _ := strconv.ParseUint(structValue, 10, 64)
				req.FieldByName(fieldType.Name).SetUint(typedV)
			case reflect.Bool:
				req.FieldByName(fieldType.Name).SetBool(structValue == "true")
			default:
				log.Printf("field type %s not support yet", fieldType)
			}
		}
		var inValue []reflect.Value
		if num == 1 {
			inValue = []reflect.Value{req}
		} else {
			inValue = []reflect.Value{reflect.ValueOf(ctx), req}
		}
		result := reflect.ValueOf(f).Call(inValue)
		data := result[0].Interface()
		err, _ := result[1].Interface().(error)
		if err == nil {
			if data == nil {
				utils.ResponseNormal(ctx)
				return
			}
			utils.ResponseWithData(ctx, data)
			return
		} else {
			utils.EndWithError(ctx, err)
		}
	}
}

func newReqInstance(t reflect.Type) interface{} {
	switch t.Kind() {
	case reflect.Ptr, reflect.Interface:
		return newReqInstance(t.Elem())
	default:
		return reflect.New(t).Interface()
	}
}
