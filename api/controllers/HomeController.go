package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/michael-kagnew/portfolio/api/models"
)

type HomeController struct {
}

func GetHome(ctx *gin.Context) {
	models.GetProjects()
}
