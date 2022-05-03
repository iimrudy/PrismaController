package handlers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/app"
	"io/fs"
	"net/http"
)

func InitRoutes(prisma *app.PrismaController) error {

	// session handler
	prisma.Gin.Use(sessions.Sessions("pc_session", app.Get().Store))

	// Authorize
	prisma.Gin.POST("/authorize", AuthorizationHandler)

	// icons
	prisma.Gin.Static("/icons", prisma.Path+"icons")

	// Get Commands
	prisma.Gin.GET("/commands/get", CommandsGet)

	// Handle Command
	prisma.Gin.POST("/commands/execute", CommandsExecute)

	// Static files
	if sub, err := fs.Sub(prisma.Static, "static"); err == nil {
		prisma.Gin.Use(func() gin.HandlerFunc {
			fss := http.FS(sub)
			fileserver := http.StripPrefix("/", http.FileServer(fss))
			return func(c *gin.Context) {
				fileserver.ServeHTTP(c.Writer, c.Request)
				c.Header("Cache-Control", "public, max-age=31536000")
			}
		}())

	} else {
		return err
	}

	return nil
}
