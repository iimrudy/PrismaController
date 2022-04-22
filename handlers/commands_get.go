package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/app"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
)

func CommandsGet(ctx *gin.Context) {
	var pass structures.PasswordRequest
	ctx.BindJSON(&pass)
	success := false
	var message interface{}
	message = "Invalid password."
	if pass.Password == app.Get().Configuration.PASSWORD {
		success = true
		var mcmds []structures.MinifiedButton
		for _, cmd := range app.Get().Configuration.BUTTONS {
			mcmds = append(mcmds, utils.ButtonToMinifiedButton(cmd))
		}
		message = mcmds
	}

	rep := gin.H{
		"success": success,
		"message": message,
	}
	ctx.JSON(200, rep)
}
