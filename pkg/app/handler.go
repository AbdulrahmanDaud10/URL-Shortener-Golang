package app

import (
	"net/http"
	"sync"

	"github.com/AbdulrahmanDaud10/url-shortner/pkg/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var cache = &sync.Map{}

func Redirect(db *gorm.DB, ctx *gin.Context) {
	shortURL := ctx.Param("shortURL")
	id := api.Base62Decode(shortURL)
	// Try to get the long URL from the cache
	if longURL, ok := cache.Load(shortURL); ok {
		ctx.Redirect(http.StatusMovedPermanently, longURL.(string))
		return
	}
	// If it's not in the cache, get it from the database
	err := db.Exec("SELECT long_url FROM urls WHERE id=$1", id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var url api.URL
	// Store the long URL in the cache
	cache.Store(shortURL, url.LongURL)
	ctx.Redirect(http.StatusMovedPermanently, url.LongURL)
}
