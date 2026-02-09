package generator

import (
	"embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"
)

// Config хранит данные для генерации
type Config struct {
	ModuleName  string
	ProjectName string
	GoVersion   string
	UseDB       bool
	IsMono      bool
	TemplatesFS embed.FS // Вшиваем шаблоны в бинарник
}

func Generate(cfg Config) error {
	// 1. Создаем корневую папку проекта
	if err := os.MkdirAll(cfg.ProjectName, 0755); err != nil {
		return fmt.Errorf("failed to create root dir: %w", err)
	}

	// 2. Определяем пути (учитываем Monorepo)
	// Если Mono, то Go код летит в /backend, иначе в корень
	basePath := cfg.ProjectName
	if cfg.IsMono {
		basePath = filepath.Join(cfg.ProjectName, "backend")
	}

	// 3. Создаем структуру папок
	dirs := []string{
		filepath.Join(basePath, "cmd", "api"),
		filepath.Join(basePath, "internal", "config"),
		filepath.Join(basePath, "deployments"),
	}
	for _, d := range dirs {
		if err := os.MkdirAll(d, 0755); err != nil {
			return err
		}
	}

	// 4. Генерируем файлы из шаблонов
	files := map[string]string{
		"templates/go.mod.tmpl":     filepath.Join(basePath, "go.mod"),
		"templates/main.go.tmpl":    filepath.Join(basePath, "cmd", "api", "main.go"),
		"templates/Dockerfile.tmpl": filepath.Join(basePath, "deployments", "Dockerfile"),
		"templates/gitignore..tmpl": filepath.Join(basePath, ".gitignore"),
		// Docker-compose всегда в корне (даже при mono)
		"templates/docker-compose.yml.tmpl": filepath.Join(cfg.ProjectName, "docker-compose.yml"),
	}

	for tmplPath, targetPath := range files {
		if err := generateFile(tmplPath, targetPath, cfg); err != nil {
			return err
		}
	}

	// 5. Финализация (go mod tidy)
	fmt.Println("📦 Downloading dependencies...")
	cmd := exec.Command("go", "mod", "tidy")
	cmd.Dir = basePath
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Warning: 'go mod tidy' failed: %v\n", err)
	}

	return nil
}

func generateFile(tmplPath, targetPath string, data Config) error {
	// Читаем шаблон из embed FS
	// tmplPath содержит "templates/go.mod.tmpl", но в embed FS файлы находятся без префикса "templates/"
	// embed.FS в templates/templates.go использует //go:embed *, поэтому пути: go.mod.tmpl, main.go.tmpl и т.д.
	embedPath := strings.TrimPrefix(tmplPath, "templates/")

	// Пробуем разные варианты путей
	var content []byte
	var err error

	// Вариант 1: без префикса templates/ (правильный для embed)
	content, err = data.TemplatesFS.ReadFile(embedPath)
	if err != nil {
		// Вариант 2: с префиксом templates/ (на случай если структура другая)
		content, err = data.TemplatesFS.ReadFile(tmplPath)
		if err != nil {
			// Вариант 3: только имя файла
			content, err = data.TemplatesFS.ReadFile(filepath.Base(tmplPath))
			if err != nil {
				return fmt.Errorf("read template error %s (tried: %s, %s, %s): %w",
					tmplPath, embedPath, tmplPath, filepath.Base(tmplPath), err)
			}
		}
	}

	// Парсим
	tmpl, err := template.New(filepath.Base(tmplPath)).Parse(string(content))
	if err != nil {
		return fmt.Errorf("parse template error: %w", err)
	}

	// Создаем файл
	f, err := os.Create(targetPath)
	if err != nil {
		return fmt.Errorf("create file error: %w", err)
	}
	defer f.Close()

	// Исполняем
	return tmpl.Execute(f, data)
}
