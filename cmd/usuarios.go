package cmd

// Implementación de comandos mediante cobra:
// https://github.com/spf13/cobra/blob/main/user_guide.md

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/surdevjalisco/goslack/pkg/goslack"
)

func getEmails(cmd *cobra.Command, args []string) error {
	emails, err := goslack.GetUsersEmails()
	if err != nil {
		return fmt.Errorf("Obteniendo emails: %w", err)
	}

	for _, email := range emails {
		fmt.Println(email)
	}

	return nil
}

func createUsersEmailsCommand() *cobra.Command {
	uEmailCommand := &cobra.Command{
		Use:   "emails",
		Short: "Obtiene la lista de correos electrónicos de los usuarios.",
		RunE:  getEmails,
	}
	return uEmailCommand
}

func createUsersCommand() *cobra.Command {
	uCommand := &cobra.Command{
		Use:   "usuarios",
		Short: "Operaciones sobre la lista de usuarios en Slack",
	}

	uCommand.AddCommand(createUsersEmailsCommand())

	return uCommand
}
