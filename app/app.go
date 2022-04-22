package app

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
	"gopkg.in/yaml.v2"
	"log"
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

var instance *PrismaController

func Init(path string, static embed.FS) (*PrismaController, error) {
	if instance != nil {
		return instance, nil
	}
	log.Println("Using path" + path)
	content, err := utils.ReadFileToString(path + "config.yml")
	if err != nil {
		return nil, err
	}
	config := &structures.Configuration{}
	err = yaml.Unmarshal([]byte(*content), config)
	if err != nil {
		return nil, err
	}
	prisma := &PrismaController{
		Path:          path,
		Static:        static,
		Address:       fmt.Sprintf("%s:%s", config.HOST, config.PORT),
		Configuration: config,
		Gin:           gin.Default(),
	}
	prisma.Server = &http.Server{
		Addr:    prisma.Address,
		Handler: prisma.Gin,
	}

	instance = prisma
	return prisma, nil
}

func Get() *PrismaController {
	return instance
}
