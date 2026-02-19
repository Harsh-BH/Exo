package exo

import (
	"fmt"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
)

func generateHelm(cwd string, data config.TemplateData, dryRun, force bool) error {
	chartsDir := filepath.Join(cwd, "charts", data.AppName)
	tmplsDir := filepath.Join(chartsDir, "templates")

	files := []struct{ tmpl, out string }{
		{"helm/Chart.yaml.tmpl", filepath.Join(chartsDir, "Chart.yaml")},
		{"helm/values.yaml.tmpl", filepath.Join(chartsDir, "values.yaml")},
		{"helm/templates/deployment.yaml.tmpl", filepath.Join(tmplsDir, "deployment.yaml")},
		{"helm/templates/service.yaml.tmpl", filepath.Join(tmplsDir, "service.yaml")},
		{"helm/templates/ingress.yaml.tmpl", filepath.Join(tmplsDir, "ingress.yaml")},
	}

	var stop func(error)
	if !dryRun {
		stop = startSpinner(fmt.Sprintf("Generating Helm chart → charts/%s/", data.AppName))
	}

	var genErr error
	for _, f := range files {
		tmpl := filepath.Join("templates", f.tmpl)
		if err := renderFile(tmpl, f.out, data, dryRun, force); err != nil {
			genErr = fmt.Errorf("helm: %w", err)
			break
		}
	}

	if !dryRun {
		stop(genErr)
	}
	return genErr
}
