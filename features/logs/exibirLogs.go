package logs

import (
	"fmt"
	"os"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/utility"
)

func ExibirLogs() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.LOGS_CARREGADOS)
	exibirLog(constants.PATH_DB_LOG)
}

func exibirLog(path string) {
	arquivo, err := os.ReadFile(path)
	if err != nil {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println("Erro ao ler aquivo de log: ", err)
	}
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(string(arquivo))
}
