package main

import (
	"flag"
	"log"
	"os"

	"github.com/getlantern/systray"
	"github.com/gobuffalo/packr/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
	"gopkg.in/yaml.v2"
)

var (
	Config       structures.Configuration
	fiberAPP     *fiber.App
	StaicDataBox *packr.Box
	Path         string
)

func init() {
	flag.StringVar(&Path, "path", "./", "Configuration & Icons location")
	flag.Parse() // parse args

	file, err := os.OpenFile(Path+"logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

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
}

func main() {
	StaicDataBox = packr.New("Static Assets", "/static")

	go systray.Run(onReady, onExit)
	log.Println("PrismaController Started")

	fiberAPP = fiber.New()

	fiberAPP.Post("/commands/execute", func(c *fiber.Ctx) error {
		r := new(structures.ClickButtonRequest) // new instance

		c.BodyParser(r) // parse json
		success := false
		message := "command not found"
		if r.Password == Config.PASSWORD {
			for _, cmd := range Config.BUTTONS {
				if cmd.Name == r.CommandName {
					go func() {
						err := utils.RunCommand(cmd)
						if err != nil {
							log.Println(err.Error())
						}
					}()
					success = true
					message = "executed"
					break
				}
			}
		} else {
			success = false
			message = "Invalid password."
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": success, "message": message})
	})

	fiberAPP.Use("/", filesystem.New(filesystem.Config{
		Root: StaicDataBox,
	}))

	fiberAPP.Static("/icons", Path+"icons")

	fiberAPP.Post("/commands/get", func(c *fiber.Ctx) error {
		r := new(structures.PasswordRequest) // new instance

		c.BodyParser(r) // parse json
		success := false
		var message interface{}
		message = "Invalid password."
		if r.Password == Config.PASSWORD {
			success = true
			mcmds := []structures.MinifiedButton{}
			for _, cmd := range Config.BUTTONS {
				mcmds = append(mcmds, utils.ButtonToMinifiedButton(cmd))
			}
			message = mcmds
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": success, "message": message})
	})

	fiberAPP.Post("/authorize", func(c *fiber.Ctx) error {
		r := new(structures.PasswordRequest) // new instance

		c.BodyParser(r) // parse json
		success := false
		var message interface{}
		message = "Invalid password."
		if r.Password == Config.PASSWORD {
			success = true
			message = "Valid password."
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": success, "message": message})
	})
	fiberAPP.Listen(Config.HOST + ":" + Config.PORT)
}

func onReady() {
	s, err := StaicDataBox.FindString("favicon.ico")
	if err != nil {
		log.Panicln(err.Error())
	}

	systray.SetIcon([]byte(s))
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
	fiberAPP.Shutdown()
	os.Exit(0)
}
