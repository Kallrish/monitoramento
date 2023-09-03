package sessao

import (
	"fmt"
	"os"

	"github.com/Kallrish/monitoramento/constants"
	"github.com/Kallrish/monitoramento/utility"
)

func SairDoPrograma() {
	utility.GeraLinhaDeTextoVazio()
	fmt.Println(constants.SAINDO_PROGRAMA)
	fmt.Println(constants.DESPEDIDA)
	utility.GeraLinhaDeTextoVazio()
	os.Exit(0)
}
