// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"
	"context"

	"github.com/gin-gonic/gin"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	_ "github.com/madasatya6/gin-gonic/internal/entity"
	"github.com/madasatya6/gin-gonic/pkg/pagination"
)

func (r *translationRoutes) GETPagination(c *gin.Context) {
	//contoh penggunaan validation
	//gunakan name inputan '_csrf'
	//kalo di header gunakan 'X-CSRF-TOKEN'
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	repo := r.t.Setting()
	settings, err := repo.AllWithPagination(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
		return
	}

	perpage := 4
	datas := pagination.Paginate(c, settings, perpage)
	
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Daftar settings",
		"data": gin.H{
			"settings": datas["posts"],
		},
	})
}

func (r *translationRoutes) GETPaginationView(c *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	PageNum := c.Query("p")
	if PageNum == "" {
		PageNum = "1"
	}

	repository := r.t.Setting()
	result, err := repository.AllWithPagination(ctx)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": err.Error(),
		})
	}

	perpage := 4
	settings := pagination.Paginate(c, result, perpage)

	var data = gin.H{
		"context": c,
		"settings": settings,
		"title": "Learn Pagination",
		"PageNum": PageNum,
	}
	c.HTML(http.StatusOK, "example.pagination", data)
} 




