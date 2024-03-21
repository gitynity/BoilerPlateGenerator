package cmd

import (
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/*
var embeddedTemplates embed.FS

func generateFiles() {
	// Define file templates
	files := map[string]string{
		"cmd/api/main.go":       "main.go.tmpl",
		"cmd/api/router.go":     "router.go.tmpl",
		"cmd/api/get-api.go":    "get-api.go.tmpl",
		"cmd/api/post-api.go":   "post-api.go.tmpl",
		"DBLayer/connection.go": "connection.go.tmpl",
		"DBLayer/dbutils.go":    "dbutils.go.tmpl",
		"Dockerfile.api":        "Dockerfile.api.tmpl",
		"Dockerfile.db":         "Dockerfile.db.tmpl",
	}

	for file, templateFile := range files {
		templatePath := filepath.Join("templates", templateFile)
		content, err := embeddedTemplates.ReadFile(templatePath)
		if err != nil {
			fmt.Println("Error reading template:", err)
			return
		}

		tmpl, err := template.New(filepath.Base(templatePath)).Parse(string(content))
		if err != nil {
			fmt.Println("Error parsing template:", err)
			return
		}

		if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
			fmt.Println("Error creating directory:", err)
			return
		}

		f, err := os.Create(file)
		if err != nil {
			fmt.Println("Error creating file:", err)
			return
		}
		defer f.Close()

		if err := tmpl.Execute(f, nil); err != nil {
			fmt.Println("Error executing template:", err)
			return
		}
	}
}
