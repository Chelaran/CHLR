package main

import (
	"fmt"
	"os"

	"github.com/Chelaran/CHLR/cmd"
	"github.com/Chelaran/CHLR/templates"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "chlr",
	Short: "Chelaran CLI - Scaffolding tool for Full-Cycle projects",
	Long: `CHLR (Chelaran CLI) — инструмент скаффолдинга и автоматизации разработки 
Full-Cycle проектов. Генерирует production-ready архитектуру, соответствующую 
инженерным стандартам агентства Chelaran.`,
	Version: "0.1.2",
}

func main() {
	// Инициализируем команды
	cmd.Init(rootCmd, templates.TemplatesFS)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "❌ Error: %v\n", err)
		os.Exit(1)
	}
}
