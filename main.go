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

	m := new(structures.Configuration)

	content, err := utils.ReadFileToString("config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	err = yaml.Unmarshal([]byte(*content), &m)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	//fmt.Printf("--- m:\n%v\n\n", m)

	//fmt.Println(utils.RunCommand(m.COMMANDS[0]))

	app := fiber.New()

	app.Post("/commands/execute", func(c *fiber.Ctx) error {
		r := new(structures.CommandRequest) // new instance

		c.BodyParser(r) // parse json
		success := false
		message := "command not found"
		if r.Password == m.PASSWORD {
			for _, cmd := range m.COMMANDS {
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
		if r.Password == m.PASSWORD {
			success = true
			mcmds := []structures.MinifiedCommand{}
			for _, cmd := range m.COMMANDS {
				mcmds = append(mcmds, utils.CommandToMiniCommand(cmd))
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
		if r.Password == m.PASSWORD {
			success = true
			message = "Valid password."
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{"success": success, "message": message})
	})

	app.Listen(m.HOST + ":" + m.PORT)
}
