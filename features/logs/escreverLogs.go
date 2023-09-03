package logs

import (
	"fmt"
	"os"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/utility"
)

func EscreveLog(site string, status bool) {
	arquivo, err := os.OpenFile(constants.PATH_DB_LOG, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	data := utility.RetornaData()
	hora := utility.RetornaHora()
	var statusDoSite string

	if status == true {
		statusDoSite = constants.ONLINE
	} else {
		statusDoSite = constants.OFFLINE
	}

	if err != nil {
		utility.GeraLinhaDeTextoVazio()
		fmt.Println("Erro ao abrir o arquivo:", err)
	}

	arquivo.WriteString(
		constants.DATA + data + constants.HYPHEN_WITH_SPACES + constants.HORA +
			hora + constants.HYPHEN_WITH_SPACES + constants.SITE + site +
			constants.HYPHEN_WITH_SPACES + constants.STATUS + statusDoSite + "\n")
	arquivo.Close()
}
