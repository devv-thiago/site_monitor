package main

import (
	"fmt";
	"os"
)

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func getCommand() int{
	var comando int

	fmt.Scan(&comando) //Ponteiro

	return comando
}

func main() {
	fmt.Println("Bem vindo ao monitorador de sites - Versão 1.0.0")

	showMenu()

	switch getCommand() {
	case (1):
		fmt.Println("Iniciando Monitoramento...")
	case (2):
		fmt.Println("Exibindo Logs...")
	case (3):
		fmt.Println("Saindo do Programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço esse comando.")
	}

}
