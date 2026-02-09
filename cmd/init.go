package cmd

import (
	"fmt"
	"log"
	"runtime"
	"strings"

	"github.com/bambutcha/chlr/internal/generator"
	"github.com/spf13/cobra"
)

var (
	isMono bool
	dbType string
)

var initCmd = &cobra.Command{
	Use:   "init <project-name>",
	Short: "Initialize a new project (Chelaran Standard)",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]
		// Имя модуля = имя папки (для простоты MVP)
		moduleName := projectName

		// Автодетект версии Go (отрезаем "go" в начале, например "go1.22.1" -> "1.22.1")
		goVer := strings.TrimPrefix(runtime.Version(), "go")

		cfg := generator.Config{
			ProjectName: projectName,
			ModuleName:  moduleName,
			GoVersion:   goVer,
			IsMono:      isMono,
			UseDB:       dbType == "postgres",
			TemplatesFS: getTemplatesFS(),
		}

		fmt.Printf("🚀 Initializing project '%s'...\n", projectName)
		fmt.Printf("⚙️  Stack: Go %s | DB: %s | Mono: %v\n", goVer, dbType, isMono)

		if err := generator.Generate(cfg); err != nil {
			log.Fatalf("❌ Error: %v", err)
		}

		fmt.Println("✅ Done! Happy coding.")
		if isMono {
			fmt.Printf("👉 cd %s/backend && go run cmd/api/main.go\n", projectName)
		} else {
			fmt.Printf("👉 cd %s && go run cmd/api/main.go\n", projectName)
		}
	},
}

func init() {
	// Регистрируем флаги
	initCmd.Flags().BoolVar(&isMono, "mono", false, "Enable monorepo structure")
	initCmd.Flags().StringVar(&dbType, "db", "none", "Database type (postgres, none)")
}

// RegisterInit регистрирует команду init в rootCmd
func RegisterInit() {
	if rootCmd != nil {
		rootCmd.AddCommand(initCmd)
	}
}
