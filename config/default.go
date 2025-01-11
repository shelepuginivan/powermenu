package config

import (
	"path/filepath"

	"github.com/adrg/xdg"
)

var defaultConfig = Config{
	Icons: IconsConfig{
		Hibernate: filepath.Join(xdg.ConfigHome, "powermenu", "icons", "hibernate.svg"),
		Lock:      filepath.Join(xdg.ConfigHome, "powermenu", "icons", "lock.svg"),
		Logout:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "logout.svg"),
		PowerOff:  filepath.Join(xdg.ConfigHome, "powermenu", "icons", "poweroff.svg"),
		Reboot:    filepath.Join(xdg.ConfigHome, "powermenu", "icons", "reboot.svg"),
		Suspend:   filepath.Join(xdg.ConfigHome, "powermenu", "icons", "suspend.svg"),
	},
}
