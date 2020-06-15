package files

import (
	"context"

	"github.com/hashicorp/waypoint/sdk/component"
	"github.com/hashicorp/waypoint/sdk/terminal"
)

// Builder validates the files in a directory or in the filesystem
type Builder struct {
	config BuilderConfig
}

// BuildFunc implements component.Builder
func (b *Builder) BuildFunc() interface{} {
	return b.Build
}

// BuilderConfig is the configuration structure for the builder.
type BuilderConfig struct{}

// Config implements Configurable
func (b *Builder) Config() (interface{}, error) {
	return &b.config, nil
}

// Build
func (b *Builder) Build(
	ctx context.Context,
	ui terminal.UI,
	src *component.Source,
) (*Files, error) {
	return &Files{
		Path: src.Path,
	}, nil
}
