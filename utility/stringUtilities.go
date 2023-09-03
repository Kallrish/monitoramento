package utility

import (
	"fmt"

	"github.com/Kallrish/monitoramento/constants"
)

func GeraLinhaDeTextoVazio() {
	fmt.Println(constants.EMPTY_TEXT)
}

func ExibeErroOpcaoInvalida() {
	fmt.Println(constants.OPCAO_INVALIDA_DIGITE)
}
