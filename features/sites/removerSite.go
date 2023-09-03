package sites

import (
	"fmt"
	"log"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/utility"
	"github.com/Kallrish/monitoramento/vars"
)

func RemoverSites() {
	utility.GeraLinhaDeTextoVazio()
	ExibirListaDeSitesCadastrados(vars.ListaDeSites)
	exibeIntroducaoRemocaoDeSites()
	indiceSite, err_one := recebeIndiceSiteParaRemocao()
	confirmaRemocao, err_two := recebeConfirmacaoRemocaoSiteDigitado(indiceSite)

	if err_one != nil {
		log.Println(err_one)
		fmt.Println(constants.OPCAO_INVALIDA_DIGITE)
	} else {
		if err_two != nil {
			log.Println(err_two)
			fmt.Println(constants.OPCAO_INVALIDA_DIGITE)
		} else if indiceSite <= len(vars.ListaDeSites)-constants.NUMBER_ONE_INT {
			switch confirmaRemocao {
			case 1:
				vars.ListaDeSites = removeElementoDoSlice(vars.ListaDeSites, indiceSite)
				utility.GeraLinhaDeTextoVazio()
				fmt.Println(constants.SITE_REMOVIDO_SUCESSO)
			case 2:
				utility.GeraLinhaDeTextoVazio()
				fmt.Println(constants.CANCELAR_RETORNAR_MENU_ANTERIOR)
			default:
				utility.GeraLinhaDeTextoVazio()
				fmt.Println(constants.OPCAO_INVALIDA_RETORNA)
			}
		} else {
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.CANCELAR_RETORNAR_MENU_ANTERIOR)
		}
	}
}

func removeElementoDoSlice(listaDeSites []string, index int) []string {
	return append(listaDeSites[:index], listaDeSites[index+1:]...)
}

func exibeIntroducaoRemocaoDeSites() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.DIGITE_INDICE_REMOCAO)
}

func recebeIndiceSiteParaRemocao() (int, error) {
	var indiceDigitado int
	_, err := fmt.Scanf("%d\n", &indiceDigitado)
	if indiceDigitado <= len(vars.ListaDeSites)-constants.NUMBER_ONE_INT {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println(constants.INDICE_DIGITADO_REMOCAO, vars.ListaDeSites[indiceDigitado])
	} else {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println(constants.INDICE_DIGITADO_REMOCAO, constants.STATUS_DESCONHECIDO)
	}
	return indiceDigitado, err
}

func recebeConfirmacaoRemocaoSiteDigitado(opcao int) (int, error) {
	var opcaoDigitada int
	var err error
	if opcao <= len(vars.ListaDeSites)-constants.NUMBER_ONE_INT {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println(constants.DIGITE_OPCAO)
		fmt.Println(constants.REALIZA_REMOCAO_SITE)
		fmt.Println(constants.CANCELAR_REMOCAO_SITE)
		_, err = fmt.Scanf("%d\n", &opcaoDigitada)
	} else {
		utility.GeraLinhaDeTextoVazio()
	}
	return opcaoDigitada, err
}
