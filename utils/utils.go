package utils

import "github.com/iimrudy/prismacontroller/structures"

func CommandToMiniCommand(command structures.Command) structures.MinifiedCommand {
	return structures.MinifiedCommand{Name: command.Name, Logo: command.Logo, DisplayName: command.DisplayName}
}
