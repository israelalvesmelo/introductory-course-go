package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	exibeIntroducao()
	for {
		exibeMenu()
		command := leComando()

		switch command {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
		case 0:
			saiDoPrograma()
		default:
			fmt.Println("Não conheço este comando !")
		}
	}
}

func exibeIntroducao() {
	name := "Israel" // ou var name string = """
	var version float32 = 1.1

	fmt.Println("Olá , sr.", name)
	fmt.Println("Este programa está na versão ", version)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do Programa")
}

func leComando() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("")

	return command
}

func saiDoPrograma() {
	fmt.Println("Saindo do Programa...")
	os.Exit(0)
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	// sites := []string{"https://www.alura.com.br", "https://www.google.com/", "https://www.caelum.com.br"}
	sites := leSitesDoArquivo()

	for j := 0; j <= monitoramentos; j++ {
		for i, site := range sites {
			fmt.Println("Testando site: ", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")

	}

	fmt.Println("")
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site: ", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site: ", site, "esta com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		if err == io.EOF {
			break
		}

		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()

}

func imprimeLogs() {
	fmt.Println("Exibindo Logs...")
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
}
