package utils

import (
	"errors"
	"log"
	"os/exec"
	"strings"

	"github.com/iimrudy/prismacontroller/structures"
	"github.com/micmonay/keybd_event"
)

func RunCommand(command structures.Button) {
	var err error
	if len(command.Keys) > 0 {
		var kb keybd_event.KeyBonding
		keys := []int{}
		for _, b := range command.Keys {
			x := GetKeyCode(b)
			if x != -1 {
				keys = append(keys, x)
			}
		}

		kb, err = keybd_event.NewKeyBonding()
		if err != nil {
			log.Fatal(err)
		}

		kb.SetKeys(keys...)

		kb.HasCTRL(command.HasCtrl)
		kb.HasCTRLR(command.HasRCtrl)

		kb.HasALT(command.HasAlt)
		kb.HasALTGR(command.HasRAlt)

		kb.HasSHIFT(command.HasShift)
		kb.HasSHIFTR(command.HasRShift)

		err = kb.Launching()
	} else if len(command.ShellCommand) > 0 {
		splitted := strings.Split(command.ShellCommand, " ")
		cmd_ := splitted[0]
		cmd := exec.Command(cmd_, remove(splitted, 0)...)
		err = cmd.Run()
	} else {
		err = errors.New("invalid Command '" + command.Name + "', no ShellCommand and no KeyCombinations")
	}

	if err != nil {
		log.Println("Error occurred while executing a command: " + err.Error())
	}
}

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}
