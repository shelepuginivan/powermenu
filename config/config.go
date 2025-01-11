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

type Option struct {
	Command *Command `yaml:"command"`
	Enabled bool     `yaml:"enabled"`
	Icon    string   `yaml:"icon"`
	Order   int      `yaml:"order"`

	GetDefaultIcon func() []byte
}

// Config represents configuration for powermenu.
type Config struct {
	// gtk-layer-shell anchors.
	Anchors *AnchorsConfig `yaml:"anchors"`

	// gtk-layer-shell margins.
	Margins *MarginsConfig `yaml:"margins"`

	Lock      *Option `yaml:"lock"`
	Logout    *Option `yaml:"logout"`
	Suspend   *Option `yaml:"suspend"`
	Hibernate *Option `yaml:"hibernate"`
	PowerOff  *Option `yaml:"poweroff"`
	Reboot    *Option `yaml:"reboot"`
}
