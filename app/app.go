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
	"os"
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

	if len(config.SESSION_SECRET) == 0 {
		config.SESSION_SECRET = utils.RandomString(64)
		bytes, err := yaml.Marshal(config)
		if err != nil {
			return nil, err
		}
		if err := utils.WriteStringToFile(path+"config.yml", string(bytes)); err != nil {
			return nil, err
		}
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
