package main

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
)

var (
	dbType    string
	cacheType string
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "api-generator",
		Short: "Generate boilerplate code for Mux-based API",
		Run:   generateCode,
	}

	rootCmd.Flags().StringVarP(&dbType, "db", "d", "mysql", "Database type (mysql, postgresql, etc)")
	rootCmd.Flags().StringVarP(&cacheType, "cache", "c", "redis", "Cache type (redis, memcached, etc)")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func generateCode(cmd *cobra.Command, args []string) {
	// Create directory structure
	directories := []string{
		"cmd/api",
		"DBLayer",
	}

	for _, dir := range directories {
		if err := os.MkdirAll(dir, 0755); err != nil {
			fmt.Printf("Error creating directory %s: %v\n", dir, err)
			return
		}
	}

	// Generate files
	if err := generateFiles(); err != nil {
		fmt.Println("Error generating files:", err)
		return
	}

	// Generate Dockerfiles
	if err := generateDockerfiles(); err != nil {
		fmt.Println("Error generating Dockerfiles:", err)
		return
	}

	fmt.Println("Boilerplate code generated successfully!")
}

func generateFiles() error {
	// Define file templates
	fileTemplates := map[string]string{
		"cmd/api/main.go":       mainTemplate,
		"cmd/api/router.go":     routerTemplate,
		"cmd/api/get-api.go":    getAPITemplate,
		"cmd/api/post-api.go":   postAPITemplate,
		"DBLayer/connection.go": connectionTemplate,
		"DBLayer/dbutils.go":    dbUtilsTemplate,
	}

	for file, templateStr := range fileTemplates {
		if err := generateFile(file, templateStr); err != nil {
			return err
		}
	}

	return nil
}

func generateFile(filePath, templateStr string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	tmpl, err := template.New(filepath.Base(filePath)).Parse(templateStr)
	if err != nil {
		return err
	}

	if err := tmpl.Execute(file, nil); err != nil {
		return err
	}

	return nil
}

func generateDockerfiles() error {
	// Define Dockerfile templates
	dockerfileTemplates := map[string]string{
		"Dockerfile.api": dockerfileAPITemplate,
		"Dockerfile.db":  dockerfileDBTemplate,
	}

	for file, templateStr := range dockerfileTemplates {
		if err := generateFile(file, templateStr); err != nil {
			return err
		}
	}

	return nil
}

// Define templates
const (
	mainTemplate = `package main

import (
	"log"
	"net/http"
)

func main() {
	r := setupRouter()
	log.Fatal(http.ListenAndServe(":8080", r))
}`

	routerTemplate = `package main

import (
	"github.com/gorilla/mux"
)

func setupRouter() *mux.Router {
	r := mux.NewRouter()

	// Define routes
	// Example: r.HandleFunc("/api/resource", getResourceHandler).Methods("GET")

	return r
}`

	getAPITemplate = `package main

import (
	"net/http"
)

func getResourceHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to handle GET request for resource
}`

	postAPITemplate = `package main

import (
	"net/http"
)

func createResourceHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logic to handle POST request to create resource
}`

	connectionTemplate = `package DBLayer

import (
	"database/sql"
)

func connectDB() (*sql.DB, error) {
	// Implement logic to connect to {{.DBType}} database
	return nil, nil
}`

	dbUtilsTemplate = `package DBLayer

import (
	"database/sql"
)

func execQuery(db *sql.DB, query string) error {
	// Implement logic to execute SQL query
	return nil
}`
	dockerfileAPITemplate = `# Dockerfile for API
# Add necessary instructions here`

	dockerfileDBTemplate = `# Dockerfile for DB
# Add necessary instructions here`
)
