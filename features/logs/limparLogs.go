package logs

import (
	"fmt"
	"os"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/utility"
)

func LimparLogs() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.ELIMINANDO_LOGS)
	limpaLog(constants.PATH_DB_LOG)
}

func limpaLog(path string) {
	err := os.Remove(path)
	if err != nil {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println("Erro ao limpar aquivo de log: ", err)
	} else {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println("Limpeza de logs concluída com sucesso!")
	}
}
