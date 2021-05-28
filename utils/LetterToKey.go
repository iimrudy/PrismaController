package utils

import (
	"github.com/micmonay/keybd_event"
)

/*
	VK_ESC = keybd_event.VK_
	VK_keybd_event.VK_   = 2
	VK_2   = 3
	VK_3   = 4
	VK_4   = 5
	VK_5   = 6
	VK_6   = 7
	VK_7   = 8
	VK_8   = 9
	VK_9   = keybd_event.VK_0
	VK_0   = keybd_event.VK_keybd_event.VK_
	VK_Q   = keybd_event.VK_6
	VK_W   = keybd_event.VK_7
	VK_E   = keybd_event.VK_8
	VK_R   = keybd_event.VK_9
	VK_T   = 20
	VK_Y   = 2keybd_event.VK_
	VK_U   = 22
	VK_I   = 23
	VK_O   = 24
	VK_P   = 25
	VK_A   = 30
	VK_S   = 3keybd_event.VK_
	VK_D   = 32
	VK_F   = 33
	VK_G   = 34
	VK_H   = 35
	VK_J   = 36
	VK_K   = 37
	VK_L   = 38
	VK_Z   = 44
	VK_X   = 45
	VK_C   = 46
	VK_V   = 47
	VK_B   = 48
	VK_N   = 49
	VK_M   = 50
	VK_Fkeybd_event.VK_  = 59
	VK_F2  = 60
	VK_F3  = 6keybd_event.VK_
	VK_F4  = 62
	VK_F5  = 63
	VK_F6  = 64
	VK_F7  = 65
	VK_F8  = 66
	VK_F9  = 67
	VK_Fkeybd_event.VK_0 = 68
	VK_Fkeybd_event.VK_keybd_event.VK_ = 87
	VK_Fkeybd_event.VK_2 = 88
*/

func GetKeyCode(a string) int {
	switch a {
	case "esc":
		return keybd_event.VK_ESC
	case "win":
		return 0x5B + 0xFFF
	case "canc":
		return 0x03 + 0xFFF
	case "0":
		return keybd_event.VK_0
	case "1":
		return keybd_event.VK_1
	case "2":
		return keybd_event.VK_2
	case "3":
		return keybd_event.VK_3
	case "4":
		return keybd_event.VK_4
	case "5":
		return keybd_event.VK_5
	case "6":
		return keybd_event.VK_6
	case "7":
		return keybd_event.VK_7
	case "8":
		return keybd_event.VK_8
	case "a":
		return keybd_event.VK_A
	case "b":
		return keybd_event.VK_B
	case "c":
		return keybd_event.VK_C
	case "d":
		return keybd_event.VK_D
	case "e":
		return keybd_event.VK_E
	case "f":
		return keybd_event.VK_F
	case "g":
		return keybd_event.VK_G
	case "h":
		return keybd_event.VK_H
	case "i":
		return keybd_event.VK_I
	case "j":
		return keybd_event.VK_J
	case "k":
		return keybd_event.VK_K
	case "l":
		return keybd_event.VK_L
	case "m":
		return keybd_event.VK_M
	case "n":
		return keybd_event.VK_N
	case "o":
		return keybd_event.VK_O
	case "p":
		return keybd_event.VK_P
	case "q":
		return keybd_event.VK_Q
	case "r":
		return keybd_event.VK_R
	case "s":
		return keybd_event.VK_S
	case "t":
		return keybd_event.VK_T
	case "u":
		return keybd_event.VK_U
	case "w":
		return keybd_event.VK_W
	case "x":
		return keybd_event.VK_X
	case "y":
		return keybd_event.VK_Y
	case "z":
		return keybd_event.VK_Z
	default:
		return -1
	}

}
