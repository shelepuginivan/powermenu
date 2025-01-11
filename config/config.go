// Package config provides configuration for powermenu.
package config

type AnchorsConfig struct {
	Top    bool `yaml:"top"`
	Bottom bool `yaml:"bottom"`
	Left   bool `yaml:"left"`
	Right  bool `yaml:"right"`
}

type MarginsConfig struct {
	Top    int `yaml:"top"`
	Bottom int `yaml:"bottom"`
	Left   int `yaml:"left"`
	Right  int `yaml:"right"`
}

// Config represents configuration for powermenu.
type Config struct {
	// gtk-layer-shell anchors.
	Anchors AnchorsConfig `yaml:"anchors"`

	// gtk-layer-shell margins.
	Margins MarginsConfig `yaml:"margins"`
}
