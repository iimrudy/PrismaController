package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/app"
	"github.com/iimrudy/prismacontroller/structures"
)

func AuthorizationHandler(ctx *gin.Context) {
	var pass structures.PasswordRequest
	ctx.BindJSON(&pass)

	var success bool = false
	var message interface{} = "Invalid password."
	sx := sessions.Default(ctx)

	if sx.Get("authorized") == true {
		success = true
		message = "Already authorized."
	} else {
		if pass.Password == app.Get().Configuration.PASSWORD {
			success = true
			message = "Valid password."
			sx.Set("authorized", true)
			sx.Save()
		} else {
			sx.Clear()
			sx.Save()
		}
	}

	ctx.JSON(200, gin.H{
		"success": success,
		"message": message,
	})
}
