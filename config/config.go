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

type Command struct {
	Name string   `yaml:"name"`
	Args []string `yaml:"args"`
}

type CommandsConfig struct {
	Lock      Command `yaml:"lock"`
	Logout    Command `yaml:"logout"`
	Suspend   Command `yaml:"suspend"`
	Hibernate Command `yaml:"hibernate"`
	PowerOff  Command `yaml:"poweroff"`
	Reboot    Command `yaml:"reboot"`
}

// Config represents configuration for powermenu.
type Config struct {
	// gtk-layer-shell anchors.
	Anchors AnchorsConfig `yaml:"anchors"`

	// gtk-layer-shell margins.
	Margins MarginsConfig `yaml:"margins"`

	// Power management commands.
	Commands CommandsConfig `yaml:"commands"`
}
