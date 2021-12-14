package language

import (
	"github.com/madasatya6/gin-gonic/pkg/structure"
	"github.com/madasatya6/gin-gonic/pkg/language/english"
	"github.com/madasatya6/gin-gonic/pkg/language/indonesia"
)

//set language
var Language *string

func WithLang(text string) string {
	
	if Language == nil {
		var x = "id"
		Language = &x
	} else if (*Language) == "" {
		var x = "id"
		Language = &x
	}

	var lang  = map[string]structure.Language{
		"en": english.Language,
		"id": indonesia.Language,
	}
	
	return lang[*Language].Get(text)
}
