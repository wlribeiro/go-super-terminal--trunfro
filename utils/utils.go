package utils

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"runtime"
	"super-trunfo/internal/cards"
)

func GetCardAttributeValue(card cards.Card, attribute string) (int, error) {
	fieldValue := reflect.ValueOf(card).FieldByName(attribute)
	if !fieldValue.IsValid() {
		return 0, fmt.Errorf("no such field: %s in card", attribute)
	}

	if fieldValue.Kind() != reflect.Int {
		return 0, fmt.Errorf("field %s is not of type int", attribute)
	}

	// Return the field value as int
	return int(fieldValue.Int()), nil
}

// clearScreen limpa a tela do terminal
func ClearScreen() {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux", "darwin": // macOS também usa "darwin"
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		panic("Unsupported platform")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}

func AwaitRead() {
	fmt.Println("Pressione enter para continuar...")
	fmt.Scanln()
}
