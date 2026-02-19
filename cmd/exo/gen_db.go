package exo

import (
	"fmt"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
)

func generateDB(cwd string, data config.TemplateData, dryRun, force bool) error {
	db := data.DB
	if db == "" || db == "none" {
		return fmt.Errorf("no database set; use --db postgres|mysql|mongo|redis or run 'exo init' first")
	}
	tmplMap := map[string]string{
		"postgres": "postgres.tmpl",
		"mysql":    "mysql.tmpl",
		"mongo":    "mongo.tmpl",
		"redis":    "redis.tmpl",
	}
	tmplFile, ok := tmplMap[db]
	if !ok {
		return fmt.Errorf("unknown database %q (postgres, mysql, mongo, redis)", db)
	}

	tmpl := filepath.Join("templates", "db", tmplFile)
	out := filepath.Join(cwd, fmt.Sprintf("docker-compose.%s.yml", db))
	if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
		return fmt.Errorf("db: %w", err)
	}
	if !dryRun {
		fmt.Printf("  ✓  %s → docker-compose.%s.yml\n", db, db)
	}
	return nil
}
