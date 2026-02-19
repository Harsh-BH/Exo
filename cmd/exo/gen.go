package exo

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/Harsh-BH/Exo/internal/config"
	"github.com/Harsh-BH/Exo/internal/detector"
	"github.com/Harsh-BH/Exo/internal/renderer"
	"github.com/spf13/cobra"
)

// renderFile is a shared helper that renders a template to an output path,
// honouring --dry-run and --force flags.
func renderFile(tmplPath, outPath string, data interface{}, dryRun, force bool) error {
	if dryRun {
		fmt.Printf("  [dry-run] would write → %s\n", outPath)
		return nil
	}
	if !force {
		if _, err := os.Stat(outPath); err == nil {
			fmt.Printf("  ⚠ %s already exists (use --force to overwrite)\n", filepath.Base(outPath))
			return nil
		}
	}
	return renderer.RenderTemplate(tmplPath, outPath, data)
}

// loadTemplateData builds a TemplateData from flags, falling back to .exo.yaml,
// then to the auto-detector.
func loadTemplateData(cmd *cobra.Command, cwd string) config.TemplateData {
	// Try loading persisted config first
	var base config.TemplateData
	if cfg, err := config.Load(cwd); err == nil {
		base = cfg.ToTemplateData()
	} else {
		// Auto-detect language/framework if no config
		if info, err := detector.Detect(cwd); err == nil {
			base.Language = info.Language
			base.Framework = info.Framework
		}
	}
	if base.Port == 0 {
		base.Port = 8080
	}

	// Flag overrides beat .exo.yaml
	if cmd.Flags().Changed("name") {
		v, _ := cmd.Flags().GetString("name")
		base.AppName = v
	}
	if base.AppName == "" {
		base.AppName = filepath.Base(cwd)
	}
	if cmd.Flags().Changed("lang") {
		v, _ := cmd.Flags().GetString("lang")
		base.Language = v
	}
	if cmd.Flags().Changed("provider") {
		v, _ := cmd.Flags().GetString("provider")
		base.Provider = v
	}
	if cmd.Flags().Changed("db") {
		v, _ := cmd.Flags().GetString("db")
		base.DB = v
	}
	if cmd.Flags().Changed("monitoring") {
		v, _ := cmd.Flags().GetString("monitoring")
		base.Monitoring = v
	}
	return base
}

var genCmd = &cobra.Command{
	Use:   "gen [type]",
	Short: "Generate a DevOps asset",
	Long: `Generate a specific DevOps asset for your project.

Types:
  docker          Dockerfile (language-aware)
  infra           Terraform infrastructure
  k8s             Kubernetes manifests
  helm            Helm chart
  ci              CI/CD pipeline
  db              Database docker-compose
  makefile        Makefile
  env             .env.example
  gitignore       .gitignore
  grafana         Grafana dashboard JSON
  alerts          Prometheus alert rules
  docker-compose  Full docker-compose.yml (app + db + monitoring)
  readme          README.md
  pre-commit      .pre-commit-config.yaml
  devcontainer    .devcontainer/devcontainer.json
  renovate        renovate.json
  license         LICENSE file
  dependabot      .github/dependabot.yml
  sonarqube       sonar-project.properties
  sbom            Software Bill of Materials (CycloneDX JSON, uses syft if available)`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		genType := args[0]
		cwd, err := os.Getwd()
		if err != nil {
			return err
		}

		// --output-dir overrides the destination directory
		if outDir, _ := cmd.Flags().GetString("output-dir"); outDir != "" {
			if !filepath.IsAbs(outDir) {
				outDir = filepath.Join(cwd, outDir)
			}
			if err := os.MkdirAll(outDir, 0o755); err != nil {
				return fmt.Errorf("creating output dir: %w", err)
			}
			cwd = outDir
		}

		dryRun, _ := cmd.Flags().GetBool("dry-run")
		force, _ := cmd.Flags().GetBool("force")
		data := loadTemplateData(cmd, cwd)

		if dryRun {
			fmt.Println("  [dry-run mode — no files will be written]")
		}

		switch genType {
		case "docker":
			return generateDockerfile(cwd, data, dryRun, force)
		case "infra":
			return generateInfra(cwd, data, dryRun, force)
		case "k8s":
			return generateK8s(cwd, data, dryRun, force)
		case "helm":
			return generateHelm(cwd, data, dryRun, force)
		case "ci":
			return generateCI(cwd, data, dryRun, force)
		case "db":
			return generateDB(cwd, data, dryRun, force)
		case "makefile":
			return generateMakefile(cwd, data, dryRun, force)
		case "env":
			return generateEnv(cwd, data, dryRun, force)
		case "gitignore":
			return generateGitignore(cwd, data, dryRun, force)
		case "grafana":
			return generateGrafana(cwd, data, dryRun, force)
		case "alerts":
			return generateAlerts(cwd, data, dryRun, force)
		case "docker-compose":
			return generateDockerCompose(cwd, data, dryRun, force)
		case "readme":
			return generateReadme(cwd, data, dryRun, force)
		case "pre-commit":
			return generatePreCommit(cwd, data, dryRun, force)
		case "devcontainer":
			return generateDevcontainer(cwd, data, dryRun, force)
		case "renovate":
			return generateRenovate(cwd, data, dryRun, force)
		case "license":
			licenseType, _ := cmd.Flags().GetString("license-type")
			return generateLicense(cwd, data, licenseType, dryRun, force)
		case "dependabot":
			return generateDependabot(cwd, data, dryRun, force)
		case "sonarqube":
			return generateSonarqube(cwd, data, dryRun, force)
		case "sbom":
			return generateSBOM(cwd, data, dryRun, force)
		default:
			return fmt.Errorf("unknown generation type: %s\n\nRun 'exo gen --help' for available types", genType)
		}
	},
}

func init() {
	rootCmd.AddCommand(genCmd)
	genCmd.Flags().StringP("name", "n", "", "Application name (defaults to .exo.yaml or directory name)")
	genCmd.Flags().StringP("lang", "l", "", "Language override (go, node, python, java, rust)")
	genCmd.Flags().StringP("provider", "p", "", "Cloud provider override (aws, gcp, azure)")
	genCmd.Flags().String("db", "", "Database override (postgres, mysql, mongo, redis)")
	genCmd.Flags().String("monitoring", "", "Monitoring override (prometheus, none)")
	genCmd.Flags().Bool("dry-run", false, "Preview what would be generated without writing files")
	genCmd.Flags().Bool("force", false, "Overwrite existing files without prompting")
	genCmd.Flags().String("license-type", "mit", "License type for 'exo gen license' (mit, apache2, gpl3)")
	genCmd.Flags().StringP("output-dir", "o", "", "Write generated files into this directory instead of the current directory")
}
