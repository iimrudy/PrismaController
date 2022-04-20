package prismacontroller

import (
	"embed"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/structures"
	"net/http"
)

type PrismaController struct {
	Static        embed.FS
	Address       string
	Gin           *gin.Engine
	Configuration *structures.Configuration
	Server        *http.Server
	Path          string
}

func (pc *PrismaController) Listen() error {
	return pc.Server.ListenAndServe()
}
