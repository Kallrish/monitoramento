package monitoramento

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/features/logs"
	"github.com/Kallrish/monitoramento/utility"
	"github.com/Kallrish/monitoramento/vars"
)

func IniciarMonitoramento() {
	listaDeResponses := buscaResponseDosSites(vars.ListaDeSites)
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.MONITORAMENTO_INICIADO)
	utility.GeraLinhaDeTextoVazio()
	fmt.Print(constants.TAMANHO_LISTA_SITES, len(vars.ListaDeSites), constants.PIPE_WITH_SPACES)
	fmt.Println(constants.TAMANHO_LISTA_RESPONSES, len(listaDeResponses))
	exibeStatusDoSite(
		vars.ListaDeSites,
		listaDeResponses,
		vars.QuantidadeDeMonitoramentos,
		vars.DelayEntreMonitoramentos)
}

func buscaResponseDosSites(listaSites []string) []*http.Response {
	var listaDeResponses []*http.Response
	for index := range listaSites {
		res, _ := http.Get(listaSites[index])
		listaDeResponses = append(listaDeResponses, res)
	}

	return listaDeResponses
}

func exibeStatusDoSite(
	lista []string,
	listaDeResponses []*http.Response,
	quantidadeRepeticoes int,
	delay time.Duration) {
	for i := 1; i <= quantidadeRepeticoes; i++ {
		for index, response := range listaDeResponses {
			if response != nil {
				if response.StatusCode == 200 {
					utility.GeraLinhaDeTextoVazio()
					fmt.Println(constants.SITE, index, constants.COLON, lista[index])
					fmt.Println(constants.HTTP_STATUS, response.StatusCode)
					fmt.Println(constants.SITE_SUCESSO)
					logs.EscreveLog(lista[index], true)
				} else {
					utility.GeraLinhaDeTextoVazio()
					fmt.Println(constants.SITE, index, constants.COLON, lista[index])
					fmt.Println(constants.HTTP_STATUS, response.StatusCode)
					fmt.Println(constants.SITE_FALHA)
					logs.EscreveLog(lista[index], false)
				}
			} else {
				utility.GeraLinhaDeTextoVazio()
				fmt.Println(constants.SITE, index, constants.COLON, lista[index])
				fmt.Println(constants.HTTP_STATUS, constants.STATUS_DESCONHECIDO)
				fmt.Println(constants.SITE_FALHA)
				logs.EscreveLog(lista[index], false)
			}
		}
		if quantidadeRepeticoes > constants.NUMBER_ONE_INT {
			defineTempoEntreMonitoramentos(delay)
		}

	}
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.MONITORAMENTO_FINALIZADO)
}

func defineTempoEntreMonitoramentos(d time.Duration) {
	time.Sleep(d * time.Second)
}
