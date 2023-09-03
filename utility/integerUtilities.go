package utility

import (
	"fmt"

	"github.com/Kallrish/monitoramento/vars"
)

func RecebeEntradaDeNumeroInteiro() (int, error) {
	_, err := fmt.Scanf("%d\n", &vars.OpcaoSelecionada)
	return vars.OpcaoSelecionada, err
}
