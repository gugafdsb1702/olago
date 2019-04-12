package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	numero := 3
	fmt.Print("O numero ", numero, " se escreve assim:")
	switch numero {
	case 1:
		fmt.Println("um.")
	case 2:
		fmt.Println("dois.")
	case 3:
		fmt.Println("três.")
	}

	fmt.Println("Você é da família do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Println("SIM!")
	default:
		fmt.Println("Deixa essa perrgunta para lá...")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("OK, você pode dormir até mais tarde.")
	default:
		fmt.Println("Vamos lá que hoje é dia de trabalhar!")
	}

	numero = 10
	fmt.Println("Este numero cabe num digito?")
	switch {
	case numero < 10:
		fmt.Println("SIM!")
	case numero >= 10:
		fmt.Println("Serve dois digitos...")
	case numero >= 100:
		fmt.Println("Não dá o numero é muito grande!")
	}

}
