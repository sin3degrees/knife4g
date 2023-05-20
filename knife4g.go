package knife4g

import (
	"embed"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"

	"github.com/kataras/iris/v12"
)

var (
	//go:embed front
	front   embed.FS
	docJson []byte
	s       service
)

type Config struct {
	RelativePath string
}

type service struct {
	Name           string `json:"name"`
	Url            string `json:"url"`
	SwaggerVersion string `json:"swaggerVersion"`
	Location       string `json:"location"`
}

func init() {
	var err error
	docJson, err = os.ReadFile("./docs/swagger.json")
	if err != nil {
		log.Println("no swagger.json found in ./docs")
	}
}

func Handler(config Config) iris.Handler {
	docJsonPath := config.RelativePath + "/docJson"
	servicesPath := config.RelativePath + "/front/service"
	docPath := config.RelativePath + "/index"
	appjsPath := config.RelativePath + "/front/webjars/js/app.42aa019b.js"

	s.Url = "/docJson"
	s.Location = "/docJson"
	s.Name = "API Documentation"
	s.SwaggerVersion = "2.0"

	appjsTemplate, err := template.New("app.42aa019b.js").
		Delims("{[(", ")]}").
		ParseFS(front, "front/webjars/js/app.42aa019b.js")
	if err != nil {
		log.Println(err)
	}
	docTemplate, err := template.New("doc.html").
		Delims("{[(", ")]}").
		ParseFS(front, "front/doc.html")
	if err != nil {
		log.Println(err)
	}

	return func(ctx iris.Context) {
		if ctx.Method() != http.MethodGet {
			ctx.StopWithStatus(http.StatusMethodNotAllowed)
			return
		}
		switch ctx.Path() {
		case appjsPath:
			err := appjsTemplate.Execute(ctx.ResponseWriter(), config)
			if err != nil {
				log.Println(err)
			}
		case servicesPath:
			ctx.JSON([]service{s})
		case docPath:
			err := docTemplate.Execute(ctx.ResponseWriter(), config)
			if err != nil {
				log.Println(err)
			}
		case docJsonPath:
			ctx.ContentType("application/json; charset=utf-8")
			ctx.StatusCode(http.StatusOK)
			ctx.Write(docJson)
		default:
			// 一下由*gin.Context.FileFromFS()修改而来
			filepath := strings.TrimPrefix(ctx.Path(), config.RelativePath)
			fs := http.FS(front)
			defer func(old string) {
				ctx.Request().URL.Path = old
			}(ctx.Request().URL.Path)
			ctx.Request().URL.Path = filepath
			http.FileServer(fs).ServeHTTP(ctx.ResponseWriter(), ctx.Request())
		}

	}
}
