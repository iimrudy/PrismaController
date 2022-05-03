package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/app"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
	"log"
	"sync"
)

type CommandInExecution struct {
	sync.Mutex
	inExec []string
}

func (c *CommandInExecution) Contains(name string) bool {
	c.Lock()
	defer c.Unlock()
	for _, v := range c.inExec {
		if v == name {
			return true
		}
	}
	return false
}

func (c *CommandInExecution) Remove(name string) {
	c.Lock()
	defer c.Unlock()
	for i, v := range c.inExec {
		if v == name {
			c.inExec = append(c.inExec[:i], c.inExec[i+1:]...)
			return
		}
	}
}

func (c *CommandInExecution) Add(name string) {
	c.Lock()
	defer c.Unlock()
	c.inExec = append(c.inExec, name)
}

var inexec = CommandInExecution{}

func CommandsExecute(ctx *gin.Context) {
	var rq structures.ClickButtonRequest
	ctx.BindJSON(&rq)

	success := false
	message := "command not found"
	sx := sessions.Default(ctx)
	if sx.Get("authorized") != nil && sx.Get("authorized").(bool) {
		for _, cmd := range app.Get().Configuration.BUTTONS {
			if cmd.Name == rq.CommandName {
				success = true
				message = "executed"
				if !inexec.Contains(cmd.Name) {
					inexec.Add(cmd.Name)
					if err := utils.RunCommand(cmd); err != nil {
						success = false
						message = err.Error()
						log.Println(err.Error())
					}
					inexec.Remove(cmd.Name)
				} else {
					success = false
					message = "already in execution"
				}
				break
			}
		}
	} else {
		success = false
		message = "Unauthorized."
		sx.Clear()
		sx.Save()
	}

	ctx.JSON(200, gin.H{"success": success, "message": message})
}
