package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ereminiu/url-shortener/pkg/models"
	"github.com/ereminiu/url-shortener/pkg/service"
	"github.com/gin-gonic/gin"
)

type Link struct {
	Link string `json:"link" binding:"required"`
}

func CreateLink(c *gin.Context) {
	var lnk Link
	if err := c.BindJSON(&lnk); err != nil {
		log.Fatalf("Failed to read link %s\n", err)
		return
	}

	// check whether link is valid
	if !service.ValidLink(lnk.Link) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"messange": "Unvalid link",
		})
		return
	}

	h, err := service.CreateLink(lnk.Link)
	if err != nil {
		log.Fatalf("Failed to create link %s\n", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hash": h,
	})
}

type HashCode struct {
	Hash string `json:"code" binding:"required"`
}

func GetLink(c *gin.Context) {
	var code HashCode
	if err := c.BindJSON(&code); err != nil {
		log.Fatalf("Failed to read code %s\n", err)
		return
	}

	link, err := service.GetLink(code.Hash)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"messange": "Hashcode not found",
		})
		log.Printf("Failed to get link %s\n", err)
		return
	}

	c.Redirect(http.StatusFound, link)
}

func CreateCustomLink(c *gin.Context) {
	var lnk models.CustomLink
	if err := c.BindJSON(&lnk); err != nil {
		log.Printf("Failed to read link %s\n", err)
		return
	}

	// check whether link is valid
	if !service.ValidLink(lnk.Link) {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"messange": "Unvalid link",
		})
		return
	}

	// check whether customCode is aleady used for other link
	if service.LinkExist(lnk.CustomCode, "customlinks", "custom") {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"messange": "This custom link is already exist",
		})
		return
	}

	err := service.CreateCustomLink(lnk)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"messange": "Failed to create your custom links: Your code should not exceed 40 characters",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"messange": "Your custom link created successfuly",
	})
}

type CCode struct {
	CustomCode string `json:"custom_code" binding:"required"`
}

func GetCustomLink(c *gin.Context) {
	var code CCode
	if err := c.BindJSON(&code); err != nil {
		log.Fatalf("Failed to read code %s\n", err)
		return
	}

	fmt.Println(code.CustomCode)

	link, err := service.GetCustomLink(code.CustomCode)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"messange": "Custom code not found",
		})
		log.Printf("Failed to get link %s\n", err)
		return
	}

	c.Redirect(http.StatusFound, link)
}
