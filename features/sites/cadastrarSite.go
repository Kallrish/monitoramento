package sites

import (
	"fmt"
	"log"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/utility"
	"github.com/Kallrish/monitoramento/vars"
)

func ExibeOpcaoEscolhidaMenuCadastrarSites(opcao int) {
	utility.GeraLinhaDeTextoVazio()
	switch opcao {
	case 1:
		fmt.Println(constants.CADASTRAR_OUTRO_SITE)
	case 2:
		fmt.Println(constants.RETORNAR_MENU_ANTERIOR)
	default:
		fmt.Println(constants.OPCAO_INVALIDA_RETORNA)
	}
}

func VerificaOpcaoEscolhidaCadastrarSites(opcao int, err error) {
	if err != nil {
		log.Println(err)
		fmt.Println(constants.OPCAO_INVALIDA_DIGITE)
	} else {
		switch opcao {
		case 1:
			ExibeOpcaoEscolhidaMenuCadastrarSites(opcao)
			CadastrarSite()
		case 2:
			ExibeOpcaoEscolhidaMenuCadastrarSites(opcao)
		default:
			ExibeOpcaoEscolhidaMenuCadastrarSites(opcao)
		}
	}
}

func CadastrarSite() {
	ExibirListaDeSitesCadastrados(vars.ListaDeSites)
	exibeIntroducaoCadastroDeSites()

	//tratar tamanho do site(string)?
	cadastroSiteDigitado := recebeSiteDigitadoParaCadastro()
	confirmaCadastro, err := recebeConfirmacaoCadastroSiteDigitado()
	if err != nil {
		log.Println(err)
		fmt.Println(constants.OPCAO_INVALIDA_DIGITE)
	} else {
		switch confirmaCadastro {
		case 1:
			vars.ListaDeSites = append(vars.ListaDeSites, cadastroSiteDigitado)
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.SITE_CADASTRADO_SUCESSO)
		case 2:
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.CANCELAR_RETORNAR_MENU_ANTERIOR)
		default:
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.OPCAO_INVALIDA_RETORNA)
		}
	}
}

func exibeIntroducaoCadastroDeSites() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.DIGITE_SITE)
	fmt.Println(constants.EXEMPLO_SITE)
}

func recebeSiteDigitadoParaCadastro() string {
	var siteDigitado string
	fmt.Scanf("%s\n", &siteDigitado)
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.SITE_DIGITADO_CADASTRO, siteDigitado)

	return siteDigitado
}

func recebeConfirmacaoCadastroSiteDigitado() (int, error) {
	var opcaoDigitada int

	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.DIGITE_OPCAO)
	fmt.Println(constants.REALIZA_CADASTRO_SITE)
	fmt.Println(constants.CANCELAR_CADASTRO_SITE)
	_, err := fmt.Scanf("%d\n", &opcaoDigitada)

	return opcaoDigitada, err
}

// Ideia para futura implementacao
// func leSitesDeArquivo() []string {
// 	var sites []string
// 	arquivo, err := os.Open(PATH_DB_SITE)
// 	if err != nil {
// 		geraLinhaDeTextoVazio()
// 		fmt.Println("Erro ao abrir o arquivo:", err)
// 	}
// 	leitor := bufio.NewReader(arquivo)
// 	for {
// 		linha, erro := leitor.ReadString('\n')
// 		linha = strings.TrimSpace(linha)
// 		sites = append(sites, linha)

// 		if erro == io.EOF {
// 			break
// 		}
// 	}
// 	fmt.Println(sites)
// 	arquivo.Close()

// 	return sites
// }
