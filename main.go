package main

import (
	"log"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/iimrudy/prismacontroller/structures"
	"github.com/iimrudy/prismacontroller/utils"
	"gopkg.in/yaml.v2"
)

func main() {

	config := new(structures.Configuration)

	content, err := utils.ReadFileToString("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(*content), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- m:\n%v\n\n", m)

	//fmt.Println(utils.RunCommand(m.COMMANDS[0]))

	app := fiber.New()

	app.Post("/commands/execute", func(c *fiber.Ctx) error {
		r := new(structures.ClickButtonRequest) // new instance

		c.BodyParser(r) // parse json
		success := false
		message := "command not found"
		if r.Password == config.PASSWORD {
			for _, cmd := range config.BUTTONS {
				if cmd.Name == r.CommandName {
					go utils.RunCommand(cmd)
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

	app.Use("/", filesystem.New(filesystem.Config{
		Root: rice.MustFindBox("static").HTTPBox(),
	}))

	app.Static("/icons", "./icons")

	app.Post("/commands/get", func(c *fiber.Ctx) error {
		r := new(structures.PasswordRequest) // new instance

		c.BodyParser(r) // parse json
		success := false
		var message interface{}
		message = "Invalid password."
		if r.Password == config.PASSWORD {
			success = true
			mcmds := []structures.MinifiedButton{}
			for _, cmd := range config.BUTTONS {
				mcmds = append(mcmds, utils.ButtonToMinifiedButton(cmd))
			}
			message = mcmds
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": success, "message": message})
	})

	app.Post("/authorize", func(c *fiber.Ctx) error {
		r := new(structures.PasswordRequest) // new instance

		c.BodyParser(r) // parse json
		success := false
		var message interface{}
		message = "Invalid password."
		if r.Password == config.PASSWORD {
			success = true
			message = "Valid password."
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": success, "message": message})
	})

	app.Listen(config.HOST + ":" + config.PORT)
}
