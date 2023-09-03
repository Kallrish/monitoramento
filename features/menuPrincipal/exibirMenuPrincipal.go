package menuPrincipal

import (
	"fmt"
	"log"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/features/logs"
	"github.com/Kallrish/monitoramento/features/monitoramento"
	"github.com/Kallrish/monitoramento/features/sessao"
	"github.com/Kallrish/monitoramento/features/sites"
	"github.com/Kallrish/monitoramento/utility"
	"github.com/Kallrish/monitoramento/vars"
)

func ExibeIntroducao() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.BEM_VINDO)
	fmt.Println(constants.VERSAO, constants.VERSAO_APP)
}

func ExibeMenuPrincipal() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.MENU_PRINCIPAL)
	fmt.Println(constants.INICIAR_MONITORAMENTO)
	fmt.Println(constants.EXIBIR_LOGS)
	fmt.Println(constants.LIMPAR_LOGS)
	fmt.Println(constants.EXIBIR_LISTA_SITES)
	fmt.Println(constants.CADASTRAR_SITE)
	fmt.Println(constants.REMOVER_SITE)
	fmt.Println(constants.CONFIGURACOES_MONITORAMENTO)
	fmt.Println(constants.SAIR_PROGRAMA)
	utility.GeraLinhaDeTextoVazio()
}

func ExibeOpcaoEscolhidaMenuPrincipal(opcao int) {
	utility.GeraLinhaDeTextoVazio()
	switch opcao {
	case 1:
		fmt.Println(constants.INICIAR_MONITORAMENTO)
	case 2:
		fmt.Println(constants.EXIBIR_LOGS)
	case 3:
		fmt.Println(constants.LIMPAR_LOGS)
	case 4:
		fmt.Println(constants.EXIBIR_LISTA_SITES)
	case 5:
		fmt.Println(constants.CADASTRAR_SITE)
	case 6:
		fmt.Println(constants.REMOVER_SITE)
	case 7:
		fmt.Println(constants.CONFIGURACOES_MONITORAMENTO)
	case 8:
		fmt.Println(constants.SAIR_PROGRAMA)
	}
}

func VerificaOpcaoEscolhidaMenuPrincipal(opcao int, err error) {
	if err != nil {
		log.Println(err)
		fmt.Println(constants.OPCAO_INVALIDA_DIGITE)
	} else {
		switch opcao {
		case 1:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			monitoramento.IniciarMonitoramento()
		case 2:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			logs.ExibirLogs()
		case 3:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			logs.LimparLogs()
		case 4:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			sites.ExibirListaDeSitesCadastrados(vars.ListaDeSites)
		case 5:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			sites.CadastrarSite()
		case 6:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			sites.RemoverSites()
		case 7:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			monitoramento.ConfiguracaoDoMonitoramento()
		case 8:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			sessao.SairDoPrograma()
		default:
			ExibeOpcaoEscolhidaMenuPrincipal(opcao)
			utility.ExibeErroOpcaoInvalida()
		}
	}
}
