package core

import (
	zhcn "github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type App struct {
	Engine *gin.Engine
	Config *AppConfig
}
type AppConfig struct {
	Name string
	Port int
	Db DbConfig
	Casbin string
}
type DbConfig struct {
	Username string
	Password string
	Host string
	DataBase string
	Driver string
}
func NewApp() *App {
	e := gin.New()
	return &App{
		Engine: e,
	}
}

type Resp struct {
	Code int
	Msg  string
	Data interface{}
}

func Ok(data interface{}) *Resp {
	return &Resp{
		Code: 20000,
		Msg:  "ok",
		Data: data,
	}
}
func Err(code int,msg string) *Resp {
	return &Resp{
		Code: code,
		Msg:  msg,
	}
}
type Handler func(h BizHandler) gin.HandlerFunc
type BizHandler func(ctx *gin.Context) *Resp

var Do = func() Handler {
	return func(h BizHandler) gin.HandlerFunc {
		return func(ctx *gin.Context) {
			res := h(ctx)
			ctx.JSON(http.StatusOK, gin.H{
				"code": res.Code,
				"msg":  res.Msg,
				"data": res.Data,
			})
		}
	}
}()

type Router struct {
	e *gin.RouterGroup
}

func (r *Router) GET(path string, handler BizHandler) {
	r.e.GET(path, Do(handler))
}
func (r *Router) POST(path string, handler BizHandler) {
	r.e.POST(path, Do(handler))
}
func (r *Router) PUT(path string, handler BizHandler) {
	r.e.PUT(path, Do(handler))
}
func (r *Router) DELETE(path string, handler BizHandler) {
	r.e.DELETE(path, Do(handler))
}
func (r *Router) PATCH(path string, handler BizHandler) {
	r.e.PATCH(path, Do(handler))
}
func (r *Router) HEAD(path string, handler BizHandler) {
	r.e.HEAD(path, Do(handler))
}
func (r *Router) OPTIONS(path string, handler BizHandler) {
	r.e.OPTIONS(path, Do(handler))
}
func NewRouter(e *gin.RouterGroup) *Router {
	return &Router{
		e: e,
	}
}
var trans ut.Translator
var zhTran= zhcn.New()
var uni =ut.New(zhTran,zhTran)

var validate *validator.Validate

func UseValidate()  {
	validate = validator.New()
	validate.RegisterValidation("tname", customFunc)
	trans, _ = uni.GetTranslator("zh")
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		return
	}
}

func customFunc(fl validator.FieldLevel) bool {
	return false
}

func Valid(obj interface{}) string {
	err := validate.Struct(obj)
	if err!=nil {
		errs := err.(validator.ValidationErrors)
		e := errs.Translate(trans)
		for k,v:=range e {
			log.Println(k)
			return v
		}
 	}
	return ""
}
