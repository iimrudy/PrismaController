package app

import (
	"embed"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/prismacontroller"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
	"gopkg.in/yaml.v2"
	"log"
)

var instance *prismacontroller.PrismaController

func Init(path string, static embed.FS) (*prismacontroller.PrismaController, error) {
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
	prisma := &prismacontroller.PrismaController{
		Path:          path,
		Static:        static,
		Address:       fmt.Sprintf("%s:%s", config.HOST, config.PORT),
		Configuration: config,
		Gin:           gin.Default(),
	}

	instance = prisma
	return prisma, nil
}

func Get() *prismacontroller.PrismaController {
	return instance
}
