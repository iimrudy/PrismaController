package main

import (
	"context"
	"embed"
	"flag"
	"github.com/getlantern/systray"
	"github.com/gin-gonic/gin"
	"github.com/iimrudy/prismacontroller/app"
	"github.com/iimrudy/prismacontroller/handlers"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

var (
	Config structures.Configuration
	Path   string
)

//go:embed static
var static embed.FS

func init() {
	flag.StringVar(&Path, "path", "./", "Configuration & Icons location")
	flag.Parse() // parse args

	/*file, err := os.OpenFile(Path+"logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}*

	//log.SetOutput(file)*/
	log.SetPrefix("[PrismaController] ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	content, err := utils.ReadFileToString(Path + "config.yml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	err = yaml.Unmarshal([]byte(*content), &Config)
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println("Using path" + Path)
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	log.Println("PrismaController Started")
	pc, err := app.Init(Path, static)
	if err != nil {
		log.Fatal(err)
	}

	if e := handlers.InitRoutes(pc); e != nil {
		log.Fatal(e)
	}

	go func() {
		if err := pc.Listen(); err != nil {
			log.Println(err)
		}
	}()
	systray.Run(onReady, onExit)
}

func onReady() {
	file, err := static.ReadFile("static/favicon.ico")
	if err != nil {
		log.Panicln(err.Error())
	}

	systray.SetIcon(file)
	systray.SetTitle("PrismaController")
	systray.SetTooltip("PrismaController Menu")
	quitBtn := systray.AddMenuItem("Quit", "Quit PrismaController")
	go func() {
		for {
			select {
			case <-quitBtn.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()
	// Sets the icon of a menu item. Only available on Mac and Windows.
}

func onExit() {
	log.Println("Exiting now.")
	pc := app.Get()
	if pc != nil {
		pc.Server.Shutdown(context.Background())
	}
	os.Exit(0)
}
