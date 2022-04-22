package utils

import (
	"github.com/iimrudy/prismacontroller/structures"
	"math/rand"
)

func ButtonToMinifiedButton(command structures.Button) structures.MinifiedButton {
	return structures.MinifiedButton{Name: command.Name, Logo: command.Logo, DisplayName: command.DisplayName}
}

func RandomString(length int) string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result)
}
