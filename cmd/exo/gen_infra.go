package exo

import (
	"fmt"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
)

func generateInfra(cwd string, data config.TemplateData, dryRun, force bool) error {
	prov := data.Provider
	if prov == "" || prov == "none" {
		return fmt.Errorf("no cloud provider set; use --provider aws|gcp|azure or run 'exo init' first")
	}
	validProviders := map[string]bool{"aws": true, "gcp": true, "azure": true}
	if !validProviders[prov] {
		return fmt.Errorf("unsupported provider %q (aws, gcp, azure)", prov)
	}

	infraDir := filepath.Join(cwd, "infra", prov)

	var stop func(error)
	if !dryRun {
		stop = startSpinner(fmt.Sprintf("Generating Terraform (%s) → infra/%s/", prov, prov))
	}

	var genErr error
	for _, f := range []string{"main.tf", "variables.tf", "provider.tf"} {
		tmpl := filepath.Join("templates", "terraform", prov, f+".tmpl")
		out := filepath.Join(infraDir, f)
		if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
			genErr = fmt.Errorf("infra/%s/%s: %w", prov, f, err)
			break
		}
	}

	if !dryRun {
		stop(genErr)
	}
	return genErr
}
