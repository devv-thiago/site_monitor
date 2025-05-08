package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const delayDuration int = 10
const reps int = 10

func showMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}

func getCommand() int {
	var comando int

	fmt.Scan(&comando) //Ponteiro

	return comando
}
func initMonitor() {
	services := []string{"https://www.alura.com.br/", "https://www.linkedin.com/", "https://gmail.com/"}
	for range reps {
		for i, service := range services {
			resp, _ := http.Get(service)
			if resp.StatusCode == 200 {
				fmt.Println(i, service, "is working!")
			} else {
				fmt.Println(service, "is not working! Please verify")
			}
		}
		fmt.Println()
		time.Sleep(time.Duration(delayDuration * time.Now().Second()))
	}

}

func main() {
	fmt.Println("Bem vindo ao monitorador de sites - Versão 1.0.0")

	for {
		showMenu()

		switch getCommand() {
		case (1):
			fmt.Println("Iniciando Monitoramento...")
			initMonitor()
		case (2):
			fmt.Println("Exibindo Logs...")
		case (3):
			fmt.Println("Saindo do Programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando.")
			os.Exit(-1)
		}
	}

}
