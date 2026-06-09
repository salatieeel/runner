package cmd

import (
	"fmt"
	"os"
)

// Execute despacha subcomandos ou abre o menu interativo.
func Execute() {
	args := os.Args[1:]
	if len(args) == 0 {
		runInteractiveMenu()
		return
	}

	switch args[0] {
	case "menu":
		runInteractiveMenu()
	case "sign":
		if err := runSign(args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
			os.Exit(1)
		}
	case "validate":
		if err := runValidate(args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
			os.Exit(1)
		}
	case "version":
		fmt.Println("assinatura versão 1.0.0")
	case "build":
		if err := runBuild(args[1:]); err != nil {
			fmt.Fprintf(os.Stderr, "Erro: %v\n", err)
			os.Exit(1)
		}
	default:
		fmt.Fprintf(os.Stderr, "Comando desconhecido: %s\n", args[0])
		fmt.Fprintln(os.Stderr, "Use sem argumentos para o menu ou: sign, validate, version, build")
		os.Exit(1)
	}
}
