package view

import (
	"path/filepath"
	"github.com/gin-contrib/multitemplate"
	"github.com/madasatya6/gin-gonic/pkg/funcmap"
)

//reference https://github.com/gin-contrib/multitemplate
var ViewLocation = "app/views/"

func Loc(file string) string {
	return filepath.Join(ViewLocation, file)
}

func New() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("example", Loc("example/partial/base.html"), Loc("example/partial/content.html"), Loc("example/partial/footer.html"))
	r.AddFromFilesFuncs("example.validation", funcmap.FuncMap, Loc("example/form-validation.html"))
	r.AddFromFilesFuncs("example.language", funcmap.FuncMap, Loc("example/language.html"))
	r.AddFromFilesFuncs("example.pagination", funcmap.FuncMap, Loc("example/pagination.html"))
	r.AddFromFilesFuncs("example.upload", funcmap.FuncMap, Loc("example/form-upload.html"))
	r.AddFromFilesFuncs("example.upload.multiple", funcmap.FuncMap, Loc("example/form-upload-multiple.html"))
	r.AddFromFilesFuncs("example.captcha", funcmap.FuncMap, Loc("example/form-captcha.html"))
	r.AddFromFilesFuncs("example.validation.multiple", funcmap.FuncMap, Loc("example/form-multiple.html"))
	return r
}
