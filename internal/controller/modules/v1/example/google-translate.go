// Package v1 implements routing paths. Each services in own file.
package example

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/madasatya6/gin-gonic/internal/usecase/webapi"
	_ "github.com/madasatya6/gin-gonic/internal/usecase/repo"
	"github.com/madasatya6/gin-gonic/internal/entity"
)

func (r *translationRoutes) GoogleTranslate(c *gin.Context) {
	//contoh penggunaan validation
	//gunakan name inputan '_csrf'
	//kalo di header gunakan 'X-CSRF-TOKEN'
	g := webapi.New()
	makanTrans := entity.Translation{
		Source: "auto",
		Destination: "en",
		Original: "Sedang makan nasi goreng.",
	}
	makan, err := g.Translate(makanTrans) 
	if err != nil {
		c.PureJSON(http.StatusInternalServerError, gin.H{
			"status": false,
			"message": err.Error(),
			"data": nil,
		})
	}

	//use logger
	r.l.Info("Berhasil ditranslate oleh google translate: ", makan.Translation)
	c.PureJSON(http.StatusOK, gin.H{
		"status": true,
		"message": "Translation: Google translate",
		"data": gin.H{
			"makan": makan,
		},
	})
}


