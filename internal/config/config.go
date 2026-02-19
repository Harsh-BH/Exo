package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const ConfigFileName = ".exo.yaml"

// ExoConfig mirrors the wizard's ProjectData for persistence.
type ExoConfig struct {
	Name       string `yaml:"name"`
	Language   string `yaml:"language"`
	Framework  string `yaml:"framework,omitempty"`
	Provider   string `yaml:"provider"`
	CI         string `yaml:"ci"`
	Monitoring string `yaml:"monitoring"`
	DB         string `yaml:"db,omitempty"`
	Port       int    `yaml:"port,omitempty"`
	Registry   string `yaml:"registry,omitempty"`
}

// Save writes the config to .exo.yaml in the given directory.
func Save(dir string, cfg *ExoConfig) error {
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}
	path := filepath.Join(dir, ConfigFileName)
	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config: %w", err)
	}
	return nil
}

// Load reads .exo.yaml from the given directory.
func Load(dir string) (*ExoConfig, error) {
	path := filepath.Join(dir, ConfigFileName)
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("no .exo.yaml found: %w", err)
	}
	var cfg ExoConfig
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("failed to parse .exo.yaml: %w", err)
	}
	return &cfg, nil
}

// Exists returns true if .exo.yaml exists in the given directory.
func Exists(dir string) bool {
	_, err := os.Stat(filepath.Join(dir, ConfigFileName))
	return err == nil
}

// TemplateData is the unified data structure passed to every template.
// All fields are optional — templates should use {{if .Field}} guards.
type TemplateData struct {
	AppName    string
	Language   string // go | node | python | java | rust
	Framework  string // gin | echo | express | fastapi | …
	Port       int    // default 8080
	DB         string // postgres | mysql | mongo | redis | none
	Provider   string // aws | gcp | azure | none
	CI         string // github-actions | gitlab-ci | none
	Monitoring string // prometheus | none
	Registry   string // docker registry URL, optional
}

// ToTemplateData converts a saved ExoConfig into a TemplateData ready for rendering.
func (c *ExoConfig) ToTemplateData() TemplateData {
	port := c.Port
	if port == 0 {
		port = 8080
	}
	return TemplateData{
		AppName:    c.Name,
		Language:   c.Language,
		Framework:  c.Framework,
		Port:       port,
		DB:         c.DB,
		Provider:   c.Provider,
		CI:         c.CI,
		Monitoring: c.Monitoring,
		Registry:   c.Registry,
	}
}
