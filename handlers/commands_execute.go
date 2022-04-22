package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/app"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
	"log"
)

func CommandsExecute(ctx *gin.Context) {
	var rq structures.ClickButtonRequest
	ctx.BindJSON(&rq)

	success := false
	message := "command not found"
	if rq.Password == app.Get().Configuration.PASSWORD {
		for _, cmd := range app.Get().Configuration.BUTTONS {
			if cmd.Name == rq.CommandName {

				err := utils.RunCommand(cmd)
				if err != nil {
					log.Println(err.Error())
				}

				success = true
				message = "executed"
				break
			}
		}
	} else {
		success = false
		message = "Invalid password."
	}

	ctx.JSON(200, gin.H{"success": success, "message": message})
}
