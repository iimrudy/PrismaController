package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/app"
	"github.com/iimrudy/prismacontroller/structures"
)

func AuthorizationHandler(ctx *gin.Context) {
	var pass structures.PasswordRequest
	ctx.BindJSON(&pass)
	fmt.Printf("PASS --> ", pass)

	var success bool = false
	var message interface{} = "Invalid password."

	if pass.Password == app.Get().Configuration.PASSWORD {
		success = true
		message = "Valid password."
	}
	ctx.JSON(200, gin.H{
		"success": success,
		"message": message,
	})
}
