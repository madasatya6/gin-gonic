package pagination

import (
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/utils/pagination"
	"github.com/flosch/pongo2"
)

var (
	paginator = &pagination.Paginator{}
	data = pongo2.Context{}
)

func NewSlice(start, count, step int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = start
		start += step
	}
	return s
}

func Paginate(c *gin.Context, data []interface{}, postsPerPage int) pongo2.Context {
	/*
	switch f := data.(type){
		case *users.UsersModel :
			model := f
	} */
	paginator = pagination.NewPaginator(c.Request, postsPerPage, len(data))
	idrange := NewSlice(paginator.Offset(), postsPerPage, 1)
	paginationdatas := make([]interface{}, 0)
	for _, num := range idrange {
		if num <= len(data)-1{
			datanow := data[num]
			paginationdatas = append(paginationdatas, datanow)
		}
	}
	result := pongo2.Context{"paginator":paginator, "posts":paginationdatas}

	return result
}