package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
	"github.com/shelepuginivan/powermenu/icons"
)

var defaultConfig = Config{
	Anchors: &AnchorsConfig{
		Top:    false,
		Left:   false,
		Right:  false,
		Bottom: false,
	},
	Margins: &MarginsConfig{
		Top:    0,
		Left:   0,
		Right:  0,
		Bottom: 0,
	},
	Hibernate: &Option{
		Command: &Command{
			Name: "systemctl",
			Args: []string{"hibernate"},
		},
		Enabled: true,
		Icon:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "hibernate.svg"),
		Order:   1,

		GetDefaultIcon: func() []byte { return icons.Hibernate },
	},
	Lock: &Option{
		Command: &Command{},
		Enabled: true,
		Icon:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "lock.svg"),
		Order:   2,

		GetDefaultIcon: func() []byte { return icons.Lock },
	},
	Logout: &Option{
		Command: &Command{
			Name: "pkill",
			Args: []string{"-KILL", "-u", os.ExpandEnv("$USER")},
		},
		Enabled: true,
		Icon:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "logout.svg"),
		Order:   3,

		GetDefaultIcon: func() []byte { return icons.Logout },
	},
	PowerOff: &Option{
		Command: &Command{
			Name: "poweroff",
		},
		Enabled: true,
		Icon:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "poweroff.svg"),
		Order:   4,

		GetDefaultIcon: func() []byte { return icons.PowerOff },
	},
	Reboot: &Option{
		Command: &Command{
			Name: "reboot",
		},
		Enabled: true,
		Icon:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "reboot.svg"),
		Order:   5,

		GetDefaultIcon: func() []byte { return icons.Reboot },
	},
	Suspend: &Option{
		Command: &Command{
			Name: "systemctl",
			Args: []string{"suspend"},
		},
		Enabled: true,
		Icon:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "suspend.svg"),
		Order:   6,

		GetDefaultIcon: func() []byte { return icons.Suspend },
	},
}
