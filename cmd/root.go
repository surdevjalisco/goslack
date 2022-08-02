package cmd

// Implementación de comandos mediante cobra:
// https://github.com/spf13/cobra/blob/main/user_guide.md

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// comando root
var rootCmd = &cobra.Command{
	Use:   "goslack",
	Short: "Un simple programa que consume las API de Slack desde Golang.",
	Long: `Programa para consolas que muestra la integración con Slack desde una aplicación desarrollada con Go.
				  Ver mas en https://github.com/surdevjalisco/goslack`,
}

// La ejecución del comando principal añade la lista de sub comandos y la configuración al comando root
func Execute() {

	rootCmd.AddCommand(createUsersCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
