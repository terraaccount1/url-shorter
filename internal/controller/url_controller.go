package controller

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type iUrlService interface {
	SetUrl(ctx context.Context, url string, lenght int) (string, error)
	GetUrl(ctx context.Context, uri string) (string, error)
}

type UrlController struct {
	svc iUrlService
}

func NewUrlController(svc iUrlService) *UrlController {
	return &UrlController{
		svc: svc,
	}
}

func (ctrl *UrlController) SetUrl(c *gin.Context) {
	type data struct {
		URL string `json:"url"`
	}

	in := data{}
	if err := c.BindJSON(&in); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "failed to bind with struct",
		})
	}

	uri, err := ctrl.svc.SetUrl(context.Background(), in.URL, 8)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   true,
			"message": "failed to create shorted URL",
		})
	}

	url := c.Request.Host + "/" + uri
	c.JSON(http.StatusOK, gin.H{"url": url})
}

func (ctrl *UrlController) GetUrl(c *gin.Context) {
	uri := c.Param("uri")
	url, err := ctrl.svc.GetUrl(context.Background(), uri)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	c.Redirect(http.StatusPermanentRedirect, url)
}
