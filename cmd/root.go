package cmd

import (
	"embed"

	"github.com/spf13/cobra"
)

var rootCmd *cobra.Command
var templatesFS embed.FS

// Init инициализирует команды и регистрирует их в rootCmd
func Init(r *cobra.Command, fs embed.FS) {
	rootCmd = r
	templatesFS = fs
	// Регистрируем все команды
	RegisterInit()
}

// getTemplatesFS возвращает embed FS с шаблонами
func getTemplatesFS() embed.FS {
	return templatesFS
}
