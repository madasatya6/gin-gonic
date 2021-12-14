package funcmap

import (
    "html/template"
	"fmt"
	"time"
	"net/url"
	"strings"
	"encoding/json"
	"log"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
	"github.com/madasatya6/gin-gonic/pkg/language"
	"github.com/madasatya6/gin-gonic/pkg/encryption"
	//pagination template
	"github.com/flosch/pongo2"
	"github.com/astaxie/beego/utils/pagination"
)

/***
* Funcmap 
* @author madasatya6
* @framework gin
*/
var Ctx *gin.Context
func SetContext(c *gin.Context){
	Ctx = c 
}

var FuncMap = template.FuncMap{
    "sumi": func(a,b int) int {
		return a+b
	},
	"sumf": func(a,b float64) float64 {
		return a+b
	},
	"mini": func(a,b int) int {
		return a-b
	},
	"minf": func(a,b float64) float64 {
		return a-b
	},
	"xi": func(a,b int) int {
		return a*b
	},
	"xf": func(a,b float64) float64 {
		return a*b
	},
	"di": func(a,b int) int {
		return a/b
	},
	"df": func(a,b float64) float64 {
		return a/b
	},
	"Text": func(text string) string {
		//use language
		return language.WithLang(text)
	},
	"SlcToString": func(data interface{}) string {
		var result string
		var temp []string 
		var val = reflect.ValueOf(data)

		switch reflect.TypeOf(data).Kind() {
			case reflect.Slice:
				for i:=0; i<val.Len(); i++ {
					temp = append(temp, val.Index(i).Interface().(string))
				}
		}

		result = strings.Join(temp, ",")
		return result
	},
	"avg": func(n ...int) int {
		var total = 0
		for _, angka := range n {
			total += angka
		}
		return total/len(n)
	},
	"unescape": func(s string) template.HTML{
		return template.HTML(s)
	},
	"TimetoYmd": func(tm time.Time) string {
		if tm.IsZero() {
			return ""
		} else {
			return fmt.Sprintf("%v-%v-%v %v:%v:%v", tm.Year(), tm.Month(), tm.Day(), tm.Hour(), tm.Minute(), tm.Second())
		}
		
	},
	"Iterate" : func(start uint, sign string, count uint) []uint {
		var i uint 
		var Items []uint 
		if sign == "<" {
			for i=start; i < count; i++ {
				Items = append(Items, i)
			}
		} else if sign == "<=" {
			for i=start; i <= count; i++ {
				Items = append(Items, i)
			}
		} else if sign == ">" {
			for i=start; i > count; i++ {
				Items = append(Items, i)
			}
		} else if sign == ">=" {
			for i=start; i >= count; i++ {
				Items = append(Items, i)
			}
		}
		return Items
	},
	"IndexString" : func(data []string, i uint) string {
		if len(data) > 0 {
			var index = int(i)
			if len(data) >= index { 
				return data[index]
			} else {
				return ""
			}
		} else { 
			return ""
		}
	},
	"Date" : func(date time.Time) string {
		format := fmt.Sprintf("%v %v %v", date.Day(), date.Month(), date.Year())
		return format
	},
	"Time" : func(date time.Time) string {
		format := date.Format("15:04")
		return format
	},
	"DateTime" : func(date time.Time) string {
		format := date.Format("2006-01-02 15:04")
		return format
	},
	"GetFlashdata" : func(c *gin.Context, key string) string {
		session := sessions.Default(c)
		var data = ArraytoString(session.Flashes(key))
		session.Save()
		return data
	},
	"UnmarshalJSONSession" : func(c *gin.Context, data string) (*map[string]interface{}) {
		session := sessions.Default(c)
		var sess = ArraytoString(session.Flashes(data))
		session.Save()
		var temp = new(map[string]interface{})
		err := json.Unmarshal([]byte(sess), &temp)
		if err != nil {
			log.Println(err.Error())
		}
		return temp
	},
	"FormError" : func(c *gin.Context, key string) string {
		var keyMsg = key+"-msg"
		var session = sessions.Default(c)
		var sess = session.Flashes(keyMsg)
		session.Save()
		var data = ArraytoString(sess)
		
		return data
	},
	"SetFlashdata" : func(c *gin.Context, key string, value string) bool {
		session := sessions.Default(c)
		session.AddFlash(value, key)
		session.Save()
		return true
	},
	"GETSession" : func(c *gin.Context, key string) string {
		session := sessions.Default(c)
		result := session.Get(key).(string)
		return result 
	},
	"QueryParam" : func(c *gin.Context, key string) string {
		var param = c.Query(key)
		return param
	},
	"Param" : func(c *gin.Context, key string) string {
		var param = c.Param(key)
		return param
	},
	"FindString" : func(sentence, key string) bool {
		var status = strings.Index(sentence, key)
		if status == -1 {
			return false
		} else {
			return true
		}
	},
	"IntToString" : func(num int) string {
		return fmt.Sprintf("%d", num)
	},
	"StringToInt" : func(num string) int {
		result, _ := strconv.Atoi(num)
		return result
	},
	"PaginatorIsActive": func(page int) bool {
		urlpath := ""
		u, _ := url.Parse(urlpath)
		values, _ := url.ParseQuery(u.RawQuery)
		p := values.Get("p")
		pageNow := fmt.Sprintf("%d", page)
		if pageNow == p {
			return true
		} else {
			return false
		}
	},
	"equal" : func(param1 string, param2 string) bool {
		if param1 != param2 {
			return false
		} else {
			return true
		}
	},
	"in_array": func(val interface{}, array interface{}) (exists bool) {
		exists = false
		switch reflect.TypeOf(array).Kind() {
			case reflect.Slice :
				s := reflect.ValueOf(array)
				for i:=0; i < s.Len(); i++ {
					if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
						exists = true 
						return 
					}
				}
		}
		return
	},
	"mkSlice": func(array ...interface{}) []interface{} {
		return array
	},
	"ShowPagination": func(c *gin.Context, style string, data pongo2.Context) template.HTML {
		//butuh css : <link rel="stylesheet" href="https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0/css/bootstrap.min.css" integrity="sha384-Gn5384xqQ1aoWXA+058RXPxPg6fy4IWvTNh0E263XmFcJlSAwiGgFAW/dAiS6JXm" crossorigin="anonymous">
		//contoh penggunaan di view : {{ShowPagination .Context "style1" .Blogs}}
		var result string
		var pageNum = c.Query("p")
		if pageNum == "" {
			pageNum = "1"
		}

		if style == "style1" {
			result = `<nav aria-label="Page navigation example">`
			paginator := data["paginator"].(*pagination.Paginator)
			if paginator.HasPages() {
				result += `<ul class="pagination justify-content-center">`
				if paginator.HasPrev() {
					pageLinkFirst := fmt.Sprintf("%v", paginator.PageLinkFirst())
					pageLinkPrev := fmt.Sprintf("%v", paginator.PageLinkPrev())
					result += `<li class="page-item"><a class="page-link" href="` + pageLinkFirst + `">First</a></li>`
					result += `<li class="page-item"><a class="page-link" href="` + pageLinkPrev + `">&lt;</a></li>`
				} else {
					result += `<li class="page-item disabled"><a class="page-link"> First</a></li>`
					result += `<li class="page-item disabled"><a class="page-link">&lt;</a></li>`
				}

				//implementasi nomor page
				for _, page := range paginator.Pages() {
					pageStr := fmt.Sprintf("%v", page)
					active := ``
					if pageStr == pageNum {
						active = `active`
					}
					pageLink := fmt.Sprintf("%v", paginator.PageLink(page))
					result += `<li class="page-item ` + active + `">`
					result += `<a class="page-link" href="` + pageLink + `">`
					result += pageStr
					result += `</a>`
				  	result += `</li>`
				}

				if paginator.HasNext() {
					pageLinkNext := fmt.Sprintf("%v", paginator.PageLinkNext())
					pageLinkLast := fmt.Sprintf("%v", paginator.PageLinkLast())
					result += `<li class="page-item"><a class="page-link" href="` + pageLinkNext + `">&gt;</a></li>`
					result += `<li class="page-item"><a class="page-link" href="` + pageLinkLast + `">Last</a></li>`
				} else {
					result += `<li class="page-item disabled"><a class="page-link">&gt;</a></li>`
					result += `<li class="page-item disabled"><a class="page-link">Last</a></li>`
				}

				result += `</ul>`
			}
			result += `</nav>`
		}
		return template.HTML(result)
	},
	//enkripsi
	"MD5": func(key interface{}) string {
		var keyStr string
		switch reflect.TypeOf(key).Kind(){
			case reflect.String :
				keyStr = fmt.Sprintf("%s", reflect.ValueOf(key).Interface().(string))
				break
			case reflect.Int :	
				keyStr = fmt.Sprintf("%d", reflect.ValueOf(key).Interface().(int))
				break
		}
		result := encryption.CreateMD5(keyStr)
		return result
	},
}	

func ArraytoString(array interface{}) string {
	var result string
	switch reflect.ValueOf(array).Kind() {
		case reflect.Slice:
			val := reflect.ValueOf(array)
			for i:=0; i<val.Len(); i++ {
				if i == 0 {
					result = fmt.Sprintf("%v", val.Index(i).Interface())
				} else {
					result = fmt.Sprintf("%v, %v", result, val.Index(i).Interface())
				}
			}
	}
	return result
}
