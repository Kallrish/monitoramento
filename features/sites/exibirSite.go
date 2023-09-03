package sites

import (
	"fmt"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/utility"
	"github.com/Kallrish/monitoramento/vars"
)

func ExibirListaDeSitesCadastrados(lista []string) {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.LISTA_SITES_CADASTRADOS)

	for index, site := range lista {
		fmt.Println(index, constants.HYPHEN_WITH_SPACES, site)
	}
}

func InicializaListaDeSites() {
	vars.ListaDeSites = append(
		vars.ListaDeSites,
		"https://www.google.com.br",
		"https://www.mercadolivre.com.br",
		"https://www.alelo.com.br")
}
