package repository

import (
	"net/http"
	"net/url"

	"github.com/AbdulrahmanDaud10/url-shortner/pkg/api"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Create(db *gorm.DB, ctx *gin.Context) {
	var url api.URL
	if err := ctx.ShouldBindJSON(&url); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}

	// Validate the URL
	if !URLValidityCheck(url.LongURL) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid URL"})
		return
	}

	// Insert the long URL into the database and get the generated ID
	var Id int
	err := db.Raw("INSERT INTO urls(long_url) VALUES($1) RETURNING id", url.LongURL).Scan(&Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	// Convert the ID to a short URL
	url.ShortURL = api.Base62Encode(Id)

	// Update the record with the short URL
	err = db.Exec("UPDATE urls SET short_url = $1 WHERE id = $2", url.ShortURL, Id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	ctx.JSON(http.StatusOK, url)
}

// Checking if the URL if Valid
func URLValidityCheck(urlString string) bool {
	u, err := url.Parse(urlString)
	return err == nil && u.Scheme != "" && u.Host != ""
}
