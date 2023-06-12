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
	// Try to get the long URL from the cache
	if longURL, ok := cache.Load(shortURL); ok {
		ctx.Redirect(http.StatusMovedPermanently, longURL.(string))
		return
	}

	// TODO: If it's not in the cache, get it from the database

	var url api.URL
	// Store the long URL in the cache
	cache.Store(shortURL, url.LongURL)
	ctx.Redirect(http.StatusMovedPermanently, url.LongURL)
}
