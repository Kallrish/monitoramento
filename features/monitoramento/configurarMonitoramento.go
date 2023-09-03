package monitoramento

import (
	"fmt"
	"log"
	"time"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/utility"
	"github.com/Kallrish/monitoramento/vars"
)

func ConfiguracaoDoMonitoramento() {
	var opcaoDigitada int
	exibeConfigucaoDoMonitoramento()
	exibeIntroducaoConfiguracoes()
	opcaoDigitada, err := utility.RecebeEntradaDeNumeroInteiro()
	if err != nil {
		log.Println(err)
		fmt.Println(constants.OPCAO_INVALIDA_DIGITE)
	} else {
		verificaOpcaoEscolhidaConfiguracoes(opcaoDigitada)
	}
}

func exibeConfigucaoDoMonitoramento() {
	exibeConfiguracaoDelay()
	exibeConfiguracaoRepeticoes()
}

func exibeConfiguracaoDelay() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.DELAY_MONITORAMENTO)
	if vars.DelayEntreMonitoramentos > 1 {
		fmt.Println(int(vars.DelayEntreMonitoramentos), constants.SEGUNDOS)
	} else {
		fmt.Println(int(vars.DelayEntreMonitoramentos), constants.SEGUNDO)
	}
}

func exibeConfiguracaoRepeticoes() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.QUANTIDADE_REPETICOES_MONITORAMENTO)
	if vars.QuantidadeDeMonitoramentos > 1 {
		fmt.Println(vars.QuantidadeDeMonitoramentos, constants.VEZES)
	} else {
		fmt.Println(vars.QuantidadeDeMonitoramentos, constants.VEZ)
	}
}

func exibeIntroducaoConfiguracoes() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.DIGITE_OPCAO)
	fmt.Println(constants.ALTERAR_DELAY)
	fmt.Println(constants.ALTERAR_REPETICOES)
	fmt.Println(constants.RETORNAR_CONFIGURACOES)
}

func verificaOpcaoEscolhidaConfiguracoes(opcao int) {
	switch opcao {
	case 1:
		exibeOpcaoEscolhidaConfiguracoes(opcao)
		alteraDelay()
	case 2:
		exibeOpcaoEscolhidaConfiguracoes(opcao)
		alteraRepeticoes()
	case 3:
		exibeOpcaoEscolhidaConfiguracoes(opcao)
	default:
		exibeOpcaoEscolhidaConfiguracoes(opcao)
	}
}

func exibeOpcaoEscolhidaConfiguracoes(opcao int) {
	utility.GeraLinhaDeTextoVazio()
	switch opcao {
	case 1:
		fmt.Println(constants.ALTERAR_DELAY)
	case 2:
		fmt.Println(constants.ALTERAR_REPETICOES)
	case 3:
		fmt.Println(constants.RETORNAR_CONFIGURACOES)
	default:
		fmt.Println(constants.OPCAO_INVALIDA_RETORNA)
	}
}

func alteraDelay() {
	exibeIntroducaoAlteraDelay()
	novoDelay := recebeInteiroParaAlteracaoDaConfiguracao()
	if novoDelay <= 4 {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println(constants.VALOR_MINIMO_DELAY)
	} else if novoDelay == int(vars.DelayEntreMonitoramentos) {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println(constants.VALOR_REPETIDO_DELAY)
	} else {
		validaAlteracao := recebeConfirmacaoNovaConfiguracao()
		switch validaAlteracao {
		case 1:
			vars.DelayEntreMonitoramentos = time.Duration(novoDelay)
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.ALTERACAO_DELAY_SUCESSO)
			fmt.Println(constants.NOVO_VALOR_CONFIGURACAO, int(vars.DelayEntreMonitoramentos), constants.SEGUNDOS)
		case 2:
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.CANCELAR_RETORNAR_MENU_ANTERIOR)
		default:
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.OPCAO_INVALIDA_RETORNA)
		}
	}

}

func exibeIntroducaoAlteraDelay() {
	exibeConfiguracaoDelay()
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.DIGITE_VALOR_DELAY)
}

func alteraRepeticoes() {
	exibeIntroducaoAlteraRepeticoes()
	novaQuantidade := recebeInteiroParaAlteracaoDaConfiguracao()

	if vars.QuantidadeDeMonitoramentos < 1 {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println(constants.VALOR_MINIMO_QUANTIDADE)
	} else if novaQuantidade == vars.QuantidadeDeMonitoramentos {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println(constants.VALOR_REPETIDO_QUANTIDADE)
	} else {
		validaAlteracao := recebeConfirmacaoNovaConfiguracao()
		switch validaAlteracao {
		case 1:
			vars.QuantidadeDeMonitoramentos = novaQuantidade
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.ALTERACAO_QUANTIDADE_SUCESSO)
			if vars.QuantidadeDeMonitoramentos < 2 {
				fmt.Println(constants.NOVO_VALOR_CONFIGURACAO, vars.QuantidadeDeMonitoramentos, constants.VEZ)
			} else {
				fmt.Println(constants.NOVO_VALOR_CONFIGURACAO, vars.QuantidadeDeMonitoramentos, constants.VEZES)
			}
		case 2:
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.CANCELAR_RETORNAR_MENU_ANTERIOR)
		default:
			utility.GeraLinhaDeTextoVazio()
			fmt.Println(constants.OPCAO_INVALIDA_RETORNA)
		}
	}

}

func exibeIntroducaoAlteraRepeticoes() {
	exibeConfiguracaoRepeticoes()
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.DIGITE_QUANTIDADE_REPETICOES)
}

func recebeInteiroParaAlteracaoDaConfiguracao() int {
	var inteiroDigitado int
	fmt.Scanf("%d\n", &inteiroDigitado)

	return inteiroDigitado
}

func recebeConfirmacaoNovaConfiguracao() int {
	var opcaoDigitada int

	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.DIGITE_OPCAO)
	fmt.Println(constants.CONFIRMA_ALTERACAO_CONFIGURACAO)
	fmt.Println(constants.CANCELA_ALTERACAO_CONFIGURACAO)
	fmt.Scanf("%d\n", &opcaoDigitada)

	return opcaoDigitada
}
