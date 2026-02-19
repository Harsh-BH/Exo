package exo

import (
	"fmt"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
)

func generateK8s(cwd string, data config.TemplateData, dryRun, force bool) error {
	k8sDir := filepath.Join(cwd, "k8s")

	var stop func(error)
	if !dryRun {
		stop = startSpinner("Generating Kubernetes manifests → k8s/")
	}

	var genErr error
	for _, f := range []string{"deployment.yaml", "service.yaml", "ingress.yaml"} {
		tmpl := filepath.Join("templates", "k8s", f+".tmpl")
		out := filepath.Join(k8sDir, f)
		if err := renderFile(tmpl, out, data, dryRun, force); err != nil {
			genErr = fmt.Errorf("k8s/%s: %w", f, err)
			break
		}
	}

	if !dryRun {
		stop(genErr)
	}
	return genErr
}
