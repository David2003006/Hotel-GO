package menu

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func Tablero() {
	opcion := true
	for {
		fmt.Printf("------Hotel villa del Mar-------\n")
		fmt.Printf("opciones:\n1.Darse de alta.\n2.Reservar una habitacion.\n3.Pagar estadía.\n4.Actualizar Datos\n5.Salir\nOpcion a escoger:")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error al ingresar el valor.")
		}

		input = strings.TrimSpace(input)
		//imprimirTipoDato(input)

		if input != "5" {
			opciones(input)
		} else {
			break
		}

		if Salir(opcion) == false {
			break
		}
		limpiarTerminal()
	}
}

func opciones(input string) {

	switch input {
	case "1":
		limpiarTerminal()
		fmt.Printf("Esto proximamente va a hacer darse de alta.\n")
		break
	case "2":
		limpiarTerminal()
		fmt.Printf("Esto proximamente va a hacer Reservar una habitacion.\n")
		break
	case "3":
		limpiarTerminal()
		fmt.Printf("Esto proxiamente va a hacer pagar estadía.\n")
		break
	case "4":
		limpiarTerminal()
		fmt.Printf("Esto proximamnete va a hacer actualizar datos.\n")
		break
	default:
		limpiarTerminal()
		fmt.Printf("Opcion no existe en el menu\n")
		break
	}

}

func Salir(op bool) bool {
	fmt.Print("¿Desea hacer otro movimiento? (SI=0, NO=1): ")
	reader := bufio.NewReader(os.Stdin)

	// Leer una línea (hasta el carácter de nueva línea)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error al leer la entrada:", err)
		return op
	}

	// Eliminar espacios en blanco y caracteres de nueva línea
	input = strings.TrimSpace(input)

	if input == "1" {
		op = false
		return op
	} else {
		return op
	}
}

func limpiarTerminal() {
	var cmd *exec.Cmd

	sistemaOperativo := runtime.GOOS
	switch sistemaOperativo {
	case "darwin", "linux":
		cmd = exec.Command("clear")
	case "windows":
		cmd = exec.Command("cmd", "/c", "cls")
	default:
		fmt.Println("No se puede limpiar la terminal en este sistema operativo.")
		return
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
