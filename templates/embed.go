// Package templates embeds all EXO template files into the binary.
// This allows `exo` to work from any directory without needing
// the templates/ directory to be present at runtime.
package templates

import "embed"

// FS is the embedded filesystem containing all templates.
//
//go:embed docker/* k8s/* ci/* monitoring/* terraform/aws/* terraform/gcp/* terraform/azure/* db/*
var FS embed.FS
