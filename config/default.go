package config

import (
	"os"
	"path/filepath"

	"github.com/adrg/xdg"
)

var defaultConfig = Config{
	Commands: CommandsConfig{
		Hibernate: Command{
			Name: "systemctl",
			Args: []string{"hibernate"},
		},
		Logout: Command{
			Name: "pkill",
			Args: []string{"-KILL", "-u", os.ExpandEnv("$USER")},
		},
		PowerOff: Command{
			Name: "poweroff",
		},
		Reboot: Command{
			Name: "reboot",
		},
		Suspend: Command{
			Name: "systemctl",
			Args: []string{"suspend"},
		},
	},
	Icons: IconsConfig{
		Hibernate: filepath.Join(xdg.ConfigHome, "powermenu", "icons", "hibernate.svg"),
		Lock:      filepath.Join(xdg.ConfigHome, "powermenu", "icons", "lock.svg"),
		Logout:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "logout.svg"),
		PowerOff:  filepath.Join(xdg.ConfigHome, "powermenu", "icons", "poweroff.svg"),
		Reboot:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "reboot.svg"),
		Suspend:   filepath.Join(xdg.ConfigHome, "powermenu", "icons", "suspend.svg"),
	},
}
