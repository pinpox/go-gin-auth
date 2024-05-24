package controllers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go-gin-auth/utils"
)

func Index(c *gin.Context) {
	var username = ""
	session := sessions.Default(c)
	username, _ = session.Get("user").(string)
	utils.RenderTemplate(c, "index.tmpl", gin.H{"Title": "Notes", "user": username})
}
