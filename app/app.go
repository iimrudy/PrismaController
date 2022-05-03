package app

import (
	"embed"
	"fmt"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
	"gopkg.in/yaml.v2"
	"log"
	"net/http"
	"os"
)

type PrismaController struct {
	Static        embed.FS
	Address       string
	Gin           *gin.Engine
	Configuration *structures.Configuration
	Server        *http.Server
	Path          string
	Store         cookie.Store
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

	if ct, err := utils.ReadFileToString(path + "key.prismac"); err != nil {
		if os.IsNotExist(err) {
			log.Println("No key found, generating one")
			if err := utils.WriteStringToFile(path+"key.prismac", utils.RandomString(64)); err != nil {
				return nil, err
			} else {
				return Init(path, static)
			}
		} else {
			config.SESSION_SECRET = *ct
		}
	}

	prisma := &PrismaController{
		Path:          path,
		Static:        static,
		Address:       fmt.Sprintf("%s:%s", config.HOST, config.PORT),
		Configuration: config,
		Gin:           gin.Default(),
		Store:         cookie.NewStore([]byte(config.SESSION_SECRET)),
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
