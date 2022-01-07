package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const monitoringQty = 3
const delay = 5

func main() {

	showIntroduction()

	for {
		showMenu()
		command := getCommand()

		switch command {
		case 1:
			startMonitoring()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func showIntroduction() {
	name := "Thiago"
	version := 1.0
	fmt.Println("Olá, sr.", name)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func getCommand() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O valor da variável comando é:", command)

	return command
}

func startMonitoring() {
	fmt.Println("Monitorando...")

	sites := readSitesFromFile()

	for i := 0; i < monitoringQty; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func readSitesFromFile() []string {
	var sites []string

	file, fileErr := os.Open("sites.txt")
	if fileErr != nil {
		fmt.Println("Ocorreu um erro:", fileErr)
	}

	reader := bufio.NewReader(file)

	for {
		line, lineErr := reader.ReadString('\n')
		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if lineErr == io.EOF {
			break
		}
	}

	file.Close()

	return sites
}

func testSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
	}

}
