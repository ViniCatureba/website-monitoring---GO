package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	//"io/ioutil"
	"bufio"
)

const monitoramentos = 3
const delay = 10

func main() {

	exibeIntroducao()
	for {
		exibeMenu()

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Ola, sr.", nome)
	fmt.Println("Esse programa esta na versao", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Mnitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitoriando...")

	/* var sites [4]string
	sites[0] = "https://random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br" */ //arrays

	//slices
	/* sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br", "https://www.caelum.com.br"} */

	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		fmt.Printf("--------------------------------------------------Monitoramento(%v)--------------------------------------------------\n\n1", i)
		for i, site := range sites {
			fmt.Print("Testando site: ", i, ": ", site, "\n\n")
			testaSite(site)
		}
		time.Sleep(monitoramentos * time.Second)

	}
}

func testaSite(site string) {
	resposta, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resposta.StatusCode == 200 {
		fmt.Print("Site: ", site, " foi carregado com sucesso\n\n")
	} else {
		fmt.Print("Site: ", site, " esta com problemas. Status:", resposta.StatusCode, "\n\n")
	}
}

func lerSitesDoArquivo() []string {
	arquivos, err := os.Open("sites.txt") //imprime local na memoria
	/*arquivo, err := ioutil.ReadFile("site.txt") imprime tudo em bytes */
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivos)
	lina, err := leitor.ReadString('\n')

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	return
}
