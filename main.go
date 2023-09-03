package main

import (
	"github.com/Kallrish/monitoramento/features/menuPrincipal"
	"github.com/Kallrish/monitoramento/features/sites"
	"github.com/Kallrish/monitoramento/utility"
)

func main() {
	sites.InicializaListaDeSites()
	menuPrincipal.ExibeIntroducao()
	for {
		menuPrincipal.ExibeMenuPrincipal()
		opcaoSelecionadaMenuPrincipal, err := utility.RecebeEntradaDeNumeroInteiro()
		menuPrincipal.VerificaOpcaoEscolhidaMenuPrincipal(opcaoSelecionadaMenuPrincipal, err)
	}
}
