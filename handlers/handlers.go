package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/prismacontroller"
	"io/fs"
	"net/http"
)

func InitRoutes(prisma *prismacontroller.PrismaController) error {
	// Static files

	// icons
	prisma.Gin.Static("/icons", prisma.Path+"icons")

	// Get Commands
	prisma.Gin.POST("/commands/get", CommandsGet)

	// Handle Command
	prisma.Gin.POST("/commands/execute", CommandsExecute)

	// Authorize
	prisma.Gin.POST("/authorize", AuthorizationHandler)
	prisma.Server = &http.Server{
		Addr:    prisma.Address,
		Handler: prisma.Gin,
	}

	if sub, err := fs.Sub(prisma.Static, "static"); err == nil {
		prisma.Gin.Use(func() gin.HandlerFunc {
			fss := http.FS(sub)
			fileserver := http.StripPrefix("/", http.FileServer(fss))
			return func(c *gin.Context) {
				fileserver.ServeHTTP(c.Writer, c.Request)
				c.Abort()
			}
		}())

	} else {
		return err
	}

	return nil
}
