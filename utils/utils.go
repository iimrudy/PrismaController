package utils

import "github.com/iimrudy/prismacontroller/structures"

func ButtonToMinifiedButton(command structures.Button) structures.MinifiedButton {
	return structures.MinifiedButton{Name: command.Name, Logo: command.Logo, DisplayName: command.DisplayName}
}
