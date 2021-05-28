package utils

import (
	"log"

	"github.com/iimrudy/prismacontroller/structures"
	"github.com/micmonay/keybd_event"
)

func RunCommand(command structures.Command) error {
	keys := []int{}
	for _, b := range command.Buttons {
		x := GetKeyCode(b)
		if x != -1 {
			keys = append(keys, x)
		}
	}

	kb, err := keybd_event.NewKeyBonding()
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

	//launch
	return kb.Launching()
}
